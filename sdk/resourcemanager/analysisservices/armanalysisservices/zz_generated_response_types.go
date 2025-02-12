//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armanalysisservices

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsClientListResult contains the result from method OperationsClient.List.
type OperationsClientListResult struct {
	OperationListResult
}

// ServersClientCheckNameAvailabilityResponse contains the response from method ServersClient.CheckNameAvailability.
type ServersClientCheckNameAvailabilityResponse struct {
	ServersClientCheckNameAvailabilityResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersClientCheckNameAvailabilityResult contains the result from method ServersClient.CheckNameAvailability.
type ServersClientCheckNameAvailabilityResult struct {
	CheckServerNameAvailabilityResult
}

// ServersClientCreatePollerResponse contains the response from method ServersClient.Create.
type ServersClientCreatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ServersClientCreatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ServersClientCreatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ServersClientCreateResponse, error) {
	respType := ServersClientCreateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Server)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ServersClientCreatePollerResponse from the provided client and resume token.
func (l *ServersClientCreatePollerResponse) Resume(ctx context.Context, client *ServersClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ServersClient.Create", token, client.pl)
	if err != nil {
		return err
	}
	poller := &ServersClientCreatePoller{
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

// ServersClientCreateResponse contains the response from method ServersClient.Create.
type ServersClientCreateResponse struct {
	ServersClientCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersClientCreateResult contains the result from method ServersClient.Create.
type ServersClientCreateResult struct {
	Server
}

// ServersClientDeletePollerResponse contains the response from method ServersClient.Delete.
type ServersClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ServersClientDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ServersClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ServersClientDeleteResponse, error) {
	respType := ServersClientDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ServersClientDeletePollerResponse from the provided client and resume token.
func (l *ServersClientDeletePollerResponse) Resume(ctx context.Context, client *ServersClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ServersClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &ServersClientDeletePoller{
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

// ServersClientDeleteResponse contains the response from method ServersClient.Delete.
type ServersClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersClientDissociateGatewayResponse contains the response from method ServersClient.DissociateGateway.
type ServersClientDissociateGatewayResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersClientGetDetailsResponse contains the response from method ServersClient.GetDetails.
type ServersClientGetDetailsResponse struct {
	ServersClientGetDetailsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersClientGetDetailsResult contains the result from method ServersClient.GetDetails.
type ServersClientGetDetailsResult struct {
	Server
}

// ServersClientListByResourceGroupResponse contains the response from method ServersClient.ListByResourceGroup.
type ServersClientListByResourceGroupResponse struct {
	ServersClientListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersClientListByResourceGroupResult contains the result from method ServersClient.ListByResourceGroup.
type ServersClientListByResourceGroupResult struct {
	Servers
}

// ServersClientListGatewayStatusResponse contains the response from method ServersClient.ListGatewayStatus.
type ServersClientListGatewayStatusResponse struct {
	ServersClientListGatewayStatusResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersClientListGatewayStatusResult contains the result from method ServersClient.ListGatewayStatus.
type ServersClientListGatewayStatusResult struct {
	GatewayListStatusLive
}

// ServersClientListOperationResultsResponse contains the response from method ServersClient.ListOperationResults.
type ServersClientListOperationResultsResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersClientListOperationStatusesResponse contains the response from method ServersClient.ListOperationStatuses.
type ServersClientListOperationStatusesResponse struct {
	ServersClientListOperationStatusesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersClientListOperationStatusesResult contains the result from method ServersClient.ListOperationStatuses.
type ServersClientListOperationStatusesResult struct {
	OperationStatus
}

// ServersClientListResponse contains the response from method ServersClient.List.
type ServersClientListResponse struct {
	ServersClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersClientListResult contains the result from method ServersClient.List.
type ServersClientListResult struct {
	Servers
}

// ServersClientListSKUsForExistingResponse contains the response from method ServersClient.ListSKUsForExisting.
type ServersClientListSKUsForExistingResponse struct {
	ServersClientListSKUsForExistingResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersClientListSKUsForExistingResult contains the result from method ServersClient.ListSKUsForExisting.
type ServersClientListSKUsForExistingResult struct {
	SKUEnumerationForExistingResourceResult
}

// ServersClientListSKUsForNewResponse contains the response from method ServersClient.ListSKUsForNew.
type ServersClientListSKUsForNewResponse struct {
	ServersClientListSKUsForNewResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersClientListSKUsForNewResult contains the result from method ServersClient.ListSKUsForNew.
type ServersClientListSKUsForNewResult struct {
	SKUEnumerationForNewResourceResult
}

// ServersClientResumePollerResponse contains the response from method ServersClient.Resume.
type ServersClientResumePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ServersClientResumePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ServersClientResumePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ServersClientResumeResponse, error) {
	respType := ServersClientResumeResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ServersClientResumePollerResponse from the provided client and resume token.
func (l *ServersClientResumePollerResponse) Resume(ctx context.Context, client *ServersClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ServersClient.Resume", token, client.pl)
	if err != nil {
		return err
	}
	poller := &ServersClientResumePoller{
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

// ServersClientResumeResponse contains the response from method ServersClient.Resume.
type ServersClientResumeResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersClientSuspendPollerResponse contains the response from method ServersClient.Suspend.
type ServersClientSuspendPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ServersClientSuspendPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ServersClientSuspendPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ServersClientSuspendResponse, error) {
	respType := ServersClientSuspendResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ServersClientSuspendPollerResponse from the provided client and resume token.
func (l *ServersClientSuspendPollerResponse) Resume(ctx context.Context, client *ServersClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ServersClient.Suspend", token, client.pl)
	if err != nil {
		return err
	}
	poller := &ServersClientSuspendPoller{
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

// ServersClientSuspendResponse contains the response from method ServersClient.Suspend.
type ServersClientSuspendResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersClientUpdatePollerResponse contains the response from method ServersClient.Update.
type ServersClientUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ServersClientUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ServersClientUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ServersClientUpdateResponse, error) {
	respType := ServersClientUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Server)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ServersClientUpdatePollerResponse from the provided client and resume token.
func (l *ServersClientUpdatePollerResponse) Resume(ctx context.Context, client *ServersClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ServersClient.Update", token, client.pl)
	if err != nil {
		return err
	}
	poller := &ServersClientUpdatePoller{
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

// ServersClientUpdateResponse contains the response from method ServersClient.Update.
type ServersClientUpdateResponse struct {
	ServersClientUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersClientUpdateResult contains the result from method ServersClient.Update.
type ServersClientUpdateResult struct {
	Server
}
