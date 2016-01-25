package roman

import "testing"

func TestSomething(t *testing.T) {
	var n Numeral

	if n.value != "" {
		t.Errorf("something went wrong")
	}
}

func TestCreateNumeralFromIntegerWhenOne(t *testing.T) {
	n, err := NewNumeral(1)
	if err != nil {
		t.Fatal(err)
	}
	if n.value != "I" {
		t.Fatalf("string representation should be I, was %s", n.value)
	}
}

func TestCreateNumeralFromIntegerWhenFive(t *testing.T) {
	n, err := NewNumeral(5)
	if err != nil {
		t.Fatal(err)
	}
	if n.value != "V" {
		t.Fatalf("string representation should be V, was %s", n.value)
	}
}

func TestCreateNumeralFromIntegerWhenTen(t *testing.T) {
	n, err := NewNumeral(10)
	if err != nil {
		t.Fatal(err)
	}
	if n.value != "X" {
		t.Fatalf("string representation should be X, was %s", n.value)
	}
}

func TestCreateNumeralFromIntegerWhenFifty(t *testing.T) {
	n, err := NewNumeral(50)
	if err != nil {
		t.Fatal(err)
	}
	if n.value != "L" {
		t.Fatalf("string representation should be L, was %s", n.value)
	}
}

func TestCreateNumeralFromIntegerWhenHundred(t *testing.T) {
	n, err := NewNumeral(100)
	if err != nil {
		t.Fatal(err)
	}
	if n.value != "C" {
		t.Fatalf("string representation should be C, was %s", n.value)
	}
}

func TestCreateNumeralFromIntegerWhen257(t *testing.T) {
	n, err := NewNumeral(257)
	if err != nil {
		t.Fatal(err)
	}
	if n.value != "CCLVII" {
		t.Fatalf("string representation should be CCLVII, was %s", n.value)
	}
}

func TestCreateNumeralFromIntegerWhen4(t *testing.T) {
	n, err := NewNumeral(4)
	if err != nil {
		t.Fatal(err)
	}
	if n.value != "IIII" {
		t.Fatalf("string representation should be CCLVII, was %s", n.value)
	}
}

func TestCreateNumeralFromString(t *testing.T) {
	n, err := NewNumeral("I")
	if err != nil {
		t.Fatal(err)
	}
	if n.value != "I" {
		t.Fatalf("string representation should be I, was %s", n.value)
	}
}

func TestCreateNumeralWithInvalidString(t *testing.T) {
	_, err := NewNumeral("A")
	if err == nil {
		t.Fatal("We expected an error with A")
	}
}

func TestShouldBeSameIfSameNumeral(t *testing.T) {
	a, _ := NewNumeral("I")
	b, _ := NewNumeral("I")
	if a.SameValueAs(b) == false {
		t.Fatal("Not same value as")
	}
}

func TestShouldBeSameIfIntegerEquivalent(t *testing.T) {
	a, _ := NewNumeral(5)
	b, _ := NewNumeral("V")
	if a.SameValueAs(b) == false {
		t.Fatal("Not same value as")
	}
}

func TestShouldCompareTwoNumeralsAsNotSame(t *testing.T) {
	a, _ := NewNumeral("I")
	b, _ := NewNumeral("X")
	if a.SameValueAs(b) == true {
		t.Fatal("Shouldn't be same value")
	}
}

func TestShouldParseTwo(t *testing.T) {
	v, _ := NewNumeral(2)
	if v.value != "II" {
		t.Fatalf("Unexpected string value %s, want %s", v.value, "II")
	}
}
