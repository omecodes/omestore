package router

import "context"

type ObjectsRouter interface {
	// GetRoute returns a sequence of handler
	GetRoute(opts ...RouteOption) ObjectsHandler
}

type ObjectsRouterProvider interface {
	//GetRouter returns a router
	GetRouter(ctx context.Context) ObjectsRouter
}

type ObjectsRouterProvideFunc func(ctx context.Context) ObjectsRouter

func (f ObjectsRouterProvideFunc) GetRouter(ctx context.Context) ObjectsRouter {
	return f(ctx)
}

type ObjectsRouteProviderFunc func(opts ...RouteOption) ObjectsHandler

func (f ObjectsRouteProviderFunc) GetRoute(opts ...RouteOption) ObjectsHandler {
	return f(opts...)
}

func DefaultRouter() ObjectsRouter {
	return ObjectsRouteProviderFunc(getObjectsRoute)
}

func getObjectsRoute(opts ...RouteOption) (handler ObjectsHandler) {
	routes := routesOptions{}

	for _, o := range opts {
		o(&routes)
	}

	handler = &ObjectsExecHandler{}

	if !routes.skipPolicies {
		handler = &ObjectsPolicyHandler{ObjectsBaseHandler: ObjectsBaseHandler{
			next: handler,
		}}
	}

	if !routes.skipParams {
		handler = &ObjectsParamsHandler{
			ObjectsBaseHandler{next: handler},
		}
	}
	return
}
