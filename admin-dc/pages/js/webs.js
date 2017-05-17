/**
 * Created by waitfish on 15/1/23.
 */
'use strict';

angular.module('myApp.webs', ['ngRoute','restangular'])



.config(['$routeProvider', function($routeProvider) {
    $routeProvider.when('/pages/webs', {
    templateUrl: '/pages/webs.html',
    controller: 'WebCtrl'
    });

    $routeProvider.when('/pages/addweb', {
    templateUrl: '/pages/addweb.html',
    controller: 'AddWebCtrl'
    });
    $routeProvider.when('/pages/webs_status_history', {
    templateUrl: '/pages/web_status_history.html',
    controller: 'WebStatusCtrl'
    });
    $routeProvider.when('/pages/webs_status_now', {
    templateUrl: '/pages/web_status_now.html',
    controller: 'WebStatusNowCtrl'
    });
}])

.controller('WebCtrl', ['$scope','Restangular',function($scope,Restangular) {
    var baseWebs = Restangular.all('api/web');
    $scope.webs=baseWebs.getList().$object;

    $scope.del=function(web,uuid){
    Restangular.one('api/web',uuid).remove().then(function(){
        var index=$scope.webs.indexOf(web);
        if(index > -1) $scope.webs.splice(index,1);
    })
}

}])

.controller('AddWebCtrl', ['$scope','Restangular',function($scope,Restangular) {
    var baseWebs = Restangular.all('api/web');
    $scope.master = {};

    $scope.update = function(web) {
        $scope.master = angular.copy(web);
        baseWebs.post($scope.master);
    };

    $scope.reset = function() {
        $scope.host = angular.copy($scope.master);
    };
}])

.controller('WebStatusNowCtrl', ['$scope','Restangular',function($scope,Restangular) {
    $scope.results=Restangular.all('api/web_status/now').getList().$object;
}])

.controller('WebStatusCtrl', ['$scope','Restangular',function($scope,Restangular) {
    var baseRes = Restangular.all('api/web_status');
    $scope.results=baseRes.getList().$object;


}]);