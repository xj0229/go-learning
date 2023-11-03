package condition

import (
	"runtime"
	"testing"
)

func TestIfMultiSec(t *testing.T) {
	if a := 1 == 1; a {
		t.Log("1==1")
	}

	//if v, err = someFunc(); err== nil {
	//	t.Log("1==1")
	//}else{
	//
	//}

}

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}

func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}

func TestSwitchRuntime(t *testing.T) {
	switch os := runtime.GOOS; os {
	case "darwin":
		t.Log("OS X.")
	case "linux":
		t.Log("Linux.")
	default:
		t.Log(os)
	}
}

func TestSwitchScope(t *testing.T) {
	for i := 0; i < 10; i++ {
		switch {
		case 0 <= i && i <= 4:
			t.Log("we should discard this number!")
		case 5 <= i && i <= 9:
			t.Log("we should level up!")
		default:
			t.Log("we do not know what we should do!")
		}
	}
}
