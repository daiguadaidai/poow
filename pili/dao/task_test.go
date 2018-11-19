package dao

import (
	"fmt"
	"github.com/daiguadaidai/poow/utils"
	"testing"
)

func TestTaskDao_QueryByProgramID(t *testing.T) {
	InitDBConfig()

	pg := &utils.Paginator{
		Offset: 0,
		Limit:  1,
	}

	pk := 1
	tasks, err := NewTaskDao().QueryByProgramID(int64(pk), pg)
	if err != nil {
		t.Fatalf(err.Error())
	}

	for _, t := range tasks {
		fmt.Println(t)
	}
}
