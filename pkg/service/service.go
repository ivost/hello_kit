package service

import "context"

// HelloKitService describes the service.
type HelloKitService interface {
	// Add your methods here
	Foo(ctx context.Context, s string) (rs string, err error)
}

type basicHelloKitService struct{}

func (b *basicHelloKitService) Foo(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of Foo
	rs = s
	return rs, err
}

// NewBasicHelloKitService returns a naive, stateless implementation of HelloKitService.
func NewBasicHelloKitService() HelloKitService {
	return &basicHelloKitService{}
}

// New returns a HelloKitService with all of the expected middleware wired in.
func New(middleware []Middleware) HelloKitService {
	var svc HelloKitService = NewBasicHelloKitService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
