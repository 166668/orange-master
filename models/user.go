package models

import (
	u "code.google.com/p/go-uuid/uuid"
)

/**
 * Userinfo for login
 * UserName is not requird
 */
type UserInfo struct {
	UserId       string `bson:"userid"`
	UserName     string `bson:"username"`
	UserPassword string `bson:"userpassword"`
	UserEmail    string `bson:"useremail"`
}

/**
 * Gen userid by uuid
 * @param  method self
 * @return UserInfo.UserId
 */
func (this *UserInfo) Gen_Id() {
	this.UserId = u.New()
}
