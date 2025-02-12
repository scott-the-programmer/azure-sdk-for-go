//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsqlvirtualmachine

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

// SQLVirtualMachinesClient contains the methods for the SQLVirtualMachines group.
// Don't use this type directly, use NewSQLVirtualMachinesClient() instead.
type SQLVirtualMachinesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSQLVirtualMachinesClient creates a new instance of SQLVirtualMachinesClient with the specified values.
// subscriptionID - Subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSQLVirtualMachinesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SQLVirtualMachinesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &SQLVirtualMachinesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates a SQL virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
// Manager API or the portal.
// sqlVirtualMachineName - Name of the SQL virtual machine.
// parameters - The SQL virtual machine.
// options - SQLVirtualMachinesClientBeginCreateOrUpdateOptions contains the optional parameters for the SQLVirtualMachinesClient.BeginCreateOrUpdate
// method.
func (client *SQLVirtualMachinesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, parameters SQLVirtualMachine, options *SQLVirtualMachinesClientBeginCreateOrUpdateOptions) (SQLVirtualMachinesClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, sqlVirtualMachineName, parameters, options)
	if err != nil {
		return SQLVirtualMachinesClientCreateOrUpdatePollerResponse{}, err
	}
	result := SQLVirtualMachinesClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SQLVirtualMachinesClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return SQLVirtualMachinesClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &SQLVirtualMachinesClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a SQL virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SQLVirtualMachinesClient) createOrUpdate(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, parameters SQLVirtualMachine, options *SQLVirtualMachinesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, sqlVirtualMachineName, parameters, options)
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
func (client *SQLVirtualMachinesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, parameters SQLVirtualMachine, options *SQLVirtualMachinesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/{sqlVirtualMachineName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlVirtualMachineName == "" {
		return nil, errors.New("parameter sqlVirtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlVirtualMachineName}", url.PathEscape(sqlVirtualMachineName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes a SQL virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
// Manager API or the portal.
// sqlVirtualMachineName - Name of the SQL virtual machine.
// options - SQLVirtualMachinesClientBeginDeleteOptions contains the optional parameters for the SQLVirtualMachinesClient.BeginDelete
// method.
func (client *SQLVirtualMachinesClient) BeginDelete(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, options *SQLVirtualMachinesClientBeginDeleteOptions) (SQLVirtualMachinesClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, sqlVirtualMachineName, options)
	if err != nil {
		return SQLVirtualMachinesClientDeletePollerResponse{}, err
	}
	result := SQLVirtualMachinesClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SQLVirtualMachinesClient.Delete", "", resp, client.pl)
	if err != nil {
		return SQLVirtualMachinesClientDeletePollerResponse{}, err
	}
	result.Poller = &SQLVirtualMachinesClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a SQL virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SQLVirtualMachinesClient) deleteOperation(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, options *SQLVirtualMachinesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, sqlVirtualMachineName, options)
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
func (client *SQLVirtualMachinesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, options *SQLVirtualMachinesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/{sqlVirtualMachineName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlVirtualMachineName == "" {
		return nil, errors.New("parameter sqlVirtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlVirtualMachineName}", url.PathEscape(sqlVirtualMachineName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets a SQL virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
// Manager API or the portal.
// sqlVirtualMachineName - Name of the SQL virtual machine.
// options - SQLVirtualMachinesClientGetOptions contains the optional parameters for the SQLVirtualMachinesClient.Get method.
func (client *SQLVirtualMachinesClient) Get(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, options *SQLVirtualMachinesClientGetOptions) (SQLVirtualMachinesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, sqlVirtualMachineName, options)
	if err != nil {
		return SQLVirtualMachinesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLVirtualMachinesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SQLVirtualMachinesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SQLVirtualMachinesClient) getCreateRequest(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, options *SQLVirtualMachinesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/{sqlVirtualMachineName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlVirtualMachineName == "" {
		return nil, errors.New("parameter sqlVirtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlVirtualMachineName}", url.PathEscape(sqlVirtualMachineName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2017-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SQLVirtualMachinesClient) getHandleResponse(resp *http.Response) (SQLVirtualMachinesClientGetResponse, error) {
	result := SQLVirtualMachinesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLVirtualMachine); err != nil {
		return SQLVirtualMachinesClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets all SQL virtual machines in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - SQLVirtualMachinesClientListOptions contains the optional parameters for the SQLVirtualMachinesClient.List method.
func (client *SQLVirtualMachinesClient) List(options *SQLVirtualMachinesClientListOptions) *SQLVirtualMachinesClientListPager {
	return &SQLVirtualMachinesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp SQLVirtualMachinesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *SQLVirtualMachinesClient) listCreateRequest(ctx context.Context, options *SQLVirtualMachinesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SQLVirtualMachinesClient) listHandleResponse(resp *http.Response) (SQLVirtualMachinesClientListResponse, error) {
	result := SQLVirtualMachinesClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResult); err != nil {
		return SQLVirtualMachinesClientListResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Gets all SQL virtual machines in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
// Manager API or the portal.
// options - SQLVirtualMachinesClientListByResourceGroupOptions contains the optional parameters for the SQLVirtualMachinesClient.ListByResourceGroup
// method.
func (client *SQLVirtualMachinesClient) ListByResourceGroup(resourceGroupName string, options *SQLVirtualMachinesClientListByResourceGroupOptions) *SQLVirtualMachinesClientListByResourceGroupPager {
	return &SQLVirtualMachinesClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp SQLVirtualMachinesClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *SQLVirtualMachinesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *SQLVirtualMachinesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *SQLVirtualMachinesClient) listByResourceGroupHandleResponse(resp *http.Response) (SQLVirtualMachinesClientListByResourceGroupResponse, error) {
	result := SQLVirtualMachinesClientListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResult); err != nil {
		return SQLVirtualMachinesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListBySQLVMGroup - Gets the list of sql virtual machines in a SQL virtual machine group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
// Manager API or the portal.
// sqlVirtualMachineGroupName - Name of the SQL virtual machine group.
// options - SQLVirtualMachinesClientListBySQLVMGroupOptions contains the optional parameters for the SQLVirtualMachinesClient.ListBySQLVMGroup
// method.
func (client *SQLVirtualMachinesClient) ListBySQLVMGroup(resourceGroupName string, sqlVirtualMachineGroupName string, options *SQLVirtualMachinesClientListBySQLVMGroupOptions) *SQLVirtualMachinesClientListBySQLVMGroupPager {
	return &SQLVirtualMachinesClientListBySQLVMGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySQLVMGroupCreateRequest(ctx, resourceGroupName, sqlVirtualMachineGroupName, options)
		},
		advancer: func(ctx context.Context, resp SQLVirtualMachinesClientListBySQLVMGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListResult.NextLink)
		},
	}
}

// listBySQLVMGroupCreateRequest creates the ListBySQLVMGroup request.
func (client *SQLVirtualMachinesClient) listBySQLVMGroupCreateRequest(ctx context.Context, resourceGroupName string, sqlVirtualMachineGroupName string, options *SQLVirtualMachinesClientListBySQLVMGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/{sqlVirtualMachineGroupName}/sqlVirtualMachines"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlVirtualMachineGroupName == "" {
		return nil, errors.New("parameter sqlVirtualMachineGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlVirtualMachineGroupName}", url.PathEscape(sqlVirtualMachineGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySQLVMGroupHandleResponse handles the ListBySQLVMGroup response.
func (client *SQLVirtualMachinesClient) listBySQLVMGroupHandleResponse(resp *http.Response) (SQLVirtualMachinesClientListBySQLVMGroupResponse, error) {
	result := SQLVirtualMachinesClientListBySQLVMGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResult); err != nil {
		return SQLVirtualMachinesClientListBySQLVMGroupResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates a SQL virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
// Manager API or the portal.
// sqlVirtualMachineName - Name of the SQL virtual machine.
// parameters - The SQL virtual machine.
// options - SQLVirtualMachinesClientBeginUpdateOptions contains the optional parameters for the SQLVirtualMachinesClient.BeginUpdate
// method.
func (client *SQLVirtualMachinesClient) BeginUpdate(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, parameters Update, options *SQLVirtualMachinesClientBeginUpdateOptions) (SQLVirtualMachinesClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, sqlVirtualMachineName, parameters, options)
	if err != nil {
		return SQLVirtualMachinesClientUpdatePollerResponse{}, err
	}
	result := SQLVirtualMachinesClientUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SQLVirtualMachinesClient.Update", "", resp, client.pl)
	if err != nil {
		return SQLVirtualMachinesClientUpdatePollerResponse{}, err
	}
	result.Poller = &SQLVirtualMachinesClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates a SQL virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SQLVirtualMachinesClient) update(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, parameters Update, options *SQLVirtualMachinesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, sqlVirtualMachineName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *SQLVirtualMachinesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, parameters Update, options *SQLVirtualMachinesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/{sqlVirtualMachineName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlVirtualMachineName == "" {
		return nil, errors.New("parameter sqlVirtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlVirtualMachineName}", url.PathEscape(sqlVirtualMachineName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}
