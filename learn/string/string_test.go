package string_test

import (
	"testing"
)


func stringSlice(s string, t *testing.T) {
	for _,v := range s {
		t.Logf("%[1]c %[1]x", v)
	}
}

func TestStringInit(t *testing.T) {
	var s string = "问君能有几多愁"
	num := len([]rune(s))
	t.Log(num)

	stringSlice(s, t)
}
