package exact

type ExtensionManager interface {
	List() ([]Extension, error)
	Install(extensions []Extension, versioning Versioning) error
}

type Extension struct {
	Name    string
	Repo    string
	Version string
}

func (e Extension) IsLocal() bool {
	// TODO: Do this at parse time into a struct
	return e.Name != "" && e.Repo == "" && e.Version == ""
}
