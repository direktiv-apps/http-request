package operations

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/direktiv/apps/go/pkg/apps"
	"github.com/go-openapi/runtime/middleware"
	// custom function imports
	// end

	"app/models"
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
	Commands    []interface{}
	DirektivDir string
}

type accParamsTemplate struct {
	models.PostParamsBody
	Commands    []interface{}
	DirektivDir string
}

type ctxInfo struct {
	cf        context.CancelFunc
	cancelled bool
}

func PostDirektivHandle(params PostParams) middleware.Responder {
	fmt.Printf("params in: %+v\n", params)
	var resp interface{}

	var (
		err  error
		ret  interface{}
		cont bool
	)

	ri, err := apps.RequestinfoFromRequest(params.HTTPRequest)
	if err != nil {
		return generateError(riErr, err)
	}

	ctx, cancel := context.WithCancel(params.HTTPRequest.Context())

	sm.Store(*params.DirektivActionID, &ctxInfo{
		cancel,
		false,
	})

	defer sm.Delete(*params.DirektivActionID)

	var responses []interface{}

	var paramsCollector []interface{}
	accParams := accParams{
		params,
		nil,
		ri.Dir(),
	}
	ret, err = runCommand0(ctx, accParams, ri)

	responses = append(responses, ret)

	// if foreach returns an error there is no continue
	//
	// default we do not continue
	cont = convertTemplateToBool("<no value>", accParams, false)
	// cont = convertTemplateToBool("<no value>", accParams, true)
	//

	if err != nil && !cont {

		errName := cmdErr

		// if the delete function added the cancel tag
		ci, ok := sm.Load(*params.DirektivActionID)
		if ok {
			cinfo, ok := ci.(*ctxInfo)
			if ok && cinfo.cancelled {
				errName = "direktiv.actionCancelled"
				err = fmt.Errorf("action got cancel request")
			}
		}

		return generateError(errName, err)
	}

	paramsCollector = append(paramsCollector, ret)
	accParams.Commands = paramsCollector

	responseBytes, err := json.Marshal(responses)
	if err != nil {
		return generateError(outErr, err)
	}
	err = json.Unmarshal(responseBytes, &resp)
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
		*params.Body,
		params.Commands,
		params.DirektivDir,
	}

	ir := make(map[string]interface{})
	ir[successKey] = false

	type baseRequest struct {
		url, method, user, password string
		insecure, err200, debug     bool
	}

	baseInfo := func(paramsIn interface{}, dir string) (*baseRequest, error) {

		u, err := templateString(`{{ .URL }}
{{- if .Params }}?
{{- range $i,$e := .Params }}{{ urlquery $i }}={{ urlquery $e }}{{ if $i }}&{{ end }}{{- end }}
{{- end }}`, paramsIn, dir)
		if err != nil {
			return nil, err
		}

		method, err := templateString(`{{ default "get" .Method }}`, paramsIn, dir)
		if err != nil {
			return nil, err
		}

		user, err := templateString(`{{ default "" .Username }}`, paramsIn, dir)
		if err != nil {
			return nil, err
		}

		password, err := templateString(`{{ default "" .Password }}`, paramsIn, dir)
		if err != nil {
			return nil, err
		}

		return &baseRequest{
			url:      u,
			method:   method,
			user:     user,
			password: password,
			err200:   convertTemplateToBool(`{{ .Error200 }}`, paramsIn, true),
			insecure: convertTemplateToBool(`{{ .Insecure }}`, paramsIn, false),
			debug:    convertTemplateToBool(`{{ .Debug }}`, paramsIn, false),
		}, nil

	}
	br, err := baseInfo(at, ri.Dir())
	if err != nil {
		ir[resultKey] = err.Error()
		return ir, err
	}

	headers := make(map[string]string)

	for k, v := range params.Body.Headers {
		headers[k] = v
	}

	var data []byte

	attachData := func(paramsIn interface{}, ri *apps.RequestInfo, dir string) ([]byte, error) {

		kind, err := templateString(`{{- if .Content }}
{{- default "json" .Content.Kind }}
{{- end }}`, paramsIn, dir)
		if err != nil {
			return nil, err
		}

		d, err := templateString(`{{- if .Content }}
{{ $content := (default "" .Content.Value) }}
{{- if eq (deref (default "json" .Content.Kind)) "json" }}
{{- $content | toJson }}
{{- else }}
{{- $content }}
{{- end }}
{{- end }}`, paramsIn, dir)
		if err != nil {
			return nil, err
		}

		if string(d) == "empty" {
			return []byte("{}"), nil
		}

		if kind == "file" {
			return os.ReadFile(filepath.Join(ri.Dir(), d))
		} else if kind == "base64" {
			return base64.StdEncoding.DecodeString(d)
		}

		return []byte(d), nil

	}

	data, err = attachData(at, ri, ri.Dir())
	if err != nil {
		ir[resultKey] = err.Error()
		return ir, err
	}

	if br.debug {
		ri.Logger().Infof("Payload:")
		ri.Logger().Infof(string(data))
	}

	if br.debug {
		ri.Logger().Infof("requesting %v", br.url)
	}

	return doHttpRequest(br.debug, br.method, br.url, br.user, br.password,
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
