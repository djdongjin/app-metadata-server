# Coding Exercise - Application Metadata API Server

## Requirements

* Build a Golang RESTful API server for application metadata.
* An endpoint to persist application metadata (In memory is fine). The API must support YAML as a valid payload format.
* An endpoint to search application metadata and retrieve a list that matches the query parameters.
* Include tests if you feel it’s appropriate.

We've provided example yaml data payloads. Two that should persist, and two that should error due to missing fields.

## "Rules"

Use golang for the server, but any other software or open source libraries are fair game to help you solve this problem. The response from the server as well as the structure of the query endpoint is intentionally vague to allow latitude in your solution.

## Advice

This exercise is an opportunity to show off your passion and the craftsmanship of your solution. Optimize your solution for quality and reliability. If you feel your solution is missing a cool feature and you have time, have fun and add it. Make the solution your own, and show off your skills.

## What about the database?

It's recommended that you don't use a database. Integrating with a database driver or ORM gives you less room to shine, and us less ability to evaluate your work.

## Example payloads

All fields in the payload are required. For illustration purposes, we have a few example payloads. One example payload where the maintainer email is not a valid email and another where the version is missing that should fail on submit and two that should be valid.

### Invalid Payloads

```yaml
title: App w/ Invalid maintainer email
version: 1.0.1
maintainers:
- name: Firstname Lastname
  email: apptwohotmail.com
company: Upbound Inc.
website: https://upbound.io
source: https://github.com/upbound/repo
license: Apache-2.0
description: |
 ### blob of markdown
 More markdown
```

```yaml
title: App w/ missing version
maintainers:
- name: first last
  email: email@hotmail.com
- name: first last
  email: email@gmail.com
company: Company Inc.
website: https://website.com
source: https://github.com/company/repo
license: Apache-2.0
description: |
 ### blob of markdown
 More markdown
```

### Valid Payloads

```yaml
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
```

```yaml
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
```
