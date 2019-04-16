package modifier

import "github.com/devopsfaith/krakend/logging"

type LoggerAdder interface {
	SetLogger(l logging.Logger)
}
