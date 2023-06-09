package middleware

import "github.com/gsxhnd/tower/src/utils"

type Middlewarer interface {
}
type middleware struct{
	logger utils.Logger
}

func NewMiddleware(l utils.Logger) Middlewarer {
	return &middleware{
		logger: l,
	}
}
