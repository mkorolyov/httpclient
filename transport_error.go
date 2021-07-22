package httpclient

import (
	"fmt"
	"net/http"
)

// ErrorDoerDecorator is a net/http.Transport decorator which wraps raised error with requested host context.
func ErrorTransportDecorator() TransportDecoratorFunc {
	return func(rt http.RoundTripper) http.RoundTripper {
		return TransportFunc(func(r *http.Request) (*http.Response, error) {
			resp, err := rt.RoundTrip(r)
			if err != nil {
				return nil, fmt.Errorf("httpclient_transport_error [%s]: %s", r.URL.Host, err)
			}

			return resp, nil
		})
	}
}
