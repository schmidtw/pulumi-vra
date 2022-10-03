// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace schmidtw.Vra
{
    public static class GetFabricCompute
    {
        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### S
        /// 
        /// This is an example of how to lookup fabric computes.
        /// 
        /// **Fabric compute data source by Id:**
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Vra.GetFabricCompute.Invoke(new()
        ///     {
        ///         Id = @var.Fabric_compute_id,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// **Fabric compute data source by filter query:**
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Vra.GetFabricCompute.Invoke(new()
        ///     {
        ///         Filter = $"name eq '{@var.Fabric_compute_name}'",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// A fabric compute data source supports the following arguments:
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetFabricComputeResult> InvokeAsync(GetFabricComputeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFabricComputeResult>("vra:index/getFabricCompute:getFabricCompute", args ?? new GetFabricComputeArgs(), options.WithDefaults());

        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### S
        /// 
        /// This is an example of how to lookup fabric computes.
        /// 
        /// **Fabric compute data source by Id:**
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Vra.GetFabricCompute.Invoke(new()
        ///     {
        ///         Id = @var.Fabric_compute_id,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// **Fabric compute data source by filter query:**
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Vra.GetFabricCompute.Invoke(new()
        ///     {
        ///         Filter = $"name eq '{@var.Fabric_compute_name}'",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// A fabric compute data source supports the following arguments:
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetFabricComputeResult> Invoke(GetFabricComputeInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetFabricComputeResult>("vra:index/getFabricCompute:getFabricCompute", args ?? new GetFabricComputeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFabricComputeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Search criteria to narrow down the fabric compute resource instance. Only one of 'id' or 'filter' must be specified.
        /// </summary>
        [Input("filter")]
        public string? Filter { get; set; }

        /// <summary>
        /// The id of the fabric compute resource instance. Only one of 'id' or 'filter' must be specified.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        [Input("tags")]
        private List<Inputs.GetFabricComputeTagArgs>? _tags;

        /// <summary>
        /// A set of tag keys and optional values that were set on this resource:
        /// </summary>
        public List<Inputs.GetFabricComputeTagArgs> Tags
        {
            get => _tags ?? (_tags = new List<Inputs.GetFabricComputeTagArgs>());
            set => _tags = value;
        }

        public GetFabricComputeArgs()
        {
        }
        public static new GetFabricComputeArgs Empty => new GetFabricComputeArgs();
    }

    public sealed class GetFabricComputeInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Search criteria to narrow down the fabric compute resource instance. Only one of 'id' or 'filter' must be specified.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// The id of the fabric compute resource instance. Only one of 'id' or 'filter' must be specified.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("tags")]
        private InputList<Inputs.GetFabricComputeTagInputArgs>? _tags;

        /// <summary>
        /// A set of tag keys and optional values that were set on this resource:
        /// </summary>
        public InputList<Inputs.GetFabricComputeTagInputArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.GetFabricComputeTagInputArgs>());
            set => _tags = value;
        }

        public GetFabricComputeInvokeArgs()
        {
        }
        public static new GetFabricComputeInvokeArgs Empty => new GetFabricComputeInvokeArgs();
    }


    [OutputType]
    public sealed class GetFabricComputeResult
    {
        /// <summary>
        /// Date when the entity was created. The date is in ISO 8601 and UTC.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// A list of key value pair of custom properties for the fabric compute resource.
        /// </summary>
        public readonly ImmutableDictionary<string, object> CustomProperties;
        /// <summary>
        /// A human-friendly description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The id of the external entity on the provider side.
        /// </summary>
        public readonly string ExternalId;
        /// <summary>
        /// The external region id of the fabric compute.
        /// </summary>
        public readonly string ExternalRegionId;
        /// <summary>
        /// The external zone id of the fabric compute.
        /// </summary>
        public readonly string ExternalZoneId;
        public readonly string? Filter;
        public readonly string Id;
        /// <summary>
        /// Lifecycle status of the compute instance.
        /// </summary>
        public readonly string LifecycleState;
        /// <summary>
        /// HATEOAS of the entity.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFabricComputeLinkResult> Links;
        /// <summary>
        /// A human-friendly name used as an identifier for the fabric compute resource instance.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The id of the organization this entity belongs to.
        /// </summary>
        public readonly string OrgId;
        /// <summary>
        /// Email of the user that owns the entity.
        /// </summary>
        public readonly string Owner;
        /// <summary>
        /// Power state of fabric compute instance.
        /// </summary>
        public readonly string PowerState;
        /// <summary>
        /// A set of tag keys and optional values that were set on this resource:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFabricComputeTagResult> Tags;
        /// <summary>
        /// Type of the fabric compute instance.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Date when the entity was last updated. The date is ISO 8601 and UTC.
        /// </summary>
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetFabricComputeResult(
            string createdAt,

            ImmutableDictionary<string, object> customProperties,

            string description,

            string externalId,

            string externalRegionId,

            string externalZoneId,

            string? filter,

            string id,

            string lifecycleState,

            ImmutableArray<Outputs.GetFabricComputeLinkResult> links,

            string name,

            string orgId,

            string owner,

            string powerState,

            ImmutableArray<Outputs.GetFabricComputeTagResult> tags,

            string type,

            string updatedAt)
        {
            CreatedAt = createdAt;
            CustomProperties = customProperties;
            Description = description;
            ExternalId = externalId;
            ExternalRegionId = externalRegionId;
            ExternalZoneId = externalZoneId;
            Filter = filter;
            Id = id;
            LifecycleState = lifecycleState;
            Links = links;
            Name = name;
            OrgId = orgId;
            Owner = owner;
            PowerState = powerState;
            Tags = tags;
            Type = type;
            UpdatedAt = updatedAt;
        }
    }
}