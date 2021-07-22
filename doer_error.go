package httpclient

import (
	"fmt"
	"net/http"
)

// ErrorDoerDecorator is a Doer decorator which wraps raised error with service name context.
func ErrorDoerDecorator(serviceName string) DoerDecoratorFunc {
	return func(c Doer) Doer {
		return DoerFunc(func(r *http.Request) (*http.Response, error) {
			resp, err := c.Do(r)
			if err != nil {
				err = fmt.Errorf("httpclient_doer_error [%s]: %w", serviceName, err)
			}

			return resp, err
		})
	}
}
