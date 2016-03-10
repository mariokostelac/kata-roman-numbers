package romans

import "testing"

func TestSimple1(t *testing.T) {
	if convert(1) != "I" {
		t.Errorf("I is wrong")
	}
	if convert(5) != "V" {
		t.Errorf("V is wrong")
	}
	if convert(10) != "X" {
		t.Errorf("X is wrong")
	}
	if convert(50) != "L" {
		t.Errorf("L is wrong")
	}
	if convert(100) != "C" {
		t.Errorf("C is wrong")
	}
	if convert(500) != "D" {
		t.Errorf("D is wrong")
	}
	if convert(1000) != "M" {
		t.Errorf("M is wrong")
	}
}

func TestComplicated(t *testing.T) {
	if convert(2) != "II" {
		t.Errorf("2 is wrong")
	}
	if convert(4) != "IV" {
		t.Errorf("4 is wrong")
	}
	if convert(6) != "VI" {
		t.Errorf("6 is wrong")
	}
	if convert(9) != "IX" {
		t.Errorf("9 is wrong")
	}
	if convert(11) != "XI" {
		t.Errorf("11 is wrong")
	}
	if convert(99) != "XCIX" {
		t.Errorf("99 is wrong")
	}
	if convert(101) != "CI" {
		t.Errorf("101 is wrong -> %s", convert(101))
	}
	if convert(1986) != "MCMLXXXVI" {
		t.Errorf("1986 is wrong -> %s", convert(1986))
	}
	if convert(2016) != "MMXVI" {
		t.Errorf("2016 is wrong -> %s", convert(2016))
	}
	if convert(3999) != "MMMCMXCIX" {
		t.Errorf("3999 is wrong -> %s", convert(3999))
	}
}
