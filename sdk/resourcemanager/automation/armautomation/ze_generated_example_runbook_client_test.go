//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/publishRunbook.json
func ExampleRunbookClient_BeginPublish() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armautomation.NewRunbookClient("<subscription-id>", cred, nil)
	poller, err := client.BeginPublish(ctx,
		"<resource-group-name>",
		"<automation-account-name>",
		"<runbook-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/getRunbookContent.json
func ExampleRunbookClient_GetContent() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armautomation.NewRunbookClient("<subscription-id>", cred, nil)
	_, err = client.GetContent(ctx,
		"<resource-group-name>",
		"<automation-account-name>",
		"<runbook-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/getRunbook.json
func ExampleRunbookClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armautomation.NewRunbookClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<automation-account-name>",
		"<runbook-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.RunbookClientGetResult)
}

// x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/createOrUpdateRunbook.json
func ExampleRunbookClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armautomation.NewRunbookClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<automation-account-name>",
		"<runbook-name>",
		armautomation.RunbookCreateOrUpdateParameters{
			Name:     to.StringPtr("<name>"),
			Location: to.StringPtr("<location>"),
			Properties: &armautomation.RunbookCreateOrUpdateProperties{
				Description:      to.StringPtr("<description>"),
				LogActivityTrace: to.Int32Ptr(1),
				LogProgress:      to.BoolPtr(true),
				LogVerbose:       to.BoolPtr(false),
				PublishContentLink: &armautomation.ContentLink{
					ContentHash: &armautomation.ContentHash{
						Algorithm: to.StringPtr("<algorithm>"),
						Value:     to.StringPtr("<value>"),
					},
					URI: to.StringPtr("<uri>"),
				},
				RunbookType: armautomation.RunbookTypeEnum("PowerShellWorkflow").ToPtr(),
			},
			Tags: map[string]*string{
				"tag01": to.StringPtr("value01"),
				"tag02": to.StringPtr("value02"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.RunbookClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/updateRunbook.json
func ExampleRunbookClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armautomation.NewRunbookClient("<subscription-id>", cred, nil)
	_, err = client.Update(ctx,
		"<resource-group-name>",
		"<automation-account-name>",
		"<runbook-name>",
		armautomation.RunbookUpdateParameters{
			Properties: &armautomation.RunbookUpdateProperties{
				Description:      to.StringPtr("<description>"),
				LogActivityTrace: to.Int32Ptr(1),
				LogProgress:      to.BoolPtr(true),
				LogVerbose:       to.BoolPtr(false),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/deleteRunbook.json
func ExampleRunbookClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armautomation.NewRunbookClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<automation-account-name>",
		"<runbook-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/listRunbooksByAutomationAccount.json
func ExampleRunbookClient_ListByAutomationAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armautomation.NewRunbookClient("<subscription-id>", cred, nil)
	pager := client.ListByAutomationAccount("<resource-group-name>",
		"<automation-account-name>",
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}
