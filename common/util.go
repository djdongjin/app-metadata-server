package common

import (
	yaml "gopkg.in/yaml.v2"
)

func YamlStringToMetadata(ym string) (Metadata, error) {
	var metadata Metadata

	err := yaml.Unmarshal([]byte(ym), &metadata)
	if err != nil {
		return metadata, err
	}

	return metadata, nil
}

func MetadataToYamlString(p Metadata) string {
	ym, err := yaml.Marshal(p)
	if err != nil {
		return ""
	}

	return string(ym)
}
