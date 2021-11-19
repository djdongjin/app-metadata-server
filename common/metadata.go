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

func (m Maintainer) Validate() error {
	errs := []error{
		NotEmpty("Name", m.Name),
		ValidEmail("Email", m.Email),
	}

	for _, err := range errs {
		if err != nil {
			return err
		}
	}

	return nil
}

func (m Metadata) Validate() error {
	errs := []error{
		// Validate string existance
		NotEmpty("Title", m.Title),
		NotEmpty("Version", m.Version),
		NotEmpty("Company", m.Company),
		NotEmpty("Website", m.Website),
		NotEmpty("Source", m.Source),
		NotEmpty("License", m.License),
		NotEmpty("Description", m.Description),
		// Validate urls
		ValidUrl("Website", m.Website),
		ValidUrl("Source", m.Source),
	}

	for _, err := range errs {
		if err != nil {
			return err
		}
	}

	// Validate maintainers
	if len(m.Maintainers) == 0 {
		return fmt.Errorf("Maintainers is empty.")
	}

	for _, maintainer := range m.Maintainers {
		if err := maintainer.Validate(); err != nil {
			return err
		}
	}

	return nil
}
