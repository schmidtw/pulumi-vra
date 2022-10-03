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
    /// <summary>
    /// ## Example Usage
    /// ### S
    /// This is an example of how to create an image profile resource.
    /// 
    /// **Image profile:**
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Vra = schmidtw.Vra;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @this = new Vra.ImageProfile("this", new()
    ///     {
    ///         Description = "test image profile",
    ///         RegionId = data.Vra_region.This.Id,
    ///         ImageMappings = new[]
    ///         {
    ///             new Vra.Inputs.ImageProfileImageMappingArgs
    ///             {
    ///                 Name = "centos",
    ///                 ImageId = data.Vra_image.Centos.Id,
    ///                 Constraints = new[]
    ///                 {
    ///                     new Vra.Inputs.ImageProfileImageMappingConstraintArgs
    ///                     {
    ///                         Mandatory = true,
    ///                         Expression = "!env:Test",
    ///                     },
    ///                     new Vra.Inputs.ImageProfileImageMappingConstraintArgs
    ///                     {
    ///                         Mandatory = false,
    ///                         Expression = "foo:bar",
    ///                     },
    ///                 },
    ///             },
    ///             new Vra.Inputs.ImageProfileImageMappingArgs
    ///             {
    ///                 Name = "photon",
    ///                 ImageId = data.Vra_image.Photon.Id,
    ///                 CloudConfig = "runcmd echo 'Hello'",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// An image profile resource supports the following arguments:
    /// </summary>
    [VraResourceType("vra:index/imageProfile:ImageProfile")]
    public partial class ImageProfile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Date when the entity was created. The date is in ISO 6801 and UTC.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// A human-friendly description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The external regionId of the resource.
        /// </summary>
        [Output("externalRegionId")]
        public Output<string> ExternalRegionId { get; private set; } = null!;

        /// <summary>
        /// Image mapping defined for the corresponding region.
        /// </summary>
        [Output("imageMappings")]
        public Output<ImmutableArray<Outputs.ImageProfileImageMapping>> ImageMappings { get; private set; } = null!;

        /// <summary>
        /// A human-friendly name used as an identifier in APIs that support this option.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Email of the user that owns the entity.
        /// </summary>
        [Output("owner")]
        public Output<string> Owner { get; private set; } = null!;

        /// <summary>
        /// The id of the region for which this profile is defined as in vRealize Automation(vRA).
        /// </summary>
        [Output("regionId")]
        public Output<string> RegionId { get; private set; } = null!;

        /// <summary>
        /// Date when the entity was last updated. The date is ISO 8601 and UTC.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a ImageProfile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ImageProfile(string name, ImageProfileArgs args, CustomResourceOptions? options = null)
            : base("vra:index/imageProfile:ImageProfile", name, args ?? new ImageProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ImageProfile(string name, Input<string> id, ImageProfileState? state = null, CustomResourceOptions? options = null)
            : base("vra:index/imageProfile:ImageProfile", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/schmidtw/pulumi-vra/releases/download/v${VERSION}",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ImageProfile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ImageProfile Get(string name, Input<string> id, ImageProfileState? state = null, CustomResourceOptions? options = null)
        {
            return new ImageProfile(name, id, state, options);
        }
    }

    public sealed class ImageProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A human-friendly description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("imageMappings")]
        private InputList<Inputs.ImageProfileImageMappingArgs>? _imageMappings;

        /// <summary>
        /// Image mapping defined for the corresponding region.
        /// </summary>
        public InputList<Inputs.ImageProfileImageMappingArgs> ImageMappings
        {
            get => _imageMappings ?? (_imageMappings = new InputList<Inputs.ImageProfileImageMappingArgs>());
            set => _imageMappings = value;
        }

        /// <summary>
        /// A human-friendly name used as an identifier in APIs that support this option.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The id of the region for which this profile is defined as in vRealize Automation(vRA).
        /// </summary>
        [Input("regionId", required: true)]
        public Input<string> RegionId { get; set; } = null!;

        public ImageProfileArgs()
        {
        }
        public static new ImageProfileArgs Empty => new ImageProfileArgs();
    }

    public sealed class ImageProfileState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Date when the entity was created. The date is in ISO 6801 and UTC.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// A human-friendly description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The external regionId of the resource.
        /// </summary>
        [Input("externalRegionId")]
        public Input<string>? ExternalRegionId { get; set; }

        [Input("imageMappings")]
        private InputList<Inputs.ImageProfileImageMappingGetArgs>? _imageMappings;

        /// <summary>
        /// Image mapping defined for the corresponding region.
        /// </summary>
        public InputList<Inputs.ImageProfileImageMappingGetArgs> ImageMappings
        {
            get => _imageMappings ?? (_imageMappings = new InputList<Inputs.ImageProfileImageMappingGetArgs>());
            set => _imageMappings = value;
        }

        /// <summary>
        /// A human-friendly name used as an identifier in APIs that support this option.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Email of the user that owns the entity.
        /// </summary>
        [Input("owner")]
        public Input<string>? Owner { get; set; }

        /// <summary>
        /// The id of the region for which this profile is defined as in vRealize Automation(vRA).
        /// </summary>
        [Input("regionId")]
        public Input<string>? RegionId { get; set; }

        /// <summary>
        /// Date when the entity was last updated. The date is ISO 8601 and UTC.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public ImageProfileState()
        {
        }
        public static new ImageProfileState Empty => new ImageProfileState();
    }
}