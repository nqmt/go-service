package product

import (
	"github.com/nqmt/goerror"
)

var (
	ErrUnableCreateProduct = goerror.DefineInternalServerError("UNABLE_CREATE_PRODUCT", "unable create product")
)
