package tioerrors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	StatusMsg = map[int]string{
		http.StatusBadRequest:                    "bad_request",
		http.StatusPaymentRequired:               "payment_required",
		http.StatusForbidden:                     "forbidden",
		http.StatusNotFound:                      "not_found",
		http.StatusMethodNotAllowed:              "method_not_allowed",
		http.StatusNotAcceptable:                 "method_not_allowed",
		http.StatusProxyAuthRequired:             "proxy_auth_required",
		http.StatusRequestTimeout:                "request_timeout",
		http.StatusConflict:                      "conflict",
		http.StatusGone:                          "gone",
		http.StatusLengthRequired:                "length_required",
		http.StatusPreconditionFailed:            "precondition_failed",
		http.StatusRequestEntityTooLarge:         "payload_too_large",
		http.StatusRequestURITooLong:             "uri_too_long",
		http.StatusUnsupportedMediaType:          "unsupported_media_type",
		http.StatusRequestedRangeNotSatisfiable:  "range_not_satisfiable",
		http.StatusExpectationFailed:             "expectation_failed",
		http.StatusTeapot:                        "im_a_teapot",
		http.StatusMisdirectedRequest:            "misdirected_request",
		http.StatusUnprocessableEntity:           "unprocessable_entity",
		http.StatusLocked:                        "locked",
		http.StatusFailedDependency:              "failed_dependency",
		http.StatusTooEarly:                      "too_early",
		http.StatusUpgradeRequired:               "upgrade_required",
		http.StatusPreconditionRequired:          "precondition_required",
		http.StatusTooManyRequests:               "too_many_requests",
		http.StatusRequestHeaderFieldsTooLarge:   "request_header_fields_too_large",
		http.StatusUnavailableForLegalReasons:    "unavailable_for_legal_reasons",
		http.StatusInternalServerError:           "internal_server_error",
		http.StatusNotImplemented:                "not_implemented",
		http.StatusBadGateway:                    "bad_gateway",
		http.StatusServiceUnavailable:            "service_unavailable",
		http.StatusGatewayTimeout:                "gateway_timeout",
		http.StatusHTTPVersionNotSupported:       "http_version_not_supported",
		http.StatusVariantAlsoNegotiates:         "variant_also_negotiates",
		http.StatusInsufficientStorage:           "insufficient_storage",
		http.StatusLoopDetected:                  "loop_detected",
		http.StatusNotExtended:                   "not_extended",
		http.StatusNetworkAuthenticationRequired: "network_authentication_required"}
)

type RestErr interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

type restErr struct {
	ErrMessage string        `json:"message" xml:"message" bson:"message" protobuf:"message" yaml:"message" datastore:"message" schema:"message" asn:"message" `
	ErrStatus  int           `json:"status" xml:"status" bson:"status" protobuf:"status" yaml:"status" datastore:"status" schema:"status" asn:"status" `
	ErrError   string        `json:"error" xml:"error" bson:"error" protobuf:"error" yaml:"error" datastore:"error" schema:"error" asn:"error" `
	ErrCauses  []interface{} `json:"causes" xml:"causes" bson:"causes" protobuf:"causes" yaml:"causes" datastore:"causes" schema:"causes" asn:"causes" `
}

func (e restErr) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: %v",
		e.ErrMessage, e.ErrStatus, e.ErrError, e.ErrCauses)
}

func (e restErr) Message() string {
	return e.ErrMessage
}

func (e restErr) Status() int {
	return e.ErrStatus
}

func (e restErr) Causes() []interface{} {
	return e.ErrCauses
}

func HttpStatusMaker(r restErr) *restErr {
	if errMsg, ok := StatusMsg[r.ErrStatus]; ok {
		r.ErrError = errMsg
	}
	return &r
}

func NewRestError(message string, status int, causes []interface{}) restErr {
	r := restErr{
		ErrMessage: message,
		ErrStatus:  status,
		ErrCauses:  causes,
	}
	return *HttpStatusMaker(r)
}

func NewRestErrorFromBytes(bytes []byte) (RestErr, error) {
	var apiErr restErr
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return apiErr, nil
}
