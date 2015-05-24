package basics

import "testing"

func TestTruth(t *testing.T) {
	if true != true {
		t.Error("everything I know is wrong")
	}
}
