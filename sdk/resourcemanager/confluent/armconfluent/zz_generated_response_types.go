//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconfluent

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// MarketplaceAgreementsClientCreateResponse contains the response from method MarketplaceAgreementsClient.Create.
type MarketplaceAgreementsClientCreateResponse struct {
	MarketplaceAgreementsClientCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MarketplaceAgreementsClientCreateResult contains the result from method MarketplaceAgreementsClient.Create.
type MarketplaceAgreementsClientCreateResult struct {
	AgreementResource
}

// MarketplaceAgreementsClientListResponse contains the response from method MarketplaceAgreementsClient.List.
type MarketplaceAgreementsClientListResponse struct {
	MarketplaceAgreementsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MarketplaceAgreementsClientListResult contains the result from method MarketplaceAgreementsClient.List.
type MarketplaceAgreementsClientListResult struct {
	AgreementResourceListResponse
}

// OrganizationClientCreatePollerResponse contains the response from method OrganizationClient.Create.
type OrganizationClientCreatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *OrganizationClientCreatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l OrganizationClientCreatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (OrganizationClientCreateResponse, error) {
	respType := OrganizationClientCreateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.OrganizationResource)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a OrganizationClientCreatePollerResponse from the provided client and resume token.
func (l *OrganizationClientCreatePollerResponse) Resume(ctx context.Context, client *OrganizationClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("OrganizationClient.Create", token, client.pl)
	if err != nil {
		return err
	}
	poller := &OrganizationClientCreatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// OrganizationClientCreateResponse contains the response from method OrganizationClient.Create.
type OrganizationClientCreateResponse struct {
	OrganizationClientCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OrganizationClientCreateResult contains the result from method OrganizationClient.Create.
type OrganizationClientCreateResult struct {
	OrganizationResource
}

// OrganizationClientDeletePollerResponse contains the response from method OrganizationClient.Delete.
type OrganizationClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *OrganizationClientDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l OrganizationClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (OrganizationClientDeleteResponse, error) {
	respType := OrganizationClientDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a OrganizationClientDeletePollerResponse from the provided client and resume token.
func (l *OrganizationClientDeletePollerResponse) Resume(ctx context.Context, client *OrganizationClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("OrganizationClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &OrganizationClientDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// OrganizationClientDeleteResponse contains the response from method OrganizationClient.Delete.
type OrganizationClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OrganizationClientGetResponse contains the response from method OrganizationClient.Get.
type OrganizationClientGetResponse struct {
	OrganizationClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OrganizationClientGetResult contains the result from method OrganizationClient.Get.
type OrganizationClientGetResult struct {
	OrganizationResource
}

// OrganizationClientListByResourceGroupResponse contains the response from method OrganizationClient.ListByResourceGroup.
type OrganizationClientListByResourceGroupResponse struct {
	OrganizationClientListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OrganizationClientListByResourceGroupResult contains the result from method OrganizationClient.ListByResourceGroup.
type OrganizationClientListByResourceGroupResult struct {
	OrganizationResourceListResult
}

// OrganizationClientListBySubscriptionResponse contains the response from method OrganizationClient.ListBySubscription.
type OrganizationClientListBySubscriptionResponse struct {
	OrganizationClientListBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OrganizationClientListBySubscriptionResult contains the result from method OrganizationClient.ListBySubscription.
type OrganizationClientListBySubscriptionResult struct {
	OrganizationResourceListResult
}

// OrganizationClientUpdateResponse contains the response from method OrganizationClient.Update.
type OrganizationClientUpdateResponse struct {
	OrganizationClientUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OrganizationClientUpdateResult contains the result from method OrganizationClient.Update.
type OrganizationClientUpdateResult struct {
	OrganizationResource
}

// OrganizationOperationsClientListResponse contains the response from method OrganizationOperationsClient.List.
type OrganizationOperationsClientListResponse struct {
	OrganizationOperationsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OrganizationOperationsClientListResult contains the result from method OrganizationOperationsClient.List.
type OrganizationOperationsClientListResult struct {
	OperationListResult
}

// ValidationsClientValidateOrganizationResponse contains the response from method ValidationsClient.ValidateOrganization.
type ValidationsClientValidateOrganizationResponse struct {
	ValidationsClientValidateOrganizationResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ValidationsClientValidateOrganizationResult contains the result from method ValidationsClient.ValidateOrganization.
type ValidationsClientValidateOrganizationResult struct {
	OrganizationResource
}
