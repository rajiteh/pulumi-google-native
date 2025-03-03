// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudTasks.V2Beta3.Outputs
{

    /// <summary>
    /// App Engine HTTP queue. The task will be delivered to the App Engine application hostname specified by its AppEngineHttpQueue and AppEngineHttpRequest. The documentation for AppEngineHttpRequest explains how the task's host URL is constructed. Using AppEngineHttpQueue requires [`appengine.applications.get`](https://cloud.google.com/appengine/docs/admin-api/access-control) Google IAM permission for the project and the following scope: `https://www.googleapis.com/auth/cloud-platform`
    /// </summary>
    [OutputType]
    public sealed class AppEngineHttpQueueResponse
    {
        /// <summary>
        /// Overrides for the task-level app_engine_routing. If set, `app_engine_routing_override` is used for all tasks in the queue, no matter what the setting is for the task-level app_engine_routing.
        /// </summary>
        public readonly Outputs.AppEngineRoutingResponse AppEngineRoutingOverride;

        [OutputConstructor]
        private AppEngineHttpQueueResponse(Outputs.AppEngineRoutingResponse appEngineRoutingOverride)
        {
            AppEngineRoutingOverride = appEngineRoutingOverride;
        }
    }
}
