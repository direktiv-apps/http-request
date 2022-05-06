package operations

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/direktiv/apps/go/pkg/apps"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	"http-request/models"
)

const (
	successKey = "success"
	resultKey  = "result"

	// http related
	statusKey  = "status"
	codeKey    = "code"
	headersKey = "headers"
)

var sm sync.Map

const (
	cmdErr = "io.direktiv.command.error"
	outErr = "io.direktiv.output.error"
	riErr  = "io.direktiv.ri.error"
)

type accParams struct {
	PostParams
	Commands []interface{}
}

type accParamsTemplate struct {
	PostBody
	Commands []interface{}
}

func PostDirektivHandle(params PostParams) middleware.Responder {
	fmt.Printf("params in: %+v", params)
	resp := &PostOKBody{}

	var (
		err error
		ret interface{}
	)

	ri, err := apps.RequestinfoFromRequest(params.HTTPRequest)
	if err != nil {
		return generateError(riErr, err)
	}

	ctx, cancel := context.WithCancel(params.HTTPRequest.Context())
	sm.Store(*params.DirektivActionID, cancel)
	defer sm.Delete(params.DirektivActionID)

	var responses []interface{}

	var paramsCollector []interface{}
	accParams := accParams{
		params,
		nil,
	}

	ret, err = runCommand0(ctx, accParams, ri)
	responses = append(responses, ret)

	if err != nil && true {
		errName := cmdErr
		return generateError(errName, err)
	}

	paramsCollector = append(paramsCollector, ret)
	accParams.Commands = paramsCollector

	fmt.Printf("object going in output template: %+v\n", responses)

	s, err := templateString(`{{ index . 0 | toJson }}`, responses)
	if err != nil {
		return generateError(outErr, err)
	}
	fmt.Printf("object from output template: %+v\n", s)

	responseBytes := []byte(s)

	// validate

	resp.UnmarshalBinary(responseBytes)
	err = resp.Validate(strfmt.Default)

	if err != nil {
		fmt.Printf("error parsing output template: %+v\n", err)
		return generateError(outErr, err)
	}

	return NewPostOK().WithPayload(resp)
}

// http request
func runCommand0(ctx context.Context,
	params accParams, ri *apps.RequestInfo) (map[string]interface{}, error) {

	ri.Logger().Infof("running http request")

	at := accParamsTemplate{
		params.Body,
		params.Commands,
	}

	ir := make(map[string]interface{})
	ir[successKey] = false

	type baseRequest struct {
		url, method, user, password string
		insecure, err200            bool
	}

	baseInfo := func(paramsIn interface{}) (*baseRequest, error) {

		u, err := templateString(`{{ .URL }}
{{- if .Params }}?
{{- range $i,$e := .Params }}{{ urlquery $i }}={{ urlquery $e }}{{- end }}
{{- end }}`, paramsIn)
		if err != nil {
			return nil, err
		}

		method, err := templateString(`{{ default "get" .Method }}`, paramsIn)
		if err != nil {
			return nil, err
		}

		user, err := templateString(`{{ default "" .Username }}`, paramsIn)
		if err != nil {
			return nil, err
		}

		password, err := templateString(`{{ default "" .Password }}`, paramsIn)
		if err != nil {
			return nil, err
		}

		return &baseRequest{
			url:      u,
			method:   method,
			user:     user,
			password: password,
			err200:   convertTemplateToBool(`{{ .Error200 }}`, paramsIn, true),
			insecure: convertTemplateToBool(`<no value>`, paramsIn, false),
		}, nil

	}
	br, err := baseInfo(at)
	if err != nil {
		ir[resultKey] = err.Error()
		return ir, err
	}

	headers := make(map[string]string)

	for k, v := range params.Body.Headers {
		headers[k] = v
	}

	var data []byte

	attachData := func(paramsIn interface{}, ri *apps.RequestInfo) ([]byte, error) {

		kind, err := templateString(`{{ default "string" .Content.Kind }}`, paramsIn)
		if err != nil {
			return nil, err
		}

		d, err := templateString(`{{ .Content.Value }}`, paramsIn)
		if err != nil {
			return nil, err
		}

		if kind == "file" {
			return os.ReadFile(filepath.Join(ri.Dir(), d))
		} else if kind == "base64" {
			return base64.StdEncoding.DecodeString(d)
		}

		return []byte(d), nil

	}

	if params.Body.Content != nil {
		data, err = attachData(at, ri)
		if err != nil {
			ir[resultKey] = err.Error()
			return ir, err
		}
	}

	ri.Logger().Infof("requesting %v", br.url)
	return doHttpRequest(br.method, br.url, br.user, br.password,
		headers, br.insecure, br.err200, data)

}

// end commands

func generateError(code string, err error) *PostDefault {

	d := NewPostDefault(0).WithDirektivErrorCode(code).
		WithDirektivErrorMessage(err.Error())

	errString := err.Error()

	errResp := models.Error{
		ErrorCode:    &code,
		ErrorMessage: &errString,
	}

	d.SetPayload(&errResp)

	return d
}

func HandleShutdown() {
	// nothing for generated functions
}
