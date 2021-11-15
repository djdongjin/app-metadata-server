package metadataserver

import (
	"github.com/djdongjin/app-metadata-server/common"

	yaml "gopkg.in/yaml.v2"
)

func YamlStringToMetadata(ym string) (common.Metadata, error) {
	var metadata common.Metadata

	err := yaml.Unmarshal([]byte(ym), &metadata)
	if err != nil {
		return metadata, err
	}

	return metadata, nil
}

func MetadataToYamlString(p common.Metadata) string {
	ym, err := yaml.Marshal(p)
	if err != nil {
		return ""
	}

	return string(ym)
}
