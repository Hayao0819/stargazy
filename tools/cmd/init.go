package cmd

import (
	"bytes"
	"log/slog"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

func install_lint() error {
	slog.Info("Installing golangci-lint")

	// Check path
	if _, err := exec.LookPath("golangci-lint"); err == nil {
		slog.Info("golangci-lint has been installed. Skip installing it.")
		return nil
	}

	// Setup command
	installer := exec.Command("sh", "-c", "curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin")
	installer.Stdout = os.Stderr
	installer.Stderr = os.Stderr

	// Run command
	err := installer.Run()
	if err == nil {
		slog.Info("golangci-lint has been successfully installed")
	} else {
		slog.Error("Installation of golangci-lint has failed")
	}
	return err
}

func setup_githook() error {
	slog.Info("Set up githook")

	// Check config
	configData := bytes.Buffer{}
	gitrefercmd := exec.Command("git", "config", "--local", "core.hooksPath")
	gitrefercmd.Stdout = &configData
	if err := gitrefercmd.Run(); err != nil {
		slog.Warn("Failed to get current the value of core.hooksPath")
	} else {
		if strings.TrimSpace(configData.String()) == ".githooks" {
			slog.Info("Git Hooks directory has already been .githooks")
			return nil
		}
	}

	// Setup command
	gitcmd := exec.Command("git", "config", "--local", "core.hooksPath", ".githooks")
	gitcmd.Stdout = os.Stderr
	gitcmd.Stderr = os.Stderr

	// Cun command
	err := gitcmd.Run()
	if err == nil {
		slog.Info("Git Hooks directory has been changed to .githooks")
	} else {
		slog.Error("Changing Git Hooks directory has failed")
	}
	return err
}

func InitCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "init",
		Args:  cobra.NoArgs,
		Short: "Setup a development enviromnemt",
		Long:  "Install and configure the necessary tools.",
		RunE: func(cmd *cobra.Command, args []string) error {
			funcs := []func() error{install_lint, setup_githook}
			for _, f := range funcs {
				if err := f(); err != nil {
					return err
				}
			}
			return nil
		},
	}

	return &cmd
}
