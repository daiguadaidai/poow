package utils

import (
	"fmt"
	"testing"
)

func TestGetUUID(t *testing.T) {
	uuid := GetUUID()
	fmt.Println(uuid, len(uuid))
}
