package dao

import (
	"fmt"
	"github.com/daiguadaidai/poow/utils"
	"testing"
)

func TestProgramDao_GetByName(t *testing.T) {
	InitDBConfig()

	name := "test_01.py"
	p, err := NewProgramDao().GetByName(name, []string{"id", "have_dedicate"})
	if err != nil {
		t.Fatalf("获取命令出错: %s. %v", name, err)
	}
	fmt.Println(p)
}

func TestProgramDao_Query(t *testing.T) {
	InitDBConfig()

	pg := &utils.Paginator{
		Offset: 0,
		Limit:  100,
	}
	list, err := NewProgramDao().Query(pg)
	if err != nil {
		t.Fatalf(err.Error())
	}
	for _, p := range list {
		fmt.Println(p)
	}
}

func TestProgramDao_CountByFileName(t *testing.T) {
	InitDBConfig()
	cnt, err := NewProgramDao().CountByFileName("test_01.py")
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println("count:", cnt)
}
