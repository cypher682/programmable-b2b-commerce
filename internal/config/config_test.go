package config

import "testing"

func TestLoadDefaults(t *testing.T) {
	t.Setenv("PORT", "")
	t.Setenv("SERVICE_NAME", "")
	t.Setenv("LOG_LEVEL", "")
	t.Setenv("DATABASE_URL", "")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if cfg.Port != 8081 {
		t.Fatalf("Port = %d, want 8081", cfg.Port)
	}
	if cfg.ServiceName != "b2b" {
		t.Fatalf("ServiceName = %q, want b2b", cfg.ServiceName)
	}
}

func TestLoadRejectsInvalidPort(t *testing.T) {
	t.Setenv("PORT", "0")

	if _, err := Load(); err == nil {
		t.Fatal("Load() = nil error, want invalid PORT error")
	}
}
