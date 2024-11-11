package leetcode

import "testing"

func Test_strStr(t *testing.T) {
	caseOne := strStr("sadbutsad", "sad")
	if caseOne != 0 {
		t.Errorf("expected 0, got %d", caseOne)
	}

	caseTwo := strStr("leetcode", "leeto")
	if caseTwo != -1 {
		t.Errorf("expected -1, got %v", caseTwo)
	}
}
