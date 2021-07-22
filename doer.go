package httpclient

import (
	"net/http"
)

// DefaultDoer returns a new http.Client with similar default values to
// http.Client, but with a non-shared Transport.
// Intended to be reused for requests to the same host.
func DefaultDoer() *http.Client {
	return &http.Client{
		Transport: DefaultTransport(),
	}
}

// Doer is an interface supported by http.Client.
// Use this interface whenever you need to abstract from http.Client implementation and to use decorated client.
type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

// DoerFunc is a function type that implements the Doer interface.
type DoerFunc func(*http.Request) (*http.Response, error)

// Do supports Doer interface
func (f DoerFunc) Do(r *http.Request) (*http.Response, error) {
	return f(r)
}

// DoerDecoratorFunc wraps a Doer with extra behaviour.
type DoerDecoratorFunc func(Doer) Doer

// DecorateDoer decorates a Doer c with all the given Decorators, in order.
func DecorateDoer(c Doer, ds ...DoerDecoratorFunc) Doer {
	result := c
	for _, decorate := range ds {
		result = decorate(result)
	}
	return result
}

// DecorateDoerByDefault decorates a Doer c with default set of decorators:
// 1. Error
func DecorateDoerByDefault(c Doer, serviceName string) Doer {
	return DecorateDoer(c,
		ErrorDoerDecorator(serviceName),
	)
}
