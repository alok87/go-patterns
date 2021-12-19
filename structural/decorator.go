package structural

import (
	"fmt"
	"net/http"
	"strings"
)

/*

Summary:
Decaorator pattern helps in achieving single responsibility
Decorator pattern helps in adding additional pre and post
functionality to some function without actually touching the actual function
and making it complicated.

Example:
Basic examples demonstrated
then Real world examples of middlewares demostrated.

Benefit: Readability, Maintainability and Single Responsiblity.
1. It can be used when some pattern of pre and post functionality is required
over some functionality. Instead of making changes to original function we
decorate it.
2. It helps from complicating existing functions and helps them be small
3. Unit testing functionalities which are small is easy. Decorators can keep
things small.
*/

/* Basic Examples below */

type Hello func(s string) string

// DecorateV1 decorates hello function with some pre post prints
// v1 style: function as a type
// Decorated functon is returned as a new function and
// then called with parameters
func DecorateV1(f Hello) Hello {
	return func(s string) string {
		fmt.Printf("pre ")
		v := f(s)
		fmt.Printf("post:%v", v)
		return v
	}
}

// DecorateV2 decorates function
// v2 style: functions with complete signatures (more readable?)
// Decorated functon is returned as a new function and
// then called with parameters
func DecorateV2(f func(s string) string) func(string) string {
	return func(s string) string {
		fmt.Printf("pre ")
		v := f(s)
		fmt.Printf("post:%v", v)
		return v
	}
}

// DecorateV3 is another style to wrap functions in one call
// Decorated functon is passed as a parameter and also the parameters
// required for the core function is passed along. In this call only
// everything happens.
func DecorateV3(f func(s string) string, param string) string {
	fmt.Printf("pre ")
	v := f(param)
	fmt.Printf("post:%v", v)
	return v
}

/* Real work examples below */

// NewRelicMiddleWare decorates newrelic: middleware
// Uses DecoratorV1 style
func NewRelicMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(rw http.ResponseWriter, r *http.Request) {
			fmt.Printf("newrelic start ")
			next.ServeHTTP(rw, r)
			fmt.Printf("newrelic end ")
		},
	)
}

// GraylogMiddleWare is the Decaorator for graylog http handler
// request logging middle ware
// Uses DecoratorV1 style
func GraylogMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(rw http.ResponseWriter, r *http.Request) {
			fmt.Printf("graylogging start ")
			next.ServeHTTP(rw, r)
			fmt.Printf("graylogging end ")
		},
	)
}

// AuthorizeMiddleware decorates and terminates early if the
// the request is not having the right token
// Uses DecoratorV1 style
func AuthorizeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(rw http.ResponseWriter, r *http.Request) {
			value := r.Header.Get("Authorization")
			token := strings.Split(value, "Bearer ")
			if len(token) != 2 || token[1] == "usertoken" {
				http.Error(rw, "access denied", http.StatusForbidden)
				return
			}
			fmt.Printf("access allowed ")
			next.ServeHTTP(rw, r)
		},
	)
}
