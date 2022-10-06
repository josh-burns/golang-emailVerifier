package main

import "testing"

func Test_EmailVerify(t *testing.T) {

	expected := "_________PASS_______"

	result := emailVerify("hello@hello.com")

	if result == expected {
		t.Logf("PASS")
	} else {
		t.Errorf("FAIL", expected, result)
	}
}
