// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package service

// Config represents the configuration necessary
// to perform service related requests with Vela.
type Config struct {
	Action  string
	Org     string
	Repo    string
	Build   int
	Number  int
	Page    int
	PerPage int
	Output  string
}
