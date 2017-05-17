/**
 * Created by waitfish on 15/1/23.
 */
'use strict';
myApp.controller('HostListController', ['$scope','hostService','Restangular',function($scope,hostService,Restangular) {
    
    $scope.hosts=hostService.GetHostList()
    $scope.del=function(host,uuid){
        Restangular.one('api/host',uuid).post().then(function(){
            var index=$scope.hosts.indexOf(host);
            if(index > -1) $scope.hosts.splice(index,1);
        })
        }
}])

myApp.controller('HostAddController', ['$scope','Restangular',function($scope,Restangular) {

    $scope.master = {};

    $scope.update = function(host) {
        $scope.master = angular.copy(host);
        Restangular.all('api/host').post($scope.master);
    };

    $scope.reset = function() {
        $scope.host = angular.copy($scope.master);
    };
}])

