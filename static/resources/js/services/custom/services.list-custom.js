(function() {

  'use strict';

  angular
    .module('harbor.services.custom')
    .factory('ListCustomService', ListCustomService);

  ListCustomService.$inject = ['$http', '$log'];

  function ListCustomService($http, $log) {
    return ListCustom;

    function ListCustom(projectId, q) {
      $log.info('list repositories:' + projectId + ', q:' + q);
      //如果传入projectid，就根据projectId查找，否则查询全部
      return $http
        .get('/api/customer', {
          'params':{
            'project_id': projectId,
            'q': q
          }
      });
    }
  }

})();
