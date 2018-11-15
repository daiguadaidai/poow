package dao

import (
	"fmt"
	"testing"
)

// 有专用机器
func TestHostDao_GetByProgramIDAndDedicate_HaveDedicate(t *testing.T) {
	InitDBConfig()

	cols := []string{"hosts.id", "hosts.host"}
	proID := int64(1)
	h, err := NewHostDao().GetByProgramIDAndDedicate(proID, true, cols)
	if err != nil {
		t.Fatalf("%v", err)
	}

	fmt.Println(h)
}

// 有共用机器
func TestHostDao_GetByProgramIDAndDedicate_HaveNotDedicate(t *testing.T) {
	InitDBConfig()

	cols := []string{"hosts.id", "hosts.host"}
	proID := int64(1)
	h, err := NewHostDao().GetByProgramIDAndDedicate(proID, false, cols)
	if err != nil {
		t.Fatalf("%v", err)
	}

	fmt.Println(h)
}
