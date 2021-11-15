package common

import (
	"testing"
)

func compareSubQuerySlice(s1, s2 []SubQuery) bool {
	if len(s1) != len(s2) {
		return false
	}

	for idx, q1 := range s1 {
		if q1.Key != s2[idx].Key || q1.Op != s2[idx].Op || q1.Value != s2[idx].Value {
			return false
		}
	}

	return true
}

func TestParseQuery(t *testing.T) {
	valid := "title = Valid App 1,version != 0.0.2"
	res := []SubQuery{
		{"title", "=", "Valid App 1"},
		{"version", "!=", "0.0.2"},
	}

	subs, err := ParseQuery(valid)
	if err != nil {
		t.Fatalf("ParseQuery failed, no err should be returned. Error: %v", err.Error())
	}
	if !compareSubQuerySlice(subs, res) {
		t.Fatalf("ParseQuery failed, expected: %v, actual: %v.", res, subs)
	}

	invalid := "title = Valid App 1,versionNumber != 0.0.2"
	subs, err = ParseQuery(invalid)
	if err == nil {
		t.Fatalf("ParseQuery should fail on query %v", invalid)
	}
}
