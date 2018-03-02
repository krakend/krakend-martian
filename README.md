# KrakenD martian

The `krakend-martian` package integrates the martian project into the KrakenD framework.

## How to use it

Add your martian DSL definition under the "github.com/krakend-martian/proxy" namespace of the backend section of the config file

	"extra_config": {
		"github.com/krakend-martian/proxy": {}
	}

More details here: https://github.com/google/martian#modifiers-all-the-way-down

Check the [example](github.com/krakend-martian/tree/master/example) forlder for a complete demo.

## Example

Build and run the example

	$ cd example
	$ go build
	$ ./example -c krakend.json -d

Send a request to the configured endpoint

	$ curl -i 127.0.0.1:8080/supu
	HTTP/1.1 200 OK
	Cache-Control: public, max-age=0
	Content-Type: application/json; charset=utf-8
	X-Krakend: Version undefined
	Date: Sun, 15 Oct 2017 10:14:20 GMT
	Content-Length: 18

	{"message":"pong"}

And check the logs of the KrakenD: the request modifiers has done their job!

	[KRAKEND] 2017/10/15 - 12:14:20.871 ▶ DEBU Method: GET
	[KRAKEND] 2017/10/15 - 12:14:20.871 ▶ DEBU URL: /__debug/supu
	[KRAKEND] 2017/10/15 - 12:14:20.871 ▶ DEBU Query: map[]
	[KRAKEND] 2017/10/15 - 12:14:20.871 ▶ DEBU Params: [{param /supu}]
	[KRAKEND] 2017/10/15 - 12:14:20.871 ▶ DEBU Headers: map[X-Martian:[ouh yeah!] Accept-Encoding:[gzip] User-Agent:[KrakenD Version undefined] Content-Length:[19] Content-Type:[application/json] X-Forwarded-For:[127.0.0.1]]
	[KRAKEND] 2017/10/15 - 12:14:20.871 ▶ DEBU Body: {"msg":"you rock!"}
	[GIN] 2017/10/15 - 12:14:20 | 200 |     226.159µs |       127.0.0.1 | GET      /__debug/supu
	[GIN] 2017/10/15 - 12:14:20 | 200 |    2.316784ms |       127.0.0.1 | GET      /supu
