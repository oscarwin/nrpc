package encoding

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CodecTestSuite struct {
	suite.Suite
	c Codec
}

func (s *CodecTestSuite) SetupTest() {
	s.c = GetCodec(CodecName)
}

func (s *CodecTestSuite) TestEncode() {
	data := []byte("hello")
	reqTypeOp := WithReqType(1)
	encodeData, err := s.c.Encode(data, reqTypeOp)
	assert.Nil(s.T(), err)
	s.T().Logf("%s", encodeData)
	decodeData, err := s.c.Decode(encodeData)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), "hello", string(decodeData))
}

func TestCodecSuite(t *testing.T) {
	suite.Run(t, new(CodecTestSuite))
}
