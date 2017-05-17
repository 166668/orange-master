'use strict';			

// myApp.controller('RegisterController', ['$scope', function($scope) {
//       $scope.master = {};

//       $scope.update = function(user) {
//         $scope.master = angular.copy(user);
//       };

//       $scope.reset = function() {
//         $scope.user = angular.copy($scope.master);
//       };

//       $scope.reset();
//     }]);

myApp.controller('RegisterController',
    function ($scope, $rootScope, $location, authService, notifyService) {
        $rootScope.pageTitle = "Register";


        $scope.register = function(userData) {
            authService.register(userData,
                function success() {
                    notifyService.showInfo("User registered successfully");
                    // $.notify("register done!") 
                    $location.path("/");
                },
                function error(err) {
                    notifyService.showError(err);
                     $location.path("#");
                }
            );
        };
    }
);
