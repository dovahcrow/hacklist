package models

import (
	. "github.com/doomsplayer/tools"
	"labix.org/v2/mgo/bson"
	"net/url"
	"time"
)

type Attack struct {
	Id             bson.ObjectId `bson:"_id"` //此次攻击在数据库中的ID
	WebSite        *url.URL      //被攻击的网址
	DefacedPage    []*url.URL    //被篡改的页面
	Level          int           //威胁等级
	AttackTimes    []time.Time   //攻击时间 数组
	FirstFoundTime time.Time     //首次发现时间
	HackerIds      []bson.ObjectId
	Hackers        []*Hacker   `bson:"-"`
	Positions      []*Position //攻击源地址
}

//monotonic type
type atk struct{}

var Atk atk

func (atk) InsertAttack(attack *Attack) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	if attack.Id == `` {
		attack.Id = bson.NewObjectId()
	}
	err = AttackCollection.Insert(attack)
	E("Insert Attack Error", true, err)
	return
}

func (atk) GetAllAttacks() (atks []*Attack, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	err = AttackCollection.Find(nil).All(&atks)
	E("Get All Attacks Error", true, err)
	for _, v := range atks {

		HackerCollection.Find(bson.M{"attackids": v.Id}).All(&v.Hackers)

	}
	return
}

func (atk) GetAttackNum() (i int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	i, err = AttackCollection.Count()
	E("Get Attack Count Error", true, err)
	return
}

func (atk) GetAttackRange(offset, num int) (atks []*Attack, err error) {
	err = AttackCollection.Find(nil).Skip(offset).Limit(num).All(&atks)
	for _, v := range atks {

		HackerCollection.Find(bson.M{"attackids": v.Id}).All(&v.Hackers)

	}
	return
}
func (atk) GetAttackById(id string) (atk *Attack, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	err = AttackCollection.FindId(bson.ObjectIdHex(id)).One(atk)

	HackerCollection.Find(bson.M{"attackids": atk.Id}).All(&atk.Hackers)

	return
}
