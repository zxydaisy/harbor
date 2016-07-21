(function() {
  'use strict';

  angular
    .module('harbor.services.custom')
    .factory('AddCustomService', AddCustomService);

  AddCustomService.$inject = ['$http', '$log'];

  function AddCustomService($http, $log) {

    return AddCustom;

    function AddCustom(repoName, labelName) {
      var params = (labelName === '') ? {'repo_name' : repoName} : {'repo_name': repoName, 'label_name': labelName};
      return $http
        .post('/api/custom', params);
    }
  }

})();
