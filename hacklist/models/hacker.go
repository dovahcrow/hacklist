package models

import (
	"fmt"
	. "github.com/doomsplayer/tools"
	"labix.org/v2/mgo/bson"
	"strings"
	"time"
)

type Hacker struct {
	Id               bson.ObjectId     `bson:"_id" form:"-"`
	AttackIds        []bson.ObjectId   `form:"-"`
	Attacks          []*Attack         `bson:"-" form:"-"`
	Nick             string            `form:"nick"`
	Gender           string            `form:"gender"`
	RealName         string            `form:"realname"`
	PersonalIdentify string            `form:"ID"`
	Birthday         time.Time         `form:"birthday"`
	Contact          map[string]string `form:"-"`
	Positions        []*Position       `form:"-"`
}

type hkr struct{}

var Hkr hkr

func (hkr) NewHacker() *Hacker {
	h := new(Hacker)
	h.Contact = make(map[string]string)
	return h
}
func (hkr) InsertHacker(hacker *Hacker) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	if hacker.Id == `` {
		hacker.Id = bson.NewObjectId()
	}

	hacker.Gender = strings.TrimSpace(hacker.Gender)
	if hacker.Gender != "male" && hacker.Gender != "female" && hacker.Gender != "unknown" {
		return fmt.Errorf("gender is illegal")
	}

	hacker.Nick = strings.TrimSpace(hacker.Nick)
	if hacker.Nick == `` {
		return fmt.Errorf("nick is null")
	}

	hacker.RealName = strings.TrimSpace(hacker.RealName)
	if hacker.RealName == `` {
		return fmt.Errorf("realname is null")
	}

	err = HackerCollection.Insert(hacker)
	E("Insert Hacker Error", true, err)
	return
}

func (hkr) GetAllHackers() (hkr []*Hacker, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	err = HackerCollection.Find(nil).All(&hkr)
	E("Get All Hackers Error", true, err)
	for _, v := range hkr {

		AttackCollection.Find(bson.M{"hackerids": v.Id}).All(&v.Attacks)

	}
	return
}
func (hkr) GetHackerRange(offset, num int) (hkrs []*Hacker, err error) {
	err = HackerCollection.Find(nil).Skip(offset).Limit(num).All(&hkrs)
	for _, v := range hkrs {

		AttackCollection.Find(bson.M{"hackerids": v.Id}).All(&v.Attacks)

	}
	return
}
func (hkr) GetHackerNum() (i int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	i, err = HackerCollection.Count()
	E("Get Hacker Count Error", true, err)
	return
}

func (hkr) GetHackerById(id string) (hkr *Hacker, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	err = HackerCollection.FindId(bson.ObjectIdHex(id)).One(hkr)

	AttackCollection.Find(bson.M{"hackerids": hkr.Id}).All(&hkr.Attacks)

	return
}
func (hkr) GetHackers(query bson.M) (hkrs []*Hacker, err error) {
	err = HackerCollection.Find(query).All(&hkrs)
	for _, v := range hkrs {

		AttackCollection.Find(bson.M{"hackerids": v.Id}).All(&v.Attacks)

	}
	return
}
func (hkr) GetHacker(query bson.M) (hkr *Hacker, err error) {
	hkr = new(Hacker)
	err = HackerCollection.Find(query).One(hkr)

	AttackCollection.Find(bson.M{"hackerids": hkr.Id}).All(&hkr.Attacks)

	return
}
func (hkr) DeleteHackerById(id string) (err error) {
	err = HackerCollection.RemoveId(bson.ObjectIdHex(id))
	return
}
