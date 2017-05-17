


// Restangular service that uses Bing
myApp.factory('hostService', function(Restangular) {
	return {
	   	GetHostList:function(){
			return Restangular.all('api/host').getList().$object;	

	   	},
	   	DelHost:function(host,uuid){
	   		Restangular.one('api/host',uuid).post().then(function(){
	   	    var index=$scope.hosts.indexOf(host);
	   	    if(index > -1) $scope.hosts.splice(index,1);
	   		})
		},
}})