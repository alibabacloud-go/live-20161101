// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveDomainStagingConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetLiveDomainStagingConfigResponseBody
	GetRequestId() *string
}

type SetLiveDomainStagingConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetLiveDomainStagingConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetLiveDomainStagingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetLiveDomainStagingConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetLiveDomainStagingConfigResponseBody) SetRequestId(v string) *SetLiveDomainStagingConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetLiveDomainStagingConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
