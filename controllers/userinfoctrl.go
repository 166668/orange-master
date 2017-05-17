package controllers

import (
	// "log"

	"github.com/ant0ine/go-json-rest/rest"
	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"

	"coding.net/waitfish/orange/models"
	"net/http"
)

func CheckLoginApi(w rest.ResponseWriter, r *rest.Request) {
	var this models.LiuBei
	var login_user models.UserInfo

	err := r.DecodeJsonPayload(&login_user)

	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := make(map[string]interface{})

	if this.Check_Login(login_user) {
		data["login"] = true
		data["useremail"] = login_user.UserEmail
		w.WriteJson(data)
		return
	}
	rest.Error(w, "Login failed!", 403)

}

func AddUserToDB(w rest.ResponseWriter, r *rest.Request) {
	var this models.LiuBei
	var user models.UserInfo

	err := r.DecodeJsonPayload(&user)
	// log.Println(r)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user.UserEmail == "" {
		rest.Error(w, "Email is required", 400)
		return
	}
	//if ok, _ := this.Check_Exist_By_Hostname(host.Hostname); ok {
	//	rest.Error(w, "Hostname is already exist in database,pls give another name.", 400)
	//	return
	//}
	if user.UserPassword == "" {
		rest.Error(w, "Password is required", 400)
		return
	}

	user.Gen_Id()
	this.Add_User_For_Api(user)
	w.WriteJson(&user)
}

/**
 * [QueryUserApi Query all user info from db for api]
 */
func QueryUserApi(w rest.ResponseWriter, r *rest.Request) {
	var this models.LiuBei
	var res []models.UserInfo
	res = this.Query_User_From_Db()

	if len(res) != 0 {
		w.WriteJson(res)
		return
	}
	null := []map[string]bool{{"ok": false}}
	w.WriteJson(null)
}
