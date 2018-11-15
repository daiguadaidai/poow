package dao

import (
	"fmt"
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
