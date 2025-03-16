package mytypes

import (
	"strings"
)

type MyInt int

func (i *MyInt) Double() {
	*i *= 2
}

type StringUppercaser struct {
	Contents strings.Builder
}

func (ms StringUppercaser) ToUpper() string {
	return strings.ToUpper(ms.Contents.String())
}

type MyBuilder struct {
	Contents strings.Builder
}

func (m MyBuilder) Hello() string {
	return "hello"
}

type MyType int

func (m MyType) Twice() MyType {
	return m * 2
}

type MyString string

func (ms MyString) Lenght() int {
	return len(string(ms))
}
