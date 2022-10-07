// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The `vra.image.getImage` data source can be used to discover the lookup machine images with cloud accounts. This can then be used with resource that require an image. For example, to create an image profile using the `vra.image.Profile` resource.
 *
 * ## Example Usage
 *
 * This is an example of how to lookup images from a vSphere cloud account.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vra from "@pulumi/vra";
 * import * as vra from "@pulumiverse/vra";
 *
 * const thisVSphere = vra.cloudaccount.getVSphere({
 *     name: _var.cloud_account,
 * });
 * const thisRegion = thisVSphere.then(thisVSphere => vra.region.getRegion({
 *     cloudAccountId: thisVSphere.id,
 *     region: _var.region,
 * }));
 * const image0 = thisVSphere.then(thisVSphere => vra.image.getImage({
 *     filter: `name eq '${_var.image_name_0}' and cloudAccountId eq '${thisVSphere.id}' and externalRegionId eq '${_var.region}'`,
 * }));
 * const image1 = thisVSphere.then(thisVSphere => vra.image.getImage({
 *     filter: `name eq '${_var.image_name_1}' and cloudAccountId eq '${thisVSphere.id}' and externalRegionId eq '${_var.region}'`,
 * }));
 * const thisProfile = new vra.image.Profile("thisProfile", {
 *     description: _var.image_profile_description,
 *     regionId: thisRegion.then(thisRegion => thisRegion.id),
 *     imageMappings: [
 *         {
 *             name: _var.image_name_0,
 *             imageId: image0.then(image0 => image0.id),
 *         },
 *         {
 *             name: _var.image_name_1,
 *             imageId: image1.then(image1 => image1.id),
 *         },
 *     ],
 * });
 * ```
 */
export function getImage(args: GetImageArgs, opts?: pulumi.InvokeOptions): Promise<GetImageResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vra:image/getImage:getImage", {
        "filter": args.filter,
    }, opts);
}

/**
 * A collection of arguments for invoking getImage.
 */
export interface GetImageArgs {
    /**
     * Search criteria to narrow down the image discovery.
     */
    filter: string;
}

/**
 * A collection of values returned by getImage.
 */
export interface GetImageResult {
    /**
     * A human-friendly description of the image.
     */
    readonly description: string;
    /**
     * External entity id on the provider side.
     */
    readonly externalId: string;
    readonly filter: string;
    /**
     * The id of the image.
     */
    readonly id: string;
    /**
     * A human-friendly name used as an identifier in APIs that support this option.
     */
    readonly name: string;
    /**
     * Indicates whether this image is private. For vSphere, private images are templates and snapshots and public images are content library items.
     */
    readonly private: boolean;
    /**
     * The regionId of the image. For a vSphere cloud account, it is the `externalRegionId` such as `Datacenter:datacenter-2` and for an AWS cloud account, it is region name such as `us-east-1`, etc.
     */
    readonly region: string;
}

export function getImageOutput(args: GetImageOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetImageResult> {
    return pulumi.output(args).apply(a => getImage(a, opts))
}

/**
 * A collection of arguments for invoking getImage.
 */
export interface GetImageOutputArgs {
    /**
     * Search criteria to narrow down the image discovery.
     */
    filter: pulumi.Input<string>;
}