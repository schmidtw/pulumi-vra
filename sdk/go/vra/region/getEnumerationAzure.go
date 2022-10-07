// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package region

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
// ### S
//
// This is an example of how to lookup a region enumeration data source for Azure cloud account.
//
// **Region enumeration data source for Azure**
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vra/sdk/go/vra/region"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-vra/sdk/go/vra/region"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := region.GetEnumerationAzure(ctx, &region.GetEnumerationAzureArgs{
//				ApplicationId:  _var.Application_id,
//				ApplicationKey: _var.Application_key,
//				SubscriptionId: _var.Subscription_id,
//				TenantId:       _var.Tenant_id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// The region enumeration data source for Azure cloud account supports the following arguments:
func GetEnumerationAzure(ctx *pulumi.Context, args *GetEnumerationAzureArgs, opts ...pulumi.InvokeOption) (*GetEnumerationAzureResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetEnumerationAzureResult
	err := ctx.Invoke("vra:region/getEnumerationAzure:getEnumerationAzure", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEnumerationAzure.
type GetEnumerationAzureArgs struct {
	// Azure Client Application ID
	ApplicationId string `pulumi:"applicationId"`
	// Azure Client Application Secret Key
	ApplicationKey string `pulumi:"applicationKey"`
	// Azure Subscription ID
	SubscriptionId string `pulumi:"subscriptionId"`
	// Azure Tenant ID
	TenantId string `pulumi:"tenantId"`
}

// A collection of values returned by getEnumerationAzure.
type GetEnumerationAzureResult struct {
	ApplicationId  string `pulumi:"applicationId"`
	ApplicationKey string `pulumi:"applicationKey"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A set of Region names to enable provisioning on. Example: northamerica-northeast1
	Regions        []string `pulumi:"regions"`
	SubscriptionId string   `pulumi:"subscriptionId"`
	TenantId       string   `pulumi:"tenantId"`
}

func GetEnumerationAzureOutput(ctx *pulumi.Context, args GetEnumerationAzureOutputArgs, opts ...pulumi.InvokeOption) GetEnumerationAzureResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetEnumerationAzureResult, error) {
			args := v.(GetEnumerationAzureArgs)
			r, err := GetEnumerationAzure(ctx, &args, opts...)
			var s GetEnumerationAzureResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetEnumerationAzureResultOutput)
}

// A collection of arguments for invoking getEnumerationAzure.
type GetEnumerationAzureOutputArgs struct {
	// Azure Client Application ID
	ApplicationId pulumi.StringInput `pulumi:"applicationId"`
	// Azure Client Application Secret Key
	ApplicationKey pulumi.StringInput `pulumi:"applicationKey"`
	// Azure Subscription ID
	SubscriptionId pulumi.StringInput `pulumi:"subscriptionId"`
	// Azure Tenant ID
	TenantId pulumi.StringInput `pulumi:"tenantId"`
}

func (GetEnumerationAzureOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEnumerationAzureArgs)(nil)).Elem()
}

// A collection of values returned by getEnumerationAzure.
type GetEnumerationAzureResultOutput struct{ *pulumi.OutputState }

func (GetEnumerationAzureResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEnumerationAzureResult)(nil)).Elem()
}

func (o GetEnumerationAzureResultOutput) ToGetEnumerationAzureResultOutput() GetEnumerationAzureResultOutput {
	return o
}

func (o GetEnumerationAzureResultOutput) ToGetEnumerationAzureResultOutputWithContext(ctx context.Context) GetEnumerationAzureResultOutput {
	return o
}

func (o GetEnumerationAzureResultOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v GetEnumerationAzureResult) string { return v.ApplicationId }).(pulumi.StringOutput)
}

func (o GetEnumerationAzureResultOutput) ApplicationKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetEnumerationAzureResult) string { return v.ApplicationKey }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetEnumerationAzureResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetEnumerationAzureResult) string { return v.Id }).(pulumi.StringOutput)
}

// A set of Region names to enable provisioning on. Example: northamerica-northeast1
func (o GetEnumerationAzureResultOutput) Regions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEnumerationAzureResult) []string { return v.Regions }).(pulumi.StringArrayOutput)
}

func (o GetEnumerationAzureResultOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v GetEnumerationAzureResult) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o GetEnumerationAzureResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v GetEnumerationAzureResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEnumerationAzureResultOutput{})
}