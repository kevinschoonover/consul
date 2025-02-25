// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package fetcher

import (
	"context"
	"fmt"

	"github.com/hashicorp/consul/internal/mesh/internal/controllers/sidecarproxy/cache"
	"github.com/hashicorp/consul/internal/mesh/internal/types"
	"github.com/hashicorp/consul/internal/resource"
	pbcatalog "github.com/hashicorp/consul/proto-public/pbcatalog/v2beta1"
	pbmesh "github.com/hashicorp/consul/proto-public/pbmesh/v2beta1"
	pbmulticluster "github.com/hashicorp/consul/proto-public/pbmulticluster/v2beta1"
	"github.com/hashicorp/consul/proto-public/pbresource"
	"google.golang.org/protobuf/proto"
)

type Fetcher struct {
	client pbresource.ResourceServiceClient
	cache  *cache.Cache
}

func New(client pbresource.ResourceServiceClient, cache *cache.Cache) *Fetcher {
	return &Fetcher{
		client: client,
		cache:  cache,
	}
}

// FetchMeshGateway fetches a service resource from the resource service.
// This will panic if the type field in the ID argument is not a MeshGateway type.
func (f *Fetcher) FetchMeshGateway(ctx context.Context, id *pbresource.ID) (*types.DecodedMeshGateway, error) {
	assertResourceType(pbmesh.MeshGatewayType, id.Type)

	dec, err := resource.GetDecodedResource[*pbmesh.MeshGateway](ctx, f.client, id)
	if err != nil {
		return nil, err
	} else if dec == nil {
		return nil, nil
	}

	return dec, nil
}

// FetchProxyStateTemplate fetches a service resource from the resource service.
// This will panic if the type field in the ID argument is not a ProxyStateTemplate type.
func (f *Fetcher) FetchProxyStateTemplate(ctx context.Context, id *pbresource.ID) (*types.DecodedProxyStateTemplate, error) {
	assertResourceType(pbmesh.ProxyStateTemplateType, id.Type)

	dec, err := resource.GetDecodedResource[*pbmesh.ProxyStateTemplate](ctx, f.client, id)
	if err != nil {
		return nil, err
	} else if dec == nil {
		return nil, nil
	}

	return dec, nil
}

// FetchWorkload fetches a service resource from the resource service.
// This will panic if the type field in the ID argument is not a Workload type.
func (f *Fetcher) FetchWorkload(ctx context.Context, id *pbresource.ID) (*types.DecodedWorkload, error) {
	assertResourceType(pbcatalog.WorkloadType, id.Type)

	dec, err := resource.GetDecodedResource[*pbcatalog.Workload](ctx, f.client, id)
	if err != nil {
		return nil, err
	} else if dec == nil {
		return nil, nil
	}

	return dec, nil
}

// FetchComputedExportedServices fetches a service resource from the resource service.
// This will panic if the type field in the ID argument is not a ComputedExportedServices type.
func (f *Fetcher) FetchComputedExportedServices(ctx context.Context, id *pbresource.ID) (*types.DecodedComputedExportedServices, error) {
	assertResourceType(pbmulticluster.ComputedExportedServicesType, id.Type)

	dec, err := resource.GetDecodedResource[*pbmulticluster.ComputedExportedServices](ctx, f.client, id)
	if err != nil {
		return nil, err
	} else if dec == nil {
		return nil, nil
	}

	return dec, nil
}

// FetchService fetches a service resource from the resource service.
// This will panic if the type field in the ID argument is not a Service type.
func (f *Fetcher) FetchService(ctx context.Context, id *pbresource.ID) (*types.DecodedService, error) {
	assertResourceType(pbcatalog.ServiceType, id.Type)

	dec, err := resource.GetDecodedResource[*pbcatalog.Service](ctx, f.client, id)
	if err != nil {
		return nil, err
	} else if dec == nil {
		return nil, nil
	}

	return dec, nil
}

// this is a helper function to ensure that the resource type we are querying for is the type we expect
func assertResourceType(expected, actual *pbresource.Type) {
	if !proto.Equal(expected, actual) {
		// this is always a programmer error so safe to panic
		panic(fmt.Sprintf("expected a query for a type of %q, you provided a type of %q", expected, actual))
	}
}
