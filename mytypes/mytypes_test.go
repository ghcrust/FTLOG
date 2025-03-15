package mytypes_test

import (
	"testing"
	"mytypes"
)

func TestMyType(t *testing.T) {
	t.Parallel()
	m := mytypes.MyType(8)
	want := mytypes.MyType(16)
	got := m.Twice()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}

}

func TestMyStringLen(t *testing.T) {
	t.Parallel()
	ms := mytypes.MyString("Lol")
	want := 3
	got := ms.Lenght()
	if want != got {
		t.Errorf("%q: want %v got %v",ms,want,got)
	}
}
