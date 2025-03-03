// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.Inputs
{

    /// <summary>
    /// Indicates which analysis completed successfully. Multiple types of analysis can be performed on a single resource.
    /// </summary>
    public sealed class AnalysisCompletedArgs : global::Pulumi.ResourceArgs
    {
        [Input("analysisType")]
        private InputList<string>? _analysisType;
        public InputList<string> AnalysisType
        {
            get => _analysisType ?? (_analysisType = new InputList<string>());
            set => _analysisType = value;
        }

        public AnalysisCompletedArgs()
        {
        }
        public static new AnalysisCompletedArgs Empty => new AnalysisCompletedArgs();
    }
}
