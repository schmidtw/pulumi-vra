// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Vra.Image
{
    public static class GetImage
    {
        /// <summary>
        /// The `vra.image.getImage` data source can be used to discover the lookup machine images with cloud accounts. This can then be used with resource that require an image. For example, to create an image profile using the `vra.image.Profile` resource.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// This is an example of how to lookup images from a vSphere cloud account.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// using Vra = Pulumiverse.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var thisVSphere = Vra.Cloudaccount.GetVSphere.Invoke(new()
        ///     {
        ///         Name = @var.Cloud_account,
        ///     });
        /// 
        ///     var thisRegion = Vra.Region.GetRegion.Invoke(new()
        ///     {
        ///         CloudAccountId = thisVSphere.Apply(getVSphereResult =&gt; getVSphereResult.Id),
        ///         Region = @var.Region,
        ///     });
        /// 
        ///     var image0 = Vra.Image.GetImage.Invoke(new()
        ///     {
        ///         Filter = $"name eq '{@var.Image_name_0}' and cloudAccountId eq '{thisVSphere.Apply(getVSphereResult =&gt; getVSphereResult.Id)}' and externalRegionId eq '{@var.Region}'",
        ///     });
        /// 
        ///     var image1 = Vra.Image.GetImage.Invoke(new()
        ///     {
        ///         Filter = $"name eq '{@var.Image_name_1}' and cloudAccountId eq '{thisVSphere.Apply(getVSphereResult =&gt; getVSphereResult.Id)}' and externalRegionId eq '{@var.Region}'",
        ///     });
        /// 
        ///     var thisProfile = new Vra.Image.Profile("thisProfile", new()
        ///     {
        ///         Description = @var.Image_profile_description,
        ///         RegionId = thisRegion.Apply(getRegionResult =&gt; getRegionResult.Id),
        ///         ImageMappings = new[]
        ///         {
        ///             new Vra.Image.Inputs.ProfileImageMappingArgs
        ///             {
        ///                 Name = @var.Image_name_0,
        ///                 ImageId = image0.Apply(getImageResult =&gt; getImageResult.Id),
        ///             },
        ///             new Vra.Image.Inputs.ProfileImageMappingArgs
        ///             {
        ///                 Name = @var.Image_name_1,
        ///                 ImageId = image1.Apply(getImageResult =&gt; getImageResult.Id),
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetImageResult> InvokeAsync(GetImageArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetImageResult>("vra:image/getImage:getImage", args ?? new GetImageArgs(), options.WithDefaults());

        /// <summary>
        /// The `vra.image.getImage` data source can be used to discover the lookup machine images with cloud accounts. This can then be used with resource that require an image. For example, to create an image profile using the `vra.image.Profile` resource.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// This is an example of how to lookup images from a vSphere cloud account.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// using Vra = Pulumiverse.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var thisVSphere = Vra.Cloudaccount.GetVSphere.Invoke(new()
        ///     {
        ///         Name = @var.Cloud_account,
        ///     });
        /// 
        ///     var thisRegion = Vra.Region.GetRegion.Invoke(new()
        ///     {
        ///         CloudAccountId = thisVSphere.Apply(getVSphereResult =&gt; getVSphereResult.Id),
        ///         Region = @var.Region,
        ///     });
        /// 
        ///     var image0 = Vra.Image.GetImage.Invoke(new()
        ///     {
        ///         Filter = $"name eq '{@var.Image_name_0}' and cloudAccountId eq '{thisVSphere.Apply(getVSphereResult =&gt; getVSphereResult.Id)}' and externalRegionId eq '{@var.Region}'",
        ///     });
        /// 
        ///     var image1 = Vra.Image.GetImage.Invoke(new()
        ///     {
        ///         Filter = $"name eq '{@var.Image_name_1}' and cloudAccountId eq '{thisVSphere.Apply(getVSphereResult =&gt; getVSphereResult.Id)}' and externalRegionId eq '{@var.Region}'",
        ///     });
        /// 
        ///     var thisProfile = new Vra.Image.Profile("thisProfile", new()
        ///     {
        ///         Description = @var.Image_profile_description,
        ///         RegionId = thisRegion.Apply(getRegionResult =&gt; getRegionResult.Id),
        ///         ImageMappings = new[]
        ///         {
        ///             new Vra.Image.Inputs.ProfileImageMappingArgs
        ///             {
        ///                 Name = @var.Image_name_0,
        ///                 ImageId = image0.Apply(getImageResult =&gt; getImageResult.Id),
        ///             },
        ///             new Vra.Image.Inputs.ProfileImageMappingArgs
        ///             {
        ///                 Name = @var.Image_name_1,
        ///                 ImageId = image1.Apply(getImageResult =&gt; getImageResult.Id),
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetImageResult> Invoke(GetImageInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetImageResult>("vra:image/getImage:getImage", args ?? new GetImageInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetImageArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Search criteria to narrow down the image discovery.
        /// </summary>
        [Input("filter", required: true)]
        public string Filter { get; set; } = null!;

        public GetImageArgs()
        {
        }
        public static new GetImageArgs Empty => new GetImageArgs();
    }

    public sealed class GetImageInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Search criteria to narrow down the image discovery.
        /// </summary>
        [Input("filter", required: true)]
        public Input<string> Filter { get; set; } = null!;

        public GetImageInvokeArgs()
        {
        }
        public static new GetImageInvokeArgs Empty => new GetImageInvokeArgs();
    }


    [OutputType]
    public sealed class GetImageResult
    {
        /// <summary>
        /// A human-friendly description of the image.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// External entity id on the provider side.
        /// </summary>
        public readonly string ExternalId;
        public readonly string Filter;
        /// <summary>
        /// The id of the image.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A human-friendly name used as an identifier in APIs that support this option.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Indicates whether this image is private. For vSphere, private images are templates and snapshots and public images are content library items.
        /// </summary>
        public readonly bool Private;
        /// <summary>
        /// The regionId of the image. For a vSphere cloud account, it is the `externalRegionId` such as `Datacenter:datacenter-2` and for an AWS cloud account, it is region name such as `us-east-1`, etc.
        /// </summary>
        public readonly string Region;

        [OutputConstructor]
        private GetImageResult(
            string description,

            string externalId,

            string filter,

            string id,

            string name,

            bool @private,

            string region)
        {
            Description = description;
            ExternalId = externalId;
            Filter = filter;
            Id = id;
            Name = name;
            Private = @private;
            Region = region;
        }
    }
}