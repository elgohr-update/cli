// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

// nolint: dupl // ignore similar code among actions
package pipeline

import (
	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/sdk-go/vela"

	"github.com/sirupsen/logrus"
)

// View inspects a pipeline based off the provided configuration.
func (c *Config) View(client *vela.Client) error {
	logrus.Debug("executing view for pipeline configuration")

	logrus.Tracef("inspecting pipeline %s/%s@%s", c.Org, c.Repo, c.Ref)

	// set the pipeline options for the call
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#PipelineOptions
	opts := &vela.PipelineOptions{
		Output: c.Output,
		Ref:    c.Ref,
	}

	// send API call to capture a pipeline
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#PipelineService.Get
	pipeline, _, err := client.Pipeline.Get(c.Org, c.Repo, opts)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the pipeline in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(pipeline)
	case output.DriverJSON:
		// output the pipeline in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(pipeline)
	case output.DriverSpew:
		// output the pipeline in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(pipeline)
	case output.DriverYAML:
		// output the pipeline in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(pipeline)
	default:
		// output the pipeline in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(pipeline)
	}
}
