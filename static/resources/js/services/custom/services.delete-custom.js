(function() {

  'use strict';

  angular
    .module('harbor.services.custom')
    .factory('DeleteCustomService', DeleteCustomService);

  DeleteCustomService.$inject = ['$http', '$log'];

  function DeleteCustomService($http, $log) {

    return DeleteCustom;

    function DeleteCustom(repoName, label) {
      var params = (label === '') ? {'repo_name' : repoName} : {'repo_name': repoName, 'label': label};
      return $http
        .delete('/api/custom', {
          'params': params
        });
    }
  }

})();
