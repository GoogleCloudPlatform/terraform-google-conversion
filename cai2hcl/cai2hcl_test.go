package cai2hcl

import (
	"testing"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"

	"github.com/google/go-cmp/cmp"
)

func TestGroupAssets(t *testing.T) {
	tests := []struct {
		name   string
		assets []*caiasset.Asset
		want   map[string][]*caiasset.Asset
	}{
		{
			name: "nil and empty value",
			assets: []*caiasset.Asset{
				nil,
				{
					Type: "cloudbilling.googleapis.com/ProjectBillingInfo",
				},
				{
					Name: "//cloudresourcemanager.googleapis.com/projects/example-project",
				},
			},
			want: map[string][]*caiasset.Asset{},
		},
		{
			name: "same name is overridden",
			assets: []*caiasset.Asset{
				nil,
				{
					Name: "name",
					Type: "t1",
				},
				{
					Name: "name",
					Type: "t2",
				},
			},
			want: map[string][]*caiasset.Asset{
				"name": {
					{
						Name: "name",
						Type: "t2",
					},
				},
			},
		},
		{
			name: "project",
			assets: []*caiasset.Asset{
				{
					Name: "//cloudresourcemanager.googleapis.com/projects/example-project",
					Type: "cloudresourcemanager.googleapis.com/Project",
				},
				{
					Name: "//cloudbilling.googleapis.com/projects/example-project/billingInfo",
					Type: "cloudbilling.googleapis.com/ProjectBillingInfo",
				},
			},
			want: map[string][]*caiasset.Asset{
				"//cloudresourcemanager.googleapis.com/projects/example-project": []*caiasset.Asset{
					{
						Name: "//cloudresourcemanager.googleapis.com/projects/example-project",
						Type: "cloudresourcemanager.googleapis.com/Project",
					},
					{
						Name: "//cloudbilling.googleapis.com/projects/example-project/billingInfo",
						Type: "cloudbilling.googleapis.com/ProjectBillingInfo",
					},
				},
			},
		},
		{
			name: "bigtable",
			assets: []*caiasset.Asset{
				{
					Name: "//bigtable.googleapis.com/projects/example-project/instances/tf-instance/clusters/placeholder-foobar",
					Type: "bigtableadmin.googleapis.com/Cluster",
				},
				{
					Name: "//bigtable.googleapis.com/projects/example-project/instances/tf-instance",
					Type: "bigtableadmin.googleapis.com/Instance",
				},
			},
			want: map[string][]*caiasset.Asset{
				"//bigtable.googleapis.com/projects/example-project/instances/tf-instance": {
					{
						Name: "//bigtable.googleapis.com/projects/example-project/instances/tf-instance",
						Type: "bigtableadmin.googleapis.com/Instance",
					},
					{
						Name: "//bigtable.googleapis.com/projects/example-project/instances/tf-instance/clusters/placeholder-foobar",
						Type: "bigtableadmin.googleapis.com/Cluster",
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := groupAssets(test.assets)
			if diff := cmp.Diff(test.want, got, cmp.AllowUnexported(caiasset.Asset{})); diff != "" {
				t.Errorf("groupAssets() returned unexpected diff (-want +got):\n%s", diff)
			}
		})
	}
}
