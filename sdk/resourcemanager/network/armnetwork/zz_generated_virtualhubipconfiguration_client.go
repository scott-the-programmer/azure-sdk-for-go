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

// VirtualHubIPConfigurationClient contains the methods for the VirtualHubIPConfiguration group.
// Don't use this type directly, use NewVirtualHubIPConfigurationClient() instead.
type VirtualHubIPConfigurationClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVirtualHubIPConfigurationClient creates a new instance of VirtualHubIPConfigurationClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVirtualHubIPConfigurationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *VirtualHubIPConfigurationClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &VirtualHubIPConfigurationClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Creates a VirtualHubIpConfiguration resource if it doesn't exist else updates the existing VirtualHubIpConfiguration.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the VirtualHub.
// virtualHubName - The name of the VirtualHub.
// ipConfigName - The name of the ipconfig.
// parameters - Hub Ip Configuration parameters.
// options - VirtualHubIPConfigurationClientBeginCreateOrUpdateOptions contains the optional parameters for the VirtualHubIPConfigurationClient.BeginCreateOrUpdate
// method.
func (client *VirtualHubIPConfigurationClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, ipConfigName string, parameters HubIPConfiguration, options *VirtualHubIPConfigurationClientBeginCreateOrUpdateOptions) (VirtualHubIPConfigurationClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, virtualHubName, ipConfigName, parameters, options)
	if err != nil {
		return VirtualHubIPConfigurationClientCreateOrUpdatePollerResponse{}, err
	}
	result := VirtualHubIPConfigurationClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualHubIPConfigurationClient.CreateOrUpdate", "azure-async-operation", resp, client.pl)
	if err != nil {
		return VirtualHubIPConfigurationClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &VirtualHubIPConfigurationClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates a VirtualHubIpConfiguration resource if it doesn't exist else updates the existing VirtualHubIpConfiguration.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualHubIPConfigurationClient) createOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, ipConfigName string, parameters HubIPConfiguration, options *VirtualHubIPConfigurationClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, virtualHubName, ipConfigName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualHubIPConfigurationClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, ipConfigName string, parameters HubIPConfiguration, options *VirtualHubIPConfigurationClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/ipConfigurations/{ipConfigName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	if ipConfigName == "" {
		return nil, errors.New("parameter ipConfigName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipConfigName}", url.PathEscape(ipConfigName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes a VirtualHubIpConfiguration.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the VirtualHubBgpConnection.
// virtualHubName - The name of the VirtualHub.
// ipConfigName - The name of the ipconfig.
// options - VirtualHubIPConfigurationClientBeginDeleteOptions contains the optional parameters for the VirtualHubIPConfigurationClient.BeginDelete
// method.
func (client *VirtualHubIPConfigurationClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualHubName string, ipConfigName string, options *VirtualHubIPConfigurationClientBeginDeleteOptions) (VirtualHubIPConfigurationClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, virtualHubName, ipConfigName, options)
	if err != nil {
		return VirtualHubIPConfigurationClientDeletePollerResponse{}, err
	}
	result := VirtualHubIPConfigurationClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualHubIPConfigurationClient.Delete", "location", resp, client.pl)
	if err != nil {
		return VirtualHubIPConfigurationClientDeletePollerResponse{}, err
	}
	result.Poller = &VirtualHubIPConfigurationClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a VirtualHubIpConfiguration.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualHubIPConfigurationClient) deleteOperation(ctx context.Context, resourceGroupName string, virtualHubName string, ipConfigName string, options *VirtualHubIPConfigurationClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualHubName, ipConfigName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VirtualHubIPConfigurationClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, ipConfigName string, options *VirtualHubIPConfigurationClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/ipConfigurations/{ipConfigName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	if ipConfigName == "" {
		return nil, errors.New("parameter ipConfigName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipConfigName}", url.PathEscape(ipConfigName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Retrieves the details of a Virtual Hub Ip configuration.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the VirtualHub.
// virtualHubName - The name of the VirtualHub.
// ipConfigName - The name of the ipconfig.
// options - VirtualHubIPConfigurationClientGetOptions contains the optional parameters for the VirtualHubIPConfigurationClient.Get
// method.
func (client *VirtualHubIPConfigurationClient) Get(ctx context.Context, resourceGroupName string, virtualHubName string, ipConfigName string, options *VirtualHubIPConfigurationClientGetOptions) (VirtualHubIPConfigurationClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualHubName, ipConfigName, options)
	if err != nil {
		return VirtualHubIPConfigurationClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualHubIPConfigurationClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualHubIPConfigurationClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualHubIPConfigurationClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, ipConfigName string, options *VirtualHubIPConfigurationClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/ipConfigurations/{ipConfigName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	if ipConfigName == "" {
		return nil, errors.New("parameter ipConfigName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipConfigName}", url.PathEscape(ipConfigName))
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

// getHandleResponse handles the Get response.
func (client *VirtualHubIPConfigurationClient) getHandleResponse(resp *http.Response) (VirtualHubIPConfigurationClientGetResponse, error) {
	result := VirtualHubIPConfigurationClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.HubIPConfiguration); err != nil {
		return VirtualHubIPConfigurationClientGetResponse{}, err
	}
	return result, nil
}

// List - Retrieves the details of all VirtualHubIpConfigurations.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the VirtualHub.
// virtualHubName - The name of the VirtualHub.
// options - VirtualHubIPConfigurationClientListOptions contains the optional parameters for the VirtualHubIPConfigurationClient.List
// method.
func (client *VirtualHubIPConfigurationClient) List(resourceGroupName string, virtualHubName string, options *VirtualHubIPConfigurationClientListOptions) *VirtualHubIPConfigurationClientListPager {
	return &VirtualHubIPConfigurationClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, virtualHubName, options)
		},
		advancer: func(ctx context.Context, resp VirtualHubIPConfigurationClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListVirtualHubIPConfigurationResults.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *VirtualHubIPConfigurationClient) listCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, options *VirtualHubIPConfigurationClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/ipConfigurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
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
func (client *VirtualHubIPConfigurationClient) listHandleResponse(resp *http.Response) (VirtualHubIPConfigurationClientListResponse, error) {
	result := VirtualHubIPConfigurationClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListVirtualHubIPConfigurationResults); err != nil {
		return VirtualHubIPConfigurationClientListResponse{}, err
	}
	return result, nil
}
