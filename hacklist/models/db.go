package models

import (
	"github.com/astaxie/beego"
	"hacklist/models/permission"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"os"
	"time"
)

var (
	db               *mgo.Database
	HackerCollection *mgo.Collection
	AttackCollection *mgo.Collection
	UserCollection   *mgo.Collection
	MetaCollection   *mgo.Collection
)
var (
	ConnString = "localhost"
)

func init() {
	if s := beego.AppConfig.String("connString"); s != `` {
		ConnString = s
	}

	mgosession, err := mgo.Dial(ConnString)
	if err != nil {
		log.Fatalf("Open database error: %v", err)
		os.Exit(1)
	}
	db = mgosession.DB("hacklist")
	HackerCollection = db.C("hackers")
	MetaCollection = db.C("meta")
	AttackCollection = db.C("attack")
	UserCollection = db.C(`user`)

	if beego.AppConfig.String("runmode") != "dev" {

		f := log.New(os.Stderr, "", log.LstdFlags)
		mgo.SetLogger(f)
		mgo.SetDebug(true)
	}

	idx := mgo.Index{
		Background: true,
		Key:        []string{`account`},
		Sparse:     true,
		Unique:     true,
	}
	UserCollection.EnsureIndex(idx)
	idx = mgo.Index{
		Background: true,
		Key:        []string{`name`},
		Sparse:     true,
		Unique:     true,
	}
	UserCollection.EnsureIndex(idx)

	i, _ := UserCollection.Find(bson.M{`account`: `admin`}).Count()
	if i == 0 {
		u := new(User)
		u.Account = `admin`
		u.CreateTime = time.Now()
		u.Email = "a@b.c"
		u.LastLoginTime = time.Now()
		u.Name = "admin"
		u.Password = "admin"
		u.Permission = permission.AdminGod
		u.Id = bson.NewObjectId()
		UserCollection.Insert(u)
	}

}
