package dto

// var _ StrictServerInterface = (*impl)(nil)

// type impl struct {
// 	usecase     usecase.Usecase
// 	tokenParser token.Parser
// }

// func New(usecase usecase.Usecase, tokenParser token.Parser) (ServerInterface, error) {
// 	strictHandler := impl{
// 		usecase, tokenParser,
// 	}

// 	spec, err := GetSwagger()
// 	if err != nil {
// 		return nil, err
// 	}

// 	authMiddleware := middleware.OapiRequestValidatorWithOptions(spec,
// 		&middleware.Options{
// 			Options: openapi3filter.Options{
// 				AuthenticationFunc: nil,
// 			},
// 		})

// 	handler := NewStrictHandler(&strictHandler, []StrictMiddlewareFunc{authMiddleware})
// 	return handler
// }

// // GetPvz implements StrictServerInterface.
// func (i *impl) GetPvz(ctx context.Context, req GetPvzRequestObject) (GetPvzResponseObject, error) {
// 	panic("unimplemented")
// }

// // PostDummyLogin implements StrictServerInterface.
// func (i *impl) PostDummyLogin(ctx context.Context, req PostDummyLoginRequestObject) (PostDummyLoginResponseObject, error) {
// 	role, err := entity.ParseUserRole(string(req.Body.Role))
// 	if err != nil {
// 		return PostDummyLogin400JSONResponse{err.Error()}, nil
// 	}

// 	payload := i.usecase.DummyLogin(role)

// 	token, err := i.tokenParser.Pack(payload)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return PostDummyLogin200JSONResponse(token), nil
// }

// // PostLogin implements StrictServerInterface.
// func (i *impl) PostLogin(ctx context.Context, req PostLoginRequestObject) (PostLoginResponseObject, error) {
// 	panic("unimplemented")
// }

// // PostProducts implements StrictServerInterface.
// func (i *impl) PostProducts(ctx context.Context, req PostProductsRequestObject) (PostProductsResponseObject, error) {
// 	panic("unimplemented")
// }

// // PostPvz implements StrictServerInterface.
// func (i *impl) PostPvz(ctx context.Context, req PostPvzRequestObject) (PostPvzResponseObject, error) {
// 	panic("unimplemented")
// }

// // PostPvzPvzIdCloseLastReception implements StrictServerInterface.
// func (i *impl) PostPvzPvzIdCloseLastReception(ctx context.Context, req PostPvzPvzIdCloseLastReceptionRequestObject) (PostPvzPvzIdCloseLastReceptionResponseObject, error) {
// 	panic("unimplemented")
// }

// // PostPvzPvzIdDeleteLastProduct implements StrictServerInterface.
// func (i *impl) PostPvzPvzIdDeleteLastProduct(ctx context.Context, req PostPvzPvzIdDeleteLastProductRequestObject) (PostPvzPvzIdDeleteLastProductResponseObject, error) {
// 	panic("unimplemented")
// }

// // PostReceptions implements StrictServerInterface.
// func (i *impl) PostReceptions(ctx context.Context, req PostReceptionsRequestObject) (PostReceptionsResponseObject, error) {
// 	panic("unimplemented")
// }

// // PostRegister implements StrictServerInterface.
// func (i *impl) PostRegister(ctx context.Context, req PostRegisterRequestObject) (PostRegisterResponseObject, error) {
// 	panic("unimplemented")
// }
