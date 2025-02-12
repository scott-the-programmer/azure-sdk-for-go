//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// BgpServiceCommunitiesClient contains the methods for the BgpServiceCommunities group.
// Don't use this type directly, use NewBgpServiceCommunitiesClient() instead.
type BgpServiceCommunitiesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewBgpServiceCommunitiesClient creates a new instance of BgpServiceCommunitiesClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewBgpServiceCommunitiesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *BgpServiceCommunitiesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &BgpServiceCommunitiesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// List - Gets all the available bgp service communities.
// If the operation fails it returns an *azcore.ResponseError type.
// options - BgpServiceCommunitiesClientListOptions contains the optional parameters for the BgpServiceCommunitiesClient.List
// method.
func (client *BgpServiceCommunitiesClient) List(options *BgpServiceCommunitiesClientListOptions) *BgpServiceCommunitiesClientListPager {
	return &BgpServiceCommunitiesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp BgpServiceCommunitiesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.BgpServiceCommunityListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *BgpServiceCommunitiesClient) listCreateRequest(ctx context.Context, options *BgpServiceCommunitiesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/bgpServiceCommunities"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BgpServiceCommunitiesClient) listHandleResponse(resp *http.Response) (BgpServiceCommunitiesClientListResponse, error) {
	result := BgpServiceCommunitiesClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BgpServiceCommunityListResult); err != nil {
		return BgpServiceCommunitiesClientListResponse{}, err
	}
	return result, nil
}
