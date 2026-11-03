package input

import (
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/ffuf/ffuf/v2/pkg/ffuf"
)

func TestCommandPositionEnvironmentIsIsolated(t *testing.T) {
	t.Setenv("FFUF_NUM", "parent-fixture")
	command := `printf '%s' "$FFUF_NUM"`
	if runtime.GOOS == "windows" {
		command = "echo %FFUF_NUM%"
	}
	provider, err := NewCommandInput("FUZZ", command, &ffuf.Config{})
	if err != nil {
		t.Fatal(err)
	}
	provider.SetPosition(37)
	if got := strings.TrimSpace(string(provider.Value())); got != "37" {
		t.Fatalf("child position = %q, want 37", got)
	}
	if got := os.Getenv("FFUF_NUM"); got != "parent-fixture" {
		t.Fatalf("child changed parent environment to %q", got)
	}
}
