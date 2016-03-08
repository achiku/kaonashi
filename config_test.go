package kaonashi

import "testing"

func TestConfig(t *testing.T) {
	testConfigFilePath := "./conf/config_test.toml"
	config, err := NewAppConfig(testConfigFilePath)
	if err != nil {
		t.Fatalf("failed to create config: %s", err)
	}
	t.Logf("%#v", config)
	if config.DatabasePath != ":memory:" {
		t.Errorf("expected :memory:, but got %s", config.DatabasePath)
	}
}
