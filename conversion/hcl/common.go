package hcl

import (
	tfvResource "github.com/GoogleCloudPlatform/terraform-validator/converters/google/resources"
)

type Asset tfvResource.Asset

type ConvertFunc func(asset *Asset) (string, map[string]interface{}, error)

type Converter struct {
	TFResourceName string
	Convert        ConvertFunc
}

func Converters() map[string]*Converter {
	return map[string]*Converter{
		tfvResource.ComputeInstanceAssetType: NewComputeInstanceConverter(),
	}
}
