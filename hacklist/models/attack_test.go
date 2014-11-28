package models

import (
	"net/url"
	"testing"
	"time"
)

func Test_InsertAttack(t *testing.T) {
	t.SkipNow()
	atk := new(Attack)
	atk.AttackTimes = []time.Time{time.Now(), time.Now(), time.Now()}
	atk.FirstFoundTime = time.Now()
	atk.Level = 3
	u, _ := url.Parse("http://www.baidu.com/123")
	atk.WebSite = u
	err := Atk.InsertAttack(atk)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_GetAllAttacks(t *testing.T) {
	atks, err := Atk.GetAllAttacks()
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range atks {
		t.Logf("%+v\n", v)
	}

}
