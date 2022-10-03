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
// This is an example of how to create a storage profile aws resource.
//
// **Storage profile aws data source by its id:**
//
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
//			_, err := vra.LookupStorageProfileAws(ctx, &GetStorageProfileAwsArgs{
//				Id: pulumi.StringRef(vra_storage_profile_aws.This.Id),
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
// **Vra storage profile data source filter by external region id:**
//
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
//			_, err := vra.LookupStorageProfileAws(ctx, &GetStorageProfileAwsArgs{
//				Filter: pulumi.StringRef("externalRegionId eq 'foobar'"),
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
// A storage profile aws data source supports the following arguments:
func LookupStorageProfileAws(ctx *pulumi.Context, args *LookupStorageProfileAwsArgs, opts ...pulumi.InvokeOption) (*LookupStorageProfileAwsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupStorageProfileAwsResult
	err := ctx.Invoke("vra:index/getStorageProfileAws:getStorageProfileAws", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStorageProfileAws.
type LookupStorageProfileAwsArgs struct {
	// Filter query string that is supported by vRA multi-cloud IaaS API. Example: regionId eq '<regionId>' and cloudAccountId eq '<cloudAccountId>'.
	Filter *string `pulumi:"filter"`
	// The id of the image profile instance.
	Id *string `pulumi:"id"`
	// A set of tag keys and optional values that were set on this Network Profile.
	// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
	Tags []GetStorageProfileAwsTag `pulumi:"tags"`
}

// A collection of values returned by getStorageProfileAws.
type LookupStorageProfileAwsResult struct {
	CloudAccountId string `pulumi:"cloudAccountId"`
	// Date when the entity was created. The date is in ISO 6801 and UTC.
	CreatedAt string `pulumi:"createdAt"`
	// Indicates if this storage profile is a default profile.
	DefaultItem bool   `pulumi:"defaultItem"`
	Description string `pulumi:"description"`
	DeviceType  string `pulumi:"deviceType"`
	// The id of the region as seen in the cloud provider for which this profile is defined.
	ExternalRegionId string  `pulumi:"externalRegionId"`
	Filter           *string `pulumi:"filter"`
	Id               string  `pulumi:"id"`
	Iops             string  `pulumi:"iops"`
	// HATEOAS of the entity
	Links []GetStorageProfileAwsLink `pulumi:"links"`
	// A human-friendly name used as an identifier in APIs that support this option.
	Name  string `pulumi:"name"`
	OrgId string `pulumi:"orgId"`
	// Email of the user that owns the entity.
	Owner              string                    `pulumi:"owner"`
	SupportsEncryption bool                      `pulumi:"supportsEncryption"`
	Tags               []GetStorageProfileAwsTag `pulumi:"tags"`
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt  string `pulumi:"updatedAt"`
	VolumeType string `pulumi:"volumeType"`
}

func LookupStorageProfileAwsOutput(ctx *pulumi.Context, args LookupStorageProfileAwsOutputArgs, opts ...pulumi.InvokeOption) LookupStorageProfileAwsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStorageProfileAwsResult, error) {
			args := v.(LookupStorageProfileAwsArgs)
			r, err := LookupStorageProfileAws(ctx, &args, opts...)
			var s LookupStorageProfileAwsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStorageProfileAwsResultOutput)
}

// A collection of arguments for invoking getStorageProfileAws.
type LookupStorageProfileAwsOutputArgs struct {
	// Filter query string that is supported by vRA multi-cloud IaaS API. Example: regionId eq '<regionId>' and cloudAccountId eq '<cloudAccountId>'.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// The id of the image profile instance.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// A set of tag keys and optional values that were set on this Network Profile.
	// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
	Tags GetStorageProfileAwsTagArrayInput `pulumi:"tags"`
}

func (LookupStorageProfileAwsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageProfileAwsArgs)(nil)).Elem()
}

// A collection of values returned by getStorageProfileAws.
type LookupStorageProfileAwsResultOutput struct{ *pulumi.OutputState }

func (LookupStorageProfileAwsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageProfileAwsResult)(nil)).Elem()
}

func (o LookupStorageProfileAwsResultOutput) ToLookupStorageProfileAwsResultOutput() LookupStorageProfileAwsResultOutput {
	return o
}

func (o LookupStorageProfileAwsResultOutput) ToLookupStorageProfileAwsResultOutputWithContext(ctx context.Context) LookupStorageProfileAwsResultOutput {
	return o
}

func (o LookupStorageProfileAwsResultOutput) CloudAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageProfileAwsResult) string { return v.CloudAccountId }).(pulumi.StringOutput)
}

// Date when the entity was created. The date is in ISO 6801 and UTC.
func (o LookupStorageProfileAwsResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageProfileAwsResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// Indicates if this storage profile is a default profile.
func (o LookupStorageProfileAwsResultOutput) DefaultItem() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupStorageProfileAwsResult) bool { return v.DefaultItem }).(pulumi.BoolOutput)
}

func (o LookupStorageProfileAwsResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageProfileAwsResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupStorageProfileAwsResultOutput) DeviceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageProfileAwsResult) string { return v.DeviceType }).(pulumi.StringOutput)
}

// The id of the region as seen in the cloud provider for which this profile is defined.
func (o LookupStorageProfileAwsResultOutput) ExternalRegionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageProfileAwsResult) string { return v.ExternalRegionId }).(pulumi.StringOutput)
}

func (o LookupStorageProfileAwsResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageProfileAwsResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

func (o LookupStorageProfileAwsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageProfileAwsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStorageProfileAwsResultOutput) Iops() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageProfileAwsResult) string { return v.Iops }).(pulumi.StringOutput)
}

// HATEOAS of the entity
func (o LookupStorageProfileAwsResultOutput) Links() GetStorageProfileAwsLinkArrayOutput {
	return o.ApplyT(func(v LookupStorageProfileAwsResult) []GetStorageProfileAwsLink { return v.Links }).(GetStorageProfileAwsLinkArrayOutput)
}

// A human-friendly name used as an identifier in APIs that support this option.
func (o LookupStorageProfileAwsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageProfileAwsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupStorageProfileAwsResultOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageProfileAwsResult) string { return v.OrgId }).(pulumi.StringOutput)
}

// Email of the user that owns the entity.
func (o LookupStorageProfileAwsResultOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageProfileAwsResult) string { return v.Owner }).(pulumi.StringOutput)
}

func (o LookupStorageProfileAwsResultOutput) SupportsEncryption() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupStorageProfileAwsResult) bool { return v.SupportsEncryption }).(pulumi.BoolOutput)
}

func (o LookupStorageProfileAwsResultOutput) Tags() GetStorageProfileAwsTagArrayOutput {
	return o.ApplyT(func(v LookupStorageProfileAwsResult) []GetStorageProfileAwsTag { return v.Tags }).(GetStorageProfileAwsTagArrayOutput)
}

// Date when the entity was last updated. The date is ISO 8601 and UTC.
func (o LookupStorageProfileAwsResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageProfileAwsResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o LookupStorageProfileAwsResultOutput) VolumeType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageProfileAwsResult) string { return v.VolumeType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStorageProfileAwsResultOutput{})
}