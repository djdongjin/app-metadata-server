package metadataserver

import "fmt"

var data = `
title: Valid App 1
version: 0.0.1
maintainers:
- name: firstmaintainer app1
  email: firstmaintainer@hotmail.com
- name: secondmaintainer app1
  email: secondmaintainer@gmail.com
company: Random Inc.
website: https://website.com
source: https://github.com/random/repo
license: Apache-2.0
description: |
 ### Interesting Title
 Some application content, and description
`

var data2 = `
title: Valid App 2
version: 1.0.1
maintainers:
- name: AppTwo Maintainer
  email: apptwo@hotmail.com
company: Upbound Inc.
website: https://upbound.io
source: https://github.com/upbound/repo
license: Apache-2.0
description: |
 ### Why app 2 is the best
 Because it simply is...
`

func main() {
	metadata, _ := YamlStringToMetadata(data)
	fmt.Println(metadata)

	fmt.Printf("\n\n")
	fmt.Println(MetadataToYamlString(metadata))

	md2, _ := YamlStringToMetadata(data2)
	fmt.Println(md2)

	fmt.Printf("\n\n")
	fmt.Println(MetadataToYamlString(md2))
}
