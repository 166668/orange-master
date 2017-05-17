/**
 * Created by waitfish on 15/1/23.
 */
'use strict';

myApp.controller('WebListController', ['$scope','webService','Restangular',function($scope,webService,Restangular) {
    
    $scope.webs=webService.GetWebList();
}])


myApp.controller('WebStatusNowController', ['$scope','webService','Restangular',function($scope,webService,Restangular) {
    
    $scope.results=webService.GetWebStatusNow();
}])

myApp.controller('WebAddController', ['$scope','Restangular',function($scope,Restangular) {
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