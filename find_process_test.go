package main

import "testing"

func TestFindProcess(t *testing.T) {
	t.Run("Find the process by PID", func(t *testing.T) {
		foundPath, err := findProcess(8405)
		if err != nil {
			t.Errorf("got an error, did not want one. %q", err)
		}
		if foundPath != "/proc/8405" {
			t.Errorf("wanted %q got %q", "/proc/8405", foundPath)
		}
	})

	t.Run("Find a process that doesn't exist", func(t *testing.T) {
		foundPath, err := findProcess(123412341234)

		if err == nil {
			t.Errorf("should have gotten an error, did not")
		}
		if foundPath != "" {
			t.Errorf("foundPath should have been empty string. got %q", foundPath)
		}
	})
}
