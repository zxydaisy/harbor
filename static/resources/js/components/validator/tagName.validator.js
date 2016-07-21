(function() {

  'use strict';

  angular
    .module('harbor.validator')
    .directive('name', projectName);

  tagName.$inject = ['PROJECT_REGEXP', 'CH_REGEXP'];

  function tagName(PROJECT_REGEXP, CH_REGEXP) {
    var directive = {
      'require': 'ngModel',
      'link': link
    };
    return directive;

    function link(scope, element, attrs, ctrl) {
      ctrl.$validators.tagName = validator;

      function validator(modelValue, viewValue) {
        return PROJECT_REGEXP.test(modelValue) || CH_REGEXP.test(modelValue);
      }
    }
  }

})();
