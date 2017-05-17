/**
 * Created by waitfish on 15/1/23.
 */
'use strict';

angular.module('myApp.dash', ['ngRoute'])



.config(['$routeProvider', function($routeProvider) {
  $routeProvider.when('/pages/dash', {
    templateUrl: '/pages/dash.html',
    controller: 'DashCtrl'
  });
}])

.controller('DashCtrl', ['$scope','$http',function($scope,$http) {
    $http.get('http://127.0.0.1:8080/api/host').success(function(data){
        $scope.hosts=data;
    });
}]);