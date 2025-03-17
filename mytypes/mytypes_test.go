package mytypes_test

import (
	"mytypes"
	"strings"
	"testing"
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
		t.Errorf("%q: want %v got %v", ms, want, got)
	}
}

func TestStringBuilder(t *testing.T) {
	t.Parallel()
	var sb strings.Builder
	sb.WriteString("hello, ")
	sb.WriteString("gophers!")
	want := "hello, gophers!"
	got := sb.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
	wantLen := 15
	gotLen := sb.Len()
	if wantLen != gotLen {
		t.Errorf("want len %d, got %d", wantLen, gotLen)
	}

}

func TestMyBuilderHello(t *testing.T) {
	t.Parallel()
	var mb mytypes.MyBuilder
	want := "hello"
	got := mb.Hello()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestMyBuilder(t *testing.T) {
	t.Parallel()
	var mb mytypes.MyBuilder
	mb.Contents.WriteString("hello, ")
	mb.Contents.WriteString("gophers!")
	want := "hello, gophers!"
	got := mb.Contents.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
	len_want := 15
	len_got := mb.Contents.Len()
	if len_want != len_got {
		t.Errorf("want %d got %d", len_want, len_got)
	}

}

func TestStringUppercaser(t *testing.T) {
	t.Parallel()
	var mb mytypes.StringUppercaser
	mb.Contents.WriteString("test")
	want := "TEST"
	got := mb.ToUpper()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestDouble(t *testing.T) {
	t.Parallel()
	x := mytypes.MyInt(12)
	(x).Double() // could be &x
	want := 24
	if want != int(x) {
		t.Errorf("want %q, got %q", want, x)
	}
}
