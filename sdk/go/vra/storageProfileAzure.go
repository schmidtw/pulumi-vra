// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vra

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
// ### S
// This is an example of how to create a storage profile azure resource.
//
// **Vra storage profile azure:**
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
//			_, err := vra.NewStorageProfileAzure(ctx, "thisStorageProfileAzure", &vra.StorageProfileAzureArgs{
//				Description:        pulumi.String("Azure Storage Profile with managed disks."),
//				RegionId:           pulumi.Any(data.Vra_region.This.Id),
//				DefaultItem:        pulumi.Bool(false),
//				SupportsEncryption: pulumi.Bool(false),
//				DataDiskCaching:    pulumi.String("None"),
//				DiskType:           pulumi.String("Standard_LRS"),
//				OsDiskCaching:      pulumi.String("None"),
//				Tags: StorageProfileAzureTagArray{
//					&StorageProfileAzureTagArgs{
//						Key:   pulumi.String("foo"),
//						Value: pulumi.String("bar"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vra.NewStorageProfileAzure(ctx, "thisIndex/storageProfileAzureStorageProfileAzure", &vra.StorageProfileAzureArgs{
//				Description:        pulumi.String("Azure Storage Profile with unmanaged disks."),
//				RegionId:           pulumi.Any(data.Vra_region.This.Id),
//				DefaultItem:        pulumi.Bool(false),
//				SupportsEncryption: pulumi.Bool(false),
//				DataDiskCaching:    pulumi.String("None"),
//				OsDiskCaching:      pulumi.String("None"),
//				Tags: StorageProfileAzureTagArray{
//					&StorageProfileAzureTagArgs{
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
// A storage profile azure resource supports the following arguments:
type StorageProfileAzure struct {
	pulumi.CustomResourceState

	// Date when the entity was created. The date is in ISO 6801 and UTC.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Indicates the caching mechanism for additional disk.
	DataDiskCaching pulumi.StringOutput `pulumi:"dataDiskCaching"`
	// Indicates if this storage profile is a default profile.
	DefaultItem pulumi.BoolOutput `pulumi:"defaultItem"`
	// A human-friendly description.
	Description pulumi.StringOutput `pulumi:"description"`
	// Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed.
	DiskType pulumi.StringOutput `pulumi:"diskType"`
	// The id of the region as seen in the cloud provider for which this profile is defined.
	ExternalRegionId pulumi.StringOutput `pulumi:"externalRegionId"`
	// HATEOAS of the entity
	Links StorageProfileAzureLinkArrayOutput `pulumi:"links"`
	// A human-friendly name used as an identifier in APIs that support this option.
	Name pulumi.StringOutput `pulumi:"name"`
	// The id of the organization this entity belongs to.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// Indicates the caching mechanism for OS disk. Default policy for OS disks is Read/Write.
	OsDiskCaching pulumi.StringOutput `pulumi:"osDiskCaching"`
	// Email of the user that owns the entity.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// A link to the region that is associated with the storage profile.
	RegionId pulumi.StringOutput `pulumi:"regionId"`
	// Id of a storage account where in the disk is placed.
	StorageAccountId pulumi.StringOutput `pulumi:"storageAccountId"`
	// Indicates whether this storage policy should support encryption or not.
	SupportsEncryption pulumi.BoolOutput `pulumi:"supportsEncryption"`
	// A set of tag keys and optional values that were set on this Network Profile.
	// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
	Tags StorageProfileAzureTagArrayOutput `pulumi:"tags"`
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewStorageProfileAzure registers a new resource with the given unique name, arguments, and options.
func NewStorageProfileAzure(ctx *pulumi.Context,
	name string, args *StorageProfileAzureArgs, opts ...pulumi.ResourceOption) (*StorageProfileAzure, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultItem == nil {
		return nil, errors.New("invalid value for required argument 'DefaultItem'")
	}
	if args.RegionId == nil {
		return nil, errors.New("invalid value for required argument 'RegionId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource StorageProfileAzure
	err := ctx.RegisterResource("vra:index/storageProfileAzure:StorageProfileAzure", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageProfileAzure gets an existing StorageProfileAzure resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageProfileAzure(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageProfileAzureState, opts ...pulumi.ResourceOption) (*StorageProfileAzure, error) {
	var resource StorageProfileAzure
	err := ctx.ReadResource("vra:index/storageProfileAzure:StorageProfileAzure", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageProfileAzure resources.
type storageProfileAzureState struct {
	// Date when the entity was created. The date is in ISO 6801 and UTC.
	CreatedAt *string `pulumi:"createdAt"`
	// Indicates the caching mechanism for additional disk.
	DataDiskCaching *string `pulumi:"dataDiskCaching"`
	// Indicates if this storage profile is a default profile.
	DefaultItem *bool `pulumi:"defaultItem"`
	// A human-friendly description.
	Description *string `pulumi:"description"`
	// Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed.
	DiskType *string `pulumi:"diskType"`
	// The id of the region as seen in the cloud provider for which this profile is defined.
	ExternalRegionId *string `pulumi:"externalRegionId"`
	// HATEOAS of the entity
	Links []StorageProfileAzureLink `pulumi:"links"`
	// A human-friendly name used as an identifier in APIs that support this option.
	Name *string `pulumi:"name"`
	// The id of the organization this entity belongs to.
	OrganizationId *string `pulumi:"organizationId"`
	// Indicates the caching mechanism for OS disk. Default policy for OS disks is Read/Write.
	OsDiskCaching *string `pulumi:"osDiskCaching"`
	// Email of the user that owns the entity.
	Owner *string `pulumi:"owner"`
	// A link to the region that is associated with the storage profile.
	RegionId *string `pulumi:"regionId"`
	// Id of a storage account where in the disk is placed.
	StorageAccountId *string `pulumi:"storageAccountId"`
	// Indicates whether this storage policy should support encryption or not.
	SupportsEncryption *bool `pulumi:"supportsEncryption"`
	// A set of tag keys and optional values that were set on this Network Profile.
	// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
	Tags []StorageProfileAzureTag `pulumi:"tags"`
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type StorageProfileAzureState struct {
	// Date when the entity was created. The date is in ISO 6801 and UTC.
	CreatedAt pulumi.StringPtrInput
	// Indicates the caching mechanism for additional disk.
	DataDiskCaching pulumi.StringPtrInput
	// Indicates if this storage profile is a default profile.
	DefaultItem pulumi.BoolPtrInput
	// A human-friendly description.
	Description pulumi.StringPtrInput
	// Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed.
	DiskType pulumi.StringPtrInput
	// The id of the region as seen in the cloud provider for which this profile is defined.
	ExternalRegionId pulumi.StringPtrInput
	// HATEOAS of the entity
	Links StorageProfileAzureLinkArrayInput
	// A human-friendly name used as an identifier in APIs that support this option.
	Name pulumi.StringPtrInput
	// The id of the organization this entity belongs to.
	OrganizationId pulumi.StringPtrInput
	// Indicates the caching mechanism for OS disk. Default policy for OS disks is Read/Write.
	OsDiskCaching pulumi.StringPtrInput
	// Email of the user that owns the entity.
	Owner pulumi.StringPtrInput
	// A link to the region that is associated with the storage profile.
	RegionId pulumi.StringPtrInput
	// Id of a storage account where in the disk is placed.
	StorageAccountId pulumi.StringPtrInput
	// Indicates whether this storage policy should support encryption or not.
	SupportsEncryption pulumi.BoolPtrInput
	// A set of tag keys and optional values that were set on this Network Profile.
	// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
	Tags StorageProfileAzureTagArrayInput
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt pulumi.StringPtrInput
}

func (StorageProfileAzureState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageProfileAzureState)(nil)).Elem()
}

type storageProfileAzureArgs struct {
	// Indicates the caching mechanism for additional disk.
	DataDiskCaching *string `pulumi:"dataDiskCaching"`
	// Indicates if this storage profile is a default profile.
	DefaultItem bool `pulumi:"defaultItem"`
	// A human-friendly description.
	Description *string `pulumi:"description"`
	// Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed.
	DiskType *string `pulumi:"diskType"`
	// A human-friendly name used as an identifier in APIs that support this option.
	Name *string `pulumi:"name"`
	// Indicates the caching mechanism for OS disk. Default policy for OS disks is Read/Write.
	OsDiskCaching *string `pulumi:"osDiskCaching"`
	// A link to the region that is associated with the storage profile.
	RegionId string `pulumi:"regionId"`
	// Id of a storage account where in the disk is placed.
	StorageAccountId *string `pulumi:"storageAccountId"`
	// Indicates whether this storage policy should support encryption or not.
	SupportsEncryption *bool `pulumi:"supportsEncryption"`
	// A set of tag keys and optional values that were set on this Network Profile.
	// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
	Tags []StorageProfileAzureTag `pulumi:"tags"`
}

// The set of arguments for constructing a StorageProfileAzure resource.
type StorageProfileAzureArgs struct {
	// Indicates the caching mechanism for additional disk.
	DataDiskCaching pulumi.StringPtrInput
	// Indicates if this storage profile is a default profile.
	DefaultItem pulumi.BoolInput
	// A human-friendly description.
	Description pulumi.StringPtrInput
	// Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed.
	DiskType pulumi.StringPtrInput
	// A human-friendly name used as an identifier in APIs that support this option.
	Name pulumi.StringPtrInput
	// Indicates the caching mechanism for OS disk. Default policy for OS disks is Read/Write.
	OsDiskCaching pulumi.StringPtrInput
	// A link to the region that is associated with the storage profile.
	RegionId pulumi.StringInput
	// Id of a storage account where in the disk is placed.
	StorageAccountId pulumi.StringPtrInput
	// Indicates whether this storage policy should support encryption or not.
	SupportsEncryption pulumi.BoolPtrInput
	// A set of tag keys and optional values that were set on this Network Profile.
	// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
	Tags StorageProfileAzureTagArrayInput
}

func (StorageProfileAzureArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageProfileAzureArgs)(nil)).Elem()
}

type StorageProfileAzureInput interface {
	pulumi.Input

	ToStorageProfileAzureOutput() StorageProfileAzureOutput
	ToStorageProfileAzureOutputWithContext(ctx context.Context) StorageProfileAzureOutput
}

func (*StorageProfileAzure) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfileAzure)(nil)).Elem()
}

func (i *StorageProfileAzure) ToStorageProfileAzureOutput() StorageProfileAzureOutput {
	return i.ToStorageProfileAzureOutputWithContext(context.Background())
}

func (i *StorageProfileAzure) ToStorageProfileAzureOutputWithContext(ctx context.Context) StorageProfileAzureOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileAzureOutput)
}

// StorageProfileAzureArrayInput is an input type that accepts StorageProfileAzureArray and StorageProfileAzureArrayOutput values.
// You can construct a concrete instance of `StorageProfileAzureArrayInput` via:
//
//	StorageProfileAzureArray{ StorageProfileAzureArgs{...} }
type StorageProfileAzureArrayInput interface {
	pulumi.Input

	ToStorageProfileAzureArrayOutput() StorageProfileAzureArrayOutput
	ToStorageProfileAzureArrayOutputWithContext(context.Context) StorageProfileAzureArrayOutput
}

type StorageProfileAzureArray []StorageProfileAzureInput

func (StorageProfileAzureArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StorageProfileAzure)(nil)).Elem()
}

func (i StorageProfileAzureArray) ToStorageProfileAzureArrayOutput() StorageProfileAzureArrayOutput {
	return i.ToStorageProfileAzureArrayOutputWithContext(context.Background())
}

func (i StorageProfileAzureArray) ToStorageProfileAzureArrayOutputWithContext(ctx context.Context) StorageProfileAzureArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileAzureArrayOutput)
}

// StorageProfileAzureMapInput is an input type that accepts StorageProfileAzureMap and StorageProfileAzureMapOutput values.
// You can construct a concrete instance of `StorageProfileAzureMapInput` via:
//
//	StorageProfileAzureMap{ "key": StorageProfileAzureArgs{...} }
type StorageProfileAzureMapInput interface {
	pulumi.Input

	ToStorageProfileAzureMapOutput() StorageProfileAzureMapOutput
	ToStorageProfileAzureMapOutputWithContext(context.Context) StorageProfileAzureMapOutput
}

type StorageProfileAzureMap map[string]StorageProfileAzureInput

func (StorageProfileAzureMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StorageProfileAzure)(nil)).Elem()
}

func (i StorageProfileAzureMap) ToStorageProfileAzureMapOutput() StorageProfileAzureMapOutput {
	return i.ToStorageProfileAzureMapOutputWithContext(context.Background())
}

func (i StorageProfileAzureMap) ToStorageProfileAzureMapOutputWithContext(ctx context.Context) StorageProfileAzureMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileAzureMapOutput)
}

type StorageProfileAzureOutput struct{ *pulumi.OutputState }

func (StorageProfileAzureOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfileAzure)(nil)).Elem()
}

func (o StorageProfileAzureOutput) ToStorageProfileAzureOutput() StorageProfileAzureOutput {
	return o
}

func (o StorageProfileAzureOutput) ToStorageProfileAzureOutputWithContext(ctx context.Context) StorageProfileAzureOutput {
	return o
}

// Date when the entity was created. The date is in ISO 6801 and UTC.
func (o StorageProfileAzureOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageProfileAzure) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Indicates the caching mechanism for additional disk.
func (o StorageProfileAzureOutput) DataDiskCaching() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageProfileAzure) pulumi.StringOutput { return v.DataDiskCaching }).(pulumi.StringOutput)
}

// Indicates if this storage profile is a default profile.
func (o StorageProfileAzureOutput) DefaultItem() pulumi.BoolOutput {
	return o.ApplyT(func(v *StorageProfileAzure) pulumi.BoolOutput { return v.DefaultItem }).(pulumi.BoolOutput)
}

// A human-friendly description.
func (o StorageProfileAzureOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageProfileAzure) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed.
func (o StorageProfileAzureOutput) DiskType() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageProfileAzure) pulumi.StringOutput { return v.DiskType }).(pulumi.StringOutput)
}

// The id of the region as seen in the cloud provider for which this profile is defined.
func (o StorageProfileAzureOutput) ExternalRegionId() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageProfileAzure) pulumi.StringOutput { return v.ExternalRegionId }).(pulumi.StringOutput)
}

// HATEOAS of the entity
func (o StorageProfileAzureOutput) Links() StorageProfileAzureLinkArrayOutput {
	return o.ApplyT(func(v *StorageProfileAzure) StorageProfileAzureLinkArrayOutput { return v.Links }).(StorageProfileAzureLinkArrayOutput)
}

// A human-friendly name used as an identifier in APIs that support this option.
func (o StorageProfileAzureOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageProfileAzure) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The id of the organization this entity belongs to.
func (o StorageProfileAzureOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageProfileAzure) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// Indicates the caching mechanism for OS disk. Default policy for OS disks is Read/Write.
func (o StorageProfileAzureOutput) OsDiskCaching() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageProfileAzure) pulumi.StringOutput { return v.OsDiskCaching }).(pulumi.StringOutput)
}

// Email of the user that owns the entity.
func (o StorageProfileAzureOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageProfileAzure) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

// A link to the region that is associated with the storage profile.
func (o StorageProfileAzureOutput) RegionId() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageProfileAzure) pulumi.StringOutput { return v.RegionId }).(pulumi.StringOutput)
}

// Id of a storage account where in the disk is placed.
func (o StorageProfileAzureOutput) StorageAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageProfileAzure) pulumi.StringOutput { return v.StorageAccountId }).(pulumi.StringOutput)
}

// Indicates whether this storage policy should support encryption or not.
func (o StorageProfileAzureOutput) SupportsEncryption() pulumi.BoolOutput {
	return o.ApplyT(func(v *StorageProfileAzure) pulumi.BoolOutput { return v.SupportsEncryption }).(pulumi.BoolOutput)
}

// A set of tag keys and optional values that were set on this Network Profile.
// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
func (o StorageProfileAzureOutput) Tags() StorageProfileAzureTagArrayOutput {
	return o.ApplyT(func(v *StorageProfileAzure) StorageProfileAzureTagArrayOutput { return v.Tags }).(StorageProfileAzureTagArrayOutput)
}

// Date when the entity was last updated. The date is ISO 8601 and UTC.
func (o StorageProfileAzureOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageProfileAzure) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type StorageProfileAzureArrayOutput struct{ *pulumi.OutputState }

func (StorageProfileAzureArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StorageProfileAzure)(nil)).Elem()
}

func (o StorageProfileAzureArrayOutput) ToStorageProfileAzureArrayOutput() StorageProfileAzureArrayOutput {
	return o
}

func (o StorageProfileAzureArrayOutput) ToStorageProfileAzureArrayOutputWithContext(ctx context.Context) StorageProfileAzureArrayOutput {
	return o
}

func (o StorageProfileAzureArrayOutput) Index(i pulumi.IntInput) StorageProfileAzureOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StorageProfileAzure {
		return vs[0].([]*StorageProfileAzure)[vs[1].(int)]
	}).(StorageProfileAzureOutput)
}

type StorageProfileAzureMapOutput struct{ *pulumi.OutputState }

func (StorageProfileAzureMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StorageProfileAzure)(nil)).Elem()
}

func (o StorageProfileAzureMapOutput) ToStorageProfileAzureMapOutput() StorageProfileAzureMapOutput {
	return o
}

func (o StorageProfileAzureMapOutput) ToStorageProfileAzureMapOutputWithContext(ctx context.Context) StorageProfileAzureMapOutput {
	return o
}

func (o StorageProfileAzureMapOutput) MapIndex(k pulumi.StringInput) StorageProfileAzureOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StorageProfileAzure {
		return vs[0].(map[string]*StorageProfileAzure)[vs[1].(string)]
	}).(StorageProfileAzureOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StorageProfileAzureInput)(nil)).Elem(), &StorageProfileAzure{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageProfileAzureArrayInput)(nil)).Elem(), StorageProfileAzureArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageProfileAzureMapInput)(nil)).Elem(), StorageProfileAzureMap{})
	pulumi.RegisterOutputType(StorageProfileAzureOutput{})
	pulumi.RegisterOutputType(StorageProfileAzureArrayOutput{})
	pulumi.RegisterOutputType(StorageProfileAzureMapOutput{})
}