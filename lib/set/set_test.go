package set

import (
	"fmt"
	"testing"
)

func Test_Set(t *testing.T) {
	s := New[string]()
	fmt.Println(s)
	s.Add("one")
	s.Add("two")
	fmt.Println(s.Size())
}
