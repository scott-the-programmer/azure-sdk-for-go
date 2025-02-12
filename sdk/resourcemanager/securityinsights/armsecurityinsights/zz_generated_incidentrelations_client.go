//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsights

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

// IncidentRelationsClient contains the methods for the IncidentRelations group.
// Don't use this type directly, use NewIncidentRelationsClient() instead.
type IncidentRelationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewIncidentRelationsClient creates a new instance of IncidentRelationsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewIncidentRelationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *IncidentRelationsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &IncidentRelationsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate - Creates or updates the incident relation.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// incidentID - Incident ID
// relationName - Relation Name
// relation - The relation model
// options - IncidentRelationsClientCreateOrUpdateOptions contains the optional parameters for the IncidentRelationsClient.CreateOrUpdate
// method.
func (client *IncidentRelationsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, relationName string, relation Relation, options *IncidentRelationsClientCreateOrUpdateOptions) (IncidentRelationsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, incidentID, relationName, relation, options)
	if err != nil {
		return IncidentRelationsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IncidentRelationsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return IncidentRelationsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *IncidentRelationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, relationName string, relation Relation, options *IncidentRelationsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}/relations/{relationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if incidentID == "" {
		return nil, errors.New("parameter incidentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentId}", url.PathEscape(incidentID))
	if relationName == "" {
		return nil, errors.New("parameter relationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{relationName}", url.PathEscape(relationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, relation)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *IncidentRelationsClient) createOrUpdateHandleResponse(resp *http.Response) (IncidentRelationsClientCreateOrUpdateResponse, error) {
	result := IncidentRelationsClientCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Relation); err != nil {
		return IncidentRelationsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete the incident relation.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// incidentID - Incident ID
// relationName - Relation Name
// options - IncidentRelationsClientDeleteOptions contains the optional parameters for the IncidentRelationsClient.Delete
// method.
func (client *IncidentRelationsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, relationName string, options *IncidentRelationsClientDeleteOptions) (IncidentRelationsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, incidentID, relationName, options)
	if err != nil {
		return IncidentRelationsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IncidentRelationsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return IncidentRelationsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return IncidentRelationsClientDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IncidentRelationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, relationName string, options *IncidentRelationsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}/relations/{relationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if incidentID == "" {
		return nil, errors.New("parameter incidentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentId}", url.PathEscape(incidentID))
	if relationName == "" {
		return nil, errors.New("parameter relationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{relationName}", url.PathEscape(relationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets an incident relation.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// incidentID - Incident ID
// relationName - Relation Name
// options - IncidentRelationsClientGetOptions contains the optional parameters for the IncidentRelationsClient.Get method.
func (client *IncidentRelationsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, relationName string, options *IncidentRelationsClientGetOptions) (IncidentRelationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, incidentID, relationName, options)
	if err != nil {
		return IncidentRelationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IncidentRelationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IncidentRelationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IncidentRelationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, relationName string, options *IncidentRelationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}/relations/{relationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if incidentID == "" {
		return nil, errors.New("parameter incidentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentId}", url.PathEscape(incidentID))
	if relationName == "" {
		return nil, errors.New("parameter relationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{relationName}", url.PathEscape(relationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IncidentRelationsClient) getHandleResponse(resp *http.Response) (IncidentRelationsClientGetResponse, error) {
	result := IncidentRelationsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Relation); err != nil {
		return IncidentRelationsClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets all incident relations.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// incidentID - Incident ID
// options - IncidentRelationsClientListOptions contains the optional parameters for the IncidentRelationsClient.List method.
func (client *IncidentRelationsClient) List(resourceGroupName string, workspaceName string, incidentID string, options *IncidentRelationsClientListOptions) *IncidentRelationsClientListPager {
	return &IncidentRelationsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, workspaceName, incidentID, options)
		},
		advancer: func(ctx context.Context, resp IncidentRelationsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RelationList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *IncidentRelationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *IncidentRelationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}/relations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if incidentID == "" {
		return nil, errors.New("parameter incidentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentId}", url.PathEscape(incidentID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IncidentRelationsClient) listHandleResponse(resp *http.Response) (IncidentRelationsClientListResponse, error) {
	result := IncidentRelationsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RelationList); err != nil {
		return IncidentRelationsClientListResponse{}, err
	}
	return result, nil
}
