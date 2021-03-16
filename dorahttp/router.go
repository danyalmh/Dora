package dorahttp

// Router is a represents the router handling HTTP.
type Router struct {
	tree *Tree
}

// NewRouter creates a new router.
func NewRouter() *Router {
	return &Router{
		tree: NewTree(),
	}
}

// GET sets a route for GET method.
func (r *Router) GET(path string, handler DoraHandler) {
	r.Register(MethodGet, path, handler)
}

// POST sets a route for POST method.
func (r *Router) POST(path string, handler DoraHandler) {
	r.Register(MethodPost, path, handler)
}

// PUT sets a route for PUT method.
func (r *Router) PUT(path string, handler DoraHandler) {
	r.Register(MethodPut, path, handler)
}

// PATCH sets a route for PATCH method.
func (r *Router) PATCH(path string, handler DoraHandler) {
	r.Register(MethodPatch, path, handler)
}

// DELETE sets a route for DELETE method.
func (r *Router) DELETE(path string, handler DoraHandler) {
	r.Register(MethodDelete, path, handler)
}

// OPTION sets a route for OPTION method.
func (r *Router) OPTION(path string, handler DoraHandler) {
	r.Register(MethodOptions, path, handler)
}

// Register registers a route.
func (r *Router) Register(method string, path string, handler DoraHandler) {
	r.tree.Insert(method, path, handler)
}

// ServeHTTP dispatches the request to the handler whose
// pattern most closely matches the request URL.
func (r *Router) ServeHTTP(dctx *Dctx) []byte {
	method := dctx.Method
	path := dctx.Path
	dctx.Params = nil
	result, err := r.tree.Search(method, path)

	if err != nil {
		return dctx.Response(StatusNotImplemented, "")
	}

	if result.params == nil || len(*result.params) == 0 {
		result.handler(dctx)
	}

	dctx.Params = *result.params
	return result.handler(dctx)
}
