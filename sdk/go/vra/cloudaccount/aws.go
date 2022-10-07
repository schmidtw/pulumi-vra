// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudaccount

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a VMware vRealize Automation AWS cloud account resource.
//
// ## Example Usage
// ### S
//
// The following example shows how to create an AWS cloud account resource.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vra/sdk/go/vra/cloudaccount"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-vra/sdk/go/vra/cloudaccount"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudaccount.NewAws(ctx, "this", &cloudaccount.AwsArgs{
//				Description: pulumi.String("terraform test cloud account aws"),
//				AccessKey:   pulumi.Any(_var.Access_key),
//				SecretKey:   pulumi.Any(_var.Secret_key),
//				Regions: pulumi.StringArray{
//					pulumi.String("us-east-1"),
//					pulumi.String("us-west-1"),
//				},
//				Tags: cloudaccount.AwsTagArray{
//					&cloudaccount.AwsTagArgs{
//						Key:   pulumi.String("foo"),
//						Value: pulumi.String("bar"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # To import the AWS cloud account, use the ID as in the following example
//
// ```sh
//
//	$ pulumi import vra:cloudaccount/aws:Aws new_aws 05956583-6488-4e7d-84c9-92a7b7219a15`
//
// ```
type Aws struct {
	pulumi.CustomResourceState

	// Access key ID for AWS.
	AccessKey pulumi.StringOutput `pulumi:"accessKey"`
	// Date when entity was created. Date and time format is ISO 8601 and UTC.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Human-friendly description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// HATEOAS of entity.
	Links AwsLinkArrayOutput `pulumi:"links"`
	// Name of AWS cloud account.
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of organization that entity belongs to.
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// Email of entity owner.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// Set of region names enabled for the cloud account.
	Regions pulumi.StringArrayOutput `pulumi:"regions"`
	// AWS Secret Access Key
	SecretKey pulumi.StringOutput `pulumi:"secretKey"`
	// Set of tag keys and values to apply to the cloud account.\
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags AwsTagArrayOutput `pulumi:"tags"`
	// Date when entity was last updated. Date and time format is ISO 8601 and UTC.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewAws registers a new resource with the given unique name, arguments, and options.
func NewAws(ctx *pulumi.Context,
	name string, args *AwsArgs, opts ...pulumi.ResourceOption) (*Aws, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessKey == nil {
		return nil, errors.New("invalid value for required argument 'AccessKey'")
	}
	if args.Regions == nil {
		return nil, errors.New("invalid value for required argument 'Regions'")
	}
	if args.SecretKey == nil {
		return nil, errors.New("invalid value for required argument 'SecretKey'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Aws
	err := ctx.RegisterResource("vra:cloudaccount/aws:Aws", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAws gets an existing Aws resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAws(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AwsState, opts ...pulumi.ResourceOption) (*Aws, error) {
	var resource Aws
	err := ctx.ReadResource("vra:cloudaccount/aws:Aws", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Aws resources.
type awsState struct {
	// Access key ID for AWS.
	AccessKey *string `pulumi:"accessKey"`
	// Date when entity was created. Date and time format is ISO 8601 and UTC.
	CreatedAt *string `pulumi:"createdAt"`
	// Human-friendly description.
	Description *string `pulumi:"description"`
	// HATEOAS of entity.
	Links []AwsLink `pulumi:"links"`
	// Name of AWS cloud account.
	Name *string `pulumi:"name"`
	// ID of organization that entity belongs to.
	OrgId *string `pulumi:"orgId"`
	// Email of entity owner.
	Owner *string `pulumi:"owner"`
	// Set of region names enabled for the cloud account.
	Regions []string `pulumi:"regions"`
	// AWS Secret Access Key
	SecretKey *string `pulumi:"secretKey"`
	// Set of tag keys and values to apply to the cloud account.\
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags []AwsTag `pulumi:"tags"`
	// Date when entity was last updated. Date and time format is ISO 8601 and UTC.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type AwsState struct {
	// Access key ID for AWS.
	AccessKey pulumi.StringPtrInput
	// Date when entity was created. Date and time format is ISO 8601 and UTC.
	CreatedAt pulumi.StringPtrInput
	// Human-friendly description.
	Description pulumi.StringPtrInput
	// HATEOAS of entity.
	Links AwsLinkArrayInput
	// Name of AWS cloud account.
	Name pulumi.StringPtrInput
	// ID of organization that entity belongs to.
	OrgId pulumi.StringPtrInput
	// Email of entity owner.
	Owner pulumi.StringPtrInput
	// Set of region names enabled for the cloud account.
	Regions pulumi.StringArrayInput
	// AWS Secret Access Key
	SecretKey pulumi.StringPtrInput
	// Set of tag keys and values to apply to the cloud account.\
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags AwsTagArrayInput
	// Date when entity was last updated. Date and time format is ISO 8601 and UTC.
	UpdatedAt pulumi.StringPtrInput
}

func (AwsState) ElementType() reflect.Type {
	return reflect.TypeOf((*awsState)(nil)).Elem()
}

type awsArgs struct {
	// Access key ID for AWS.
	AccessKey string `pulumi:"accessKey"`
	// Human-friendly description.
	Description *string `pulumi:"description"`
	// Name of AWS cloud account.
	Name *string `pulumi:"name"`
	// Set of region names enabled for the cloud account.
	Regions []string `pulumi:"regions"`
	// AWS Secret Access Key
	SecretKey string `pulumi:"secretKey"`
	// Set of tag keys and values to apply to the cloud account.\
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags []AwsTag `pulumi:"tags"`
}

// The set of arguments for constructing a Aws resource.
type AwsArgs struct {
	// Access key ID for AWS.
	AccessKey pulumi.StringInput
	// Human-friendly description.
	Description pulumi.StringPtrInput
	// Name of AWS cloud account.
	Name pulumi.StringPtrInput
	// Set of region names enabled for the cloud account.
	Regions pulumi.StringArrayInput
	// AWS Secret Access Key
	SecretKey pulumi.StringInput
	// Set of tag keys and values to apply to the cloud account.\
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags AwsTagArrayInput
}

func (AwsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*awsArgs)(nil)).Elem()
}

type AwsInput interface {
	pulumi.Input

	ToAwsOutput() AwsOutput
	ToAwsOutputWithContext(ctx context.Context) AwsOutput
}

func (*Aws) ElementType() reflect.Type {
	return reflect.TypeOf((**Aws)(nil)).Elem()
}

func (i *Aws) ToAwsOutput() AwsOutput {
	return i.ToAwsOutputWithContext(context.Background())
}

func (i *Aws) ToAwsOutputWithContext(ctx context.Context) AwsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsOutput)
}

// AwsArrayInput is an input type that accepts AwsArray and AwsArrayOutput values.
// You can construct a concrete instance of `AwsArrayInput` via:
//
//	AwsArray{ AwsArgs{...} }
type AwsArrayInput interface {
	pulumi.Input

	ToAwsArrayOutput() AwsArrayOutput
	ToAwsArrayOutputWithContext(context.Context) AwsArrayOutput
}

type AwsArray []AwsInput

func (AwsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Aws)(nil)).Elem()
}

func (i AwsArray) ToAwsArrayOutput() AwsArrayOutput {
	return i.ToAwsArrayOutputWithContext(context.Background())
}

func (i AwsArray) ToAwsArrayOutputWithContext(ctx context.Context) AwsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsArrayOutput)
}

// AwsMapInput is an input type that accepts AwsMap and AwsMapOutput values.
// You can construct a concrete instance of `AwsMapInput` via:
//
//	AwsMap{ "key": AwsArgs{...} }
type AwsMapInput interface {
	pulumi.Input

	ToAwsMapOutput() AwsMapOutput
	ToAwsMapOutputWithContext(context.Context) AwsMapOutput
}

type AwsMap map[string]AwsInput

func (AwsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Aws)(nil)).Elem()
}

func (i AwsMap) ToAwsMapOutput() AwsMapOutput {
	return i.ToAwsMapOutputWithContext(context.Background())
}

func (i AwsMap) ToAwsMapOutputWithContext(ctx context.Context) AwsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsMapOutput)
}

type AwsOutput struct{ *pulumi.OutputState }

func (AwsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Aws)(nil)).Elem()
}

func (o AwsOutput) ToAwsOutput() AwsOutput {
	return o
}

func (o AwsOutput) ToAwsOutputWithContext(ctx context.Context) AwsOutput {
	return o
}

// Access key ID for AWS.
func (o AwsOutput) AccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Aws) pulumi.StringOutput { return v.AccessKey }).(pulumi.StringOutput)
}

// Date when entity was created. Date and time format is ISO 8601 and UTC.
func (o AwsOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Aws) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Human-friendly description.
func (o AwsOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Aws) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// HATEOAS of entity.
func (o AwsOutput) Links() AwsLinkArrayOutput {
	return o.ApplyT(func(v *Aws) AwsLinkArrayOutput { return v.Links }).(AwsLinkArrayOutput)
}

// Name of AWS cloud account.
func (o AwsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Aws) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of organization that entity belongs to.
func (o AwsOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *Aws) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

// Email of entity owner.
func (o AwsOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *Aws) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

// Set of region names enabled for the cloud account.
func (o AwsOutput) Regions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Aws) pulumi.StringArrayOutput { return v.Regions }).(pulumi.StringArrayOutput)
}

// AWS Secret Access Key
func (o AwsOutput) SecretKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Aws) pulumi.StringOutput { return v.SecretKey }).(pulumi.StringOutput)
}

// Set of tag keys and values to apply to the cloud account.\
// Example:[ { "key" : "vmware", "value": "provider" } ]
func (o AwsOutput) Tags() AwsTagArrayOutput {
	return o.ApplyT(func(v *Aws) AwsTagArrayOutput { return v.Tags }).(AwsTagArrayOutput)
}

// Date when entity was last updated. Date and time format is ISO 8601 and UTC.
func (o AwsOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Aws) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type AwsArrayOutput struct{ *pulumi.OutputState }

func (AwsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Aws)(nil)).Elem()
}

func (o AwsArrayOutput) ToAwsArrayOutput() AwsArrayOutput {
	return o
}

func (o AwsArrayOutput) ToAwsArrayOutputWithContext(ctx context.Context) AwsArrayOutput {
	return o
}

func (o AwsArrayOutput) Index(i pulumi.IntInput) AwsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Aws {
		return vs[0].([]*Aws)[vs[1].(int)]
	}).(AwsOutput)
}

type AwsMapOutput struct{ *pulumi.OutputState }

func (AwsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Aws)(nil)).Elem()
}

func (o AwsMapOutput) ToAwsMapOutput() AwsMapOutput {
	return o
}

func (o AwsMapOutput) ToAwsMapOutputWithContext(ctx context.Context) AwsMapOutput {
	return o
}

func (o AwsMapOutput) MapIndex(k pulumi.StringInput) AwsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Aws {
		return vs[0].(map[string]*Aws)[vs[1].(string)]
	}).(AwsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AwsInput)(nil)).Elem(), &Aws{})
	pulumi.RegisterInputType(reflect.TypeOf((*AwsArrayInput)(nil)).Elem(), AwsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AwsMapInput)(nil)).Elem(), AwsMap{})
	pulumi.RegisterOutputType(AwsOutput{})
	pulumi.RegisterOutputType(AwsArrayOutput{})
	pulumi.RegisterOutputType(AwsMapOutput{})
}