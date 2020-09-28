package main

import (
	"testing"
)

func TestArgsToString(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{"hey", "there"}, "hey there"},
		{[]string{"you are", "very", "silly"}, "you are very silly"},
		{[]string{"      so spacey", "it's crazy    "}, "so spacey it's crazy"},
		{[]string{}, ""},
	}

	for i, test := range tests {
		if output := ArgsToString(test.input); output != test.expected {
			t.Fatalf("test %d: Inputted: %v, Expected: %v, Received: %v ", i+1, test.input, test.expected, output)
		}
	}
}

func TestTableToSarcastic(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"nice", "nIcE"},
		{"I hadn't thought of it like that before.", "i HaDn'T tHoUgHt Of It LiKe ThAt BeFoRe."},
		{"that's a good idea", "tHaT's A gOoD iDeA"},
		{"", ""},
	}

	for i, test := range tests {
		if output := ToSarcastic(test.input); output != test.expected {
			t.Fatalf("test %d: Inputted: %v, Expected: %v, Received: %v ", i+1, test.input, test.expected, output)
		}
	}
}
