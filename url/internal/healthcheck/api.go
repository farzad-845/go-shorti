package healthcheck

import routing "github.com/go-ozzo/ozzo-routing/v2"

// RegisterHandlers registers the handlers that perform health checks.
func RegisterHandlers(r *routing.Router, version string) {
	r.To("GET,HEAD", "/healthcheck", healthCheck(version))
}

// healthCheck responds to a healthCheck request.
func healthCheck(version string) routing.Handler {
	return func(c *routing.Context) error {
		return c.Write("OK " + version)
	}
}
