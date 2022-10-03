// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vra

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides information about a catalog source entitlement in vRA.
//
// ## Example Usage
// ### S
//
// This is an example of how to get a vRA catalog source entitlement by its id:
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
//			_, err := vra.LookupCatalogSourceEntitlement(ctx, &GetCatalogSourceEntitlementArgs{
//				Id:        pulumi.StringRef(_var.Catalog_source_entitlement_id),
//				ProjectId: _var.Project_id,
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
// This is an example of how to get a vRA catalog source entitlement by its catalog source id:
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
//			_, err := vra.LookupCatalogSourceEntitlement(ctx, &GetCatalogSourceEntitlementArgs{
//				CatalogSourceId: pulumi.StringRef(_var.Catalog_source_id),
//				ProjectId:       _var.Project_id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupCatalogSourceEntitlement(ctx *pulumi.Context, args *LookupCatalogSourceEntitlementArgs, opts ...pulumi.InvokeOption) (*LookupCatalogSourceEntitlementResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupCatalogSourceEntitlementResult
	err := ctx.Invoke("vra:index/getCatalogSourceEntitlement:getCatalogSourceEntitlement", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCatalogSourceEntitlement.
type LookupCatalogSourceEntitlementArgs struct {
	// The id of the catalog source to find the entitlement. One of `catalogSourceId` or `id` must be provided.
	CatalogSourceId *string `pulumi:"catalogSourceId"`
	// The id of entitlement. One of `catalogSourceId` or `id` must be provided.
	Id *string `pulumi:"id"`
	// The id of the project that this entitlement belongs to.
	ProjectId string `pulumi:"projectId"`
}

// A collection of values returned by getCatalogSourceEntitlement.
type LookupCatalogSourceEntitlementResult struct {
	CatalogSourceId *string `pulumi:"catalogSourceId"`
	// Represents a catalog source that is linked to a project via an entitlement.
	Definitions []GetCatalogSourceEntitlementDefinition `pulumi:"definitions"`
	// Id of the catalog source.
	Id        *string `pulumi:"id"`
	ProjectId string  `pulumi:"projectId"`
}

func LookupCatalogSourceEntitlementOutput(ctx *pulumi.Context, args LookupCatalogSourceEntitlementOutputArgs, opts ...pulumi.InvokeOption) LookupCatalogSourceEntitlementResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCatalogSourceEntitlementResult, error) {
			args := v.(LookupCatalogSourceEntitlementArgs)
			r, err := LookupCatalogSourceEntitlement(ctx, &args, opts...)
			var s LookupCatalogSourceEntitlementResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCatalogSourceEntitlementResultOutput)
}

// A collection of arguments for invoking getCatalogSourceEntitlement.
type LookupCatalogSourceEntitlementOutputArgs struct {
	// The id of the catalog source to find the entitlement. One of `catalogSourceId` or `id` must be provided.
	CatalogSourceId pulumi.StringPtrInput `pulumi:"catalogSourceId"`
	// The id of entitlement. One of `catalogSourceId` or `id` must be provided.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The id of the project that this entitlement belongs to.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
}

func (LookupCatalogSourceEntitlementOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCatalogSourceEntitlementArgs)(nil)).Elem()
}

// A collection of values returned by getCatalogSourceEntitlement.
type LookupCatalogSourceEntitlementResultOutput struct{ *pulumi.OutputState }

func (LookupCatalogSourceEntitlementResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCatalogSourceEntitlementResult)(nil)).Elem()
}

func (o LookupCatalogSourceEntitlementResultOutput) ToLookupCatalogSourceEntitlementResultOutput() LookupCatalogSourceEntitlementResultOutput {
	return o
}

func (o LookupCatalogSourceEntitlementResultOutput) ToLookupCatalogSourceEntitlementResultOutputWithContext(ctx context.Context) LookupCatalogSourceEntitlementResultOutput {
	return o
}

func (o LookupCatalogSourceEntitlementResultOutput) CatalogSourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCatalogSourceEntitlementResult) *string { return v.CatalogSourceId }).(pulumi.StringPtrOutput)
}

// Represents a catalog source that is linked to a project via an entitlement.
func (o LookupCatalogSourceEntitlementResultOutput) Definitions() GetCatalogSourceEntitlementDefinitionArrayOutput {
	return o.ApplyT(func(v LookupCatalogSourceEntitlementResult) []GetCatalogSourceEntitlementDefinition {
		return v.Definitions
	}).(GetCatalogSourceEntitlementDefinitionArrayOutput)
}

// Id of the catalog source.
func (o LookupCatalogSourceEntitlementResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCatalogSourceEntitlementResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupCatalogSourceEntitlementResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCatalogSourceEntitlementResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCatalogSourceEntitlementResultOutput{})
}