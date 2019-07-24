package product

import "github.com/nqmt/go-service/util/errorx"

var (
	ErrUnableCreateProduct = errorx.DefineInternalServerError("UNABLE_CREATE_PRODUCT", "unable create product")
)

