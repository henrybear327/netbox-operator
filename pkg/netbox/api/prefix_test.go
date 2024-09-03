/*
Copyright 2024 Swisscom (Schweiz) AG.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package api

import (
	"testing"

	"github.com/netbox-community/go-netbox/v3/netbox/client/ipam"
	"github.com/netbox-community/go-netbox/v3/netbox/client/tenancy"
	netboxModels "github.com/netbox-community/go-netbox/v3/netbox/models"
	"github.com/netbox-community/netbox-operator/gen/mock_interfaces"
	"github.com/netbox-community/netbox-operator/pkg/netbox/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestPrefix_GetExistingPrefix(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockPrefixIpam := mock_interfaces.NewMockIpamInterface(ctrl)
	mockPrefixTenancy := mock_interfaces.NewMockTenancyInterface(ctrl)

	//tenant mock input
	tenant := "Tenant1"
	tenantListRequestInput := tenancy.NewTenancyTenantsListParams().WithName(&tenant)

	//tenant mock output
	tenantOutputId := int64(1)
	tenantOutputSlug := "tenant1"
	tenantListRequestOutput := &tenancy.TenancyTenantsListOK{
		Payload: &tenancy.TenancyTenantsListOKBody{
			Results: []*netboxModels.Tenant{
				{
					ID:   tenantOutputId,
					Name: &tenant,
					Slug: &tenantOutputSlug,
				},
			},
		},
	}

	//prefix mock input
	prefix := "10.112.140.0/24"
	prefixListRequestInput := ipam.
		NewIpamPrefixesListParams().
		WithPrefix(&prefix)

	//prefix mock output
	prefixId := int64(4)
	site := "Site"
	siteId := int64(2)
	siteSlug := "site"
	comments := "blabla"
	description := "very useful prefix"
	nestedSite := &netboxModels.NestedSite{
		ID:   siteId,
		Name: &site,
		Slug: &siteSlug,
	}
	nestedTenant := &netboxModels.NestedTenant{
		Name: &tenant,
		ID:   tenantOutputId,
		Slug: &tenantOutputSlug,
	}
	prefixListOutput := &ipam.IpamPrefixesListOK{
		Payload: &ipam.IpamPrefixesListOKBody{
			Results: []*netboxModels.Prefix{
				{
					ID:          prefixId,
					Comments:    comments,
					Description: description,
					Display:     prefix,
					Prefix:      &prefix,
					Site:        nestedSite,
					Tenant:      nestedTenant,
				},
			},
		},
	}

	mockPrefixTenancy.EXPECT().TenancyTenantsList(tenantListRequestInput, nil).Return(tenantListRequestOutput, nil).AnyTimes()
	mockPrefixIpam.EXPECT().IpamPrefixesList(prefixListRequestInput, nil).Return(prefixListOutput, nil)

	netboxClient := &NetboxClient{
		Ipam:    mockPrefixIpam,
		Tenancy: mockPrefixTenancy,
	}

	actual, err := netboxClient.GetPrefix(&models.Prefix{
		Prefix: prefix,
		Metadata: &models.NetboxMetadata{
			Tenant:      tenant,
			Comments:    comments,
			Description: description,
		},
	})

	assert.Nil(t, err)
	assert.Equal(t, prefixId, actual.Payload.Results[0].ID)
	assert.Equal(t, comments, actual.Payload.Results[0].Comments)
	assert.Equal(t, description, actual.Payload.Results[0].Description)
	assert.Equal(t, prefix, actual.Payload.Results[0].Display)
	assert.Equal(t, prefix, *actual.Payload.Results[0].Prefix)
	assert.Equal(t, *nestedTenant.Name, *actual.Payload.Results[0].Tenant.Name)
	assert.Equal(t, nestedTenant.ID, actual.Payload.Results[0].Tenant.ID)
	assert.Equal(t, *nestedTenant.Slug, *actual.Payload.Results[0].Tenant.Slug)
	assert.Equal(t, *nestedSite.Name, *actual.Payload.Results[0].Site.Name)
	assert.Equal(t, nestedSite.ID, actual.Payload.Results[0].Site.ID)
	assert.Equal(t, *nestedSite.Slug, *actual.Payload.Results[0].Site.Slug)
}

func TestPrefix_GetNonExistingPrefix(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockPrefixIpam := mock_interfaces.NewMockIpamInterface(ctrl)
	mockPrefixTenancy := mock_interfaces.NewMockTenancyInterface(ctrl)

	//tenant mock input
	tenant := "tenant1"
	tenantListRequestInput := tenancy.NewTenancyTenantsListParams().WithName(&tenant)

	//tenant mock output
	tenantOutputId := int64(1)
	tenantOutputSlug := "tenant1"
	tenantListOutput := &tenancy.TenancyTenantsListOK{
		Payload: &tenancy.TenancyTenantsListOKBody{
			Results: []*netboxModels.Tenant{
				{
					ID:   tenantOutputId,
					Name: &tenant,
					Slug: &tenantOutputSlug,
				},
			},
		},
	}

	//prefix mock input
	prefix := "10.112.140.0/24"
	prefixListRequestInput := ipam.
		NewIpamPrefixesListParams().
		WithPrefix(&prefix)

	//prefix mock output
	emptyPrefixListOutput := &ipam.IpamPrefixesListOK{
		Payload: &ipam.IpamPrefixesListOKBody{
			Results: []*netboxModels.Prefix{},
		},
	}

	mockPrefixTenancy.EXPECT().TenancyTenantsList(tenantListRequestInput, nil).Return(tenantListOutput, nil).AnyTimes()
	mockPrefixIpam.EXPECT().IpamPrefixesList(prefixListRequestInput, nil).Return(emptyPrefixListOutput, nil)

	netboxClient := &NetboxClient{
		Ipam:    mockPrefixIpam,
		Tenancy: mockPrefixTenancy,
	}

	actual, err := netboxClient.GetPrefix(&models.Prefix{
		Prefix: prefix,
		Metadata: &models.NetboxMetadata{
			Tenant: tenant,
		},
	})

	assert.Nil(t, err)
	assert.Len(t, actual.Payload.Results, 0)
}

func TestPrefix_CreatePrefix(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockPrefixIpam := mock_interfaces.NewMockIpamInterface(ctrl)

	//prefix mock input
	prefix := "10.112.140.0/24"
	siteId := int64(1)
	tenantId := int64(1)
	comment := "a comment"
	description := "very useful prefix"
	prefixToCreate := &netboxModels.WritablePrefix{
		Comments:    comment,
		Description: description,
		Prefix:      &prefix,
		Site:        &siteId,
		Tenant:      &tenantId,
	}
	createPrefixInput := ipam.
		NewIpamPrefixesCreateParams().
		WithDefaults().
		WithData(prefixToCreate)

	//prefix mock output
	createPrefixOutput := &ipam.IpamPrefixesCreateCreated{
		Payload: &netboxModels.Prefix{
			ID:          int64(1),
			Comments:    comment,
			Description: description,
			Display:     prefix,
			Prefix:      &prefix,
			Site: &netboxModels.NestedSite{
				ID: siteId,
			},
			Tenant: &netboxModels.NestedTenant{
				ID: tenantId,
			},
		},
	}

	mockPrefixIpam.EXPECT().IpamPrefixesCreate(createPrefixInput, nil).Return(createPrefixOutput, nil)

	netboxClient := &NetboxClient{
		Ipam: mockPrefixIpam,
	}

	actual, err := netboxClient.CreatePrefix(prefixToCreate)
	assert.Nil(t, err)
	assert.Greater(t, actual.ID, int64(0))
	assert.Equal(t, prefix, *actual.Prefix)
	assert.Equal(t, comment, actual.Comments)
	assert.Equal(t, description, actual.Description)
	assert.Equal(t, prefix, actual.Display)
	assert.Equal(t, siteId, actual.Site.ID)
	assert.Equal(t, tenantId, actual.Tenant.ID)
}

func TestPrefix_UpdatePrefix(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockPrefixIpam := mock_interfaces.NewMockIpamInterface(ctrl)

	//prefix mock input
	prefixId := int64(1)
	prefix := "10.112.140.0/24"
	siteId := int64(1)
	tenantId := int64(1)
	comment := "a comment"
	updatedDescription := "updated"
	prefixToUpdate := &netboxModels.WritablePrefix{
		Comments:    comment,
		Description: updatedDescription,
		Prefix:      &prefix,
		Site:        &siteId,
		Tenant:      &tenantId,
	}
	updatePrefixInput := ipam.
		NewIpamPrefixesUpdateParams().
		WithDefaults().
		WithData(prefixToUpdate).
		WithID(prefixId)

	//prefix mock output
	updatePrefixOutput := &ipam.IpamPrefixesUpdateOK{
		Payload: &netboxModels.Prefix{
			ID:          int64(1),
			Comments:    comment,
			Description: updatedDescription,
			Display:     prefix,
			Prefix:      &prefix,
			Site: &netboxModels.NestedSite{
				ID: siteId,
			},
			Tenant: &netboxModels.NestedTenant{
				ID: tenantId,
			},
		},
	}

	mockPrefixIpam.EXPECT().IpamPrefixesUpdate(updatePrefixInput, nil).Return(updatePrefixOutput, nil)

	netboxClient := &NetboxClient{
		Ipam: mockPrefixIpam,
	}

	actual, err := netboxClient.UpdatePrefix(prefixId, prefixToUpdate)
	assert.Nil(t, err)
	assert.Greater(t, actual.ID, int64(0))
	assert.Equal(t, prefix, *actual.Prefix)
	assert.Equal(t, comment, actual.Comments)
	assert.Equal(t, updatedDescription, actual.Description)
	assert.Equal(t, prefix, actual.Display)
	assert.Equal(t, siteId, actual.Site.ID)
	assert.Equal(t, tenantId, actual.Tenant.ID)
}

func TestPrefix_DeletePrefix(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockPrefixIpam := mock_interfaces.NewMockIpamInterface(ctrl)

	//prefix mock input
	prefixId := int64(4)
	deletePrefixInput := ipam.NewIpamPrefixesDeleteParams().WithID(prefixId)
	//prefix mock output
	deletePrefixOutput := &ipam.IpamPrefixesDeleteNoContent{}

	mockPrefixIpam.EXPECT().IpamPrefixesDelete(deletePrefixInput, nil).Return(deletePrefixOutput, nil)

	netboxClient := &NetboxClient{
		Ipam: mockPrefixIpam,
	}

	err := netboxClient.DeletePrefix(prefixId)
	assert.Nil(t, err)
}