/**
 * Created by waitfish on 15/1/23.
 */
'use strict';

angular.module('myApp.hosts', ['ngRoute','restangular'])



.config(['$routeProvider', function($routeProvider) {
    $routeProvider.when('/pages/hosts', {
    templateUrl: '/pages/hosts.html',
    controller: 'HostCtrl'
    });

    $routeProvider.when('/pages/addhost', {
    templateUrl: '/pages/addhost.html',
    controller: 'AddHostCtrl'
    });

    $routeProvider.when('/host/:name', {
    templateUrl: '/pages/host.html',
    controller: 'ModifyHostCtrl'
    });
}])

.controller('HostCtrl', ['$scope','Restangular',function($scope,Restangular) {
    var baseHosts = Restangular.all('api/host');
    
    $scope.hosts=baseHosts.getList().$object;

    $scope.del=function(host,uuid){
        Restangular.one('api/host',uuid).post().then(function(){
            var index=$scope.hosts.indexOf(host);
            if(index > -1) $scope.hosts.splice(index,1);
        })
        }
}])

.controller('AddHostCtrl', ['$scope','Restangular',function($scope,Restangular) {
    var baseHosts = Restangular.all('api/host');
    $scope.master = {};

    $scope.update = function(host) {
        $scope.master = angular.copy(host);
        baseHosts.post($scope.master);
    };

    $scope.reset = function() {
        $scope.host = angular.copy($scope.master);
    };
}])

.controller('ModifyHostCtrl', ['$scope','Restangular','$routeParams',function($scope,Restangular,$routeParams) {
    var baseHosts = Restangular.all('api/host');
    var Hostname =$routeParams.name;
    $scope.host=Restangular.one('api/host',Hostname).get().$object;

    //$scope.host={};
    //$scope.host.Hostname=$scope.hosted.Hostname;
    //$scope.host.Ip=$scope.hosted.Ip;
    //$scope.host.Tcp_list=$scope.hosted.Tcp_list.join(",");
    //$scope.host.Udp_list=$scope.hosted.Udp_list.join(",");

    
    $scope.master = {};

    $scope.update = function(host) {
        $scope.master = angular.copy(host);
        baseHosts.post($scope.master);
    };

    $scope.reset = function() {
        $scope.host = angular.copy($scope.master);
    };


}]);
