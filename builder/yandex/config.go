//go:generate struct-markdown
//go:generate mapstructure-to-hcl2 -type Config

package yandex

import (
	"errors"
	"fmt"

	"github.com/hashicorp/packer-plugin-sdk/common"
	"github.com/hashicorp/packer-plugin-sdk/communicator"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"
)

type Config struct {
	common.PackerConfig `mapstructure:",squash"`
	Communicator        communicator.Config `mapstructure:",squash"`
	AccessConfig        `mapstructure:",squash"`

	CommonConfig `mapstructure:",squash"`
	ImageConfig  `mapstructure:",squash"`

	// Service account identifier to assign to instance.
	ServiceAccountID string `mapstructure:"service_account_id" required:"false"`

	// The source image family to create the new image
	// from. You can also specify source_image_id instead. Just one of a source_image_id or
	// source_image_family must be specified. Example: `ubuntu-1804-lts`.
	SourceImageFamily string `mapstructure:"source_image_family" required:"true"`
	// The ID of the folder containing the source image.
	SourceImageFolderID string `mapstructure:"source_image_folder_id" required:"false"`
	// The source image ID to use to create the new image from.
	SourceImageID string `mapstructure:"source_image_id" required:"false"`
	// The source image name to use to create the new image
	// from. Name will be looked up in `source_image_folder_id`.
	SourceImageName string `mapstructure:"source_image_name"`
	// The ID of the folder to save built image in.
	// This defaults to value of 'folder_id'.
	TargetImageFolderID string `mapstructure:"target_image_folder_id" required:"false"`

	ctx interpolate.Context
}

func (c *Config) Prepare(raws ...interface{}) ([]string, error) {
	c.ctx.Funcs = TemplateFuncs
	err := config.Decode(c, &config.DecodeOpts{
		Interpolate:        true,
		InterpolateContext: &c.ctx,
	}, raws...)
	if err != nil {
		return nil, err
	}

	// Accumulate any errors
	var errs *packersdk.MultiError

	errs = packersdk.MultiErrorAppend(errs, c.AccessConfig.Prepare(&c.ctx)...)

	errs = c.CommonConfig.Prepare(errs)
	errs = c.ImageConfig.Prepare(errs)

	if c.ImageMinDiskSizeGb == 0 {
		c.ImageMinDiskSizeGb = c.DiskSizeGb
	}

	if c.ImageMinDiskSizeGb < c.DiskSizeGb {
		errs = packersdk.MultiErrorAppend(errs,
			fmt.Errorf("Invalid image_min_disk_size value (%d): Must be equal or greate than disk_size_gb (%d)",
				c.ImageMinDiskSizeGb, c.DiskSizeGb))
	}

	if c.DiskName == "" {
		c.DiskName = c.InstanceName + "-disk"
	}

	if es := c.Communicator.Prepare(&c.ctx); len(es) > 0 {
		errs = packersdk.MultiErrorAppend(errs, es...)
	}

	// Process required parameters.
	if c.SourceImageID == "" {
		if c.SourceImageFamily == "" && c.SourceImageName == "" {
			errs = packersdk.MultiErrorAppend(
				errs, errors.New("a source_image_name or source_image_family must be specified"))
		}
		if c.SourceImageFamily != "" && c.SourceImageName != "" {
			errs = packersdk.MultiErrorAppend(
				errs, errors.New("one of source_image_name or source_image_family must be specified, not both"))
		}
	}

	if c.TargetImageFolderID == "" {
		c.TargetImageFolderID = c.FolderID
	}

	// Check for any errors.
	if errs != nil && len(errs.Errors) > 0 {
		return nil, errs
	}

	return nil, nil
}
