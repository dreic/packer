// Code generated by "mapstructure-to-hcl2 -type ImageDestination"; DO NOT EDIT.

package common

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatImageDestination is an auto-generated flat version of ImageDestination.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatImageDestination struct {
	ProjectId   *string `mapstructure:"project_id" required:"false" cty:"project_id" hcl:"project_id"`
	Region      *string `mapstructure:"region" required:"false" cty:"region" hcl:"region"`
	Name        *string `mapstructure:"name" required:"false" cty:"name" hcl:"name"`
	Description *string `mapstructure:"description" required:"false" cty:"description" hcl:"description"`
}

// FlatMapstructure returns a new FlatImageDestination.
// FlatImageDestination is an auto-generated flat version of ImageDestination.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*ImageDestination) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatImageDestination)
}

// HCL2Spec returns the hcl spec of a ImageDestination.
// This spec is used by HCL to read the fields of ImageDestination.
// The decoded values from this spec will then be applied to a FlatImageDestination.
func (*FlatImageDestination) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"project_id":  &hcldec.AttrSpec{Name: "project_id", Type: cty.String, Required: false},
		"region":      &hcldec.AttrSpec{Name: "region", Type: cty.String, Required: false},
		"name":        &hcldec.AttrSpec{Name: "name", Type: cty.String, Required: false},
		"description": &hcldec.AttrSpec{Name: "description", Type: cty.String, Required: false},
	}
	return s
}
