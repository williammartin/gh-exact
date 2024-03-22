package exact

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Versioning string

const (
	PinVersions    Versioning = "pin"
	LatestVersions Versioning = "latest"
)

type App struct {
	FilePath         string
	ExtensionManager ExtensionManager
}

func (a App) Dump() error {
	extensions, err := a.ExtensionManager.List()
	if err != nil {
		return fmt.Errorf("dump: listing extensions: %v", err)
	}

	extensionDTOs := make([]extensionDTO, len(extensions))
	for i, extension := range extensions {
		extensionDTOs[i] = extensionDTO{
			Name:    extension.Name,
			Repo:    extension.Repo,
			Version: extension.Version,
		}
	}

	f, err := os.OpenFile(a.FilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("dump: opening extensions: %v", err)
	}
	defer f.Close()

	if err := yaml.NewEncoder(f).Encode(extensionDTOs); err != nil {
		return fmt.Errorf("dump: writing extensions %v", err)
	}

	return nil
}

func (a App) Restore(v Versioning) error {
	f, err := os.Open(a.FilePath)
	if err != nil {
		return fmt.Errorf("restore: opening extensions: %v", err)
	}

	var extensionDTOs []extensionDTO
	if err := yaml.NewDecoder(f).Decode(&extensionDTOs); err != nil {
		return fmt.Errorf("restore: reading extensions: %v", err)
	}

	extensions := make([]Extension, len(extensionDTOs))
	for i, extensionDTO := range extensionDTOs {
		extensions[i] = Extension{
			Name:    extensionDTO.Name,
			Repo:    extensionDTO.Repo,
			Version: extensionDTO.Version,
		}
	}

	if err := a.ExtensionManager.Install(extensions, v); err != nil {
		return fmt.Errorf("restore: installing extensions: %v", err)
	}

	return nil
}

type extensionDTO struct {
	Name    string `json:"name"`
	Repo    string `json:"repo"`
	Version string `json:"version"`
}
