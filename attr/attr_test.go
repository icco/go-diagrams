package attr

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type AttrSuite struct {
	suite.Suite
}

func TestRunAttrSuite(t *testing.T) {
	suite.Run(t, new(AttrSuite))
}

func (s AttrSuite) TestNewAttributes() {
	attrs := New()

	attrs.Set(Label("test"), Shape("box"))

	s.Equal("test", attrs["label"])
	s.Equal("box", attrs["shape"])
}


func (s AttrSuite) TestPeriodAttributes() {
	attrs := New()

	attrs.Set(Label("example.com"))

	s.Equal("example.com", attrs["label"])
}
