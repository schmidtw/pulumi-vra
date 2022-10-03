// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This resource provides a way to create a catalog source entitlement in VMware vRealize Automation.
 *
 * ## Example Usage
 * ### S
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vra from "@schmidtw/vra";
 *
 * const _this = new vra.CatalogSourceEntitlement("this", {
 *     catalogSourceId: _var.catalog_source_blueprint_id,
 *     projectId: _var.project_id,
 * });
 * ```
 * ## Attribute Reference
 *
 * * `definition` - Represents a catalog source that is linked to a project via an entitlement.
 *   
 *     * `description` - Description of the catalog source.
 *   
 *     * `iconId` - Icon id of associated catalog source.
 *   
 *     * `id` - Id of the catalog source.
 *   
 *     * `name` - Name of the catalog source.
 *   
 *     * `numberOfItems` - Number of items in the associated catalog source.
 *   
 *     * `sourceName` - Catalog source name.
 *   
 *     * `sourceType` - Catalog source type.
 *   
 *     * `type` - Content definition type.
 *
 * ## Import
 *
 * Catalog source entitlement can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import vra:index/catalogSourceEntitlement:CatalogSourceEntitlement this 05956583-6488-4e7d-84c9-92a7b7219a15`
 * ```
 */
export class CatalogSourceEntitlement extends pulumi.CustomResource {
    /**
     * Get an existing CatalogSourceEntitlement resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CatalogSourceEntitlementState, opts?: pulumi.CustomResourceOptions): CatalogSourceEntitlement {
        return new CatalogSourceEntitlement(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vra:index/catalogSourceEntitlement:CatalogSourceEntitlement';

    /**
     * Returns true if the given object is an instance of CatalogSourceEntitlement.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CatalogSourceEntitlement {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CatalogSourceEntitlement.__pulumiType;
    }

    /**
     * The id of the catalog source to create the entitlement.
     */
    public readonly catalogSourceId!: pulumi.Output<string>;
    public /*out*/ readonly definitions!: pulumi.Output<outputs.CatalogSourceEntitlementDefinition[]>;
    /**
     * The id of the project this entity belongs to.
     */
    public readonly projectId!: pulumi.Output<string>;

    /**
     * Create a CatalogSourceEntitlement resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CatalogSourceEntitlementArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CatalogSourceEntitlementArgs | CatalogSourceEntitlementState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CatalogSourceEntitlementState | undefined;
            resourceInputs["catalogSourceId"] = state ? state.catalogSourceId : undefined;
            resourceInputs["definitions"] = state ? state.definitions : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
        } else {
            const args = argsOrState as CatalogSourceEntitlementArgs | undefined;
            if ((!args || args.catalogSourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'catalogSourceId'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            resourceInputs["catalogSourceId"] = args ? args.catalogSourceId : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["definitions"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CatalogSourceEntitlement.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CatalogSourceEntitlement resources.
 */
export interface CatalogSourceEntitlementState {
    /**
     * The id of the catalog source to create the entitlement.
     */
    catalogSourceId?: pulumi.Input<string>;
    definitions?: pulumi.Input<pulumi.Input<inputs.CatalogSourceEntitlementDefinition>[]>;
    /**
     * The id of the project this entity belongs to.
     */
    projectId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CatalogSourceEntitlement resource.
 */
export interface CatalogSourceEntitlementArgs {
    /**
     * The id of the catalog source to create the entitlement.
     */
    catalogSourceId: pulumi.Input<string>;
    /**
     * The id of the project this entity belongs to.
     */
    projectId: pulumi.Input<string>;
}