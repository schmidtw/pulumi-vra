// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace schmidtw.Vra.Inputs
{

    public sealed class GetProjectMemberRoleArgs : global::Pulumi.InvokeArgs
    {
        [Input("email", required: true)]
        public string Email { get; set; } = null!;

        [Input("type")]
        public string? Type { get; set; }

        public GetProjectMemberRoleArgs()
        {
        }
        public static new GetProjectMemberRoleArgs Empty => new GetProjectMemberRoleArgs();
    }
}