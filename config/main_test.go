package config

import (
	"testing"
)

type Settings struct {
	Database map[string]interface{}
}

func TestFail(t *testing.T) {
	var settings *Settings
	settings = &Settings{}
	err := Load(settings)

	if err == nil {
		t.Logf("Should have failed.")
	}

}
