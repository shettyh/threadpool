package internal

import "testing"

var (
	set *Set
)

func TestNewSet(t *testing.T) {
	set = NewSet()
}

func TestSet_Add(t *testing.T) {
	set.Add(20)
	if ok := set.Contains(20); !ok {
		t.Fail()
	}
}

func TestSet_Remove(t *testing.T) {
	set.Remove(20)
	if ok := set.Contains(20); ok {
		t.Fail()
	}
}

func TestSet_Contains(t *testing.T) {
	set.Add(40)
	if ok := set.Contains(40); !ok {
		t.Fail()
	}
}

func TestSet_GetAll(t *testing.T) {
	data := set.GetAll()
	if len(data) != 1 {
		t.Fail()
	}
}

func TestSet_GetAllWithCap(t *testing.T) {
	set.Add(50)
	data:= set.GetAllWithCap(1)
	if len(data) != 1 {
		t.Fail()
	}
}
