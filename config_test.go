package kaonashi

import "testing"

func TestConfig(t *testing.T) {
	testConfigFilePath := "./conf/config_test.toml"
	config, err := NewAppConfig(testConfigFilePath)
	if err != nil {
		t.Fatalf("failed to create config: %s", err)
	}
	if config.DatabasePath != ":memory:" {
		t.Fatalf("expected pgtest, but got %s", config.DatabasePath)
	}
}
