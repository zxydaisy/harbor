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
(function() {

  'use strict';

  angular
    .module('harbor.repository')
    .directive('listTag', listTag);

  ListTagController.$inject = ['$scope', 'ListTagService', 'ListLabelService', 'DeleteLabelService', '$filter', 'trFilter'];

  function ListTagController($scope, ListTagService, ListLabelService, DeleteLabelService, $filter, trFilter) {
    var vm = this;

    vm.tags = [];
    vm.retrieve = retrieve;

    $scope.$watch('vm.repoName', function(current, origin) {
      if(current) {
        console.log('vm.repoName in tags:' + current);
        vm.retrieve();
      }
    });

    $scope.$on('refreshTagsAndLabels', function(e, val) {
      if(val) {
        vm.retrieve();
      }
    });


    vm.deleteTag = deleteTag;
    vm.deleteLabel = deleteLabel;
    vm.isOpen = false;
    vm.showAddLabel = showAddLabel;


    function retrieve() {
      ListTagService(vm.repoName)
        .success(getTagSuccess)
        .error(getTagFailed);

      ListLabelService(vm.repoName)
        .success(getLabelSuccess)
        .error(getLabelFailed);
    }

    function showAddLabel() {
      if(vm.isOpen){
        vm.isOpen = false;
      }else{
        vm.isOpen = true;
      }
    }


    function getLabelSuccess() {
      vm.labels = data || [];
      vm.labelCount[vm.repoName] = vm.labels.length;

      $scope.$emit('labels', vm.labels);
      $scope.$emit('labelCount', vm.labelCount);

      angular.forEach(vm.labels, function(item) {
        vm.toggleInProgress[vm.repoName + '|' + item] = false;
      });
    }

    function getLabelFailed(data) {
      $scope.$emit('modalTitle', $filter('tr')('error'));
      $scope.$emit('modalMessage', $filter('tr')('failed_to_get_tag') + data);
      $scope.$emit('raiseError', true);
      console.log('Failed to get tag:' + data);
    }

    //通过label删除代码吧
    function deleteByLabel() {
      $scope.$emit('modalTitle', $filter('tr')('alert_delete_tag_title', [vm.label]));
      var message;
      $scope.$emit('modalMessage',  $filter('tr')('alert_delete_tag', [vm.label]));

      var emitInfo = {
        'confirmOnly': false,
        'contentType': 'text/html',
        'action' : vm.deleteLabel
      };

      $scope.$emit('raiseInfo', emitInfo);
    }


    function deleteLabel() {
      console.log('Delete image, repoName:' + vm.repoName + ', label:' + vm.label);
      vm.toggleInProgress[vm.repoName + '|' + vm.label] = true;
      DeleteLabelService(vm.repoName, vm.label)
        .success(deleteLabelSuccess)
        .error(deleteLabelFailed);
    }

    function deleteLabelSuccess(data, status) {
      vm.toggleInProgress[vm.repoName + '|' + vm.label] = false;
      vm.retrieve();
    }

    function deleteLabelFailed(data, status) {
      vm.toggleInProgress[vm.repoName + '|' + vm.label] = false;

      $scope.$emit('modalTitle', $filter('tr')('error'));
      var message;
      if(status === 401) {
        message = $filter('tr')('failed_to_delete_repo_insuffient_permissions');
      }else{
        message = $filter('tr')('failed_to_delete_repo');
      }
      $scope.$emit('modalMessage', message);
      $scope.$emit('raiseError', true);
    }

    function getTagSuccess(data) {

      vm.tags = data || [];
      vm.tagCount[vm.repoName] = vm.tags.length;

      $scope.$emit('tags', vm.tags);
      $scope.$emit('tagCount', vm.tagCount);

      angular.forEach(vm.tags, function(item) {
        vm.toggleInProgress[vm.repoName + '|' + item] = false;
      });
    }

    function getTagFailed(data) {
      $scope.$emit('modalTitle', $filter('tr')('error'));
      $scope.$emit('modalMessage', $filter('tr')('failed_to_get_tag') + data);
      $scope.$emit('raiseError', true);
      console.log('Failed to get tag:' + data);
    }

    function deleteTag(e) {
      $scope.$emit('repoName', e.repoName);
      $scope.$emit('tag', e.tag);
      vm.deleteByTag();
    }

  }

  function listTag() {
    var directive = {
      'restrict': 'E',
      'templateUrl': '/static/resources/js/components/repository/list-tag.directive.html',
      'scope': {
        'tagCount': '=',
        'associateId': '=',
        'repoName': '=',
        'toggleInProgress': '=',
        'deleteByTag': '&'
      },
      'replace': true,
      'controller': ListTagController,
      'controllerAs': 'vm',
      'bindToController': true
    };
    return directive;

  }

})();