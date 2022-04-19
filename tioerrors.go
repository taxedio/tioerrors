package tioerrors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type RestErr interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

type restErr struct {
	ErrMessage string        `json:"message"`
	ErrStatus  int           `json:"status"`
	ErrError   string        `json:"error"`
	ErrCauses  []interface{} `json:"causes"`
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
	switch r.Status() {
	// 400s
	case http.StatusBadRequest: // 400
		r.ErrError = "bad_request"
	case http.StatusPaymentRequired: // 402
		r.ErrError = "payment_required"
	case http.StatusForbidden: // 403
		r.ErrError = "forbidden"
	case http.StatusNotFound: // 404
		r.ErrError = "not_found"
	case http.StatusMethodNotAllowed: // 405
		r.ErrError = "method_not_allowed"
	case http.StatusNotAcceptable: // 406
		r.ErrError = "method_not_allowed"
	case http.StatusProxyAuthRequired: // 407
		r.ErrError = "proxy_auth_required"
	case http.StatusRequestTimeout: // 408
		r.ErrError = "request_timeout"
	case http.StatusConflict: // 409
		r.ErrError = "conflict"
	case http.StatusGone: // 410
		r.ErrError = "gone"
	case http.StatusLengthRequired: // 411
		r.ErrError = "length_required"
	case http.StatusPreconditionFailed: // 412
		r.ErrError = "precondition_failed"
	case http.StatusRequestEntityTooLarge: // 413
		r.ErrError = "payload_too_large"
	case http.StatusRequestURITooLong: // 414
		r.ErrError = "uri_too_long"
	case http.StatusUnsupportedMediaType: // 415
		r.ErrError = "unsupported_media_type"
	case http.StatusRequestedRangeNotSatisfiable: // 416
		r.ErrError = "range_not_satisfiable"
	case http.StatusExpectationFailed: // 417
		r.ErrError = "expectation_failed"
	case http.StatusTeapot: // 418
		r.ErrError = "im_a_teapot"
	case http.StatusMisdirectedRequest: // 421
		r.ErrError = "misdirected_request"
	case http.StatusUnprocessableEntity: // 422
		r.ErrError = "unprocessable_entity"
	case http.StatusLocked: // 423
		r.ErrError = "locked"
	case http.StatusFailedDependency: // 424
		r.ErrError = "failed_dependency"
	case http.StatusTooEarly: // 425
		r.ErrError = "too_early"
	case http.StatusUpgradeRequired: // 426
		r.ErrError = "upgrade_required"
	case http.StatusPreconditionRequired: // 428
		r.ErrError = "precondition_required"
	case http.StatusTooManyRequests: // 429
		r.ErrError = "too_many_requests"
	case http.StatusRequestHeaderFieldsTooLarge: // 431
		r.ErrError = "request_header_fields_too_large"
	case http.StatusUnavailableForLegalReasons: // 451
		r.ErrError = "unavailable_for_legal_reasons"
	// 500s
	case http.StatusInternalServerError: // 500
		r.ErrError = "internal_server_error"
	case http.StatusNotImplemented: // 501
		r.ErrError = "not_implemented"
	case http.StatusBadGateway: // 502
		r.ErrError = "bad_gateway"
	case http.StatusServiceUnavailable: // 503
		r.ErrError = "service_unavailable"
	case http.StatusGatewayTimeout: // 504
		r.ErrError = "gateway_timeout"
	case http.StatusHTTPVersionNotSupported: // 505
		r.ErrError = "http_version_not_supported"
	case http.StatusVariantAlsoNegotiates: // 506
		r.ErrError = "variant_also_negotiates"
	case http.StatusInsufficientStorage: // 507
		r.ErrError = "insufficent_storage"
	case http.StatusLoopDetected: // 508
		r.ErrError = "loop_detected"
	case http.StatusNotExtended: // 510
		r.ErrError = "not_extended"
	case http.StatusNetworkAuthenticationRequired: // 511
		r.ErrError = "network_authentication_required"
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
