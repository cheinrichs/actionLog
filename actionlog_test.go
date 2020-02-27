package actionlog

import "testing"

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
	var expected string = ""
	var actual string = GetStats()

	if expected != actual {
		t.Errorf("Get Stats failed, got: %s, want: %s.", actual, expected)
	}
}
