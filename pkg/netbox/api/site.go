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
	"github.com/netbox-community/go-netbox/v3/netbox/client/dcim"

	"github.com/netbox-community/netbox-operator/pkg/netbox/models"
	"github.com/netbox-community/netbox-operator/pkg/netbox/utils"
)

func (r *NetboxClient) GetSiteDetails(name string) (*models.Site, error) {
	request := dcim.NewDcimSitesListParams().WithName(&name)
	response, err := r.Dcim.DcimSitesList(request, nil)
	if err != nil {
		return nil, utils.NetboxError("failed to fetch Site details", err)
	}
	if len(response.Payload.Results) == 0 {
		return nil, utils.NetboxNotFoundError("site '" + name + "'")
	}

	return &models.Site{
		Id:   response.Payload.Results[0].ID,
		Slug: *response.Payload.Results[0].Slug,
		Name: *response.Payload.Results[0].Name,
	}, nil
}