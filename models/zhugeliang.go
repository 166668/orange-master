package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

//定义数据库信息
const (
	DB_URL = "to_db"
	//DB_URL = "xmdx:171"

	DB_NAME = "orange"
)

//所有跟数据库相关的操作均由诸葛亮来完成
type ZhuGeliang struct {
	//Mob              map[string]interface{}
	Res_List_From_Db []Check_Res
}

//删除所有匹配 hostname 的 host 记录
func (this *ZhuGeliang) Remove_Host_By_Uuid(Uuid string) {
	//新建一个 session
	shouyu := this.Get_A_DB_Session()
	defer shouyu.Close()
	//选择一个 collection
	c := shouyu.DB("orange").C("hosts")

	_, err := c.RemoveAll(bson.M{"uuid": Uuid})
	if err != nil {
		log.Println(err)
	}
	log.Printf("%s is remove from database successful", Uuid)
}

//删除所有匹配 hostname 的 host 记录
func (this *ZhuGeliang) Remove_Host_By_Hostname(hostname string) {
	//新建一个 session
	shouyu := this.Get_A_DB_Session()
	defer shouyu.Close()
	//选择一个 collection
	c := shouyu.DB("orange").C("hosts")

	_, err := c.RemoveAll(bson.M{"hostname": hostname})
	if err != nil {
		log.Println(err)
	}
	log.Printf("%s is remove from database successful", hostname)
}

//检查数据库中是否存在该 host 记录，通过 hostname （每个 host 都有 hostname ） 查询，若存在，则返回 True，否则返回 false
func (this *ZhuGeliang) Check_Exist_By_Hostname(hostname string) (ok bool, host Host) {
	//新建一个 session
	shouyu := this.Get_A_DB_Session()
	defer shouyu.Close()
	//选择一个 collection
	c := shouyu.DB("orange").C("hosts")

	host_list := []Host{}
	err := c.Find(bson.M{"hostname": hostname}).All(&host_list)

	if err != nil {
		log.Println(err)
	}
	if len(host_list) != 0 {
		ok = true
		log.Printf("The hostname is existed")
		err := c.Find(bson.M{"hostname": hostname}).One(&host)
		if err != nil {
			log.Fatalln(err)
		}
		return
	}
	ok = false
	return
}

//添加主机到数据库，供 api 调用
func (this *ZhuGeliang) Add_Host_For_Api(host Host) {
	//新建一个 session
	shouyu := this.Get_A_DB_Session()
	defer shouyu.Close()
	//选择一个 collection
	c := shouyu.DB("orange").C("hosts")

	err := c.Insert(host)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Add host into db success!")
}

//根据主机名修改主机，供 api 调用
func (this *ZhuGeliang) Modify_Host_By_Name_For_Api(host Host) {
	//新建一个 session
	shouyu := this.Get_A_DB_Session()
	defer shouyu.Close()
	//选择一个 collection
	c := shouyu.DB("orange").C("hosts")

	err := c.Update(bson.M{"hostname": host.Hostname}, bson.M{})
	if err != nil {
		log.Println(host)
		log.Printf("出错啦！！！")
		log.Println(err)
	}
	log.Printf("Modify host  success!")
}

//从数据库查询主机，供 api 调用
func (this *ZhuGeliang) Query_Host_For_Api() (hosts []Host) {
	//新建一个 session
	shouyu := this.Get_A_DB_Session()
	defer shouyu.Close()
	//选择一个 collection
	c := shouyu.DB("orange").C("hosts")

	c.Find(nil).All(&hosts)
	return
}

func (this *ZhuGeliang) Query_Counts_Hosts() (num int) {
	//新建一个 session
	shouyu := this.Get_A_DB_Session()
	defer shouyu.Close()
	//选择一个 collection
	hosts := []Host{}
	c := shouyu.DB("orange").C("hosts")

	c.Find(nil).All(&hosts)
	num = len(hosts)
	return
}

//保存查询结果到数据库
func (this *ZhuGeliang) Save_Res_Check(check_res Check_Res) {
	//新建一个 session
	shouyu := this.Get_A_DB_Session()
	defer shouyu.Close()
	//选择一个 collection
	c := shouyu.DB("orange").C("res_check")

	err := c.Insert(check_res)
	if err != nil {
		log.Println(err)
	}
	// log.Printf("Result install into db success!", check_res)
}

//从数据库查询所有检查结果数据
func (this *ZhuGeliang) Get_All_Res_From_Db() {
	//新建一个 session
	shouyu := this.Get_A_DB_Session()
	defer shouyu.Close()

	c := shouyu.DB("orange").C("res_check")

	c.Find(nil).All(&this.Res_List_From_Db)
}

//根据当前数据库中主机的数量查询相应的检查结果记录
func (this *ZhuGeliang) Get_Res_Now_From_Db() (res []Check_Res) {
	//新建一个 session
	shouyu := this.Get_A_DB_Session()
	defer shouyu.Close()

	c := shouyu.DB("orange").C("res_check")

	c.Find(nil).Sort("-time").Limit(this.Query_Counts_Hosts()).All(&res)
	return
}

//从数据库查询待检查的主机，通道版，供刘备做任务使用
func (this *ZhuGeliang) Query_Host_From_DB(ch chan Host) {
	//新建一个 session
	shouyu := this.Get_A_DB_Session()
	defer shouyu.Close()
	//选择一个 collection
	c := shouyu.DB("orange").C("hosts")

	hosts := []Host{}
	c.Find(nil).All(&hosts)
	//log.Println(hosts)
	for _, host := range hosts {
		ch <- host
	}
	close(ch)
}

//根据当前数据库中主机的数量查询相应的检查结果记录
func (this *ZhuGeliang) Get_Web_Status_Now_From_Db() (res []Web_Status) {
	//新建一个 session
	shouyu := this.Get_A_DB_Session()
	defer shouyu.Close()

	c := shouyu.DB("orange").C("web_status")

	c.Find(nil).Sort("-time").Limit(this.Query_Counts_Webs()).All(&res)
	return
}

//查询 web 站点数量
func (this *ZhuGeliang) Query_Counts_Webs() (num int) {
	//新建一个 session
	shouyu := this.Get_A_DB_Session()
	defer shouyu.Close()
	//选择一个 collection
	webs := []Web{}
	c := shouyu.DB("orange").C("webs")

	c.Find(nil).All(&webs)
	num = len(webs)
	return
}

//保存站web_status到服务器
func (this *ZhuGeliang) Save_Webstatus_To_Db(res Web_Status) {
	//新建一个 session
	shouyu := this.Get_A_DB_Session()
	defer shouyu.Close()
	//选择一个 collection
	c := shouyu.DB("orange").C("web_status")

	err := c.Insert(res)
	if err != nil {
		log.Println(err)
	}
	// log.Printf("Result install into db success!", res)
}

//从数据库查询所有 Web_status
func (this *ZhuGeliang) Query_Web_Status_From_Db() (web_status []Web_Status) {
	//新建一个 session
	shouyu := this.Get_A_DB_Session()
	defer shouyu.Close()
	//选择一个 collection
	c := shouyu.DB("orange").C("web_status")

	c.Find(nil).All(&web_status)
	return
}

//从数据库查询所有站点信息
func (this *ZhuGeliang) Query_Web_From_Db() (webs []Web) {
	//新建一个 session
	shouyu := this.Get_A_DB_Session()
	defer shouyu.Close()
	//选择一个 collection
	c := shouyu.DB("orange").C("webs")

	c.Find(nil).All(&webs)
	return
}

//保存站点信息到数据库
func (this *ZhuGeliang) Save_Web_To_Db(web Web) {
	//新建一个 session
	shouyu := this.Get_A_DB_Session()
	defer shouyu.Close()
	//选择一个 collection
	c := shouyu.DB("orange").C("webs")

	err := c.Insert(web)
	if err != nil {
		log.Println(err)
	}
	// log.Printf("Web install into db success!", web)
}

//根据 uuid 删除 web
func (this *ZhuGeliang) Delete_Web_By_UUID(uuid string) {
	//新建一个 session
	shouyu := this.Get_A_DB_Session()
	defer shouyu.Close()
	//选择一个 collection
	c := shouyu.DB("orange").C("webs")

	_, err := c.RemoveAll(bson.M{"uuid": uuid})
	if err != nil {
		log.Println(err)
	}
	log.Printf("%s is remove from database successful", uuid)
}

/**
 * Check user by useremail
 * @param  {[email]} this *ZhuGeliang)  Check_Exist_By_Email(email string) (exist bool)
 * @return {[bool]}      [If the user is exist return true, else false]
 */
func (this *ZhuGeliang) Check_Exist_By_Email(email string) (exist bool, user UserInfo) {
	//新建一个 session
	shouyu := this.Get_A_DB_Session()
	defer shouyu.Close()
	//选择一个 collection
	c := shouyu.DB("orange").C("userinfo")

	log.Println(email)
	err := c.Find(bson.M{"useremail": email}).One(&user)
	log.Println(user)
	if err != nil {
		log.Println(err)
		exist = false
		return
	}
	exist = true
	return

}

/**
 * Check user login for api
 * @param  {UserInfo} this *ZhuGeliang)  Check_Login(user UserInfo) (ok bool )
 * @return {bool}      If email & password check success,return true; else false.
 */
func (this *ZhuGeliang) Check_Login(user UserInfo) (ok bool) {
	exist, user_from_db := this.Check_Exist_By_Email(user.UserEmail)
	if exist {
		if user_from_db.UserPassword == user.UserPassword {
			ok = true
			return
		}
		ok = false
		return
	}
	ok = false
	return
}

/**
 * Query users from database by Zhugeliang
 * @param  {none} this *ZhuGeliang)  Query_User_From_Db(users []UserInfo self fun
 * @return {UserInfo}      A  list of UserInfo
 */
func (this *ZhuGeliang) Query_User_From_Db() (users []UserInfo) {
	//新建一个 session
	shouyu := this.Get_A_DB_Session()
	defer shouyu.Close()
	//选择一个 collection
	c := shouyu.DB("orange").C("userinfo")

	c.Find(nil).All(&users)
	return
}

/**
 * Add userinfo for api
 * @param  {userinfo} this *ZhuGeliang)  Add_User_For_Api(user UserInfo
 * @return {none}      Insert userinfo into database
 */
func (this *ZhuGeliang) Add_User_For_Api(user UserInfo) {
	//新建一个 session
	shouyu := this.Get_A_DB_Session()
	defer shouyu.Close()
	//选择一个 collection
	c := shouyu.DB("orange").C("userinfo")

	err := c.Insert(user)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Add user into db success!")
}

//返回一个 mgo 的 session
func (z *ZhuGeliang) Get_A_DB_Session() *mgo.Session {
	session, err := mgo.Dial(DB_URL)
	if err != nil {
		log.Println(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}

//返回一个 mgo 的 collection
func (z *ZhuGeliang) Get_A_DB_Collection(collection string) *mgo.Collection {
	session, err := mgo.Dial(DB_URL)
	if err != nil {
		log.Println(err)
	}
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(DB_NAME).C(collection)
	return c
}
