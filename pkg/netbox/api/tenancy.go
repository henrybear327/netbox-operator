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
	"errors"

	"github.com/netbox-community/go-netbox/v3/netbox/client/tenancy"

	"github.com/netbox-community/netbox-operator/pkg/netbox/models"
	"github.com/netbox-community/netbox-operator/pkg/netbox/utils"
)

func (r *NetboxClient) GetTenantDetails(name string) (*models.Tenant, error) {
	request := tenancy.NewTenancyTenantsListParams().WithName(&name)
	response, err := r.Tenancy.TenancyTenantsList(request, nil)
	if err != nil {
		return nil, utils.NetboxError("failed to fetch Tenant details", err)
	}
	if len(response.Payload.Results) == 0 {
		return nil, errors.New("tenant not found")
	}

	return &models.Tenant{
		Id:   response.Payload.Results[0].ID,
		Slug: *response.Payload.Results[0].Slug,
		Name: *response.Payload.Results[0].Name,
	}, nil
}