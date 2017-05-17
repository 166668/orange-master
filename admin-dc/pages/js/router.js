myApp.config(function($stateProvider, $urlRouterProvider){
        $urlRouterProvider.otherwise("login")
        $stateProvider
            .state('index',{
                url:"/index",
                views:{
                '':{
                    templateUrl:'pages/index.html'
                },
                'nav@index':{
                    templateUrl:'pages/nav.html'
                },
                'content@index':{
                    templateUrl:'pages/hosts.html'
                },

            }                  
            })
            
            .state('login',{
                url:"/login",
                views:{
                '':{
                    templateUrl:'pages/login.html'
                },
            }                  
            })

            .state('register',{
                url:"/register",
                views:{
                '':{
                    templateUrl:'pages/register.html'
                },
            }                  
            })

            .state('index.hosts',{
                url:"/hosts",
                views:{
                'content@hosts': {
                    templateUrl:'pages/hosts.html'
                }
            }             
            })

            .state('index.addhost',{
                url:"/addhost",
                views:{
                'content@index': {
                    templateUrl:'pages/addhost.html'
                }
            }             
            })


            .state('index.res',{
                url:"/res",
                views:{
                'content@index': {
                    templateUrl:'pages/res_now.html'
                }
            }             
            })

            .state('index.reshis',{
                url:"/reshis",
                views:{
                'content@index': {
                    templateUrl:'pages/res_history.html'
                }
            }             
            })

            .state('index.webs',{
                url:"/webs",
                views:{
                'content@index': {
                    templateUrl:'pages/webs.html'
                }
            }             
            })

            .state('index.addweb',{
                url:"/addweb",
                views:{
                'content@index': {
                    templateUrl:'pages/addweb.html'
                }
            }             
            })

            .state('index.webstatusnow',{
                url:"/webstatusnow",
                views:{
                'content@index': {
                    templateUrl:'pages/web_status_now.html'
                }
            }             
            })

            .state('index.webstatushis',{
                url:"/webstatushis",
                views:{
                'content@index': {
                    templateUrl:'pages/web_status_history.html'
                }
            }             
            })
    });