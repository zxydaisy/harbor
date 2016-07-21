(function() {

  'use strict';

  angular
    .module('harbor.validator')
    .directive('tagName', tagName);

  tagName.$inject = ['TAG_REGEXP']; //只能为英文的标记名称

  function tagName(TAG_REGEXP) {
    var directive = {
      'require': 'ngModel',
      'link': link
    };
    return directive;

    function link(scope, element, attrs, ctrl) {
      ctrl.$validators.tagName = validator;

      function validator(modelValue, viewValue) {
        return TAG_REGEXP.test(modelValue);
      }
    }
  }

})();
