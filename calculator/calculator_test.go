package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
    t.Parallel()
    type TestCase struct {
        a, b, want float64
    }

    testCases := []TestCase{
        {a: 2, b: 2, want: 4},
        {a: 3, b: 3, want: 6},
    }

    for _,v := range testCases {
        got := calculator.Add(v.a,v.b)
        if v.want != got {
            t.Errorf("Want: %f Got: %f\n",v.want,got)
        }
    }
}


func TestMultiply(t *testing.T) {
    type testCase struct {
        a,b,want float64
    }
    
    testCases := []testCase{
        {a: 2,b: 2,want: 4},
        {a: 3,b:2,want:6},
    }

    for _,v := range testCases {

        got := calculator.Multiply(v.a,v.b)
        if v.want != got {
            t.Errorf("Want: %f Got: %f\n",v.want,got)
        }
    }

}

func closeEnough(a,b,tolerance float64) bool {
    if math.Abs(a-b) < tolerance {
        return true
    } else {
        return false
    }
}

func TestDivide(t *testing.T) {
    type testCase struct {
        a,b, want float64
    }

    testCases := []testCase{
        {a: 2,b: 2, want: 1},
        {a: 0,b: 1, want: 0},
        {a:1,b: 3,want: 0.33333333},
    }

    for _,tc := range testCases{
        got,err := calculator.Divide(tc.a,tc.b)

        if err != nil {
            t.Error(err)
        }

        if closeEnough(got,tc.want,0.0001) == false {
            t.Errorf("Got: %f Want: %f\n",got,tc.want)
        }
    }
}


