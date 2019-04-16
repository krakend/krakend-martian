package postParams

import "net/url"

type ModifierInterface interface {
	Execute(values url.Values)
}
