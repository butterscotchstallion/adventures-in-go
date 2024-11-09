package leetcode

import "testing"

func Test_isValid(t *testing.T) {
	validResult := isValid("()")
	if !validResult {
		t.Errorf("() expected to be valid")
	}

	validResult2 := isValid("()[]{}")
	if !validResult2 {
		t.Errorf("()[]{} expected to be valid")
	}

	validResult3 := isValid("([])")
	if !validResult3 {
		t.Errorf("([]) expected to be valid")
	}

	validResult4 := isValid("[")
	if validResult4 {
		t.Errorf("[ expected to be invalid")
	}

	validResult5 := isValid("{[]}")
	if !validResult5 {
		t.Errorf("{[]} expected to be valid")
	}

	validResult6 := isValid("[{()}]")
	if !validResult6 {
		t.Errorf("[{()}] expected to be valid")
	}

	invalidResult := isValid("(]")
	if invalidResult {
		t.Errorf("(] is expected to be invalid")
	}

	invalidResult2 := isValid("([)]")
	if invalidResult2 {
		t.Errorf("([)] is expected to be invalid")
	}

	invalidResult3 := isValid("(){}}{")
	if invalidResult3 {
		t.Errorf("(){}}{ expected to be invalid")
	}

	invalidResult4 := isValid("[([]])")
	if invalidResult4 {
		t.Errorf("[([]]) is expected to be invalid")
	}
}
