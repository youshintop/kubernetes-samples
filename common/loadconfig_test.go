package common

import "testing"

func TestLoadconfig(t *testing.T) {
	if cfg := LoadConfig(); cfg == nil {
		t.Error("Failed load kubeconfig.")
	} else {
		t.Log("LoadConfig Func 测试通过.")
	}
}
