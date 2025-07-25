// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveRecordVodConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *AddLiveRecordVodConfigRequest
	GetAppName() *string
	SetAutoCompose(v string) *AddLiveRecordVodConfigRequest
	GetAutoCompose() *string
	SetComposeVodTranscodeGroupId(v string) *AddLiveRecordVodConfigRequest
	GetComposeVodTranscodeGroupId() *string
	SetCycleDuration(v int32) *AddLiveRecordVodConfigRequest
	GetCycleDuration() *int32
	SetDomainName(v string) *AddLiveRecordVodConfigRequest
	GetDomainName() *string
	SetOnDemand(v int32) *AddLiveRecordVodConfigRequest
	GetOnDemand() *int32
	SetOwnerId(v int64) *AddLiveRecordVodConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AddLiveRecordVodConfigRequest
	GetRegionId() *string
	SetStorageLocation(v string) *AddLiveRecordVodConfigRequest
	GetStorageLocation() *string
	SetStreamName(v string) *AddLiveRecordVodConfigRequest
	GetStreamName() *string
	SetVodTranscodeGroupId(v string) *AddLiveRecordVodConfigRequest
	GetVodTranscodeGroupId() *string
}

type AddLiveRecordVodConfigRequest struct {
	// The name of the application to which the live stream belongs. You can view the application name on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page of the ApsaraVideo Live console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testName
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Specifies whether to enable automatic merging. If you set the value to **ON**, automatic merging is enabled and the ComposeVodTranscodeGroupId parameter is required. If you do not specify this parameter, automatic merging is disabled.
	//
	// >  If you enable automatic merging, the VOD files that are created from live streams are automatically merged by using the editing and production feature of ApsaraVideo VOD. For information about the billing of the feature, see [Billing of value-added services](https://help.aliyun.com/document_detail/188310.html).
	//
	// example:
	//
	// ON
	AutoCompose *string `json:"AutoCompose,omitempty" xml:"AutoCompose,omitempty"`
	// The ID of the transcoding template group in ApsaraVideo VOD that is used to transcode the video file. The video file is generated by merging the VOD files created from live streams.
	//
	// >
	//
	// 	- This parameter is required if you set the AutoCompose parameter to ON.
	//
	// 	- For more information about automatic merging and transcoding, see [FAQ about Live-to-VOD](https://help.aliyun.com/document_detail/99726.html).
	//
	// 	- For information about the billing of transcoding in ApsaraVideo VOD, see [Billing of basic services](https://help.aliyun.com/document_detail/188308.html).
	//
	// example:
	//
	// *****
	ComposeVodTranscodeGroupId *string `json:"ComposeVodTranscodeGroupId,omitempty" xml:"ComposeVodTranscodeGroupId,omitempty"`
	// The recording cycle. Unit: seconds. Valid values: **300 to 21600**. Default value: **3600**.
	//
	// example:
	//
	// 300
	CycleDuration *int32 `json:"CycleDuration,omitempty" xml:"CycleDuration,omitempty"`
	// The main streaming domain.
	//
	// >  Make sure that ApsaraVideo VOD is activated in the same region as the live center of the streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Specifies whether to enable on-demand recording. Valid values:
	//
	// 	- **0*	- (default): disables on-demand recording.
	//
	// 	- **1**: enables on-demand recording.
	//
	// example:
	//
	// 0
	OnDemand *int32  `json:"OnDemand,omitempty" xml:"OnDemand,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The storage location.
	//
	// example:
	//
	// ****-tjptr2vatm.oss-cn-shanghai.aliyuncs.com
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	// The name of the live stream. You can view the stream name on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page of the ApsaraVideo Live console.
	//
	// example:
	//
	// stream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// The ID of the transcoding template group in ApsaraVideo VOD.
	//
	// This parameter is required.
	//
	// example:
	//
	// e2d796d3bb5fd8049d32bff62f94****
	VodTranscodeGroupId *string `json:"VodTranscodeGroupId,omitempty" xml:"VodTranscodeGroupId,omitempty"`
}

func (s AddLiveRecordVodConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s AddLiveRecordVodConfigRequest) GoString() string {
	return s.String()
}

func (s *AddLiveRecordVodConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *AddLiveRecordVodConfigRequest) GetAutoCompose() *string {
	return s.AutoCompose
}

func (s *AddLiveRecordVodConfigRequest) GetComposeVodTranscodeGroupId() *string {
	return s.ComposeVodTranscodeGroupId
}

func (s *AddLiveRecordVodConfigRequest) GetCycleDuration() *int32 {
	return s.CycleDuration
}

func (s *AddLiveRecordVodConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *AddLiveRecordVodConfigRequest) GetOnDemand() *int32 {
	return s.OnDemand
}

func (s *AddLiveRecordVodConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddLiveRecordVodConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddLiveRecordVodConfigRequest) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *AddLiveRecordVodConfigRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *AddLiveRecordVodConfigRequest) GetVodTranscodeGroupId() *string {
	return s.VodTranscodeGroupId
}

func (s *AddLiveRecordVodConfigRequest) SetAppName(v string) *AddLiveRecordVodConfigRequest {
	s.AppName = &v
	return s
}

func (s *AddLiveRecordVodConfigRequest) SetAutoCompose(v string) *AddLiveRecordVodConfigRequest {
	s.AutoCompose = &v
	return s
}

func (s *AddLiveRecordVodConfigRequest) SetComposeVodTranscodeGroupId(v string) *AddLiveRecordVodConfigRequest {
	s.ComposeVodTranscodeGroupId = &v
	return s
}

func (s *AddLiveRecordVodConfigRequest) SetCycleDuration(v int32) *AddLiveRecordVodConfigRequest {
	s.CycleDuration = &v
	return s
}

func (s *AddLiveRecordVodConfigRequest) SetDomainName(v string) *AddLiveRecordVodConfigRequest {
	s.DomainName = &v
	return s
}

func (s *AddLiveRecordVodConfigRequest) SetOnDemand(v int32) *AddLiveRecordVodConfigRequest {
	s.OnDemand = &v
	return s
}

func (s *AddLiveRecordVodConfigRequest) SetOwnerId(v int64) *AddLiveRecordVodConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *AddLiveRecordVodConfigRequest) SetRegionId(v string) *AddLiveRecordVodConfigRequest {
	s.RegionId = &v
	return s
}

func (s *AddLiveRecordVodConfigRequest) SetStorageLocation(v string) *AddLiveRecordVodConfigRequest {
	s.StorageLocation = &v
	return s
}

func (s *AddLiveRecordVodConfigRequest) SetStreamName(v string) *AddLiveRecordVodConfigRequest {
	s.StreamName = &v
	return s
}

func (s *AddLiveRecordVodConfigRequest) SetVodTranscodeGroupId(v string) *AddLiveRecordVodConfigRequest {
	s.VodTranscodeGroupId = &v
	return s
}

func (s *AddLiveRecordVodConfigRequest) Validate() error {
	return dara.Validate(s)
}
