// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLiveDomainResponseBody
	GetRequestId() *string
}

type DeleteLiveDomainResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 94E3559F-7B6A-4A5E-AFFD-44E2A208A249
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveDomainResponseBody) SetRequestId(v string) *DeleteLiveDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
