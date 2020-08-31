package httpserver

// Option is the server options.
type Option func(*Server)

// WithServices registers the server services.
func WithServices(services ...Service) Option {
	return func(srv *Server) {
		for _, service := range services {
			service(srv)
		}
	}
}
