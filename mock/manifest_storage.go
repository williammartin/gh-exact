package mock

import (
	exact "github.com/williammartin/gh-exact"
)

type ManifestStorage struct {
	storedManifest exact.Manifest
}

func (m *ManifestStorage) Store(manifest exact.Manifest) error {
	m.storedManifest = manifest
	return nil
}

func (m *ManifestStorage) Load() (exact.Manifest, error) {
	return m.storedManifest, nil
}
