package roman

import (
	"errors"
	"regexp"
	"strings"
)

// ValueObject defines the interface of a simple value object
type ValueObject interface {
	SameValueAs(value ValueObject) bool
	GetValue() interface{}
}

// Numeral value object representation of a roman numeral
type Numeral struct {
	value string
}

var numbers = []uint{
	1000,
	900,
	500,
	400,
	100,
	90,
	50,
	40,
	10,
	9,
	5,
	4,
	1,
}

var letters = []string{
	"M",
	"CM",
	"D",
	"CD",
	"C",
	"XC",
	"L",
	"XL",
	"X",
	"IX",
	"V",
	"IV",
	"I",
}

func pattern() *regexp.Regexp {
	pattern := strings.Join(letters, "|")
	return regexp.MustCompile("[" + pattern + "]+")
}

// NewNumeral creates a new Numeral from a string or int
func NewNumeral(v interface{}) (Numeral, error) {
	var n Numeral
	switch t := v.(type) {
	case string:
		if !pattern().Match([]byte(t)) {
			return n, errors.New("Not a valid numeral")
		}
		n.value = t
	case uint:
		n.value = itoa(t)
	case int:
		return fromInt(int64(t))
	case int64:
		return fromInt(t)
	}

	return n, nil
}

func fromInt(in int64) (Numeral, error) {
	var n Numeral
	if in < 0 {
		return n, errors.New("can not represent negative numbers")
	}
	n.value = itoa(uint(in))
	return n, nil
}

// SameValueAs returns true if the supplied value object is equal - ValueObject interface
func (n Numeral) SameValueAs(value ValueObject) bool {
	if numeral, ok := value.(Numeral); ok {
		return n.value == numeral.GetValue()
	}
	return false
}

// GetValue returns the value of the numeral - ValueObject interface
func (n Numeral) GetValue() interface{} {
	return n.value
}

func itoa(in uint) string {
	acc := ""
	remainder := in

	for i := range numbers {
		for remainder >= numbers[i] {
			acc += letters[i]
			remainder -= numbers[i]
		}
	}

	return acc
}
