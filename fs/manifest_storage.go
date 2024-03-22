package fs

import (
	"fmt"
	"os"
	"time"

	exact "github.com/williammartin/gh-exact"
	"gopkg.in/yaml.v3"
)

type ManifestStorage struct {
	FilePath string
}

func (m ManifestStorage) Store(manifest exact.Manifest) error {
	extensionDTOs := make([]extensionDTO, len(manifest.Extensions))
	for i, extension := range manifest.Extensions {
		extensionDTOs[i] = extensionDTO{
			Name:    extension.Name,
			Repo:    extension.Repo,
			Version: extension.Version,
		}
	}

	f, err := os.OpenFile(m.FilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("store: opening manifest: %v", err)
	}
	defer f.Close()

	mDTO := manifestDTO{
		CreatedAt:  manifest.CreatedAt,
		Extensions: extensionDTOs,
	}

	if err := yaml.NewEncoder(f).Encode(mDTO); err != nil {
		return fmt.Errorf("store: writing manifest: %v", err)
	}

	return nil
}

func (m ManifestStorage) Load() (exact.Manifest, error) {
	f, err := os.Open(m.FilePath)
	if err != nil {
		return exact.Manifest{}, fmt.Errorf("load: opening manifest: %v", err)
	}

	var mDTO manifestDTO
	if err := yaml.NewDecoder(f).Decode(&mDTO); err != nil {
		return exact.Manifest{}, fmt.Errorf("load: reading manifest: %v", err)
	}

	extensions := make([]exact.Extension, len(mDTO.Extensions))
	for i, extensionDTO := range mDTO.Extensions {
		extensions[i] = exact.Extension{
			Name:    extensionDTO.Name,
			Repo:    extensionDTO.Repo,
			Version: extensionDTO.Version,
		}
	}

	return exact.Manifest{
		CreatedAt:  mDTO.CreatedAt,
		Extensions: extensions,
	}, nil
}

type manifestDTO struct {
	CreatedAt  time.Time      `yaml:"created_at"`
	Extensions []extensionDTO `yaml:"extensions"`
}

type extensionDTO struct {
	Name    string `yaml:"name"`
	Repo    string `yaml:"repo"`
	Version string `yaml:"version"`
}
