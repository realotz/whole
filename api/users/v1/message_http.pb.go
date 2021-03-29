// Code generated by protoc-gen-go-http. DO NOT EDIT.

package v1

import (
	context "context"
	http1 "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	mux "github.com/gorilla/mux"
	http "net/http"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(http.Request)
var _ = new(context.Context)
var _ = binding.MapProto
var _ = mux.NewRouter

const _ = http1.SupportPackageIsVersion1

type MessageServiceHandler interface {
	Create(context.Context, *MessageGetOption) (*Message, error)

	Delete(context.Context, *MessageDeleteOption) (*Message, error)

	Get(context.Context, *MessageGetOption) (*Message, error)

	List(context.Context, *MessageListOption) (*MessageList, error)

	Update(context.Context, *MessageUpdateOption) (*Message, error)
}

func NewMessageServiceHandler(srv MessageServiceHandler, opts ...http1.HandleOption) http.Handler {
	h := http1.DefaultHandleOptions()
	for _, o := range opts {
		o(&h)
	}
	r := mux.NewRouter()

	r.HandleFunc("/news/message", func(w http.ResponseWriter, r *http.Request) {
		var in MessageListOption

		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}
		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.List(ctx, req.(*MessageListOption))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		if err := h.Encode(w, r, out); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("GET")

	r.HandleFunc("/news/message/{id}", func(w http.ResponseWriter, r *http.Request) {
		var in MessageGetOption

		if err := binding.MapProto(&in, mux.Vars(r)); err != nil {
			h.Error(w, r, err)
			return
		}

		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}
		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Get(ctx, req.(*MessageGetOption))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		if err := h.Encode(w, r, out); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("GET")

	r.HandleFunc("/news/message", func(w http.ResponseWriter, r *http.Request) {
		var in MessageGetOption

		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}
		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Create(ctx, req.(*MessageGetOption))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		if err := h.Encode(w, r, out); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("POST")

	r.HandleFunc("/news/message", func(w http.ResponseWriter, r *http.Request) {
		var in MessageUpdateOption

		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}
		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Update(ctx, req.(*MessageUpdateOption))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		if err := h.Encode(w, r, out); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("PUT")

	r.HandleFunc("/news/message/{id}", func(w http.ResponseWriter, r *http.Request) {
		var in MessageDeleteOption

		if err := binding.MapProto(&in, mux.Vars(r)); err != nil {
			h.Error(w, r, err)
			return
		}

		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}
		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Delete(ctx, req.(*MessageDeleteOption))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		if err := h.Encode(w, r, out); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("DELETE")

	return r
}
