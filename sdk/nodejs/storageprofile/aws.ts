// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 * ### S
 * This is an example of how to create a storage profile aws resource.
 *
 * **Vra storage profile aws:**
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vra from "@pulumiverse/vra";
 *
 * // AWS storage profile using vra_storage_profile_aws resource and EBS disk type.
 * const thisAws = new vra.storageprofile.Aws("thisAws", {
 *     description: "AWS Storage Profile with instance store device type.",
 *     regionId: data.vra_region["this"].id,
 *     defaultItem: false,
 *     supportsEncryption: false,
 *     deviceType: "ebs",
 *     volumeType: "io1",
 *     iops: "1000",
 *     tags: [{
 *         key: "foo",
 *         value: "bar",
 *     }],
 * });
 * // AWS storage profile using vra_storage_profile_aws resource and instance store disk type.
 * const thisStorageprofile_awsAws = new vra.storageprofile.Aws("thisStorageprofile/awsAws", {
 *     description: "AWS Storage Profile with instance store device type.",
 *     regionId: data.vra_region["this"].id,
 *     defaultItem: false,
 *     deviceType: "instance-store",
 *     tags: [{
 *         key: "foo",
 *         value: "bar",
 *     }],
 * });
 * ```
 *
 * A storage profile aws resource supports the following arguments:
 */
export class Aws extends pulumi.CustomResource {
    /**
     * Get an existing Aws resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AwsState, opts?: pulumi.CustomResourceOptions): Aws {
        return new Aws(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vra:storageprofile/aws:Aws';

    /**
     * Returns true if the given object is an instance of Aws.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Aws {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Aws.__pulumiType;
    }

    /**
     * Date when the entity was created. The date is in ISO 6801 and UTC.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Indicates if this storage profile is a default profile.
     */
    public readonly defaultItem!: pulumi.Output<boolean>;
    /**
     * A human-friendly description.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Indicates the type of storage device.
     */
    public readonly deviceType!: pulumi.Output<string>;
    /**
     * The id of the region as seen in the cloud provider for which this profile is defined.
     */
    public /*out*/ readonly externalRegionId!: pulumi.Output<string>;
    /**
     * Indicates maximum I/O operations per second in range(1-20,000).
     */
    public readonly iops!: pulumi.Output<string>;
    /**
     * HATEOAS of the entity
     */
    public /*out*/ readonly links!: pulumi.Output<outputs.storageprofile.AwsLink[]>;
    /**
     * A human-friendly name used as an identifier in APIs that support this option.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The id of the organization this entity belongs to.
     */
    public /*out*/ readonly organizationId!: pulumi.Output<string>;
    /**
     * Email of the user that owns the entity.
     */
    public /*out*/ readonly owner!: pulumi.Output<string>;
    /**
     * A link to the region that is associated with the storage profile.
     */
    public readonly regionId!: pulumi.Output<string>;
    /**
     * Indicates whether this storage profile supports encryption or not.
     */
    public readonly supportsEncryption!: pulumi.Output<boolean>;
    /**
     * A set of tag keys and optional values that were set on this Network Profile.
     * example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
     */
    public readonly tags!: pulumi.Output<outputs.storageprofile.AwsTag[]>;
    /**
     * Date when the entity was last updated. The date is ISO 8601 and UTC.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * Indicates the type of volume associated with type of storage.
     */
    public readonly volumeType!: pulumi.Output<string>;

    /**
     * Create a Aws resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AwsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AwsArgs | AwsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AwsState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["defaultItem"] = state ? state.defaultItem : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["deviceType"] = state ? state.deviceType : undefined;
            resourceInputs["externalRegionId"] = state ? state.externalRegionId : undefined;
            resourceInputs["iops"] = state ? state.iops : undefined;
            resourceInputs["links"] = state ? state.links : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["organizationId"] = state ? state.organizationId : undefined;
            resourceInputs["owner"] = state ? state.owner : undefined;
            resourceInputs["regionId"] = state ? state.regionId : undefined;
            resourceInputs["supportsEncryption"] = state ? state.supportsEncryption : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["volumeType"] = state ? state.volumeType : undefined;
        } else {
            const args = argsOrState as AwsArgs | undefined;
            if ((!args || args.defaultItem === undefined) && !opts.urn) {
                throw new Error("Missing required property 'defaultItem'");
            }
            if ((!args || args.regionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'regionId'");
            }
            resourceInputs["defaultItem"] = args ? args.defaultItem : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["deviceType"] = args ? args.deviceType : undefined;
            resourceInputs["iops"] = args ? args.iops : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["regionId"] = args ? args.regionId : undefined;
            resourceInputs["supportsEncryption"] = args ? args.supportsEncryption : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["volumeType"] = args ? args.volumeType : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["externalRegionId"] = undefined /*out*/;
            resourceInputs["links"] = undefined /*out*/;
            resourceInputs["organizationId"] = undefined /*out*/;
            resourceInputs["owner"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Aws.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Aws resources.
 */
export interface AwsState {
    /**
     * Date when the entity was created. The date is in ISO 6801 and UTC.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Indicates if this storage profile is a default profile.
     */
    defaultItem?: pulumi.Input<boolean>;
    /**
     * A human-friendly description.
     */
    description?: pulumi.Input<string>;
    /**
     * Indicates the type of storage device.
     */
    deviceType?: pulumi.Input<string>;
    /**
     * The id of the region as seen in the cloud provider for which this profile is defined.
     */
    externalRegionId?: pulumi.Input<string>;
    /**
     * Indicates maximum I/O operations per second in range(1-20,000).
     */
    iops?: pulumi.Input<string>;
    /**
     * HATEOAS of the entity
     */
    links?: pulumi.Input<pulumi.Input<inputs.storageprofile.AwsLink>[]>;
    /**
     * A human-friendly name used as an identifier in APIs that support this option.
     */
    name?: pulumi.Input<string>;
    /**
     * The id of the organization this entity belongs to.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * Email of the user that owns the entity.
     */
    owner?: pulumi.Input<string>;
    /**
     * A link to the region that is associated with the storage profile.
     */
    regionId?: pulumi.Input<string>;
    /**
     * Indicates whether this storage profile supports encryption or not.
     */
    supportsEncryption?: pulumi.Input<boolean>;
    /**
     * A set of tag keys and optional values that were set on this Network Profile.
     * example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
     */
    tags?: pulumi.Input<pulumi.Input<inputs.storageprofile.AwsTag>[]>;
    /**
     * Date when the entity was last updated. The date is ISO 8601 and UTC.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * Indicates the type of volume associated with type of storage.
     */
    volumeType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Aws resource.
 */
export interface AwsArgs {
    /**
     * Indicates if this storage profile is a default profile.
     */
    defaultItem: pulumi.Input<boolean>;
    /**
     * A human-friendly description.
     */
    description?: pulumi.Input<string>;
    /**
     * Indicates the type of storage device.
     */
    deviceType?: pulumi.Input<string>;
    /**
     * Indicates maximum I/O operations per second in range(1-20,000).
     */
    iops?: pulumi.Input<string>;
    /**
     * A human-friendly name used as an identifier in APIs that support this option.
     */
    name?: pulumi.Input<string>;
    /**
     * A link to the region that is associated with the storage profile.
     */
    regionId: pulumi.Input<string>;
    /**
     * Indicates whether this storage profile supports encryption or not.
     */
    supportsEncryption?: pulumi.Input<boolean>;
    /**
     * A set of tag keys and optional values that were set on this Network Profile.
     * example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
     */
    tags?: pulumi.Input<pulumi.Input<inputs.storageprofile.AwsTag>[]>;
    /**
     * Indicates the type of volume associated with type of storage.
     */
    volumeType?: pulumi.Input<string>;
}