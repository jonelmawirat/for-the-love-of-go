package creditcard

import "errors"

type Card struct {
    number string
}

func New(number string) (*Card,error) {
    if number == "" {
        return &Card{},errors.New("Invalid Card Number")
    }

    c := Card{number: number}
    return &c,nil
}

func (c *Card) Number() string {
    return c.number
}
