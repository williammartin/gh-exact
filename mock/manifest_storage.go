package mock

import (
	exact "github.com/williammartin/gh-exact"
)

type ManifestStorage struct {
}

func (m ManifestStorage) Store(manifest exact.Manifest) error {
	return nil
}

func (m ManifestStorage) Load() (exact.Manifest, error) {
	return exact.Manifest{}, nil
}
