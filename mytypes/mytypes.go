package mytypes

type MyType int

func (m MyType) Twice() MyType {
	return m*2
}


type MyString string

func (ms MyString) Lenght() int {
	return len(string(ms))
}
