package tioerrors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

type codePair struct {
	i int
	s string
}

var (
	test1 = restErr{
		ErrMessage: "testmessage",
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "bad_request",
	}
	test1Bytes = []byte{123, 34, 109, 101, 115, 115, 97, 103, 101,
		34, 58, 34, 116, 101, 115, 116, 109, 101, 115, 115, 97, 103,
		101, 34, 44, 34, 115, 116, 97, 116, 117, 115, 34, 58, 52, 48,
		48, 44, 34, 101, 114, 114, 111, 114, 34, 58, 34, 98, 97, 100,
		95, 114, 101, 113, 117, 101, 115, 116, 34, 44, 34, 99, 97, 117,
		115, 101, 115, 34, 58, 110, 117, 108, 108, 125}

	test1BytesFail = []byte{123, 34, 109, 101, 115, 115, 97, 103, 101,
		34, 58, 34, 116, 101, 115, 116, 109, 101, 115, 115, 97, 103,
		101, 34, 44, 34, 115, 116, 97, 116, 117, 115, 34, 58, 52, 48,
		48, 44, 34, 101, 114, 114, 111, 114, 34, 58, 34, 98, 97, 100,
		95, 114, 101, 113, 117, 101}

	codePairTests = []codePair{
		{http.StatusBadRequest,
			"bad_request"},
		{http.StatusPaymentRequired,
			"payment_required"},
		{http.StatusForbidden, // 403 ,
			"forbidden"},
		{http.StatusNotFound, // 404,
			"not_found"},
		{http.StatusMethodNotAllowed, // 405,
			"method_not_allowed"},
		{http.StatusNotAcceptable, // 406,
			"method_not_allowed"},
		{http.StatusProxyAuthRequired, // 407,
			"proxy_auth_required"},
		{http.StatusRequestTimeout, // 408,
			"request_timeout"},
		{http.StatusConflict, // 409,
			"conflict"},
		{http.StatusGone, // 410,
			"gone"},
		{http.StatusLengthRequired, // 411,
			"length_required"},
		{http.StatusPreconditionFailed, // 412,
			"precondition_failed"},
		{http.StatusRequestEntityTooLarge, // 413,
			"payload_too_large"},
		{http.StatusRequestURITooLong, // 414,
			"uri_too_long"},
		{http.StatusUnsupportedMediaType, // 415,
			"unsupported_media_type"},
		{http.StatusRequestedRangeNotSatisfiable, // 416,
			"range_not_satisfiable"},
		{http.StatusExpectationFailed, // 417,
			"expectation_failed"},
		{http.StatusTeapot, // 418,
			"im_a_teapot"},
		{http.StatusMisdirectedRequest, // 421,
			"misdirected_request"},
		{http.StatusUnprocessableEntity, // 422,
			"unprocessable_entity"},
		{http.StatusLocked, // 423,
			"locked"},
		{http.StatusFailedDependency, // 424,
			"failed_dependency"},
		{http.StatusTooEarly, // 425,
			"too_early"},
		{http.StatusUpgradeRequired, // 426,
			"upgrade_required"},
		{http.StatusPreconditionRequired, // 428,
			"precondition_required"},
		{http.StatusTooManyRequests, // 429,
			"too_many_requests"},
		{http.StatusRequestHeaderFieldsTooLarge, // 431,
			"request_header_fields_too_large"},
		{http.StatusUnavailableForLegalReasons, // 451,
			"unavailable_for_legal_reasons"},
		{http.StatusInternalServerError, // 500,
			"internal_server_error"},
		{http.StatusNotImplemented, // 501,
			"not_implemented"},
		{http.StatusBadGateway, // 502,
			"bad_gateway"},
		{http.StatusServiceUnavailable, // 503,
			"service_unavailable"},
		{http.StatusGatewayTimeout, // 504,
			"gateway_timeout"},
		{http.StatusHTTPVersionNotSupported, // 505,
			"http_version_not_supported"},
		{http.StatusVariantAlsoNegotiates, // 506,
			"variant_also_negotiates"},
		{http.StatusInsufficientStorage, // 507,
			"insufficent_storage"},
		{http.StatusLoopDetected, // 508,
			"loop_detected"},
		{http.StatusNotExtended, // 510,
			"not_extended"},
		{http.StatusNetworkAuthenticationRequired, // 511,
			"network_authentication_required"},
	}
)

func Test_restErr_Error(t *testing.T) {
	tests := []struct {
		name string
		e    restErr
		want string
	}{
		{"test 1", test1, "message: testmessage - status: 400 - error: test1 - causes: []"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Error(); got != tt.want {
				t.Errorf("restErr.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_restErr_Message(t *testing.T) {
	tests := []struct {
		name string
		e    restErr
		want string
	}{
		{"test 1", test1, "testmessage"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("%#v %T", tt.e.ErrCauses, tt.e.ErrCauses)
			if got := tt.e.Message(); got != tt.want {
				t.Errorf("restErr.Message() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_restErr_Status(t *testing.T) {
	tests := []struct {
		name string
		e    restErr
		want int
	}{
		{"test 1", test1, http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Status(); got != tt.want {
				t.Errorf("restErr.Status() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_restErr_Causes(t *testing.T) {
	tests := []struct {
		name string
		e    restErr
		want []interface{}
	}{
		{"test 1", test1, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Causes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restErr.Causes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHttpStatusMaker(t *testing.T) {
	type args struct {
		r []codePair
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{r: codePairTests}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, v := range tt.args.r {
				inR := restErr{
					ErrMessage: "test",
					ErrStatus:  v.i,
				}
				wantR := restErr{
					ErrMessage: "test",
					ErrStatus:  v.i,
					ErrError:   v.s,
				}
				if got := HttpStatusMaker(inR); !reflect.DeepEqual(got, &wantR) {
					t.Errorf("HttpStatusMaker() = %v, want %v", got, wantR)
				}
			}

		})
	}
}

func TestNewRestError(t *testing.T) {
	type args struct {
		message string
		status  int
		causes  []interface{}
	}
	tests := []struct {
		name string
		args args
		want restErr
	}{
		{"Test 1", args{message: "testmessage", status: http.StatusBadRequest, causes: nil}, test1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRestError(tt.args.message, tt.args.status, tt.args.causes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRestError() = %v, want %v", got, tt.want)
			} else {
				by, _ := json.Marshal(got)
				fmt.Println(by)
			}
		})
	}
}

func TestNewRestErrorFromBytes(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		args    args
		want    RestErr
		wantErr bool
	}{
		{"Test 1", args{bytes: test1Bytes}, test1, false},
		{"Test 2", args{bytes: test1BytesFail}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRestErrorFromBytes(tt.args.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRestErrorFromBytes() error = %v, wantErr %v", (err != nil), tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRestErrorFromBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
