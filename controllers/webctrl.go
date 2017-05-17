package controllers

import (
	"log"

	"github.com/ant0ine/go-json-rest/rest"
	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"

	"coding.net/waitfish/orange/models"
	"net/http"
)

//添加 web 到数据库
func AddWebToDB(w rest.ResponseWriter, r *rest.Request) {
	var this models.LiuBei
	var web models.Web
	err := r.DecodeJsonPayload(&web)
	log.Println(r)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if web.Name == "" {
		rest.Error(w, "Name is required", 400)
		return
	}
	//if ok, _ := this.Check_Exist_By_Hostname(host.Hostname); ok {
	//	rest.Error(w, "Hostname is already exist in database,pls give another name.", 400)
	//	return
	//}
	if web.Url == "" {
		rest.Error(w, "Url is required", 400)
		return
	}
	if web.Pattern == "" {
		rest.Error(w, "Pattern is required", 400)
		return
	}
	web.Gen_Web_UUID()
	this.Save_Web_To_Db(web)
	w.WriteJson(&web)
}

//从数据库查询 web
func QueryWebApi(w rest.ResponseWriter, r *rest.Request) {
	var this models.LiuBei
	var res []models.Web
	res = this.Query_Web_From_Db()

	if len(res) != 0 {
		w.WriteJson(res)
		return
	}
	null := []map[string]bool{{"ok": false}}
	w.WriteJson(null)
}

//删除 web api
func DeleteWebApi(w rest.ResponseWriter, r *rest.Request) {
	var this models.ZhuGeliang
	uuid := r.PathParam("uuid")
	log.Printf(uuid)
	this.Delete_Web_By_UUID(uuid)
}

//从数据库查询所有 web_status
func QueryWebStatusApi(w rest.ResponseWriter, r *rest.Request) {
	var this models.LiuBei
	var web_status []models.Web_Status
	web_status = this.Query_Web_Status_From_Db()

	if len(web_status) != 0 {
		w.WriteJson(web_status)
		return
	}
	null := []map[string]bool{{"ok": false}}
	w.WriteJson(null)
}

//从数据库查询所有 web 数量
func QueryWebCountsApi(w rest.ResponseWriter, r *rest.Request) {
	var this models.LiuBei
	res := make(map[string]int)

	res["counts"] = this.Query_Counts_Webs()
	w.WriteJson(res)
}

//从数据库当前 web_status
func QueryWebStatusNowApi(w rest.ResponseWriter, r *rest.Request) {
	var this models.LiuBei
	res := this.Get_Web_Status_Now_From_Db()

	if len(res) != 0 {
		w.WriteJson(res)
		return
	}
	null := []map[string]bool{{"ok": false}}
	w.WriteJson(null)
}
