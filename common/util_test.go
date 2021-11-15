package common

import (
	"testing"
)

const data = `
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

func TestYamlStringToMetadata(t *testing.T) {
	_, err := YamlStringToMetadata(data)
	if err != nil {
		t.Fatalf("No error should be returned. Error: %v", err.Error())
	}
}
