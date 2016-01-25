package roman

import "testing"

func TestCreateNumeralFromInteger(t *testing.T) {
	tests := []struct {
		in  uint
		out string
	}{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{7, "VII"},
		{8, "VIII"},
		{9, "IX"},
		{10, "X"},
		{11, "XI"},
		{50, "L"},
		{54, "LIV"},
		{100, "C"},
		{257, "CCLVII"},
		{2157, "MMCLVII"},
	}

	for _, test := range tests {
		n, err := NewNumeral(test.in)
		if err != nil {
			t.Fatal(err)
		}
		if n.GetValue() != test.out {
			t.Fatalf("Unexpected string value %v, want %s", n.GetValue(), test.out)
		}
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
	if !a.SameValueAs(b) {
		t.Fatal("Not same value as")
	}
}

func TestShouldBeSameIfIntegerEquivalent(t *testing.T) {
	a, _ := NewNumeral(5)
	b, _ := NewNumeral("V")
	if !a.SameValueAs(b) {
		t.Fatal("Not same value as")
	}
}

func TestShouldCompareTwoNumeralsAsNotSame(t *testing.T) {
	a, _ := NewNumeral("I")
	b, _ := NewNumeral("X")
	if a.SameValueAs(b) {
		t.Fatal("Shouldn't be same value")
	}
}

func TestShouldCompareAsNotSameIfNotEquivalent(t *testing.T) {
	a, _ := NewNumeral(9)
	b, _ := NewNumeral("X")
	if a.SameValueAs(b) {
		t.Fatal("Shouldn't be same value")
	}
}

func TestCreateFromInt64(t *testing.T) {
	v, err := NewNumeral(int64(67))
	if err != nil {
		t.Fatal(err)
	}
	if v.GetValue() != "LXVII" {
		t.Fatalf("Unexpected value %v, want %s", v.GetValue(), "LXVII")
	}
}

func TestNegativeInt(t *testing.T) {
	_, err := NewNumeral(-1)
	if err == nil {
		t.Fatal("expected an error with negative number")
	}
}

func TestDifferentValueObjectTypeNotEqual(t *testing.T) {
	n, err := NewNumeral("I")
	if err != nil {
		t.Fatal(err)
	}
	if n.SameValueAs(testValueObject{}) {
		t.Errorf("unexpected equality with different value object type")
	}
}

type testValueObject struct{}

func (v testValueObject) SameValueAs(value ValueObject) bool {
	return false
}

func (v testValueObject) GetValue() interface{} {
	return 0
}
