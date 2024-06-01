package main

import "testing"


func TestFindNumbers(t *testing.T) {
	input_message := []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}

	got := FindNumbers(input_message)
	want := 142

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
	
}