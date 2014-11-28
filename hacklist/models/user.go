package models

import (
	"hacklist/controllers/helper"
	"hacklist/models/permission"
	"labix.org/v2/mgo/bson"
	"time"
)

type usr struct{}

var Usr usr

func init() {

}

type User struct {
	Id            bson.ObjectId         `form:"-" bson:"_id"`
	Account       string                `form:"account"`
	Password      string                `form:"password"`
	Name          string                `form:"name"`
	Email         string                `form:"email"`
	Permission    permission.Permission `form:"permission"`
	Does          []string              `form:"-"`
	CreateTime    time.Time             `form:"-"`
	LastLoginTime time.Time             `form:"-"`
}

func (usr) GetUser(account, password string) (user *User, err error) {
	user = new(User)
	err = UserCollection.Find(bson.M{"account": account, "password": password}).One(user)
	return
}
func (usr) GetUserById(id string) (user *User, err error) {
	user = new(User)
	err = UserCollection.FindId(bson.ObjectIdHex(id)).One(user)
	return
}
func (usr) IncrUserLogin(id string) (err error) {
	user := new(User)
	err = UserCollection.FindId(bson.ObjectIdHex(id)).One(user)
	if err == nil {
		user.LastLoginTime = time.Now()
		err = UserCollection.UpdateId(user.Id, user)
	}
	return
}

func (usr) CreateUser(
	account, password, name, email string, permission permission.Permission,
) (err error) {

	u := new(User)
	u.Id = bson.NewObjectId()
	u.Account = account
	u.Password = password
	u.Name = name
	u.Email = email
	u.Permission = permission
	u.CreateTime = time.Now()
	u.LastLoginTime = u.CreateTime
	err = UserCollection.Insert(u)
	return
}

func (usr) GetUserNum() (num int, err error) {
	num, err = UserCollection.Count()
	return
}
func (usr) GetUserRange(offset, num int) (user []*User, err error) {
	err = UserCollection.Find(nil).Skip(offset).Limit(num).All(&user)
	return
}

func (usr) UpdateUser(u *User) (err error) {
	m := helper.Struct2Map(u)
	delete(m, "createtime")
	delete(m, "lastlogintime")
	UserCollection.UpsertId(u.Id)
}
