package mytypes_test

import (
	"mytypes"
	"strings"
	"testing"
)

func TestTwice(t *testing.T) {
    t.Parallel()
    want := mytypes.MyInt(4)
    
    var a mytypes.MyInt = 2
    got := a.Twice()
    if want != got {
        t.Errorf("Want: %d Got: %d\n",want,got)
    }

}

func TestMyStringLen(t *testing.T) {
    a := mytypes.MyString("Jonel")
    got := a.Len()
    want := 5
    if got != want {
        t.Errorf("Got: %d Want: %d\n",got,want)
    }
}

func TestStringBuilder(t *testing.T) {
    t.Parallel()
    var sb strings.Builder
    sb.WriteString("Hello, ")
    sb.WriteString("Gophers!")
    want := "Hello, Gophers!"
    got := sb.String()
    if want != got {
        t.Errorf("Want: %s Got: %s\n",want,got)
    }

    wantLen := 15
    gotLen := sb.Len()
    if wantLen != gotLen {
        t.Errorf("Want: %d Got: %d\n",wantLen,gotLen)
    }
}

func TestMyBuilderHello(t *testing.T) {
    t.Parallel()
    var mb mytypes.MyBuilder
    want := "Hello, Gophers!"
    got := mb.Hello()
    if want != got {
        t.Errorf("Want: %s Got: %s\n",want,got)
    }
}

func TestMyBuilder(t *testing.T){
    t.Parallel()
    var mb mytypes.MyBuilder
    mb.Contents.WriteString("Hello, ")
    mb.Contents.WriteString("Gophers!")
    want := "Hello, Gophers!"
    got := mb.Contents.String()

    if want != got {
        t.Errorf("Want: %s Got: %s\n",want,got)
    } 

    wantLen := 15
    gotLen := mb.Contents.Len()
    if wantLen != gotLen {
        t.Errorf("Want: %d Got: %d\n",wantLen,gotLen)
    }
}

func TestStringUppercaser(t *testing.T){
    t.Parallel()
    var mb mytypes.StringUpperCaser
    mb.Contents.WriteString("Hello, Gophers!")
    want := "HELLO, GOPHERS!"
    got := mb.ToUpper()
    if want != got {
        t.Errorf("Want: %s Got: %s",want,got)
    }
}

func TestDouble(t *testing.T){
    t.Parallel()
    var x mytypes.MyInt = 12
    want := mytypes.MyInt(24)
    x.Double()
    if want != x {
        t.Errorf("Want %d Got: %d\n",want,x)
    }
}
