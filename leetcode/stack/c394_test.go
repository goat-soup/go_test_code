package stack

import (
	"fmt"
	"testing"
)

func TestDecodeString(t *testing.T) {
	fmt.Println()
	if "abcabccdcdcdef" != DecodeString("2[abc]3[cd]ef") {
		t.Error("not equea")
	}
}
