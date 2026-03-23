package analymodel

import (
	"fmt"
	"testing"
)

func TestGetInstance(t *testing.T) {
	s1 := GetInstance()
	s2 := GetInstance()

	if s1 != s2 {
		t.Error("Singleton instances are not the same")
	}
	fmt.Printf("s1 %d, s1: %p\n", s1.data, &s1)
	fmt.Printf("s2 %d, s2: %p\n", s2.data, &s2)
}
