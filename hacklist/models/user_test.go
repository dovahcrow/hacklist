package models

import (
	"hacklist/models/permission"
	"testing"
)

func Test_CreateUser(t *testing.T) {
	err := Usr.CreateUser(`user1`, `isopcs`, `inclaoc`, `kk,scl`, permission.AdminGod)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_GetUser(t *testing.T) {
	// err := Usr.CreateUser(`user2`, `isopcs`, `inclaoca`, `kk,scl`, permission.AdminGod)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	u, err := Usr.GetUser(`user2`, `isopcs`)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", u)
	t.Logf("%+v\n", u.Id.Hex())
}
