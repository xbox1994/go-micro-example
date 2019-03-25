package router

import (
	"github.com/micro/go-api/resolver"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/registry"
)

type Options struct {
	Namespace string
	Handler   string
	Registry  registry.Registry
	Resolver  resolver.Resolver
}

type Option func(o *Options)

func newOptions(opts ...Option) Options {
	options := Options{
		Namespace: "go.micro.api",
		Handler:   "meta",
		Registry:  *cmd.DefaultOptions().Registry,
		Resolver: &defaultResolver{
			namespace: "go.micro.api",
			handler:   "meta",
		},
	}

	for _, o := range opts {
		o(&options)
	}

	return options
}

func WithHandler(h string) Option {
	return func(o *Options) {
		o.Handler = h

		// set handler in default resolver
		r, ok := o.Resolver.(*defaultResolver)
		if ok {
			r.handler = h
		}
	}
}

func WithNamespace(ns string) Option {
	return func(o *Options) {
		o.Namespace = ns

		// set namespace in default resolver
		r, ok := o.Resolver.(*defaultResolver)
		if ok {
			r.namespace = ns
		}
	}
}

func WithRegistry(r registry.Registry) Option {
	return func(o *Options) {
		o.Registry = r
	}
}

func WithResolver(r resolver.Resolver) Option {
	return func(o *Options) {
		o.Resolver = r
	}
}
