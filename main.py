import requests
import json
import time
import multiprocessing

host = 'http://127.0.0.1:7099/'
#host = 'http://118.193.185.187:7099/'
#host = 'http://118.193.143.243/'
#host = 'http://192.168.1.28:7099/'
#host = 'https://fornax-canary.caicloud.io/'

headers = {
    "Accept": "application/json",
    "Content-type": "application/json;charset=utf-8",
}

data = {
    "name": "hello",
    "description": "test",
    "username": "test",
    "repository": {
        "url": "https://github.com/fsouza/go-dockerclient.git",
        "vcs": "git"
    },
    "build_path": ""
}

def health_check():
    url = host + 'api/v0.1/healthcheck'
    r = requests.post(url, headers=headers)
    print r.status_code, r.text

def create_service(user_id, service_name, git_repo):
    _data = {
        "name": service_name,
        "description": "test",
        "username": "test",
        "operation": "integrationpublish",
        "repository": {
            #"url": "/home/superxi/gopath/src/github.com/caicloud/console-web",
            "url": "https://github.com/superxi911/TestCase.git", 
            "vcs": "git"
            #"url": "svn://118.193.185.187/svn-demo/trunk", 
            #"vcs": "svn",
            #"username": "superxi",
            #"password": "xxxpwd"
        },
        "build_path": ""
    }
    url = host + 'api/v0.1/{user_id}/services'
    r = requests.post(url.format(**{'user_id': user_id}), headers=headers, data=json.dumps(_data))
    print r.status_code, r.text

def delete_service(user_id, service_id):
    url = host + 'api/v0.1/{user_id}/services/{service_id}'
    r = requests.delete(url.format(**{'user_id': user_id,'service_id': service_id}), headers=headers)
    print r.status_code, r.text

def get_services(user_id):
    url = host + 'api/v0.1/{user_id}/services'
    r = requests.get(url.format(**{'user_id': user_id}), headers=headers)
    print r.status_code, r.text


def create_version(uid, service_id):
    url = host + 'api/v0.1/{uid}/versions'

    _data = {
        "name": "d89u8i523okiklk0fjn9mkjdgnhafrkddud5291by",
        "description": "v3",
        "service_id": service_id
    }

    r = requests.post(url.format(uid=uid), headers=headers, data=json.dumps(_data))
    print r.status_code, r.text

def get_versions(user_id, service_id):
    url = host + 'api/v0.1/{user_id}/services/{service_id}/versions'
    r = requests.get(url.format(**{'user_id': user_id,'service_id': service_id}), headers=headers)
    print r.status_code, r.text

def worker(num):
    """thread worker function"""
    print 'Worker:', num
    while 1:
        create_service('superxi', 'test_service', 'https://github.com/superxi911/console-web.git')
        time.sleep(0.1)
    return

def create_project(user_id, project_name):
    _data = {
        "name": project_name,
        "description": "test",
    }
    url = host + 'api/v0.1/{user_id}/projects'
    r = requests.post(url.format(**{'user_id': user_id}), headers=headers, data=json.dumps(_data))
    print r.status_code, r.text

def get_project(user_id, project_id):
    url = host + 'api/v0.1/{user_id}/projects/{project_id}'
    r = requests.get(url.format(**{'user_id': user_id, 'project_id': project_id}), headers=headers)
    print r.status_code, r.text

def get_projects(user_id):
    url = host + 'api/v0.1/{user_id}/projects'
    r = requests.get(url.format(**{'user_id': user_id}), headers=headers)
    print r.status_code, r.text

def set_project(user_id, project_id):
    url = host + 'api/v0.1/{user_id}/projects/{project_id}'
    _data = {
        "services" : [
            {
                "service_id" : "288664e6-55ea-41ef-afda-c8bf7bc8d8e9", 
                "depend" : [
                    {"service_id" : "94c936c6-ebc5-43bb-bf38-2772f431ec4e"}
                ]
            }, 
            {
                "service_id" : "288664e6-55ea-41ef-afda-c8bf7bc8d8e9", 
                "depend" : [
                    {"service_id" : "94c936c6-ebc5-43bb-bf38-2772f431ec4e"}
                ]
            }, 
            {
                "service_id" : "288664e6-55ea-41ef-afda-c8bf7bc8d8e9", 
                "depend" : [
                    {"service_id" : "94c936c6-ebc5-43bb-bf38-2772f431ec4e"}
                ]
            }
        ]
    }
    r = requests.put(url.format(**{'user_id': user_id,'project_id': project_id}), headers=headers, data=json.dumps(_data))
    print r.status_code, r.text

def delete_project(user_id, project_id):
    url = host + 'api/v0.1/{user_id}/projects/{project_id}'
    r = requests.delete(url.format(**{'user_id': user_id,'project_id': project_id}), headers=headers)
    print r.status_code, r.text

def create_project_version(user_id, project_id, version_name):
    _data = {
        "project_id": project_id,
        "policy": "manual", 
        "name": version_name, 
        "description": "test",
    }
    url = host + 'api/v0.1/{user_id}/versions_project'
    r = requests.post(url.format(**{'user_id': user_id}), headers=headers, data=json.dumps(_data))
    print r.status_code, r.text

def get_project_versions(user_id, project_id):
    url = host + 'api/v0.1/{user_id}/projects/{project_id}/versions'
    r = requests.get(url.format(**{'user_id': user_id,'project_id': project_id}), headers=headers)
    print r.status_code, r.text


if __name__ == '__main__':

    #health_check()
    #create_service('superxi', 'test1', 'https://github.com/superxi911/TestCase.git')

    #get_services('c818874c-3ade-4525-a0ba-87b553fb7abd')
    #get_services('superxi')
    #delete_service("superxi", "d99946d5-bc34-43b3-8b8e-94e926f87519")

    #service_id = 'f952c3ea-7c27-4b32-8b15-e871dff064e9'
    #create_version('superxi', service_id)

    #get_versions('superxi', service_id)

    #for i in range(5):
    #    p = multiprocessing.Process(target=worker, args=(i,))
    #    p.start()

    #create_project("superxi", "test")
    #delete_project("superxi", "bf6afce1-be09-40ab-91d5-68a3129ba5c1")
    set_project("superxi", "601780f2-266d-4072-9ff5-7b6b4ab4452b")
    #create_project_version("superxi", "e22144ee-ba66-4424-a405-054ddd749830", "v1")
    get_projects("superxi")
    #get_project("superxi", "a0298c65-5237-4ec1-8af7-f4056e290adf")
    #get_project_versions("superxi", "e22144ee-ba66-4424-a405-054ddd749830")

    pass
