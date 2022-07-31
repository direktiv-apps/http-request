
Feature: Basic

# The secrects can be used in the payload with the following syntax #(mysecretname)
Background:


Scenario: direktiv

	Given url karate.properties['testURL']

	And path '/'
	And header Direktiv-ActionID = 'development'
	And header Direktiv-TempDir = '/tmp'
	And request
	"""
	{
		"url": "https://www.direktiv.io"
	}
	"""
	When method POST
	Then status 200
	And match $ ==
	"""
	[
	{
		"code": 200,
		"headers": #object,
		"result": #notnull,
		"status": "200 OK",
		"success": true
	}
	]
	"""
	
Scenario: json

	Given url karate.properties['testURL']

	And path '/'
	And header Direktiv-ActionID = 'development'
	And header Direktiv-TempDir = '/tmp'
	And request
	"""
	{
		"url": "https://api.blockchain.com/v3/exchange/l2/BTC-USD"
	}
	"""
	When method POST
	Then status 200
	And match $ ==
	"""
	[
	{
		"code": 200,
		"headers": #object,
		"result": {
		"asks": #array,
		"bids": #array,
		"symbol": "BTC-USD"
		},
		"status": "200 OK",
		"success": true
	}
	]
	"""


Scenario: notfound

	Given url karate.properties['testURL']

	And path '/'
	And header Direktiv-ActionID = 'development'
	And header Direktiv-TempDir = '/tmp'
	And request
	"""
	{
		"url": "https://www.direktiv.io/whereareyou"
	}
	"""
	When method POST
	Then status 200
	And match $ ==
	"""
	[
	{
		"code": 404,
		"headers": #object,
		"result": #notnull,
		"status": "404 Not Found",
		"success": true
	}
	]
	"""

Scenario: notfounderror

	Given url karate.properties['testURL']

	And path '/'
	And header Direktiv-ActionID = 'development'
	And header Direktiv-TempDir = '/tmp'
	And request
	"""
	{
		"url": "https://www.direktiv.io/whereareyou",
		"error200": true
	}
	"""
	When method POST
	Then status 500
	And match $ ==
	"""
	{
		"errorCode": "#notnull",
		"errorMessage": "#notnull"
	}
	"""

Scenario: urlrequired

	Given url karate.properties['testURL']

	And path '/'
	And header Direktiv-ActionID = 'development'
	And header Direktiv-TempDir = '/tmp'
	And request
	"""
	{
	}
	"""
	When method POST
	Then status 422
	And match $ ==
	"""
	{
		"errorCode": "#notnull",
		"errorMessage": "url in body is required"
	}
	"""


Scenario: paramsheaders

	Given url karate.properties['testURL']

	And path '/'
	And header Direktiv-ActionID = 'development'
	And header Direktiv-TempDir = '/tmp'
	And request
	"""
	{	
		"debug": true,
		"url": "https://www.direktiv.io",
		"params": 
		{
		"hello": "world",
		"hello2": "world2"
		},
		"headers": 
		{
		"myheader":"value",
		"myheader2": "value2"
		}

	}
	"""
	When method POST
	Then status 200
