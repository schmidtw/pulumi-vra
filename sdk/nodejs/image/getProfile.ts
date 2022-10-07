// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 * ### S
 * This is an example of how to read an image profile as data source.
 *
 * **Image profile data source by its id:**
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vra from "@pulumi/vra";
 *
 * const thisProfile = pulumi.output(vra.image.getProfile({
 *     filter: "name eq 'foobar'",
 * }));
 * ```
 *
 * **Vra image profile data source filter by region id:**
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vra from "@pulumi/vra";
 *
 * const this = vra.image.getProfile({
 *     regionId: vra_image_profile["this"].region_id,
 * });
 * ```
 *
 * An image profile data source supports the following arguments:
 */
export function getProfile(args?: GetProfileArgs, opts?: pulumi.InvokeOptions): Promise<GetProfileResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vra:image/getProfile:getProfile", {
        "description": args.description,
        "filter": args.filter,
        "id": args.id,
        "imageMappings": args.imageMappings,
        "name": args.name,
        "regionId": args.regionId,
    }, opts);
}

/**
 * A collection of arguments for invoking getProfile.
 */
export interface GetProfileArgs {
    /**
     * A human-friendly description.
     */
    description?: string;
    /**
     * Filter query string that is supported by vRA multi-cloud IaaS API. Example: regionId eq '<regionId>' and cloudAccountId eq '<cloudAccountId>'.
     */
    filter?: string;
    /**
     * The id of the image profile instance.
     */
    id?: string;
    /**
     * Image mapping defined for the corresponding region.
     */
    imageMappings?: inputs.image.GetProfileImageMapping[];
    /**
     * A human-friendly name used as an identifier in APIs that support this option.
     */
    name?: string;
    /**
     * The id of the region for which this profile is defined as in vRealize Automation(vRA).
     */
    regionId?: string;
}

/**
 * A collection of values returned by getProfile.
 */
export interface GetProfileResult {
    /**
     * Date when the entity was created. The date is in ISO 6801 and UTC.
     */
    readonly createdAt: string;
    /**
     * A human-friendly description.
     */
    readonly description?: string;
    /**
     * The external regionId of the resource.
     */
    readonly externalRegionId: string;
    readonly filter?: string;
    readonly id: string;
    /**
     * Image mapping defined for the corresponding region.
     */
    readonly imageMappings?: outputs.image.GetProfileImageMapping[];
    readonly name: string;
    /**
     * Email of the user that owns the entity.
     */
    readonly owner: string;
    readonly regionId: string;
    /**
     * Date when the entity was last updated. The date is ISO 8601 and UTC.
     */
    readonly updatedAt: string;
}

export function getProfileOutput(args?: GetProfileOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProfileResult> {
    return pulumi.output(args).apply(a => getProfile(a, opts))
}

/**
 * A collection of arguments for invoking getProfile.
 */
export interface GetProfileOutputArgs {
    /**
     * A human-friendly description.
     */
    description?: pulumi.Input<string>;
    /**
     * Filter query string that is supported by vRA multi-cloud IaaS API. Example: regionId eq '<regionId>' and cloudAccountId eq '<cloudAccountId>'.
     */
    filter?: pulumi.Input<string>;
    /**
     * The id of the image profile instance.
     */
    id?: pulumi.Input<string>;
    /**
     * Image mapping defined for the corresponding region.
     */
    imageMappings?: pulumi.Input<pulumi.Input<inputs.image.GetProfileImageMappingArgs>[]>;
    /**
     * A human-friendly name used as an identifier in APIs that support this option.
     */
    name?: pulumi.Input<string>;
    /**
     * The id of the region for which this profile is defined as in vRealize Automation(vRA).
     */
    regionId?: pulumi.Input<string>;
}