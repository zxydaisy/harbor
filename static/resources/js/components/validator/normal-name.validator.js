(function() {

  'use strict';

  angular
    .module('harbor.validator')
    .directive('normalName', normalName); //可以为中文或者英文的名称

  normalName.$inject = ['TAG_REGEXP'];

  function normalName(TAG_REGEXP, CH_REGEXP) {
    var directive = {
      'require': 'ngModel',
      'link': link
    };
    return directive;

    function link(scope, element, attrs, ctrl) {
      ctrl.$validators.tagName = validator;

      function validator(modelValue, viewValue) {
        return TAG_REGEXP.test(modelValue) || CH_REGEXP.test(modelValue);
      }
    }
  }

})();
