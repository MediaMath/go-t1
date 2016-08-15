package t1

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
)

const (
	headerMasheryError        = "X-Mashery-Error-Code"
	headerMasheryDetail       = "X-Error-Detail-Header"
	masheryErrorCodeRateLimit = "ERR_403_DEVELOPER_OVER_QPS"
)

// ErrorResponse reports one or more errors caused by an API request.
type ErrorResponse struct {
	Response *http.Response // HTTP response that caused this error
	Message  string         `json:"message"` // error message
	Errors   []Error        `json:"errors"`  // more detail on individual errors
	Meta     Meta           `json:"meta"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v %+v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Message, r.Errors)
}

// RateLimitError occurs when API returns 403 Forbidden response with an
// error header signifying over QPS
type RateLimitError struct {
	RetryAt  time.Time
	Response *http.Response // HTTP response that caused this error
	Message  string         `json:"message"` // error message
}

func (r *RateLimitError) Error() string {
	return fmt.Sprintf("%v %v: %d %v; rate reset in %v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Message, r.RetryAt.Sub(time.Now()))
}

// Error reports more details on an individual error in an ErrorResponse.
// These are example validation error codes:
type Error struct {
	Resource string `json:"resource"` // resource on which the error occurred
	Type     string `json:"type"`     // Type of error (e.g. field-error)
	Field    string `json:"field"`    // field on which the error occurred
	Code     string `json:"code"`     // validation error code
	Message  string `json:"message"`  // Message describing the error. Errors with Code == "custom" will always have this set.
}

// Error tells you what error was caused by what field and resource
func (e *Error) Error() string {
	return fmt.Sprintf("%v error caused by %v field on %v resource",
		e.Code, e.Field, e.Resource)
}

// CheckResponse checks the API response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range. API error responses are expected to have either no response
// body, or a JSON response body that maps to ErrorResponse. Any other
// response body will be silently ignored.
//
// The error type will be *RateLimitError for rate limit exceeded errors.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	switch cType := getContentType(r); cType {
	case mediaTypeMashery:
		return parseMasheryError(r)
	case mediaTypeJSON:
		return parseAdamaError(r)
	default:
		return fmt.Errorf("unknown content type: %s", cType)
	}
}

func parseMasheryError(r *http.Response) error {
	var mr masheryResponse
	if err := xml.NewDecoder(r.Body).Decode(&mr); err != nil {
		return err
	}
	if r.StatusCode == http.StatusForbidden && r.Header.Get(headerMasheryError) == masheryErrorCodeRateLimit {
		return &RateLimitError{
			RetryAt:  parseRateLimit(r),
			Response: r,
			Message:  mr.Message,
		}
	}
	return &ErrorResponse{
		Response: r,
		Message:  mr.Message,
	}
}

func parseAdamaError(r *http.Response) error {
	var er ErrorResponse
	if err := json.NewDecoder(r.Body).Decode(&er); err != nil {
		return err
	}
	if er.Message != "" {
		return &er
	}
	if len(er.Errors) == 1 {
		er.Message = er.Errors[0].Message
	} else if er.Meta.Status != "" {
		er.Message = er.Meta.Status
	}
	return &er
}

func getContentType(r *http.Response) string {
	return r.Header.Get("Content-Type")
}
