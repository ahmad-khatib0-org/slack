// This package provides aliases for the contents of "github.com/stretchr/testify/mock".
// Because external packages can't import our vendored dependencies, this is necessary
// for them to be able to fully utilize the plugintest package.
package mock

import (
	"github.com/stretchr/testify/mock"
)

const (
	Anything = mock.Anything
)

type (
	Arguments              = mock.Arguments
	AnythingOfTypeArgument = mock.AnythingOfTypeArgument
	Call                   = mock.Call
	Mock                   = mock.Mock
	TestingT               = mock.TestingT
)

func AnythingOfType(t string) AnythingOfTypeArgument {
	return mock.AnythingOfType(t)
}

func AssertExpectationsForObjects(t TestingT, testObjects ...any) bool {
	return mock.AssertExpectationsForObjects(t, testObjects...)
}

func MatchedBy(fn any) any {
	return mock.MatchedBy(fn)
}
