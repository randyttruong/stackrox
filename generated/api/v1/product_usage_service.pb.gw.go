// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: api/v1/product_usage_service.proto

/*
Package v1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package v1

import (
	"context"
	"io"
	"net/http"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = descriptor.ForMessage
var _ = metadata.Join

func request_ProductUsageService_GetCurrentSecuredUnitsUsage_0(ctx context.Context, marshaler runtime.Marshaler, client ProductUsageServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq Empty
	var metadata runtime.ServerMetadata

	msg, err := client.GetCurrentSecuredUnitsUsage(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ProductUsageService_GetCurrentSecuredUnitsUsage_0(ctx context.Context, marshaler runtime.Marshaler, server ProductUsageServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq Empty
	var metadata runtime.ServerMetadata

	msg, err := server.GetCurrentSecuredUnitsUsage(ctx, &protoReq)
	return msg, metadata, err

}

var (
	filter_ProductUsageService_GetMaxSecuredUnitsUsage_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_ProductUsageService_GetMaxSecuredUnitsUsage_0(ctx context.Context, marshaler runtime.Marshaler, client ProductUsageServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq TimeRange
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_ProductUsageService_GetMaxSecuredUnitsUsage_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetMaxSecuredUnitsUsage(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ProductUsageService_GetMaxSecuredUnitsUsage_0(ctx context.Context, marshaler runtime.Marshaler, server ProductUsageServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq TimeRange
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_ProductUsageService_GetMaxSecuredUnitsUsage_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GetMaxSecuredUnitsUsage(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterProductUsageServiceHandlerServer registers the http handlers for service ProductUsageService to "mux".
// UnaryRPC     :call ProductUsageServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterProductUsageServiceHandlerFromEndpoint instead.
func RegisterProductUsageServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server ProductUsageServiceServer) error {

	mux.Handle("GET", pattern_ProductUsageService_GetCurrentSecuredUnitsUsage_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ProductUsageService_GetCurrentSecuredUnitsUsage_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ProductUsageService_GetCurrentSecuredUnitsUsage_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_ProductUsageService_GetMaxSecuredUnitsUsage_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ProductUsageService_GetMaxSecuredUnitsUsage_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ProductUsageService_GetMaxSecuredUnitsUsage_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterProductUsageServiceHandlerFromEndpoint is same as RegisterProductUsageServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterProductUsageServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterProductUsageServiceHandler(ctx, mux, conn)
}

// RegisterProductUsageServiceHandler registers the http handlers for service ProductUsageService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterProductUsageServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterProductUsageServiceHandlerClient(ctx, mux, NewProductUsageServiceClient(conn))
}

// RegisterProductUsageServiceHandlerClient registers the http handlers for service ProductUsageService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "ProductUsageServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "ProductUsageServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "ProductUsageServiceClient" to call the correct interceptors.
func RegisterProductUsageServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client ProductUsageServiceClient) error {

	mux.Handle("GET", pattern_ProductUsageService_GetCurrentSecuredUnitsUsage_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ProductUsageService_GetCurrentSecuredUnitsUsage_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ProductUsageService_GetCurrentSecuredUnitsUsage_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_ProductUsageService_GetMaxSecuredUnitsUsage_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ProductUsageService_GetMaxSecuredUnitsUsage_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ProductUsageService_GetMaxSecuredUnitsUsage_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_ProductUsageService_GetCurrentSecuredUnitsUsage_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4}, []string{"v1", "product", "usage", "secured-units", "current"}, "", runtime.AssumeColonVerbOpt(false)))

	pattern_ProductUsageService_GetMaxSecuredUnitsUsage_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4}, []string{"v1", "product", "usage", "secured-units", "max"}, "", runtime.AssumeColonVerbOpt(false)))
)

var (
	forward_ProductUsageService_GetCurrentSecuredUnitsUsage_0 = runtime.ForwardResponseMessage

	forward_ProductUsageService_GetMaxSecuredUnitsUsage_0 = runtime.ForwardResponseMessage
)