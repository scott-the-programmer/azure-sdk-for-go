//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatalakestore_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-store/armdatalakestore"
)

// x-ms-original-file: specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/VirtualNetworkRules_ListByAccount.json
func ExampleVirtualNetworkRulesClient_ListByAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatalakestore.NewVirtualNetworkRulesClient("<subscription-id>", cred, nil)
	pager := client.ListByAccount("<resource-group-name>",
		"<account-name>",
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

// x-ms-original-file: specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/VirtualNetworkRules_CreateOrUpdate.json
func ExampleVirtualNetworkRulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatalakestore.NewVirtualNetworkRulesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<virtual-network-rule-name>",
		armdatalakestore.CreateOrUpdateVirtualNetworkRuleParameters{
			Properties: &armdatalakestore.CreateOrUpdateVirtualNetworkRuleProperties{
				SubnetID: to.StringPtr("<subnet-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.VirtualNetworkRulesClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/VirtualNetworkRules_Get.json
func ExampleVirtualNetworkRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatalakestore.NewVirtualNetworkRulesClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<virtual-network-rule-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.VirtualNetworkRulesClientGetResult)
}

// x-ms-original-file: specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/VirtualNetworkRules_Update.json
func ExampleVirtualNetworkRulesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatalakestore.NewVirtualNetworkRulesClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<virtual-network-rule-name>",
		&armdatalakestore.VirtualNetworkRulesClientUpdateOptions{Parameters: &armdatalakestore.UpdateVirtualNetworkRuleParameters{
			Properties: &armdatalakestore.UpdateVirtualNetworkRuleProperties{
				SubnetID: to.StringPtr("<subnet-id>"),
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.VirtualNetworkRulesClientUpdateResult)
}

// x-ms-original-file: specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/VirtualNetworkRules_Delete.json
func ExampleVirtualNetworkRulesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatalakestore.NewVirtualNetworkRulesClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<virtual-network-rule-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
