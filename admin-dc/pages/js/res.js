/**
 * Created by waitfish on 15/1/23.
 */
'use strict';

angular.module('myApp.res', ['ngRoute','restangular'])



.config(['$routeProvider', function($routeProvider) {
    $routeProvider.when('/pages/res', {
    templateUrl: '/pages/res_now.html',
    controller: 'ResNowCtrl'
    });
    $routeProvider.when('/pages/res_history', {
    templateUrl: '/pages/res_history.html',
    controller: 'ResCtrl'
    });
}])

.controller('ResNowCtrl', ['$scope','Restangular',function($scope,Restangular) {
    $scope.now_results=Restangular.all('api/res_list/now').getList().$object;
    $scope.counts_hosts =Restangular.one('api/host/counts').get().$object;


}])

.controller('ResCtrl', ['$scope','Restangular',function($scope,Restangular) {
    var baseRes = Restangular.all('api/res_list');
    $scope.now_results=Restangular.all('api/res_list/now').getList().$object;
    $scope.results=baseRes.getList().$object;
    $scope.counts_hosts =Restangular.one('api/host/counts').get().$object;


}]);