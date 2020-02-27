package actionlog

import (
	"errors"
	"testing"
)

func TestAddAction(t *testing.T) {
	var expected error = nil
	var actual = AddAction("{\"action\":\"jump\", \"time\":100}")

	if expected != actual {
		t.Errorf("Adding first action failed, got: %d, want: %d.", actual, expected)
	}
}

func TestSecondAddAction(t *testing.T) {
	var expected error = nil
	var actual = AddAction("{\"action\":\"run\", \"time\":75}")

	if expected != actual {
		t.Errorf("Adding second action failed, got: %s, want: %s.", actual, expected)
	}
}

func TestThirdAddAction(t *testing.T) {
	var expected error = nil
	var actual = AddAction("{\"action\":\"jump\", \"time\":200}")

	if expected != actual {
		t.Errorf("Adding third action failed, got: %s, want: %s.", actual, expected)
	}
}

func TestGetStats(t *testing.T) {
	var expected string = "[{\"action\":\"jump\",\"avg\":150},{\"action\":\"run\",\"avg\":75}]"
	var actual string = GetStats()

	if expected != actual {
		t.Errorf("Get Stats failed, got: %s, want: %s.", actual, expected)
	}
}

func TestAddActionWithInvalidJSON(t *testing.T) {
	var expected error = errors.New("invalid character '1' looking for beginning of object key string")
	var actual = AddAction("{1:\"jump\", \"time\":200}")

	if expected.Error() != actual.Error() {
		t.Errorf("Adding Invalid JSON, should fail, got: %s, want: %s.", actual, expected)
	}
}

func TestAddActionWithInvalidAction(t *testing.T) {
	var expected error = errors.New("no action provided")
	var actual = AddAction("{\"action\":\"\", \"time\":200}")

	if expected.Error() != actual.Error() {
		t.Errorf("Adding Invalid Action, should fail, got: %s, want: %s.", actual, expected)
	}
}

func TestAddActionWithInvalidTime(t *testing.T) {
	var expected error = errors.New("invalid time provided")
	var actual = AddAction("{\"action\":\"jump\", \"time\":-10}")

	if expected.Error() != actual.Error() {
		t.Errorf("Adding Invalid Time, should fail, got: %s, want: %s.", actual, expected)
	}
}
