// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1.Outputs
{

    /// <summary>
    /// A Dataproc job for running Apache Spark (https://spark.apache.org/) applications on YARN.
    /// </summary>
    [OutputType]
    public sealed class SparkJobResponse
    {
        /// <summary>
        /// Optional. HCFS URIs of archives to be extracted into the working directory of each executor. Supported file types: .jar, .tar, .tar.gz, .tgz, and .zip.
        /// </summary>
        public readonly ImmutableArray<string> ArchiveUris;
        /// <summary>
        /// Optional. The arguments to pass to the driver. Do not include arguments, such as --conf, that can be set as job properties, since a collision may occur that causes an incorrect job submission.
        /// </summary>
        public readonly ImmutableArray<string> Args;
        /// <summary>
        /// Optional. HCFS URIs of files to be placed in the working directory of each executor. Useful for naively parallel tasks.
        /// </summary>
        public readonly ImmutableArray<string> FileUris;
        /// <summary>
        /// Optional. HCFS URIs of jar files to add to the CLASSPATHs of the Spark driver and tasks.
        /// </summary>
        public readonly ImmutableArray<string> JarFileUris;
        /// <summary>
        /// Optional. The runtime log config for job execution.
        /// </summary>
        public readonly Outputs.LoggingConfigResponse LoggingConfig;
        /// <summary>
        /// The name of the driver's main class. The jar file that contains the class must be in the default CLASSPATH or specified in jar_file_uris.
        /// </summary>
        public readonly string MainClass;
        /// <summary>
        /// The HCFS URI of the jar file that contains the main class.
        /// </summary>
        public readonly string MainJarFileUri;
        /// <summary>
        /// Optional. A mapping of property names to values, used to configure Spark. Properties that conflict with values set by the Dataproc API may be overwritten. Can include properties set in /etc/spark/conf/spark-defaults.conf and classes in user code.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Properties;

        [OutputConstructor]
        private SparkJobResponse(
            ImmutableArray<string> archiveUris,

            ImmutableArray<string> args,

            ImmutableArray<string> fileUris,

            ImmutableArray<string> jarFileUris,

            Outputs.LoggingConfigResponse loggingConfig,

            string mainClass,

            string mainJarFileUri,

            ImmutableDictionary<string, string> properties)
        {
            ArchiveUris = archiveUris;
            Args = args;
            FileUris = fileUris;
            JarFileUris = jarFileUris;
            LoggingConfig = loggingConfig;
            MainClass = mainClass;
            MainJarFileUri = mainJarFileUri;
            Properties = properties;
        }
    }
}
