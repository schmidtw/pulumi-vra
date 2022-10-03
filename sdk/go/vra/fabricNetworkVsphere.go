// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vra

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Updates a VMware vRealize Automation fabricNetworkVsphere resource.
//
// ## Example Usage
// ### S
//
// You cannot create a vSphere fabric network resource, however you can import using the command specified in the import section below.
// Once a resource is imported, you can update it as shown below:
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
//			_, err := vra.NewFabricNetworkVsphere(ctx, "simple", &vra.FabricNetworkVsphereArgs{
//				Cidr:           pulumi.Any(_var.Cidr),
//				DefaultGateway: pulumi.Any(_var.Gateway),
//				Domain:         pulumi.Any(_var.Domain),
//				Tags: FabricNetworkVsphereTagArray{
//					&FabricNetworkVsphereTagArgs{
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
// # To import the vSphere fabric network resource, use the ID as in the following example
//
// ```sh
//
//	$ pulumi import vra:index/fabricNetworkVsphere:FabricNetworkVsphere new_fabric_network_vsphere 05956583-6488-4e7d-84c9-92a7b7219a15`
//
// ```
type FabricNetworkVsphere struct {
	pulumi.CustomResourceState

	// Network CIDR to be used.
	Cidr pulumi.StringOutput `pulumi:"cidr"`
	// Set of ids of the cloud accounts this entity belongs to.
	CloudAccountIds pulumi.StringArrayOutput `pulumi:"cloudAccountIds"`
	// Date when the entity was created. The date is in ISO 6801 and UTC.
	CreatedAt        pulumi.StringOutput `pulumi:"createdAt"`
	CustomProperties pulumi.MapOutput    `pulumi:"customProperties"`
	// IPv4 default gateway to be used.
	DefaultGateway pulumi.StringOutput `pulumi:"defaultGateway"`
	// IPv6 default gateway to be used.
	DefaultIpv6Gateway pulumi.StringOutput `pulumi:"defaultIpv6Gateway"`
	// List of dns search domains for the vSphere network.
	DnsSearchDomains pulumi.StringArrayOutput `pulumi:"dnsSearchDomains"`
	// A human-friendly name used as an identifier in APIs that support this option.
	DnsServerAddresses pulumi.StringArrayOutput `pulumi:"dnsServerAddresses"`
	// Domain for the vSphere network.
	Domain pulumi.StringPtrOutput `pulumi:"domain"`
	// External entity Id on the provider side.
	ExternalId pulumi.StringOutput `pulumi:"externalId"`
	// The id of the region for which this network is defined.
	ExternalRegionId pulumi.StringOutput `pulumi:"externalRegionId"`
	// Network IPv6 CIDR to be used.
	Ipv6Cidr pulumi.StringPtrOutput `pulumi:"ipv6Cidr"`
	// Indicates whether this is the default subnet for the zone.
	IsDefault pulumi.BoolPtrOutput `pulumi:"isDefault"`
	// Indicates whether the sub-network supports public IP assignment.
	IsPublic pulumi.BoolPtrOutput `pulumi:"isPublic"`
	// HATEOAS of the entity
	Links FabricNetworkVsphereLinkArrayOutput `pulumi:"links"`
	// A human-friendly name used as an identifier in APIs that support this option.
	Name  pulumi.StringOutput `pulumi:"name"`
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// ID of organization that entity belongs to.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// Set of tag keys and values to apply to the resource.
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags FabricNetworkVsphereTagArrayOutput `pulumi:"tags"`
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewFabricNetworkVsphere registers a new resource with the given unique name, arguments, and options.
func NewFabricNetworkVsphere(ctx *pulumi.Context,
	name string, args *FabricNetworkVsphereArgs, opts ...pulumi.ResourceOption) (*FabricNetworkVsphere, error) {
	if args == nil {
		args = &FabricNetworkVsphereArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FabricNetworkVsphere
	err := ctx.RegisterResource("vra:index/fabricNetworkVsphere:FabricNetworkVsphere", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFabricNetworkVsphere gets an existing FabricNetworkVsphere resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFabricNetworkVsphere(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FabricNetworkVsphereState, opts ...pulumi.ResourceOption) (*FabricNetworkVsphere, error) {
	var resource FabricNetworkVsphere
	err := ctx.ReadResource("vra:index/fabricNetworkVsphere:FabricNetworkVsphere", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FabricNetworkVsphere resources.
type fabricNetworkVsphereState struct {
	// Network CIDR to be used.
	Cidr *string `pulumi:"cidr"`
	// Set of ids of the cloud accounts this entity belongs to.
	CloudAccountIds []string `pulumi:"cloudAccountIds"`
	// Date when the entity was created. The date is in ISO 6801 and UTC.
	CreatedAt        *string                `pulumi:"createdAt"`
	CustomProperties map[string]interface{} `pulumi:"customProperties"`
	// IPv4 default gateway to be used.
	DefaultGateway *string `pulumi:"defaultGateway"`
	// IPv6 default gateway to be used.
	DefaultIpv6Gateway *string `pulumi:"defaultIpv6Gateway"`
	// List of dns search domains for the vSphere network.
	DnsSearchDomains []string `pulumi:"dnsSearchDomains"`
	// A human-friendly name used as an identifier in APIs that support this option.
	DnsServerAddresses []string `pulumi:"dnsServerAddresses"`
	// Domain for the vSphere network.
	Domain *string `pulumi:"domain"`
	// External entity Id on the provider side.
	ExternalId *string `pulumi:"externalId"`
	// The id of the region for which this network is defined.
	ExternalRegionId *string `pulumi:"externalRegionId"`
	// Network IPv6 CIDR to be used.
	Ipv6Cidr *string `pulumi:"ipv6Cidr"`
	// Indicates whether this is the default subnet for the zone.
	IsDefault *bool `pulumi:"isDefault"`
	// Indicates whether the sub-network supports public IP assignment.
	IsPublic *bool `pulumi:"isPublic"`
	// HATEOAS of the entity
	Links []FabricNetworkVsphereLink `pulumi:"links"`
	// A human-friendly name used as an identifier in APIs that support this option.
	Name  *string `pulumi:"name"`
	OrgId *string `pulumi:"orgId"`
	// ID of organization that entity belongs to.
	OrganizationId *string `pulumi:"organizationId"`
	// Set of tag keys and values to apply to the resource.
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags []FabricNetworkVsphereTag `pulumi:"tags"`
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type FabricNetworkVsphereState struct {
	// Network CIDR to be used.
	Cidr pulumi.StringPtrInput
	// Set of ids of the cloud accounts this entity belongs to.
	CloudAccountIds pulumi.StringArrayInput
	// Date when the entity was created. The date is in ISO 6801 and UTC.
	CreatedAt        pulumi.StringPtrInput
	CustomProperties pulumi.MapInput
	// IPv4 default gateway to be used.
	DefaultGateway pulumi.StringPtrInput
	// IPv6 default gateway to be used.
	DefaultIpv6Gateway pulumi.StringPtrInput
	// List of dns search domains for the vSphere network.
	DnsSearchDomains pulumi.StringArrayInput
	// A human-friendly name used as an identifier in APIs that support this option.
	DnsServerAddresses pulumi.StringArrayInput
	// Domain for the vSphere network.
	Domain pulumi.StringPtrInput
	// External entity Id on the provider side.
	ExternalId pulumi.StringPtrInput
	// The id of the region for which this network is defined.
	ExternalRegionId pulumi.StringPtrInput
	// Network IPv6 CIDR to be used.
	Ipv6Cidr pulumi.StringPtrInput
	// Indicates whether this is the default subnet for the zone.
	IsDefault pulumi.BoolPtrInput
	// Indicates whether the sub-network supports public IP assignment.
	IsPublic pulumi.BoolPtrInput
	// HATEOAS of the entity
	Links FabricNetworkVsphereLinkArrayInput
	// A human-friendly name used as an identifier in APIs that support this option.
	Name  pulumi.StringPtrInput
	OrgId pulumi.StringPtrInput
	// ID of organization that entity belongs to.
	OrganizationId pulumi.StringPtrInput
	// Set of tag keys and values to apply to the resource.
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags FabricNetworkVsphereTagArrayInput
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt pulumi.StringPtrInput
}

func (FabricNetworkVsphereState) ElementType() reflect.Type {
	return reflect.TypeOf((*fabricNetworkVsphereState)(nil)).Elem()
}

type fabricNetworkVsphereArgs struct {
	// Network CIDR to be used.
	Cidr *string `pulumi:"cidr"`
	// IPv4 default gateway to be used.
	DefaultGateway *string `pulumi:"defaultGateway"`
	// IPv6 default gateway to be used.
	DefaultIpv6Gateway *string `pulumi:"defaultIpv6Gateway"`
	// List of dns search domains for the vSphere network.
	DnsSearchDomains []string `pulumi:"dnsSearchDomains"`
	// A human-friendly name used as an identifier in APIs that support this option.
	DnsServerAddresses []string `pulumi:"dnsServerAddresses"`
	// Domain for the vSphere network.
	Domain *string `pulumi:"domain"`
	// Network IPv6 CIDR to be used.
	Ipv6Cidr *string `pulumi:"ipv6Cidr"`
	// Indicates whether this is the default subnet for the zone.
	IsDefault *bool `pulumi:"isDefault"`
	// Indicates whether the sub-network supports public IP assignment.
	IsPublic *bool `pulumi:"isPublic"`
	// Set of tag keys and values to apply to the resource.
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags []FabricNetworkVsphereTag `pulumi:"tags"`
}

// The set of arguments for constructing a FabricNetworkVsphere resource.
type FabricNetworkVsphereArgs struct {
	// Network CIDR to be used.
	Cidr pulumi.StringPtrInput
	// IPv4 default gateway to be used.
	DefaultGateway pulumi.StringPtrInput
	// IPv6 default gateway to be used.
	DefaultIpv6Gateway pulumi.StringPtrInput
	// List of dns search domains for the vSphere network.
	DnsSearchDomains pulumi.StringArrayInput
	// A human-friendly name used as an identifier in APIs that support this option.
	DnsServerAddresses pulumi.StringArrayInput
	// Domain for the vSphere network.
	Domain pulumi.StringPtrInput
	// Network IPv6 CIDR to be used.
	Ipv6Cidr pulumi.StringPtrInput
	// Indicates whether this is the default subnet for the zone.
	IsDefault pulumi.BoolPtrInput
	// Indicates whether the sub-network supports public IP assignment.
	IsPublic pulumi.BoolPtrInput
	// Set of tag keys and values to apply to the resource.
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags FabricNetworkVsphereTagArrayInput
}

func (FabricNetworkVsphereArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fabricNetworkVsphereArgs)(nil)).Elem()
}

type FabricNetworkVsphereInput interface {
	pulumi.Input

	ToFabricNetworkVsphereOutput() FabricNetworkVsphereOutput
	ToFabricNetworkVsphereOutputWithContext(ctx context.Context) FabricNetworkVsphereOutput
}

func (*FabricNetworkVsphere) ElementType() reflect.Type {
	return reflect.TypeOf((**FabricNetworkVsphere)(nil)).Elem()
}

func (i *FabricNetworkVsphere) ToFabricNetworkVsphereOutput() FabricNetworkVsphereOutput {
	return i.ToFabricNetworkVsphereOutputWithContext(context.Background())
}

func (i *FabricNetworkVsphere) ToFabricNetworkVsphereOutputWithContext(ctx context.Context) FabricNetworkVsphereOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FabricNetworkVsphereOutput)
}

// FabricNetworkVsphereArrayInput is an input type that accepts FabricNetworkVsphereArray and FabricNetworkVsphereArrayOutput values.
// You can construct a concrete instance of `FabricNetworkVsphereArrayInput` via:
//
//	FabricNetworkVsphereArray{ FabricNetworkVsphereArgs{...} }
type FabricNetworkVsphereArrayInput interface {
	pulumi.Input

	ToFabricNetworkVsphereArrayOutput() FabricNetworkVsphereArrayOutput
	ToFabricNetworkVsphereArrayOutputWithContext(context.Context) FabricNetworkVsphereArrayOutput
}

type FabricNetworkVsphereArray []FabricNetworkVsphereInput

func (FabricNetworkVsphereArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FabricNetworkVsphere)(nil)).Elem()
}

func (i FabricNetworkVsphereArray) ToFabricNetworkVsphereArrayOutput() FabricNetworkVsphereArrayOutput {
	return i.ToFabricNetworkVsphereArrayOutputWithContext(context.Background())
}

func (i FabricNetworkVsphereArray) ToFabricNetworkVsphereArrayOutputWithContext(ctx context.Context) FabricNetworkVsphereArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FabricNetworkVsphereArrayOutput)
}

// FabricNetworkVsphereMapInput is an input type that accepts FabricNetworkVsphereMap and FabricNetworkVsphereMapOutput values.
// You can construct a concrete instance of `FabricNetworkVsphereMapInput` via:
//
//	FabricNetworkVsphereMap{ "key": FabricNetworkVsphereArgs{...} }
type FabricNetworkVsphereMapInput interface {
	pulumi.Input

	ToFabricNetworkVsphereMapOutput() FabricNetworkVsphereMapOutput
	ToFabricNetworkVsphereMapOutputWithContext(context.Context) FabricNetworkVsphereMapOutput
}

type FabricNetworkVsphereMap map[string]FabricNetworkVsphereInput

func (FabricNetworkVsphereMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FabricNetworkVsphere)(nil)).Elem()
}

func (i FabricNetworkVsphereMap) ToFabricNetworkVsphereMapOutput() FabricNetworkVsphereMapOutput {
	return i.ToFabricNetworkVsphereMapOutputWithContext(context.Background())
}

func (i FabricNetworkVsphereMap) ToFabricNetworkVsphereMapOutputWithContext(ctx context.Context) FabricNetworkVsphereMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FabricNetworkVsphereMapOutput)
}

type FabricNetworkVsphereOutput struct{ *pulumi.OutputState }

func (FabricNetworkVsphereOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FabricNetworkVsphere)(nil)).Elem()
}

func (o FabricNetworkVsphereOutput) ToFabricNetworkVsphereOutput() FabricNetworkVsphereOutput {
	return o
}

func (o FabricNetworkVsphereOutput) ToFabricNetworkVsphereOutputWithContext(ctx context.Context) FabricNetworkVsphereOutput {
	return o
}

// Network CIDR to be used.
func (o FabricNetworkVsphereOutput) Cidr() pulumi.StringOutput {
	return o.ApplyT(func(v *FabricNetworkVsphere) pulumi.StringOutput { return v.Cidr }).(pulumi.StringOutput)
}

// Set of ids of the cloud accounts this entity belongs to.
func (o FabricNetworkVsphereOutput) CloudAccountIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FabricNetworkVsphere) pulumi.StringArrayOutput { return v.CloudAccountIds }).(pulumi.StringArrayOutput)
}

// Date when the entity was created. The date is in ISO 6801 and UTC.
func (o FabricNetworkVsphereOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *FabricNetworkVsphere) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o FabricNetworkVsphereOutput) CustomProperties() pulumi.MapOutput {
	return o.ApplyT(func(v *FabricNetworkVsphere) pulumi.MapOutput { return v.CustomProperties }).(pulumi.MapOutput)
}

// IPv4 default gateway to be used.
func (o FabricNetworkVsphereOutput) DefaultGateway() pulumi.StringOutput {
	return o.ApplyT(func(v *FabricNetworkVsphere) pulumi.StringOutput { return v.DefaultGateway }).(pulumi.StringOutput)
}

// IPv6 default gateway to be used.
func (o FabricNetworkVsphereOutput) DefaultIpv6Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v *FabricNetworkVsphere) pulumi.StringOutput { return v.DefaultIpv6Gateway }).(pulumi.StringOutput)
}

// List of dns search domains for the vSphere network.
func (o FabricNetworkVsphereOutput) DnsSearchDomains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FabricNetworkVsphere) pulumi.StringArrayOutput { return v.DnsSearchDomains }).(pulumi.StringArrayOutput)
}

// A human-friendly name used as an identifier in APIs that support this option.
func (o FabricNetworkVsphereOutput) DnsServerAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FabricNetworkVsphere) pulumi.StringArrayOutput { return v.DnsServerAddresses }).(pulumi.StringArrayOutput)
}

// Domain for the vSphere network.
func (o FabricNetworkVsphereOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FabricNetworkVsphere) pulumi.StringPtrOutput { return v.Domain }).(pulumi.StringPtrOutput)
}

// External entity Id on the provider side.
func (o FabricNetworkVsphereOutput) ExternalId() pulumi.StringOutput {
	return o.ApplyT(func(v *FabricNetworkVsphere) pulumi.StringOutput { return v.ExternalId }).(pulumi.StringOutput)
}

// The id of the region for which this network is defined.
func (o FabricNetworkVsphereOutput) ExternalRegionId() pulumi.StringOutput {
	return o.ApplyT(func(v *FabricNetworkVsphere) pulumi.StringOutput { return v.ExternalRegionId }).(pulumi.StringOutput)
}

// Network IPv6 CIDR to be used.
func (o FabricNetworkVsphereOutput) Ipv6Cidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FabricNetworkVsphere) pulumi.StringPtrOutput { return v.Ipv6Cidr }).(pulumi.StringPtrOutput)
}

// Indicates whether this is the default subnet for the zone.
func (o FabricNetworkVsphereOutput) IsDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FabricNetworkVsphere) pulumi.BoolPtrOutput { return v.IsDefault }).(pulumi.BoolPtrOutput)
}

// Indicates whether the sub-network supports public IP assignment.
func (o FabricNetworkVsphereOutput) IsPublic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FabricNetworkVsphere) pulumi.BoolPtrOutput { return v.IsPublic }).(pulumi.BoolPtrOutput)
}

// HATEOAS of the entity
func (o FabricNetworkVsphereOutput) Links() FabricNetworkVsphereLinkArrayOutput {
	return o.ApplyT(func(v *FabricNetworkVsphere) FabricNetworkVsphereLinkArrayOutput { return v.Links }).(FabricNetworkVsphereLinkArrayOutput)
}

// A human-friendly name used as an identifier in APIs that support this option.
func (o FabricNetworkVsphereOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FabricNetworkVsphere) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FabricNetworkVsphereOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *FabricNetworkVsphere) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

// ID of organization that entity belongs to.
func (o FabricNetworkVsphereOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *FabricNetworkVsphere) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// Set of tag keys and values to apply to the resource.
// Example:[ { "key" : "vmware", "value": "provider" } ]
func (o FabricNetworkVsphereOutput) Tags() FabricNetworkVsphereTagArrayOutput {
	return o.ApplyT(func(v *FabricNetworkVsphere) FabricNetworkVsphereTagArrayOutput { return v.Tags }).(FabricNetworkVsphereTagArrayOutput)
}

// Date when the entity was last updated. The date is ISO 8601 and UTC.
func (o FabricNetworkVsphereOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *FabricNetworkVsphere) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type FabricNetworkVsphereArrayOutput struct{ *pulumi.OutputState }

func (FabricNetworkVsphereArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FabricNetworkVsphere)(nil)).Elem()
}

func (o FabricNetworkVsphereArrayOutput) ToFabricNetworkVsphereArrayOutput() FabricNetworkVsphereArrayOutput {
	return o
}

func (o FabricNetworkVsphereArrayOutput) ToFabricNetworkVsphereArrayOutputWithContext(ctx context.Context) FabricNetworkVsphereArrayOutput {
	return o
}

func (o FabricNetworkVsphereArrayOutput) Index(i pulumi.IntInput) FabricNetworkVsphereOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FabricNetworkVsphere {
		return vs[0].([]*FabricNetworkVsphere)[vs[1].(int)]
	}).(FabricNetworkVsphereOutput)
}

type FabricNetworkVsphereMapOutput struct{ *pulumi.OutputState }

func (FabricNetworkVsphereMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FabricNetworkVsphere)(nil)).Elem()
}

func (o FabricNetworkVsphereMapOutput) ToFabricNetworkVsphereMapOutput() FabricNetworkVsphereMapOutput {
	return o
}

func (o FabricNetworkVsphereMapOutput) ToFabricNetworkVsphereMapOutputWithContext(ctx context.Context) FabricNetworkVsphereMapOutput {
	return o
}

func (o FabricNetworkVsphereMapOutput) MapIndex(k pulumi.StringInput) FabricNetworkVsphereOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FabricNetworkVsphere {
		return vs[0].(map[string]*FabricNetworkVsphere)[vs[1].(string)]
	}).(FabricNetworkVsphereOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FabricNetworkVsphereInput)(nil)).Elem(), &FabricNetworkVsphere{})
	pulumi.RegisterInputType(reflect.TypeOf((*FabricNetworkVsphereArrayInput)(nil)).Elem(), FabricNetworkVsphereArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FabricNetworkVsphereMapInput)(nil)).Elem(), FabricNetworkVsphereMap{})
	pulumi.RegisterOutputType(FabricNetworkVsphereOutput{})
	pulumi.RegisterOutputType(FabricNetworkVsphereArrayOutput{})
	pulumi.RegisterOutputType(FabricNetworkVsphereMapOutput{})
}