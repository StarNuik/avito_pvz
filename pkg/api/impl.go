package api

import (
	"context"

	"github.com/starnuik/avito_pvz/pkg/usecase"
)

var _ StrictServerInterface = (*impl)(nil)

type impl struct {
	usecase usecase.Usecase
}

func New(usecase usecase.Usecase) ServerInterface {
	strictHandler := impl{
		usecase,
	}
	handler := NewStrictHandler(&strictHandler, []StrictMiddlewareFunc{})
	return handler
}

// GetPvz implements StrictServerInterface.
func (i *impl) GetPvz(ctx context.Context, request GetPvzRequestObject) (GetPvzResponseObject, error) {
	panic("unimplemented")
}

// PostDummyLogin implements StrictServerInterface.
func (i *impl) PostDummyLogin(ctx context.Context, request PostDummyLoginRequestObject) (PostDummyLoginResponseObject, error) {
	panic("unimplemented")
}

// PostLogin implements StrictServerInterface.
func (i *impl) PostLogin(ctx context.Context, request PostLoginRequestObject) (PostLoginResponseObject, error) {
	panic("unimplemented")
}

// PostProducts implements StrictServerInterface.
func (i *impl) PostProducts(ctx context.Context, request PostProductsRequestObject) (PostProductsResponseObject, error) {
	panic("unimplemented")
}

// PostPvz implements StrictServerInterface.
func (i *impl) PostPvz(ctx context.Context, request PostPvzRequestObject) (PostPvzResponseObject, error) {
	panic("unimplemented")
}

// PostPvzPvzIdCloseLastReception implements StrictServerInterface.
func (i *impl) PostPvzPvzIdCloseLastReception(ctx context.Context, request PostPvzPvzIdCloseLastReceptionRequestObject) (PostPvzPvzIdCloseLastReceptionResponseObject, error) {
	panic("unimplemented")
}

// PostPvzPvzIdDeleteLastProduct implements StrictServerInterface.
func (i *impl) PostPvzPvzIdDeleteLastProduct(ctx context.Context, request PostPvzPvzIdDeleteLastProductRequestObject) (PostPvzPvzIdDeleteLastProductResponseObject, error) {
	panic("unimplemented")
}

// PostReceptions implements StrictServerInterface.
func (i *impl) PostReceptions(ctx context.Context, request PostReceptionsRequestObject) (PostReceptionsResponseObject, error) {
	panic("unimplemented")
}

// PostRegister implements StrictServerInterface.
func (i *impl) PostRegister(ctx context.Context, request PostRegisterRequestObject) (PostRegisterResponseObject, error) {
	panic("unimplemented")
}
