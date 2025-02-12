//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbatch

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
	"strconv"
	"strings"
)

// PrivateLinkResourceClient contains the methods for the PrivateLinkResource group.
// Don't use this type directly, use NewPrivateLinkResourceClient() instead.
type PrivateLinkResourceClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPrivateLinkResourceClient creates a new instance of PrivateLinkResourceClient with the specified values.
// subscriptionID - The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000)
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPrivateLinkResourceClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *PrivateLinkResourceClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &PrivateLinkResourceClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Get - Gets information about the specified private link resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the Batch account.
// accountName - The name of the Batch account.
// privateLinkResourceName - The private link resource name. This must be unique within the account.
// options - PrivateLinkResourceClientGetOptions contains the optional parameters for the PrivateLinkResourceClient.Get method.
func (client *PrivateLinkResourceClient) Get(ctx context.Context, resourceGroupName string, accountName string, privateLinkResourceName string, options *PrivateLinkResourceClientGetOptions) (PrivateLinkResourceClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, privateLinkResourceName, options)
	if err != nil {
		return PrivateLinkResourceClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateLinkResourceClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateLinkResourceClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateLinkResourceClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, privateLinkResourceName string, options *PrivateLinkResourceClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/privateLinkResources/{privateLinkResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if privateLinkResourceName == "" {
		return nil, errors.New("parameter privateLinkResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateLinkResourceName}", url.PathEscape(privateLinkResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateLinkResourceClient) getHandleResponse(resp *http.Response) (PrivateLinkResourceClientGetResponse, error) {
	result := PrivateLinkResourceClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResource); err != nil {
		return PrivateLinkResourceClientGetResponse{}, err
	}
	return result, nil
}

// ListByBatchAccount - Lists all of the private link resources in the specified account.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the Batch account.
// accountName - The name of the Batch account.
// options - PrivateLinkResourceClientListByBatchAccountOptions contains the optional parameters for the PrivateLinkResourceClient.ListByBatchAccount
// method.
func (client *PrivateLinkResourceClient) ListByBatchAccount(resourceGroupName string, accountName string, options *PrivateLinkResourceClientListByBatchAccountOptions) *PrivateLinkResourceClientListByBatchAccountPager {
	return &PrivateLinkResourceClientListByBatchAccountPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByBatchAccountCreateRequest(ctx, resourceGroupName, accountName, options)
		},
		advancer: func(ctx context.Context, resp PrivateLinkResourceClientListByBatchAccountResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListPrivateLinkResourcesResult.NextLink)
		},
	}
}

// listByBatchAccountCreateRequest creates the ListByBatchAccount request.
func (client *PrivateLinkResourceClient) listByBatchAccountCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *PrivateLinkResourceClientListByBatchAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/privateLinkResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	if options != nil && options.Maxresults != nil {
		reqQP.Set("maxresults", strconv.FormatInt(int64(*options.Maxresults), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByBatchAccountHandleResponse handles the ListByBatchAccount response.
func (client *PrivateLinkResourceClient) listByBatchAccountHandleResponse(resp *http.Response) (PrivateLinkResourceClientListByBatchAccountResponse, error) {
	result := PrivateLinkResourceClientListByBatchAccountResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListPrivateLinkResourcesResult); err != nil {
		return PrivateLinkResourceClientListByBatchAccountResponse{}, err
	}
	return result, nil
}
