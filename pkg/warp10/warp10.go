package warp10

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type (
	Request struct {
		Endpoint   string `json:"endpoint"`
		WarpScript string `json:"warpscript"`
	}

	Response struct {
		request    *Request
		err        error
		elapsed    int64
		fetched    int64
		operations int64
		stack      []byte
	}

	ResponseError struct {
		Message string
		Line    int64
	}
)

const (
	ElapsedHeader      = "X-Warp10-Elapsed"
	ErrorsLineHeader   = "X-Warp10-Error-Line"
	ErrorMessageHeader = "X-Warp10-Error-Message"
	FetchedHeader      = "X-Warp10-Fetched"
	OpsHeader          = "X-Warp10-Ops"
)

func NewRequest(endpoint, warpscript string) Request {
	return Request{Endpoint: endpoint, WarpScript: warpscript}
}

func (req Request) Exec() *Response {
	return Exec(req)
}

func (res Response) IsErrored() bool {
	return res.err != nil
}

func (res Response) Error() error {
	return res.err
}

func (res Response) Fetched() int64 {
	return res.fetched
}

func (res Response) Elapsed() time.Duration {
	return time.Duration(res.elapsed)
}

func (res Response) Ops() int64 {
	return res.operations
}

func (res Response) StackRaw() []byte {
	return res.stack
}

func (res Response) StackRawString() string {
	return string(res.StackRaw())
}

func (res Response) StackSlice() []interface{} {
	var s []interface{}
	_ = json.Unmarshal(res.StackRaw(), &s)
	return s
}

func (res Response) StackUnmarshal(i interface{}) error {
	return json.Unmarshal(res.stack, i)
}

func (err ResponseError) Error() string {
	return fmt.Sprintf("WarpScript#%d: %s", err.Line, err.Message)
}

func Exec(request Request) *Response {
	response := &Response{request: &request}

	req, err := http.NewRequest("POST", request.Endpoint, strings.NewReader(request.WarpScript))
	if err != nil {
		return response
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		response.err = errors.Wrap(err, "cannot perform HTTP request")
		return response
	}
	defer res.Body.Close()

	if elapsed := res.Header[ElapsedHeader]; len(elapsed) > 0 {
		response.elapsed, _ = strconv.ParseInt(elapsed[0], 10, 64)
	}

	if ops := res.Header[OpsHeader]; len(ops) > 0 {
		response.operations, _ = strconv.ParseInt(ops[0], 10, 64)
	}

	if fetched := res.Header[FetchedHeader]; len(fetched) > 0 {
		response.fetched, _ = strconv.ParseInt(fetched[0], 10, 64)
	}

	if res.StatusCode == http.StatusInternalServerError {
		werr := ResponseError{}

		if line := res.Header[ErrorsLineHeader]; len(line) > 0 {
			werr.Line, _ = strconv.ParseInt(line[0], 10, 64)
		}

		if message := res.Header[ErrorMessageHeader]; len(message) > 0 {
			werr.Message = message[0]
		}

		response.err = werr
		return response
	} else if res.StatusCode != http.StatusOK {
		response.err = fmt.Errorf("invlid status code: %s", res.Status)
		return response
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		response.err = errors.Wrap(err, "cannot read HTTP response body")
		return response
	}

	response.stack = b
	return response
}
