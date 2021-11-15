package common

type Maintainer struct {
	Name  string
	Email string
}

type Metadata struct {
	Title       string
	Version     string
	Maintainers []Maintainer
	Company     string
	Website     string
	Source      string
	License     string
	Description string
}
