// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Alpha1.Inputs
{

    /// <summary>
    /// Identifies the event that kicked off the build.
    /// </summary>
    public sealed class GoogleDevtoolsContaineranalysisV1alpha1SlsaProvenanceZeroTwoSlsaInvocationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Describes where the config file that kicked off the build came from.
        /// </summary>
        [Input("configSource")]
        public Input<Inputs.GoogleDevtoolsContaineranalysisV1alpha1SlsaProvenanceZeroTwoSlsaConfigSourceArgs>? ConfigSource { get; set; }

        [Input("environment")]
        private InputMap<string>? _environment;

        /// <summary>
        /// Any other builder-controlled inputs necessary for correctly evaluating the build.
        /// </summary>
        public InputMap<string> Environment
        {
            get => _environment ?? (_environment = new InputMap<string>());
            set => _environment = value;
        }

        [Input("parameters")]
        private InputMap<string>? _parameters;

        /// <summary>
        /// Collection of all external inputs that influenced the build on top of invocation.configSource.
        /// </summary>
        public InputMap<string> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<string>());
            set => _parameters = value;
        }

        public GoogleDevtoolsContaineranalysisV1alpha1SlsaProvenanceZeroTwoSlsaInvocationArgs()
        {
        }
        public static new GoogleDevtoolsContaineranalysisV1alpha1SlsaProvenanceZeroTwoSlsaInvocationArgs Empty => new GoogleDevtoolsContaineranalysisV1alpha1SlsaProvenanceZeroTwoSlsaInvocationArgs();
    }
}
