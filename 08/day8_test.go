package day8

import "testing"

func TestDay8(t *testing.T) {
	if p1, p2 := Day8("calib2.txt"); p1 != 6 && p2 != -1 {
		t.Fail()
	}
}
