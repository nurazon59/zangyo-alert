package main

import (
	"os"
	"path/filepath"

	"github.com/alecthomas/kong"
)

const appVersion = "v0.1.0"

var CLI struct {
	Config  string           `help:"config json path." required:"" env:"ZANGYO_CONFIG"`
	Import  string           `help:"attendance xlsx path."`
	Grades  string           `help:"grades csv path. defaults to data/social-insurance-tokyo-2026.csv next to the binary."`
	Version kong.VersionFlag `name:"version" help:"Print version information and quit."`
}

func main() {
	kong.Parse(&CLI, kong.Name("zangyo-alert"), kong.Vars{"version": appVersion})

	gradesPath := CLI.Grades
	if gradesPath == "" {
		gradesPath = defaultGradesPath()
	}

	os.Exit(Run(Args{
		ConfigPath: CLI.Config,
		ImportPath: CLI.Import,
		GradesPath: gradesPath,
	}))
}

func defaultGradesPath() string {
	candidates := []string{
		"data/social-insurance-tokyo-2026.csv",
	}
	if exe, err := os.Executable(); err == nil {
		candidates = append(candidates, filepath.Join(filepath.Dir(exe), "data", "social-insurance-tokyo-2026.csv"))
	}
	for _, p := range candidates {
		if _, err := os.Stat(p); err == nil {
			return p
		}
	}
	return candidates[0]
}
