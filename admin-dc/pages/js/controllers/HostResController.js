/**
 * Created by waitfish on 15/1/23.
 */
'use strict';
myApp.controller('HostResListNowController', ['$scope','hostResService','Restangular',function($scope,hostResService,Restangular) {

    $scope.now_results=hostResService.GetHostResNow()
    $scope.counts_hosts =hostResService.GetHostCount()
}])

myApp.controller('HostResHistoryController', ['$scope','hostResService','Restangular',function($scope,hostResService,Restangular) {

    $scope.results=hostResService.GetHostResHis()
}])