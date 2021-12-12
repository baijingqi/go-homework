// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.2

package Comment

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type CommentHTTPServer interface {
	AddComment(context.Context, *AddCommentRequest) (*AddCommentReply, error)
	CommentList(context.Context, *DelCommentRequest) (*DelCommentReply, error)
	DelComment(context.Context, *DelCommentRequest) (*DelCommentReply, error)
}

func RegisterCommentHTTPServer(s *http.Server, srv CommentHTTPServer) {
	r := s.Route("/")
	r.GET("/comment/add", _Comment_AddComment0_HTTP_Handler(srv))
	r.GET("/comment/Del", _Comment_DelComment0_HTTP_Handler(srv))
	r.GET("/comment/list", _Comment_CommentList0_HTTP_Handler(srv))
}

func _Comment_AddComment0_HTTP_Handler(srv CommentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddCommentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.Comment.Comment/AddComment")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddComment(ctx, req.(*AddCommentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AddCommentReply)
		return ctx.Result(200, reply)
	}
}

func _Comment_DelComment0_HTTP_Handler(srv CommentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DelCommentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.Comment.Comment/DelComment")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DelComment(ctx, req.(*DelCommentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DelCommentReply)
		return ctx.Result(200, reply)
	}
}

func _Comment_CommentList0_HTTP_Handler(srv CommentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DelCommentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.Comment.Comment/CommentList")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CommentList(ctx, req.(*DelCommentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DelCommentReply)
		return ctx.Result(200, reply)
	}
}

type CommentHTTPClient interface {
	AddComment(ctx context.Context, req *AddCommentRequest, opts ...http.CallOption) (rsp *AddCommentReply, err error)
	CommentList(ctx context.Context, req *DelCommentRequest, opts ...http.CallOption) (rsp *DelCommentReply, err error)
	DelComment(ctx context.Context, req *DelCommentRequest, opts ...http.CallOption) (rsp *DelCommentReply, err error)
}

type CommentHTTPClientImpl struct {
	cc *http.Client
}

func NewCommentHTTPClient(client *http.Client) CommentHTTPClient {
	return &CommentHTTPClientImpl{client}
}

func (c *CommentHTTPClientImpl) AddComment(ctx context.Context, in *AddCommentRequest, opts ...http.CallOption) (*AddCommentReply, error) {
	var out AddCommentReply
	pattern := "/comment/add"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.Comment.Comment/AddComment"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CommentHTTPClientImpl) CommentList(ctx context.Context, in *DelCommentRequest, opts ...http.CallOption) (*DelCommentReply, error) {
	var out DelCommentReply
	pattern := "/comment/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.Comment.Comment/CommentList"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CommentHTTPClientImpl) DelComment(ctx context.Context, in *DelCommentRequest, opts ...http.CallOption) (*DelCommentReply, error) {
	var out DelCommentReply
	pattern := "/comment/Del"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.Comment.Comment/DelComment"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
