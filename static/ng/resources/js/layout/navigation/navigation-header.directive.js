(function() {

  'use strict';

  angular
    .module('harbor.layout.navigation')
    .directive('navigationHeader', navigationHeader);
  
  NavigationHeaderController.$inject = ['$window', '$scope', 'currentUser', '$timeout'];
    
  function NavigationHeaderController($window, $scope, currentUser, $timeout) {
    var vm = this;
    vm.url = $window.location.pathname;    
  }
  
  function navigationHeader() {
    var directive = {
      restrict: 'E',
      templateUrl: '/ng/navigation_header',
      link: link,
      scope: true,
      controller: NavigationHeaderController,
      controllerAs: 'vm',
      bindToController: true
    };
    
    return directive;
   
    function link(scope, element, attrs, ctrl) {     
      var visited = ctrl.url;
      if (visited !== "/ng") {
         element.find('a[href*="' + visited + '"]').addClass('active'); 
      }      
      element.find('a').on('click', click);
      function click(event) {
        element.find('a').removeClass('active');
        $(event.target).not('span').addClass('active');
      }     
    }
   
  }
  
})();