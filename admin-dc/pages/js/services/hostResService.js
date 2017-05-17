


// Restangular service that uses Bing
myApp.factory('hostResService', function(Restangular) {
  return {
	   	GetHostResNow:function(){
			return Restangular.all('api/res_list/now').getList().$object;	

	   	},
	   	GetHostCount:function(){
	   		return Restangular.one('api/host/counts').get().$object;
	   	},
	   	GetHostResHis:function(){
	   		return Restangular.all('api/res_list').getList().$object;
		},
}});