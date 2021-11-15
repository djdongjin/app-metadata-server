package common

import (
	"testing"
)

const validYaml = `
title: Valid App 1
version: 0.0.1
maintainers:
- name: firstmaintainer app1
  email: firstmaintainer@hotmail.com
- name: secondmaintainer app1
  email: secondmaintainer@gmail.com
company: Random Inc.
website: https://website.com
source: https://github.com/random/repo
license: Apache-2.0
description: |
 ### Interesting Title
 Some application content, and description
`

const validQuery = "title = Valid App 1,version != 0.0.2"

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
	res := []SubQuery{
		{"title", "=", "Valid App 1"},
		{"version", "!=", "0.0.2"},
	}

	subs, err := ParseQuery(validQuery)
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

func TestExecuteQuery(t *testing.T) {
	md, _ := YamlStringToMetadata(validYaml)
	subqueries, _ := ParseQuery(validQuery)
	if !ExecuteQuery(md, subqueries) {
		t.Fatalf("ExecuteQuery failed.")
	}
}
