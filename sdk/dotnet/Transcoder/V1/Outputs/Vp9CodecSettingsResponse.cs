// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Transcoder.V1.Outputs
{

    /// <summary>
    /// VP9 codec settings.
    /// </summary>
    [OutputType]
    public sealed class Vp9CodecSettingsResponse
    {
        /// <summary>
        /// The video bitrate in bits per second. The minimum value is 1,000. The maximum value is 480,000,000.
        /// </summary>
        public readonly int BitrateBps;
        /// <summary>
        /// Target CRF level. Must be between 10 and 36, where 10 is the highest quality and 36 is the most efficient compression. The default is 21. **Note:** This field is not supported.
        /// </summary>
        public readonly int CrfLevel;
        /// <summary>
        /// The target video frame rate in frames per second (FPS). Must be less than or equal to 120. Will default to the input frame rate if larger than the input frame rate. The API will generate an output FPS that is divisible by the input FPS, and smaller or equal to the target FPS. See [Calculating frame rate](https://cloud.google.com/transcoder/docs/concepts/frame-rate) for more information.
        /// </summary>
        public readonly double FrameRate;
        /// <summary>
        /// Select the GOP size based on the specified duration. The default is `3s`. Note that `gopDuration` must be less than or equal to [`segmentDuration`](#SegmentSettings), and [`segmentDuration`](#SegmentSettings) must be divisible by `gopDuration`.
        /// </summary>
        public readonly string GopDuration;
        /// <summary>
        /// Select the GOP size based on the specified frame count. Must be greater than zero.
        /// </summary>
        public readonly int GopFrameCount;
        /// <summary>
        /// The height of the video in pixels. Must be an even integer. When not specified, the height is adjusted to match the specified width and input aspect ratio. If both are omitted, the input height is used. For portrait videos that contain horizontal ASR and rotation metadata, provide the height, in pixels, per the horizontal ASR. The API calculates the width per the horizontal ASR. The API detects any rotation metadata and swaps the requested height and width for the output.
        /// </summary>
        public readonly int HeightPixels;
        /// <summary>
        /// Pixel format to use. The default is `yuv420p`. Supported pixel formats: - `yuv420p` pixel format - `yuv422p` pixel format - `yuv444p` pixel format - `yuv420p10` 10-bit HDR pixel format - `yuv422p10` 10-bit HDR pixel format - `yuv444p10` 10-bit HDR pixel format - `yuv420p12` 12-bit HDR pixel format - `yuv422p12` 12-bit HDR pixel format - `yuv444p12` 12-bit HDR pixel format
        /// </summary>
        public readonly string PixelFormat;
        /// <summary>
        /// Enforces the specified codec profile. The following profiles are supported: * `profile0` (default) * `profile1` * `profile2` * `profile3` The available options are [WebM-compatible](https://www.webmproject.org/vp9/profiles/). Note that certain values for this field may cause the transcoder to override other fields you set in the `Vp9CodecSettings` message.
        /// </summary>
        public readonly string Profile;
        /// <summary>
        /// Specify the `rate_control_mode`. The default is `vbr`. Supported rate control modes: - `vbr` - variable bitrate
        /// </summary>
        public readonly string RateControlMode;
        /// <summary>
        /// The width of the video in pixels. Must be an even integer. When not specified, the width is adjusted to match the specified height and input aspect ratio. If both are omitted, the input width is used. For portrait videos that contain horizontal ASR and rotation metadata, provide the width, in pixels, per the horizontal ASR. The API calculates the height per the horizontal ASR. The API detects any rotation metadata and swaps the requested height and width for the output.
        /// </summary>
        public readonly int WidthPixels;

        [OutputConstructor]
        private Vp9CodecSettingsResponse(
            int bitrateBps,

            int crfLevel,

            double frameRate,

            string gopDuration,

            int gopFrameCount,

            int heightPixels,

            string pixelFormat,

            string profile,

            string rateControlMode,

            int widthPixels)
        {
            BitrateBps = bitrateBps;
            CrfLevel = crfLevel;
            FrameRate = frameRate;
            GopDuration = gopDuration;
            GopFrameCount = gopFrameCount;
            HeightPixels = heightPixels;
            PixelFormat = pixelFormat;
            Profile = profile;
            RateControlMode = rateControlMode;
            WidthPixels = widthPixels;
        }
    }
}
