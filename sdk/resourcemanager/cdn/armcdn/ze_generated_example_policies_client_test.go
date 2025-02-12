//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcdn_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
)

// x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/WafListPolicies.json
func ExamplePoliciesClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcdn.NewPoliciesClient("<subscription-id>", cred, nil)
	pager := client.List("<resource-group-name>",
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

// x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/WafPolicyGet.json
func ExamplePoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcdn.NewPoliciesClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<policy-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PoliciesClientGetResult)
}

// x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/WafPolicyCreateOrUpdate.json
func ExamplePoliciesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcdn.NewPoliciesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<policy-name>",
		armcdn.WebApplicationFirewallPolicy{
			Location: to.StringPtr("<location>"),
			Properties: &armcdn.WebApplicationFirewallPolicyProperties{
				CustomRules: &armcdn.CustomRuleList{
					Rules: []*armcdn.CustomRule{
						{
							Name:         to.StringPtr("<name>"),
							Action:       armcdn.ActionType("Block").ToPtr(),
							EnabledState: armcdn.CustomRuleEnabledState("Enabled").ToPtr(),
							MatchConditions: []*armcdn.MatchCondition{
								{
									MatchValue: []*string{
										to.StringPtr("CH")},
									MatchVariable:   armcdn.WafMatchVariable("RemoteAddr").ToPtr(),
									NegateCondition: to.BoolPtr(false),
									Operator:        armcdn.Operator("GeoMatch").ToPtr(),
									Transforms:      []*armcdn.TransformType{},
								},
								{
									MatchValue: []*string{
										to.StringPtr("windows")},
									MatchVariable:   armcdn.WafMatchVariable("RequestHeader").ToPtr(),
									NegateCondition: to.BoolPtr(false),
									Operator:        armcdn.Operator("Contains").ToPtr(),
									Selector:        to.StringPtr("<selector>"),
									Transforms:      []*armcdn.TransformType{},
								},
								{
									MatchValue: []*string{
										to.StringPtr("<?php"),
										to.StringPtr("?>")},
									MatchVariable:   armcdn.WafMatchVariable("QueryString").ToPtr(),
									NegateCondition: to.BoolPtr(false),
									Operator:        armcdn.Operator("Contains").ToPtr(),
									Selector:        to.StringPtr("<selector>"),
									Transforms: []*armcdn.TransformType{
										armcdn.TransformType("UrlDecode").ToPtr(),
										armcdn.TransformType("Lowercase").ToPtr()},
								}},
							Priority: to.Int32Ptr(2),
						}},
				},
				ManagedRules: &armcdn.ManagedRuleSetList{
					ManagedRuleSets: []*armcdn.ManagedRuleSet{
						{
							RuleGroupOverrides: []*armcdn.ManagedRuleGroupOverride{
								{
									RuleGroupName: to.StringPtr("<rule-group-name>"),
									Rules: []*armcdn.ManagedRuleOverride{
										{
											Action:       armcdn.ActionType("Redirect").ToPtr(),
											EnabledState: armcdn.ManagedRuleEnabledState("Enabled").ToPtr(),
											RuleID:       to.StringPtr("<rule-id>"),
										},
										{
											EnabledState: armcdn.ManagedRuleEnabledState("Disabled").ToPtr(),
											RuleID:       to.StringPtr("<rule-id>"),
										}},
								}},
							RuleSetType:    to.StringPtr("<rule-set-type>"),
							RuleSetVersion: to.StringPtr("<rule-set-version>"),
						}},
				},
				PolicySettings: &armcdn.PolicySettings{
					DefaultCustomBlockResponseBody:       to.StringPtr("<default-custom-block-response-body>"),
					DefaultCustomBlockResponseStatusCode: armcdn.PolicySettingsDefaultCustomBlockResponseStatusCode(200).ToPtr(),
					DefaultRedirectURL:                   to.StringPtr("<default-redirect-url>"),
				},
				RateLimitRules: &armcdn.RateLimitRuleList{
					Rules: []*armcdn.RateLimitRule{
						{
							Name:         to.StringPtr("<name>"),
							Action:       armcdn.ActionType("Block").ToPtr(),
							EnabledState: armcdn.CustomRuleEnabledState("Enabled").ToPtr(),
							MatchConditions: []*armcdn.MatchCondition{
								{
									MatchValue: []*string{
										to.StringPtr("192.168.1.0/24"),
										to.StringPtr("10.0.0.0/24")},
									MatchVariable:   armcdn.WafMatchVariable("RemoteAddr").ToPtr(),
									NegateCondition: to.BoolPtr(false),
									Operator:        armcdn.Operator("IPMatch").ToPtr(),
									Transforms:      []*armcdn.TransformType{},
								}},
							Priority:                   to.Int32Ptr(1),
							RateLimitDurationInMinutes: to.Int32Ptr(0),
							RateLimitThreshold:         to.Int32Ptr(1000),
						}},
				},
			},
			SKU: &armcdn.SKU{
				Name: armcdn.SKUName("Standard_Microsoft").ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PoliciesClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/WafPatchPolicy.json
func ExamplePoliciesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcdn.NewPoliciesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<policy-name>",
		armcdn.WebApplicationFirewallPolicyPatchParameters{
			Tags: map[string]*string{
				"foo": to.StringPtr("bar"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PoliciesClientUpdateResult)
}

// x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/WafPolicyDelete.json
func ExamplePoliciesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcdn.NewPoliciesClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<policy-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
