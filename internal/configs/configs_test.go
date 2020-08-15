package configs

import (
	"testing"
)

type data struct {
	Name  string
	Types int32
	Array []string
	B     bool
}

func TestLoadConfig(t *testing.T) {
	var v *data = &data{}
	LoadConfig(".", "config_test_data", v)

	if v.Name != "test" {
		t.Errorf("incorrect name binding. Expected %s, Got %s", "test", v.Name)
	}

	if v.Types != 1233 {
		t.Errorf("incorrect types binding. Expected %d, Got %d", 1233, v.Types)
	}

	if v.Array[0] != "test1" || v.Array[1] != "test2" || v.Array[2] != "test3" {
		t.Errorf("incorrect array binding. Expected %s, Got %v",
			"test1, test2, test3", v.Array)
	}

	if v.B != true {
		t.Errorf("incorrect types binding. Expected %v, Got %v", true, v.B)
	}
}