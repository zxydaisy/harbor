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

      return $http
        .get('/api/repositories', {
          'params':{
            'project_id': projectId,
            'q': q
          }
      });
    }
  }

})();
