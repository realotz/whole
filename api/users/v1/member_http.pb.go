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

type MemberServiceHandler interface {
	Create(context.Context, *MemberCreateOption) (*Member, error)

	Delete(context.Context, *MemberDeleteOption) (*NullReply, error)

	Get(context.Context, *MemberGetOption) (*Member, error)

	List(context.Context, *MemberListOption) (*MemberList, error)

	Login(context.Context, *MemberLogin) (*MemberLoginRes, error)

	Update(context.Context, *MemberUpdateOption) (*Member, error)
}

func NewMemberServiceHandler(srv MemberServiceHandler, opts ...http1.HandleOption) http.Handler {
	h := http1.DefaultHandleOptions()
	for _, o := range opts {
		o(&h)
	}
	r := mux.NewRouter()

	r.HandleFunc("/users/member/login", func(w http.ResponseWriter, r *http.Request) {
		var in MemberLogin

		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}
		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*MemberLogin))
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

	r.HandleFunc("/users/member", func(w http.ResponseWriter, r *http.Request) {
		var in MemberListOption

		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}
		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.List(ctx, req.(*MemberListOption))
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

	r.HandleFunc("/users/member/{id}", func(w http.ResponseWriter, r *http.Request) {
		var in MemberGetOption

		if err := binding.MapProto(&in, mux.Vars(r)); err != nil {
			h.Error(w, r, err)
			return
		}

		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}
		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Get(ctx, req.(*MemberGetOption))
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

	r.HandleFunc("/users/member", func(w http.ResponseWriter, r *http.Request) {
		var in MemberCreateOption

		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}
		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Create(ctx, req.(*MemberCreateOption))
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

	r.HandleFunc("/users/member", func(w http.ResponseWriter, r *http.Request) {
		var in MemberUpdateOption

		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}
		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Update(ctx, req.(*MemberUpdateOption))
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

	r.HandleFunc("/users/member/{id}", func(w http.ResponseWriter, r *http.Request) {
		var in MemberDeleteOption

		if err := binding.MapProto(&in, mux.Vars(r)); err != nil {
			h.Error(w, r, err)
			return
		}

		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}
		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Delete(ctx, req.(*MemberDeleteOption))
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
