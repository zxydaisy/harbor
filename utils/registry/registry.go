/*
   Copyright (c) 2016 VMware, Inc. All Rights Reserved.
   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package registry

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/vmware/harbor/utils/log"
	"github.com/vmware/harbor/utils/registry/auth"
	"github.com/vmware/harbor/utils/registry/errors"
)

// Registry holds information of a registry entity
type Registry struct {
	Endpoint *url.URL
	client   *http.Client
}

// NewRegistry returns an instance of registry
func NewRegistry(endpoint string, client *http.Client) (*Registry, error) {
	endpoint = strings.TrimRight(endpoint, "/")

	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	registry := &Registry{
		Endpoint: u,
		client:   client,
	}

	log.Debugf("initialized a registry client: %s", endpoint)

	return registry, nil
}

// NewRegistryWithUsername returns a Registry instance which will authorize the request
// according to the privileges of user
func NewRegistryWithUsername(endpoint, username string) (*Registry, error) {
	endpoint = strings.TrimRight(endpoint, "/")

	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	client, err := newClient(endpoint, username, nil, "registry", "catalog", "*")
	if err != nil {
		return nil, err
	}

	registry := &Registry{
		Endpoint: u,
		client:   client,
	}

	log.Debugf("initialized a registry client with username: %s %s", endpoint, username)

	return registry, nil
}

// Catalog ...
func (r *Registry) Catalog() ([]string, error) {
	repos := []string{}
	suffix := "/v2/_catalog?n=1000"
	var url string

	for len(suffix) > 0 {
		url = r.Endpoint.String() + suffix

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return repos, err
		}
		resp, err := r.client.Do(req)
		if err != nil {
			ok, e := isUnauthorizedError(err)
			if ok {
				return repos, e
			}
			return repos, err
		}

		defer resp.Body.Close()
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return repos, err
		}

		if resp.StatusCode == http.StatusOK {
			catalogResp := struct {
				Repositories []string `json:"repositories"`
			}{}

			if err := json.Unmarshal(b, &catalogResp); err != nil {
				return repos, err
			}

			repos = append(repos, catalogResp.Repositories...)
			//Link: </v2/_catalog?last=library%2Fhello-world-25&n=100>; rel="next"
			link := resp.Header.Get("Link")
			if strings.HasSuffix(link, `rel="next"`) && strings.Index(link, "<") >= 0 && strings.Index(link, ">") >= 0 {
				suffix = link[strings.Index(link, "<")+1 : strings.Index(link, ">")]
			} else {
				suffix = ""
			}
		} else {
			return repos, errors.Error{
				StatusCode: resp.StatusCode,
				StatusText: resp.Status,
				Message:    string(b),
			}
		}
	}
	return repos, nil
}

func newClient(endpoint, username string, credential auth.Credential,
	scopeType, scopeName string, scopeActions ...string) (*http.Client, error) {

	endpoint = strings.TrimRight(endpoint, "/")
	resp, err := http.Get(buildPingURL(endpoint))
	if err != nil {
		return nil, err
	}

	var handlers []auth.Handler
	var handler auth.Handler
	if credential != nil {
		handler = auth.NewStandardTokenHandler(credential, scopeType, scopeName, scopeActions...)
	} else {
		handler = auth.NewUsernameTokenHandler(username, scopeType, scopeName, scopeActions...)
	}

	handlers = append(handlers, handler)

	challenges := auth.ParseChallengeFromResponse(resp)
	authorizer := auth.NewRequestAuthorizer(handlers, challenges)

	transport := NewTransport(http.DefaultTransport, []RequestModifier{authorizer})
	return &http.Client{
		Transport: transport,
	}, nil
}
