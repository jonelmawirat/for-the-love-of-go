package mytypes

import "strings"
type MyInt int

func (i MyInt) Twice() MyInt{
    return i * 2
}

type MyString string

func (s MyString) Len() int {
    return len(s)
}

type MyBuilder struct {
    Contents strings.Builder
}

func (b *MyBuilder) Hello() string {
    return "Hello, Gophers!"
}

type StringUpperCaser struct {
    Contents strings.Builder
}

func (s StringUpperCaser) ToUpper() string {
    return strings.ToUpper(s.Contents.String())
}


func (a *MyInt) Double() {
  *a *= 2  
}

