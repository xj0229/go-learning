package my_map_ext

import (
	"testing"
	"unsafe"
)

func TestMapWithFunctionValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int {
		return op
	}
	m[2] = func(op int) int {
		return op * op
	}
	m[3] = func(op int) int {
		return op * op * op
	}
	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 1
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
	mySet[3] = true
	t.Log(len(mySet))

	delete(mySet, 1)
	n = 1
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
}

func TestString(t *testing.T) {
	var s string
	t.Log(s)
	s = "hello"
	t.Log(len(s))
	//s[1] = '3'
	s = "\xE4\xB8\xA5"
	s = "\xE4\xBA\xB5\xFF"
	t.Log(s)
	s = "折"
	t.Log(len(s))
	c := []rune(s)
	t.Log("rune size: ", unsafe.Sizeof(c[0]))
	t.Logf("折 unicode %x", c[0])
	t.Logf("折 UTF8 %x", s)
}
