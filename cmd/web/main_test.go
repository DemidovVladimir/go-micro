package main

import (
	"testing"

	mocks "github.com/VladimirDemidov/webservice/mocks/cmd/web"
)

func TestStub(t *testing.T) {
	pr := new(mocks.Proxy)
	pr.On("Passthrough", "test").Return(true)
	my := prpr{pr}

	my.validate("test")
	pr.AssertExpectations(t)
}
