
# http-request 1.0

Run http-request in Direktiv

---
- #### Categories: unknown
- #### Image: gcr.io/direktiv/apps/http-request 
- #### License: [Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)
- #### Issue Tracking: https://github.com/direktiv-apps/http-request/issues
- #### URL: https://github.com/direktiv-apps/http-request
- #### Maintainer: [direktiv.io](https://www.direktiv.io) 
---

## About http-request

Run http-request in Direktiv as a function

### Example(s)
  #### Function Configuration
```yaml
functions:
- id: http-request
  image: gcr.io/direktiv/apps/http-request:1.0
  type: knative-workflow
```
   #### Basic
```yaml
- id: http-request
  type: action
  action:
    function: http-request
    input: 
      debug: true
      url: 'https://www.direktiv.io'
      query:
        hello: world
        hello1: world2
```
   #### POST Request
```yaml
- id: http-request
  type: action
  action:
    function: http-request
    input: 
      url: 'https://www.direktiv.io'
      method: POST
      headers:
        header1: value1
        header2: value2
```
   #### POST Request with file
```yaml
- id: http-request
  type: action
  action:
    function: http-request
    input: 
      url: 'https://www.direktiv.io'
      method: POST
      content:
        kind: string
        value: 'This is the payload'
```
   #### Treat 404 as error
```yaml
- id: http-request
  type: action
  action:
    function: http-request
    input: 
      url: 'https://www.direktiv.io/doesnotexist'
      error200: true
  catch:
  - error: "*"                  
```

   ### Secrets


*No secrets required*







### Request



#### Request Attributes
[PostParamsBody](#post-params-body)

### Response
  List of executed commands.
#### Example Reponses
    
```json
{
  "code": 200,
  "headers": {
    "Access-Control-Allow-Origin": "*",
    "Content-Type": "text/html"
  },
  "result": "KXx8T2JqZWN0LmR...",
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
#### <span id="post-params-body"></span> postParamsBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| content | [PostParamsBodyContent](#post-params-body-content)| `PostParamsBodyContent` |  | |  |  |
| debug | boolean| `bool` |  | | Prints the full URL and headers to logs. | `true` |
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
 - file: File payload, e.g. instance or workflow variables
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| kind | string| `string` |  | `"string"`| Kind of data |  |
| value | string| `string` |  | | Value depends on `kind` value. |  |

 
