package queryParams

import "net/url"

type ModifierInterface interface {
	Execute(values url.Values)
}
