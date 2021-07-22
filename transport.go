package httpclient

import (
	"net"
	"net/http"
	"runtime"
	"time"
)

// DefaultTransport returns a new http.Transport with similar default
// values to http.DefaultTransport.
// Intended to be reused for requests to the same host.
func DefaultTransport() *http.Transport {
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		ForceAttemptHTTP2:     true,
		MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
	}
	return transport
}

// TransportFunc is a function type that implements the http.RoundTripper interface.
type TransportFunc func(*http.Request) (*http.Response, error)

// RoundTrip supports http.RoundTripper interface
func (f TransportFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

// TransportDecoratorFunc wraps a http.RoundTripper with extra behaviour.
type TransportDecoratorFunc func(http.RoundTripper) http.RoundTripper

// DecorateTransport decorates a http.RoundTripper c with all the given Decorators, in order.
func DecorateTransport(rt http.RoundTripper, ds ...TransportDecoratorFunc) http.RoundTripper {
	result := rt
	for _, decorate := range ds {
		result = decorate(result)
	}
	return result
}

// DecorateDoerTransport decorates internal Transport of http.Client
func DecorateDoerTransport(c *http.Client, ds ...TransportDecoratorFunc) *http.Client {
	if c == nil {
		c = http.DefaultClient
	}

	transport := c.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}

	for _, decorate := range ds {
		transport = decorate(transport)
	}

	resultClient := *c
	resultClient.Transport = transport

	return &resultClient
}

// DecorateDoerTransportByDefault decorates a http.Client c with default set of decorators:
// 1. Error
func DecorateDoerTransportByDefault(c *http.Client) *http.Client {
	return DecorateDoerTransport(c,
		ErrorTransportDecorator(),
	)
}
