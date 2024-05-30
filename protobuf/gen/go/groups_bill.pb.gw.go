// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: groups_bill.proto

/*
Package monify is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package monify

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_GroupsBillService_CreateGroupBill_0(ctx context.Context, marshaler runtime.Marshaler, client GroupsBillServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreateGroupBillRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.CreateGroupBill(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_GroupsBillService_CreateGroupBill_0(ctx context.Context, marshaler runtime.Marshaler, server GroupsBillServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreateGroupBillRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.CreateGroupBill(ctx, &protoReq)
	return msg, metadata, err

}

func request_GroupsBillService_DeleteGroupBill_0(ctx context.Context, marshaler runtime.Marshaler, client GroupsBillServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq DeleteGroupBillRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["bill_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "bill_id")
	}

	protoReq.BillId, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "bill_id", err)
	}

	msg, err := client.DeleteGroupBill(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_GroupsBillService_DeleteGroupBill_0(ctx context.Context, marshaler runtime.Marshaler, server GroupsBillServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq DeleteGroupBillRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["bill_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "bill_id")
	}

	protoReq.BillId, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "bill_id", err)
	}

	msg, err := server.DeleteGroupBill(ctx, &protoReq)
	return msg, metadata, err

}

func request_GroupsBillService_ModifyGroupBill_0(ctx context.Context, marshaler runtime.Marshaler, client GroupsBillServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ModifyGroupBillRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["bill_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "bill_id")
	}

	protoReq.BillId, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "bill_id", err)
	}

	msg, err := client.ModifyGroupBill(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_GroupsBillService_ModifyGroupBill_0(ctx context.Context, marshaler runtime.Marshaler, server GroupsBillServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ModifyGroupBillRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["bill_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "bill_id")
	}

	protoReq.BillId, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "bill_id", err)
	}

	msg, err := server.ModifyGroupBill(ctx, &protoReq)
	return msg, metadata, err

}

var (
	filter_GroupsBillService_GetSelfTotalSpend_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_GroupsBillService_GetSelfTotalSpend_0(ctx context.Context, marshaler runtime.Marshaler, client GroupsBillServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetSelfTotalSpendRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_GroupsBillService_GetSelfTotalSpend_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetSelfTotalSpend(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_GroupsBillService_GetSelfTotalSpend_0(ctx context.Context, marshaler runtime.Marshaler, server GroupsBillServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetSelfTotalSpendRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_GroupsBillService_GetSelfTotalSpend_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GetSelfTotalSpend(ctx, &protoReq)
	return msg, metadata, err

}

var (
	filter_GroupsBillService_GetGroupTotalSpend_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_GroupsBillService_GetGroupTotalSpend_0(ctx context.Context, marshaler runtime.Marshaler, client GroupsBillServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetGroupTotalSpendRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_GroupsBillService_GetGroupTotalSpend_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetGroupTotalSpend(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_GroupsBillService_GetGroupTotalSpend_0(ctx context.Context, marshaler runtime.Marshaler, server GroupsBillServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetGroupTotalSpendRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_GroupsBillService_GetGroupTotalSpend_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GetGroupTotalSpend(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterGroupsBillServiceHandlerServer registers the http handlers for service GroupsBillService to "mux".
// UnaryRPC     :call GroupsBillServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterGroupsBillServiceHandlerFromEndpoint instead.
func RegisterGroupsBillServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server GroupsBillServiceServer) error {

	mux.Handle("POST", pattern_GroupsBillService_CreateGroupBill_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/.GroupsBillService/CreateGroupBill", runtime.WithHTTPPathPattern("/v1/groups_bill"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_GroupsBillService_CreateGroupBill_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_GroupsBillService_CreateGroupBill_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("DELETE", pattern_GroupsBillService_DeleteGroupBill_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/.GroupsBillService/DeleteGroupBill", runtime.WithHTTPPathPattern("/v1/groups_bill/{bill_id}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_GroupsBillService_DeleteGroupBill_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_GroupsBillService_DeleteGroupBill_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("PUT", pattern_GroupsBillService_ModifyGroupBill_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/.GroupsBillService/ModifyGroupBill", runtime.WithHTTPPathPattern("/v1/groups_bill/{bill_id}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_GroupsBillService_ModifyGroupBill_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_GroupsBillService_ModifyGroupBill_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_GroupsBillService_GetSelfTotalSpend_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/.GroupsBillService/GetSelfTotalSpend", runtime.WithHTTPPathPattern("/v1/groups_bill/self_total_spend"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_GroupsBillService_GetSelfTotalSpend_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_GroupsBillService_GetSelfTotalSpend_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_GroupsBillService_GetGroupTotalSpend_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/.GroupsBillService/GetGroupTotalSpend", runtime.WithHTTPPathPattern("/v1/groups_bill/group_bill"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_GroupsBillService_GetGroupTotalSpend_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_GroupsBillService_GetGroupTotalSpend_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterGroupsBillServiceHandlerFromEndpoint is same as RegisterGroupsBillServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterGroupsBillServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.NewClient(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterGroupsBillServiceHandler(ctx, mux, conn)
}

// RegisterGroupsBillServiceHandler registers the http handlers for service GroupsBillService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterGroupsBillServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterGroupsBillServiceHandlerClient(ctx, mux, NewGroupsBillServiceClient(conn))
}

// RegisterGroupsBillServiceHandlerClient registers the http handlers for service GroupsBillService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "GroupsBillServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "GroupsBillServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "GroupsBillServiceClient" to call the correct interceptors.
func RegisterGroupsBillServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client GroupsBillServiceClient) error {

	mux.Handle("POST", pattern_GroupsBillService_CreateGroupBill_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/.GroupsBillService/CreateGroupBill", runtime.WithHTTPPathPattern("/v1/groups_bill"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_GroupsBillService_CreateGroupBill_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_GroupsBillService_CreateGroupBill_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("DELETE", pattern_GroupsBillService_DeleteGroupBill_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/.GroupsBillService/DeleteGroupBill", runtime.WithHTTPPathPattern("/v1/groups_bill/{bill_id}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_GroupsBillService_DeleteGroupBill_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_GroupsBillService_DeleteGroupBill_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("PUT", pattern_GroupsBillService_ModifyGroupBill_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/.GroupsBillService/ModifyGroupBill", runtime.WithHTTPPathPattern("/v1/groups_bill/{bill_id}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_GroupsBillService_ModifyGroupBill_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_GroupsBillService_ModifyGroupBill_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_GroupsBillService_GetSelfTotalSpend_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/.GroupsBillService/GetSelfTotalSpend", runtime.WithHTTPPathPattern("/v1/groups_bill/self_total_spend"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_GroupsBillService_GetSelfTotalSpend_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_GroupsBillService_GetSelfTotalSpend_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_GroupsBillService_GetGroupTotalSpend_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/.GroupsBillService/GetGroupTotalSpend", runtime.WithHTTPPathPattern("/v1/groups_bill/group_bill"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_GroupsBillService_GetGroupTotalSpend_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_GroupsBillService_GetGroupTotalSpend_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_GroupsBillService_CreateGroupBill_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "groups_bill"}, ""))

	pattern_GroupsBillService_DeleteGroupBill_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 4, 1, 5, 2}, []string{"v1", "groups_bill", "bill_id"}, ""))

	pattern_GroupsBillService_ModifyGroupBill_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 4, 1, 5, 2}, []string{"v1", "groups_bill", "bill_id"}, ""))

	pattern_GroupsBillService_GetSelfTotalSpend_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "groups_bill", "self_total_spend"}, ""))

	pattern_GroupsBillService_GetGroupTotalSpend_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "groups_bill", "group_bill"}, ""))
)

var (
	forward_GroupsBillService_CreateGroupBill_0 = runtime.ForwardResponseMessage

	forward_GroupsBillService_DeleteGroupBill_0 = runtime.ForwardResponseMessage

	forward_GroupsBillService_ModifyGroupBill_0 = runtime.ForwardResponseMessage

	forward_GroupsBillService_GetSelfTotalSpend_0 = runtime.ForwardResponseMessage

	forward_GroupsBillService_GetGroupTotalSpend_0 = runtime.ForwardResponseMessage
)
