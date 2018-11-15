package utils

import (
	"github.com/liudng/godump"
	"testing"
)

type UserTest struct {
	IDID  int    `json:"id_id"`
	NameA string `json:"name_a"`
	Age   *Age   `json:"age_1"`
}

type Age struct {
	age int `json:"json"`
}

func TestStruct2Map(t *testing.T) {
	ut := new(UserTest)
	ut.IDID = 1
	ut.NameA = ""

	// fmt.Println(reflect.ValueOf(ut).Elem().Interface())
	// godump.Dump(ut)
	// godump.Dump(*ut)

	data := Obj2Map(ut, false)
	godump.Dump(data)
}
