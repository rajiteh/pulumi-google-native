// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2Beta1.Inputs
{

    /// <summary>
    /// Browse Carousel Card for Actions on Google. https://developers.google.com/actions/assistant/responses#browsing_carousel
    /// </summary>
    public sealed class GoogleCloudDialogflowV2beta1IntentMessageBrowseCarouselCardArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Settings for displaying the image. Applies to every image in items.
        /// </summary>
        [Input("imageDisplayOptions")]
        public Input<Pulumi.GoogleNative.Dialogflow.V2Beta1.GoogleCloudDialogflowV2beta1IntentMessageBrowseCarouselCardImageDisplayOptions>? ImageDisplayOptions { get; set; }

        [Input("items", required: true)]
        private InputList<Inputs.GoogleCloudDialogflowV2beta1IntentMessageBrowseCarouselCardBrowseCarouselCardItemArgs>? _items;

        /// <summary>
        /// List of items in the Browse Carousel Card. Minimum of two items, maximum of ten.
        /// </summary>
        public InputList<Inputs.GoogleCloudDialogflowV2beta1IntentMessageBrowseCarouselCardBrowseCarouselCardItemArgs> Items
        {
            get => _items ?? (_items = new InputList<Inputs.GoogleCloudDialogflowV2beta1IntentMessageBrowseCarouselCardBrowseCarouselCardItemArgs>());
            set => _items = value;
        }

        public GoogleCloudDialogflowV2beta1IntentMessageBrowseCarouselCardArgs()
        {
        }
        public static new GoogleCloudDialogflowV2beta1IntentMessageBrowseCarouselCardArgs Empty => new GoogleCloudDialogflowV2beta1IntentMessageBrowseCarouselCardArgs();
    }
}
