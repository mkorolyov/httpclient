package httpclient

import (
	"net/http"
	"net/http/httputil"

	"github.com/mkorolyov/httpclient/log"
)

// DumpRequestFunc function to dump request. In most cases httputil.DumpRequestOut can be used
type DumpRequestFunc func(r *http.Request, body bool) ([]byte, error)

// DumpResponseFunc function to dump response. In most cases httputil.DumpResponse can be used
type DumpResponseFunc func(r *http.Response, body bool) ([]byte, error)

// LoggerFromRequestFunc is a function to get logger from request/to enrich logger with context meta data.
type LoggerFromRequestFunc func(r *http.Request) log.Logger

// LoggerTransportDecorator creates
func LoggerTransportDecorator(loggerFromRequestFn LoggerFromRequestFunc) TransportDecoratorFunc {
	return LoggerTransportDecoratorWithDumpers(httputil.DumpRequestOut, httputil.DumpResponse, loggerFromRequestFn)
}

// LoggerTransportDecoratorWithDumpers is a net/http.Transport decorator which logs response body to context`s logger
func LoggerTransportDecoratorWithDumpers(
	reqFn DumpRequestFunc,
	respFn DumpResponseFunc,
	loggerFromRequestFn LoggerFromRequestFunc,
) TransportDecoratorFunc {
	return func(rt http.RoundTripper) http.RoundTripper {
		return TransportFunc(func(r *http.Request) (*http.Response, error) {
			var (
				reqDump, respDump       []byte
				reqDumpErr, respDumpErr error
			)

			logger := loggerFromRequestFn(r)

			reqDump, reqDumpErr = reqFn(r, true)
			if reqDumpErr != nil {
				logger.Errorf("failed to dump request: %v", reqDumpErr)
			}

			resp, err := rt.RoundTrip(r)

			if resp != nil {
				respDump, respDumpErr = respFn(resp, true)
				if respDumpErr != nil {
					logger.Errorf("failed to dump response: %v", respDumpErr)
				}
			}

			logger.Tracef("\n--->>> request: %s\n <<<--- response: %s", reqDump, respDump)

			return resp, err
		})
	}
}
