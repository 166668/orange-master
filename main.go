package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
	"sync"
	"time"

	"coding.net/waitfish/orange/controllers"
	"coding.net/waitfish/orange/models"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		var timer_2s = time.NewTicker(time.Second * 5)
		var liubei models.LiuBei
		for {
			select {
			case <-timer_2s.C:
				liubei.Do_Port_Check()
			}
		}

		wg.Done()

	}()

	go func() {
		var timer_2s = time.NewTicker(time.Second * 5)
		var liubei models.LiuBei
		for {
			select {
			case <-timer_2s.C:
				liubei.Do_Web_Check()
			}
		}

		wg.Done()

	}()

	go func() {

		api := rest.NewApi()
		api.Use(rest.DefaultDevStack...)
		api.Use(&rest.CorsMiddleware{
			RejectNonCorsRequests: false,
			OriginValidator: func(origin string, request *rest.Request) bool {
				return origin == "http://127.0.0.1:8081"
			},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
			AllowedHeaders: []string{
				"Accept", "Content-Type", "X-Custom-Header", "Origin"},
			AccessControlAllowCredentials: true,
			AccessControlMaxAge:           3600,
		})

		// api.Use(&rest.AuthBasicMiddleware{
		//       Realm: "test zone",
		//       Authenticator: controllers.CheckLoginApi
		//   })

		router, err := rest.MakeRouter(
			&rest.Route{"GET", "/host", controllers.GetAllHosts},
			&rest.Route{"GET", "/host/counts", controllers.GetHostsCounts},
			&rest.Route{"POST", "/host", controllers.AddHostToDB},
			&rest.Route{"POST", "/host/:host", controllers.RemoveHostByUuid},
			&rest.Route{"GET", "/host/:host", controllers.GetHostByName},

			&rest.Route{"GET", "/res_list", controllers.GetAllRes_Check},
			&rest.Route{"GET", "/res_list/now", controllers.GetNowRes_Check},

			&rest.Route{"POST", "/web", controllers.AddWebToDB},
			&rest.Route{"GET", "/web", controllers.QueryWebApi},
			&rest.Route{"DELETE", "/web/:uuid", controllers.DeleteWebApi},
			&rest.Route{"GET", "/web/counts", controllers.QueryWebCountsApi},

			&rest.Route{"GET", "/web_status", controllers.QueryWebStatusApi},
			&rest.Route{"GET", "/web_status/now", controllers.QueryWebStatusNowApi},

			&rest.Route{"POST", "/user/register", controllers.AddUserToDB},
			&rest.Route{"POST", "/user/login", controllers.CheckLoginApi},
			&rest.Route{"GET", "/user", controllers.QueryUserApi},
			// &rest.Route{"POST", "/web", controllers.AddWebToDB},
		)

		if err != nil {
			log.Fatal(err)
		}

		api.SetApp(router)
		//http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./admin-dc"))))
		http.Handle("/api/", http.StripPrefix("/api", api.MakeHandler()))

		log.Fatal(http.ListenAndServe(":8080", nil))

		wg.Done()

	}()

	wg.Wait()

}
