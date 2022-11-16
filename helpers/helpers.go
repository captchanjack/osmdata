package helpers

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func JoinArrFloat64(a []float64, delim ...string) string {
	_delim := ","
	if len(delim) > 0 {
		_delim = delim[0]
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatFloat(v, 'f', -1, 64)
	}
	return strings.Join(b, _delim)
}

func JoinArrInt(a []int, delim ...string) string {
	_delim := ","
	if len(delim) > 0 {
		_delim = delim[0]
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}
	return strings.Join(b, _delim)
}

func FormatHTTPGetURL(baseURL string, queryParams map[string]string) string {
	base, err := url.Parse(baseURL)

	if err != nil {
		fmt.Println(fmt.Errorf("url was malformed: %s", err))
		return baseURL
	}

	params := url.Values{}

	for k, v := range queryParams {
		params.Add(k, v)
	}

	base.RawQuery = params.Encode()
	return base.String()
}

func ExecVaradicFunction(fn interface{}, args ...interface{}) []reflect.Value {
	// Convert arguments to reflect.Value
	vs := make([]reflect.Value, len(args))

	for n := range args {
		vs[n] = reflect.ValueOf(args[n])
	}

	// Recover message in case panic
	defer func() {
		if r := recover(); r != nil {
			panic(fmt.Sprintf("input args were incorrect %t, check if they match the signature of the varadic function: %t", args, fn))
		}
	}()

	// Call it. Note it panics if func is not callable or arguments don't match
	return reflect.ValueOf(fn).Call(vs)
}
