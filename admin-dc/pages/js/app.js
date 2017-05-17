/**
 * Created by waitfish on 15/1/23.
 */
'use strict';

// Declare app level module which depends on views, and components
var myApp = angular.module('myApp', [
    "restangular",
    'ui.router',
    'ngMessages',
    'myApp.dash',
    'myApp.hosts',
    'myApp.res',
    'myApp.webs'
])
.constant('baseServiceUrl', 'http://127.0.0.1:8080')


.config(['$locationProvider',function($locationProvider){
    $locationProvider.html5Mode(true);
}])
myApp.config(function(RestangularProvider){
    RestangularProvider.setBaseUrl('http://127.0.0.1:8080');
    // RestangularProvider.setDefaultHeaders({'Access-Control-Allow-Origin':"http://127.0.0.1:8081"});
});