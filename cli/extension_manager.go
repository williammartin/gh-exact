package cli

import (
	"encoding/csv"
	"fmt"
	"io"
	"os/exec"

	exact "github.com/williammartin/gh-exact"
)

type ExtensionManager struct {
	Out io.Writer
	Err io.Writer
}

func (e ExtensionManager) List() ([]exact.Extension, error) {
	cmd := exec.Command("gh", "extension", "list")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("list: stdout pipe: %v", err)
	}

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("list: cmd start: %v", err)
	}

	r := csv.NewReader(stdout)
	r.Comma = '\t'

	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	if err := cmd.Wait(); err != nil {
		// TODO: Capture stderr in case of err
		return nil, fmt.Errorf("list: cmd wait: %v", err)
	}

	extensions := make([]exact.Extension, len(records))
	for i, record := range records {
		extensions[i] = exact.Extension{
			Name:    record[0],
			Repo:    record[1],
			Version: record[2],
		}
	}

	return extensions, nil
}

func (e ExtensionManager) Install(extensions []exact.Extension, v exact.Versioning) error {
	for _, extension := range extensions {
		// Offer the chance to install local extensions interactively.
		if extension.IsLocal() {
			continue
		}

		args := []string{"extension", "install", extension.Repo}
		if v == exact.PinVersions {
			args = append(args, "--pin", extension.Version)
		}

		cmd := exec.Command("gh", args...)
		cmd.Stdout = e.Out
		cmd.Stderr = e.Err

		fmt.Fprintf(e.Out, "Installing %q\n", extension.Name)
		// Bail on first failure
		if err := cmd.Run(); err != nil {
			// TODO: Capture stderr in case of err
			return fmt.Errorf("install: installing %q: %v", extension.Name, err)
		}
	}

	return nil
}
