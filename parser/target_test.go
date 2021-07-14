package parser

import (
	"fmt"
	"testing"
)

func TestIsTargetHeader(t *testing.T) {
	tests := []struct {
		str    string
		expect bool
	}{
		{"aaa", false},
		{"fds:main,fds,ff", true},
		{"main: $(SRC", true},
	}

	for _, tt := range tests {
		if tt.expect != isTargetDefinedLine(tt.str) {
			t.Fatalf("%v should be %v", tt.str, tt.expect)
		}
	}
}

func arrayEq(arr1, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for k := range arr1 {
		if arr1[k] != arr2[k] {
			return false
		}
	}

	return true
}

func TestArrayEq(t *testing.T) {
	arr1 := []string{"this", "is", "an", "Array"}
	arr2 := []string{"this", "is", "an", "Array"}
	arr3 := []string{}

	if !arrayEq(arr1, arr2) {
		t.Fail()
	}
	if arrayEq(arr1, arr3) {
		t.Fail()
	}
}

func TestParseHeader(t *testing.T) {
	tests := []struct {
		header string
		name   string
		deps   []string
	}{
		{"main:all", "main", []string{"all"}},
		{"main:", "main", []string{""}},
	}

	for _, test := range tests {
		target := parseTargetNameAndDeps(test.header)

		if target.Name != test.name {
			t.Fail()
		}

		if !arrayEq(target.Deps, test.deps) {
			fmt.Printf("%#v\n", target)
			t.Fail()
		}
	}
}
