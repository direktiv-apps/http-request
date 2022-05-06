
# http-request 1.0.0

Executes HTTP requests from Direktiv

---
- #### Category: Network
- #### Image: direktiv/http-request 
- #### License: [Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)
- #### Issue Tracking: https://github.com/direktiv-apps/http-request/issues
- #### URL: https://github.com/direktiv-apps/http-request
- #### Maintainer: [direktiv.io](https://www.direktiv.io)
---

## About http-request

This function executes HTTP requests. Query parameters, data payloads and custom headers are supported. The payload can be a string, 
base64 or a file. It can be configured to return errors if the response status is not 2xx and it can be configured to ignore SSL errors in
case of self-signed certificates. 

### Example(s)
  #### Function Configuration
  ```yaml
  functions:
  - id: request
    image: direktiv/http-request
    type: knative-workflow
  ```
   #### Basic
   ```yaml
   - id: req
     type: action
     action:
       function: request
     input: 
     url: "http://www.direktiv.io"
   ```
   #### Post Request
   ```yaml
   url: http://www.direktiv.io
method: post
   ```
   #### Request with File
   ```yaml
   - id: set
     type: setter
     variables:
     - key: myfile
       scope: workflow
       mimeType: text/plain
       value: This Is The Data
     transition: send
- id: send 
     type: action
     action:
       function: post
       files:
         - key: myfile
           scope: workflow
           as: myfile.txt
       input: 
         url: "https://webhook.site"
         content: 
           kind: file
           value: myfile.txt
   ```
   #### Request with Basic Authentication
   ```yaml
   - id: getter 
     type: action
     action:
       secrets: ["mypassword"]
       function: get
       input: 
         url: "https://webhook.site
         user: admin
         password: jq(.secrets.mypassword)
   ```

### Responses
  Returns headers and content of the requested URL. If the response is JSON it will be returned as JSON. 
If it is not JSON the content is returned as base64 encoded string.
#### Reponse Types
    
  

[PostOKBody](#post-o-k-body)
#### Example Reponses
    
```json
{
  "code": 200,
  "headers": {
    "Content-Type": [
      "text/plain; charset=UTF-8"
    ]
  },
  "status": "200 OK",
  "success": true
}
```

### Errors
| Type | Description
|------|---------|
| io.direktiv.command.error | Command execution failed |
| io.direktiv.output.error | Template error for output generation of the service |
| io.direktiv.ri.error | Can not create information object from request |


### Types
#### <span id="post-o-k-body"></span> postOKBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| code | int64 (formatted integer)| `int64` |  | | http response code | `200` |
| headers | map of [[]string](#string)| `map[string][]string` |  | | Key/Value map of response headers | `{"content-type":"application/json"}` |
| result | string| `string` |  | | JSON or base64 content of the response | `{"json":"response"}` |
| status | string| `string` |  | | test representation of status | `200 OK` |
| success | boolean| `bool` |  | | Indicates successful or unsuccessful execution of the request | `true` |


#### <span id="post-params-body"></span> postParamsBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| content | [PostParamsBodyContent](#post-params-body-content)| `PostParamsBodyContent` |  | |  |  |
| error200 | boolean| `bool` |  | | If set to `true` responses with status above 299 will be treated as errors. | `true` |
| headers | map of string| `map[string]string` |  | | List of key/values send as headers with the request. | `{"myheader":"value"}` |
| insecure | boolean| `bool` |  | | Skips the verification the server certificate chain and host name. | `true` |
| method | string| `string` |  | | HTTP method. Defaults to GET. | `POST` |
| params | map of string| `map[string]string` |  | | List of key/values appended to URL as query parameters. | `{"query1":"queryvalue"}` |
| password | string| `string` |  | | If username and password are set, it will be used for basic authenitcation for the request. | `mypassword` |
| url | string| `string` | ✓ | | URL for the request. | `http://www.direktiv.io` |
| username | string| `string` |  | | If username and password are set, it will be used for basic authenitcation for the request. | `myuser` |


#### <span id="post-params-body-content"></span> postParamsBodyContent

> Defines the payload of the request. The `kind` value can have three different values: 
 - string: Plain string payload, e.g. JSON
 - base64: Will be "converted" to binary and attached
 - file: File payload
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| kind | string| `string` |  | `"string"`| Kind of data |  |
| value | string| `string` |  | | Value depends on `kind` value. |  |

 
