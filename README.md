# KrakenD martian

The `krakend-martian` package integrates the martian project into the KrakenD framework.

## How to use it

Add your martian DSL definition under the "github.com/devopsfaith/krakend-martian" namespace of the backend section of the config file

```
"extra_config": {
  "github.com/devopsfaith/krakend-martian": {}
}
```

More details here: https://github.com/google/martian#modifiers-all-the-way-down

Check the [example](github.com/krakend-martian/tree/master/example) folder for a complete demo.

## Example

The martian example is built on top of [chuknorris.io](https://api.chucknorris.io/) REST API.

* Build and run krakend edition with martian

```
$ cd example
$ go install ./...
$ GOPATH/bin/example -c example/krakend.json -d -p 8000
```

* Send a request to the configured endpoint

```
$ curl -i 127.0.0.1:8000/supu
 HTTP/1.1 200 OK
 Cache-Control: public, max-age=0
 Content-Type: application/json; charset=utf-8
 X-Krakend: Version undefined
 X-Krakend-Completed: true
 Date: Tue, 10 Jul 2018 13:15:04 GMT
 Content-Length: 80

 {"author":"Chuck Norris","joke":"People pray to God. God prays to Chuck Norris"}
```

And check the logs of the KrakenD: the request modifiers has done their job!
See how `"author":"Chuck Norris"` was added to the payload.

