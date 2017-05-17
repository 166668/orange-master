package controllers

import (
	"log"

	"github.com/ant0ine/go-json-rest/rest"
	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"

	"coding.net/waitfish/orange/models"
	"net/http"
)

//通过 url 中的 hostname 查询 host
func GetHostByName(w rest.ResponseWriter, r *rest.Request) {
	hostname := r.PathParam("host")
	var this models.LiuBei

	exist, host := this.Check_Exist_By_Hostname(hostname)
	if exist {
		w.WriteJson(host)
		return
	}
	rest.Error(w, "Host is not find in database!", 400)

}

//根据 uuid 删除所有匹配的 host
func RemoveHostByUuid(w rest.ResponseWriter, r *rest.Request) {
	uuid := r.PathParam("host")
	//log.Println(hostname)
	var this models.LiuBei
	this.Remove_Host_By_Uuid(uuid)
}

//根据 hostname 删除所有匹配的 host
// func RemoveHostByName(w rest.ResponseWriter, r *rest.Request) {
// 	hostname := r.PathParam("host")
// 	//log.Println(hostname)
// 	var this models.LiuBei
// 	exist, host := this.Check_Exist_By_Hostname(hostname)
// 	if exist {
// 		this.Remove_Host_By_Hostname(hostname)
// 		w.WriteJson(host)
// 		return
// 	}
// 	rest.Error(w, "Host is not find in database!", 400)
// }

//从数据库查询所有 host
func GetAllHosts(w rest.ResponseWriter, r *rest.Request) {
	//召唤刘备
	var liubei models.LiuBei

	hosts := []models.Host{}
	hosts = liubei.Query_Host_For_Api()
	if len(hosts) != 0 {
		w.WriteJson(hosts)
		return
	}
	null := []map[string]bool{{"ok": false}}
	w.WriteJson(null)
}

//从数据库查询 host 个数
func GetHostsCounts(w rest.ResponseWriter, r *rest.Request) {
	//召唤刘备
	var liubei models.LiuBei
	res := make(map[string]int)

	res["counts"] = liubei.Query_Counts_Hosts()
	w.WriteJson(res)

}

//从数据库查询所有检查结果
func GetAllRes_Check(w rest.ResponseWriter, r *rest.Request) {
	var this models.LiuBei
	this.Get_All_Res_From_Db()
	res := this.Res_List_From_Db
	if len(res) != 0 {
		w.WriteJson(res)
		return
	}
	null := []map[string]bool{{"ok": false}}
	w.WriteJson(null)

}

//从数据库查询当前结果
func GetNowRes_Check(w rest.ResponseWriter, r *rest.Request) {
	var this models.LiuBei
	res := this.Get_Res_Now_From_Db()
	//log.Println(res)
	if len(res) != 0 {
		w.WriteJson(res)
		return
	}
	null := []map[string]bool{{"ok": false}}
	w.WriteJson(null)
}

//添加 host 到数据库
func AddHostToDB(w rest.ResponseWriter, r *rest.Request) {
	var this models.LiuBei
	var host models.Host
	var hosted models.Host_From_Web
	err := r.DecodeJsonPayload(&hosted)
	host = hosted.Change_To_Host()
	log.Println(r)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if host.Hostname == "" {
		rest.Error(w, "Hostname is required", 400)
		return
	}
	if ok, _ := this.Check_Exist_By_Hostname(host.Hostname); ok {
		this.Modify_Host_By_Name_For_Api(host)
		// rest.Error(w, "Hostname is already exist in database,pls give another name.", 400)
		w.WriteJson(host)
		return
	}
	if host.Ip == "" {
		rest.Error(w, "IP is required", 400)
		return
	}
	host.Gen_Id()
	this.Add_Host_For_Api(host)
	w.WriteJson(&host)
}
func ModifyHost(w rest.ResponseWriter, r *rest.Request) {
	var this models.LiuBei
	var host models.Host
	err := r.DecodeJsonPayload(&host)

	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	this.Modify_Host_By_Name_For_Api(host)
	w.WriteJson(host)
}
