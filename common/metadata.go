package common

import "fmt"

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

func (m Maintainer) Validate() (err error) {
	if err = NotEmpty("Name", m.Name); err != nil {
		return
	}

	if err = ValidEmail("Email", m.Email); err != nil {
		return
	}

	return nil
}

func (m Metadata) Validate() (err error) {
	// Validate string existance
	if err = NotEmpty("Title", m.Title); err != nil {
		return
	}
	if err = NotEmpty("Version", m.Version); err != nil {
		return
	}
	if err = NotEmpty("Company", m.Company); err != nil {
		return
	}
	if err = NotEmpty("Website", m.Website); err != nil {
		return
	}
	if err = NotEmpty("Source", m.Source); err != nil {
		return
	}
	if err = NotEmpty("License", m.License); err != nil {
		return
	}
	if err = NotEmpty("Description", m.Description); err != nil {
		return
	}

	// Validate urls
	if err = ValidUrl("Website", m.Website); err != nil {
		return
	}
	if err = ValidUrl("Source", m.Source); err != nil {
		return
	}

	// Validate maintainers
	if len(m.Maintainers) == 0 {
		return fmt.Errorf("Maintainers is empty.")
	}

	for _, maintainer := range m.Maintainers {
		if err = maintainer.Validate(); err != nil {
			return err
		}
	}

	return nil
}
