package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func buildBinary(t *testing.T) string {
	t.Helper()

	dir := t.TempDir()
	bin := filepath.Join(dir, "zangyo-alert")

	cmd := exec.Command("go", "build", "-o", bin, ".")
	cmd.Dir = "."
	cmd.Env = append(os.Environ(), "CGO_ENABLED=0")

	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("build failed: %v\n%s", err, string(out))
	}
	return bin
}

func TestVersion(t *testing.T) {
	bin := buildBinary(t)

	cmd := exec.Command(bin, "--version")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("--version failed: %v\n%s", err, string(out))
	}
	if string(out) != "v0.1.0\n" {
		t.Fatalf("--version output = %q, want %q", string(out), "v0.1.0\n")
	}
}

func TestHelp(t *testing.T) {
	bin := buildBinary(t)

	cmd := exec.Command(bin, "--help")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("--help failed: %v\n%s", err, string(out))
	}

	for _, want := range [][]byte{
		[]byte("Usage: zangyo-alert"),
		[]byte("--config=STRING"),
		[]byte("--version"),
	} {
		if !bytes.Contains(out, want) {
			t.Fatalf("--help output missing %q\n%s", want, string(out))
		}
	}
}
