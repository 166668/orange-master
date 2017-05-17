
myApp.factory('webService', function(Restangular) {
   return {
   	GetWebList:function(){
   		var baseWebs = Restangular.all('api/web');
		return baseWebs.getList().$object;	

   	},
   	DelWeb:function(web,uuid){
   		Restangular.one('api/web',uuid).post().then(function(){
   	    var index=$scope.webs.indexOf(web);
   	    if(index > -1) $scope.webs.splice(index,1);
   		})
	},
	GetWebStatusNow: function(){
		return Restangular.all('api/web_status/now').getList().$object;
	},
}})