package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr1[1] = 5
	arr3 := [...]int{1, 3, 5, 7}
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr3)
	t.Log(arr1[1:len(arr1)])
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 5, 7, 9}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
	for idx, e := range arr3 {
		t.Log(idx, e)
	}
	for _, e := range arr3 {
		t.Log(e)
	}
}

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5, 6}
	arr3_sec := arr3[3:]
	t.Log(arr3_sec)
}
