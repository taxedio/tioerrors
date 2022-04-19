<img src="assets\taxediologolandscape.jpg" alt="drawing" width="200"/>

<h1 align="center">
  TIOERRORS
</h1>

<h3 align="center">
  <a href="https://taxed.io">taxed.io</a>
</h3>

![GitHub go.mod Go version (branch)](https://img.shields.io/github/go-mod/go-version/taxedio/tioerrors/main?style=for-the-badge) ![GitHub](https://img.shields.io/github/license/taxedio/tioerrors?style=for-the-badge) ![GitLab Release (custom instance)](https://img.shields.io/gitlab/v/release/taxedio/tioerrors?include_prereleases&style=for-the-badge) ![Gitlab pipeline status](https://img.shields.io/gitlab/pipeline-status/taxedio/tioerrors?branch=main&style=for-the-badge) ![Gitlab code coverage](https://img.shields.io/gitlab/coverage/taxedio/tioerrors/main?style=for-the-badge) ![GitHub contributors](https://img.shields.io/github/contributors/taxedio/tioerrors?style=for-the-badge)

[![Go Report Card](https://goreportcard.com/badge/gitlab.com/taxedio/tioerrors)](https://goreportcard.com/report/gitlab.com/taxedio/tioerrors)

# Introduction

A basic package to create structs with tags for easy API response building.

# Example 1

```GO
package main

import (
  "github.com/taxedio/tioerrors"
)

func main(){
  restErr := tioerrors.NewRestError("user input incorrect", http.StatusBadRequest, nil)
  fmt.Println(restErr)
}
```

**console**:

```stdout
message: user input incorrect - status: 400 - error: bad_request - causes: []
```

# Example 2

```GO
package main

import (
  "github.com/taxedio/tioerrors"
)

func main(){
	var (
		errTest []interface{}
	)
	errTest = append(errTest, "example error")
	errTest = append(errTest, "another example error")
	restErr := tioerrors.NewRestError("user input incorrect", http.StatusBadRequest, errTest)
  fmt.Println(restErr)
}
```

**console**:

```stdout
message: user input incorrect - status: 400 - error: bad_request - causes: [example error another example error]
```
