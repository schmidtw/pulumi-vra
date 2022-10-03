// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vra

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
// ### S
//
// This is an example of how to lookup a region enumeration data source for GCP cloud account.
//
// **Region enumeration data source for GCP**
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vra/sdk/go/vra"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/schmidtw/pulumi-vra/sdk/go/vra"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vra.GetRegionEnumerationGcp(ctx, &GetRegionEnumerationGcpArgs{
//				ClientEmail:  _var.Client_email,
//				PrivateKeyId: _var.Private_key_id,
//				PrivateKey:   _var.Private_key,
//				ProjectId:    _var.Project_id,
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
// The region enumeration data source for GCP cloud account supports the following arguments:
func GetRegionEnumerationGcp(ctx *pulumi.Context, args *GetRegionEnumerationGcpArgs, opts ...pulumi.InvokeOption) (*GetRegionEnumerationGcpResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetRegionEnumerationGcpResult
	err := ctx.Invoke("vra:index/getRegionEnumerationGcp:getRegionEnumerationGcp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRegionEnumerationGcp.
type GetRegionEnumerationGcpArgs struct {
	// Client E-mail ID.
	ClientEmail string `pulumi:"clientEmail"`
	// GCP Private key.
	PrivateKey string `pulumi:"privateKey"`
	// GCP Private key ID.
	PrivateKeyId string `pulumi:"privateKeyId"`
	// GCP Project ID.
	ProjectId string `pulumi:"projectId"`
}

// A collection of values returned by getRegionEnumerationGcp.
type GetRegionEnumerationGcpResult struct {
	ClientEmail string `pulumi:"clientEmail"`
	// The provider-assigned unique ID for this managed resource.
	Id           string `pulumi:"id"`
	PrivateKey   string `pulumi:"privateKey"`
	PrivateKeyId string `pulumi:"privateKeyId"`
	ProjectId    string `pulumi:"projectId"`
	// A set of Region names to enable provisioning on. Example: northamerica-northeast1
	Regions []string `pulumi:"regions"`
}

func GetRegionEnumerationGcpOutput(ctx *pulumi.Context, args GetRegionEnumerationGcpOutputArgs, opts ...pulumi.InvokeOption) GetRegionEnumerationGcpResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRegionEnumerationGcpResult, error) {
			args := v.(GetRegionEnumerationGcpArgs)
			r, err := GetRegionEnumerationGcp(ctx, &args, opts...)
			var s GetRegionEnumerationGcpResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRegionEnumerationGcpResultOutput)
}

// A collection of arguments for invoking getRegionEnumerationGcp.
type GetRegionEnumerationGcpOutputArgs struct {
	// Client E-mail ID.
	ClientEmail pulumi.StringInput `pulumi:"clientEmail"`
	// GCP Private key.
	PrivateKey pulumi.StringInput `pulumi:"privateKey"`
	// GCP Private key ID.
	PrivateKeyId pulumi.StringInput `pulumi:"privateKeyId"`
	// GCP Project ID.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
}

func (GetRegionEnumerationGcpOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRegionEnumerationGcpArgs)(nil)).Elem()
}

// A collection of values returned by getRegionEnumerationGcp.
type GetRegionEnumerationGcpResultOutput struct{ *pulumi.OutputState }

func (GetRegionEnumerationGcpResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRegionEnumerationGcpResult)(nil)).Elem()
}

func (o GetRegionEnumerationGcpResultOutput) ToGetRegionEnumerationGcpResultOutput() GetRegionEnumerationGcpResultOutput {
	return o
}

func (o GetRegionEnumerationGcpResultOutput) ToGetRegionEnumerationGcpResultOutputWithContext(ctx context.Context) GetRegionEnumerationGcpResultOutput {
	return o
}

func (o GetRegionEnumerationGcpResultOutput) ClientEmail() pulumi.StringOutput {
	return o.ApplyT(func(v GetRegionEnumerationGcpResult) string { return v.ClientEmail }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRegionEnumerationGcpResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRegionEnumerationGcpResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRegionEnumerationGcpResultOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetRegionEnumerationGcpResult) string { return v.PrivateKey }).(pulumi.StringOutput)
}

func (o GetRegionEnumerationGcpResultOutput) PrivateKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v GetRegionEnumerationGcpResult) string { return v.PrivateKeyId }).(pulumi.StringOutput)
}

func (o GetRegionEnumerationGcpResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetRegionEnumerationGcpResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// A set of Region names to enable provisioning on. Example: northamerica-northeast1
func (o GetRegionEnumerationGcpResultOutput) Regions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRegionEnumerationGcpResult) []string { return v.Regions }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRegionEnumerationGcpResultOutput{})
}