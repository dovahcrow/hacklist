package models

import (
	"testing"
	"time"
)

func Test_Insert(t *testing.T) {
	h := Hkr.NewHacker()
	h.Birthday = time.Now()
	h.Contact[`qq`] = `019241029`
	h.Contact[`oosadc`] = `asvasfa`
	h.Nick = "kiler"
	h.RealName = "asdcq"

	err := Hkr.InsertHacker(h)
	if err != nil {
		t.Fatal(err)
	}
}
