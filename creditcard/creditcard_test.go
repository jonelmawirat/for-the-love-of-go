package creditcard_test

import (
	"testing"
    "creditcard"
)


func TestNew(t *testing.T) {
    card,err := creditcard.New("1235")
    if err != nil {
        t.Error(err)
    }

    got := card.Number()
    want := "1235"

    if got != want {
        t.Errorf("Want: %s Got: %s",want,got)
    }
}
