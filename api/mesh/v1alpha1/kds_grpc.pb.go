// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1alpha1

import (
	context "context"
	v3 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// KumaDiscoveryServiceClient is the client API for KumaDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KumaDiscoveryServiceClient interface {
	StreamKumaResources(ctx context.Context, opts ...grpc.CallOption) (KumaDiscoveryService_StreamKumaResourcesClient, error)
}

type kumaDiscoveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKumaDiscoveryServiceClient(cc grpc.ClientConnInterface) KumaDiscoveryServiceClient {
	return &kumaDiscoveryServiceClient{cc}
}

func (c *kumaDiscoveryServiceClient) StreamKumaResources(ctx context.Context, opts ...grpc.CallOption) (KumaDiscoveryService_StreamKumaResourcesClient, error) {
	stream, err := c.cc.NewStream(ctx, &KumaDiscoveryService_ServiceDesc.Streams[0], "/kuma.mesh.v1alpha1.KumaDiscoveryService/StreamKumaResources", opts...)
	if err != nil {
		return nil, err
	}
	x := &kumaDiscoveryServiceStreamKumaResourcesClient{stream}
	return x, nil
}

type KumaDiscoveryService_StreamKumaResourcesClient interface {
	Send(*v3.DiscoveryRequest) error
	Recv() (*v3.DiscoveryResponse, error)
	grpc.ClientStream
}

type kumaDiscoveryServiceStreamKumaResourcesClient struct {
	grpc.ClientStream
}

func (x *kumaDiscoveryServiceStreamKumaResourcesClient) Send(m *v3.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *kumaDiscoveryServiceStreamKumaResourcesClient) Recv() (*v3.DiscoveryResponse, error) {
	m := new(v3.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KumaDiscoveryServiceServer is the server API for KumaDiscoveryService service.
// All implementations must embed UnimplementedKumaDiscoveryServiceServer
// for forward compatibility
type KumaDiscoveryServiceServer interface {
	StreamKumaResources(KumaDiscoveryService_StreamKumaResourcesServer) error
	mustEmbedUnimplementedKumaDiscoveryServiceServer()
}

// UnimplementedKumaDiscoveryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKumaDiscoveryServiceServer struct {
}

func (UnimplementedKumaDiscoveryServiceServer) StreamKumaResources(KumaDiscoveryService_StreamKumaResourcesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamKumaResources not implemented")
}
func (UnimplementedKumaDiscoveryServiceServer) mustEmbedUnimplementedKumaDiscoveryServiceServer() {}

// UnsafeKumaDiscoveryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KumaDiscoveryServiceServer will
// result in compilation errors.
type UnsafeKumaDiscoveryServiceServer interface {
	mustEmbedUnimplementedKumaDiscoveryServiceServer()
}

func RegisterKumaDiscoveryServiceServer(s grpc.ServiceRegistrar, srv KumaDiscoveryServiceServer) {
	s.RegisterService(&KumaDiscoveryService_ServiceDesc, srv)
}

func _KumaDiscoveryService_StreamKumaResources_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(KumaDiscoveryServiceServer).StreamKumaResources(&kumaDiscoveryServiceStreamKumaResourcesServer{stream})
}

type KumaDiscoveryService_StreamKumaResourcesServer interface {
	Send(*v3.DiscoveryResponse) error
	Recv() (*v3.DiscoveryRequest, error)
	grpc.ServerStream
}

type kumaDiscoveryServiceStreamKumaResourcesServer struct {
	grpc.ServerStream
}

func (x *kumaDiscoveryServiceStreamKumaResourcesServer) Send(m *v3.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *kumaDiscoveryServiceStreamKumaResourcesServer) Recv() (*v3.DiscoveryRequest, error) {
	m := new(v3.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KumaDiscoveryService_ServiceDesc is the grpc.ServiceDesc for KumaDiscoveryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KumaDiscoveryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kuma.mesh.v1alpha1.KumaDiscoveryService",
	HandlerType: (*KumaDiscoveryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamKumaResources",
			Handler:       _KumaDiscoveryService_StreamKumaResources_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/mesh/v1alpha1/kds.proto",
}

// GlobalKDSServiceClient is the client API for GlobalKDSService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GlobalKDSServiceClient interface {
	// StreamXDSConfigs is logically a service exposed by Zone CP so Global CP can
	// execute Config Dumps. It is however represented by bi-directional streaming
	// to leverage existing connection from Zone CP to Global CP.
	StreamXDSConfigs(ctx context.Context, opts ...grpc.CallOption) (GlobalKDSService_StreamXDSConfigsClient, error)
	// StreamStats is logically a service exposed by Zone CP so Global CP can
	// execute kuma-dp stats requests. It is however represented by bi-directional
	// streaming to leverage existing connection from Zone CP to Global CP.
	StreamStats(ctx context.Context, opts ...grpc.CallOption) (GlobalKDSService_StreamStatsClient, error)
	// StreamStats is logically a service exposed by Zone CP so Global CP can
	// execute kuma-dp clusters request. It is however represented by
	// bi-directional streaming to leverage existing connection from Zone CP to
	// Global CP.
	StreamClusters(ctx context.Context, opts ...grpc.CallOption) (GlobalKDSService_StreamClustersClient, error)
}

type globalKDSServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGlobalKDSServiceClient(cc grpc.ClientConnInterface) GlobalKDSServiceClient {
	return &globalKDSServiceClient{cc}
}

func (c *globalKDSServiceClient) StreamXDSConfigs(ctx context.Context, opts ...grpc.CallOption) (GlobalKDSService_StreamXDSConfigsClient, error) {
	stream, err := c.cc.NewStream(ctx, &GlobalKDSService_ServiceDesc.Streams[0], "/kuma.mesh.v1alpha1.GlobalKDSService/StreamXDSConfigs", opts...)
	if err != nil {
		return nil, err
	}
	x := &globalKDSServiceStreamXDSConfigsClient{stream}
	return x, nil
}

type GlobalKDSService_StreamXDSConfigsClient interface {
	Send(*XDSConfigResponse) error
	Recv() (*XDSConfigRequest, error)
	grpc.ClientStream
}

type globalKDSServiceStreamXDSConfigsClient struct {
	grpc.ClientStream
}

func (x *globalKDSServiceStreamXDSConfigsClient) Send(m *XDSConfigResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *globalKDSServiceStreamXDSConfigsClient) Recv() (*XDSConfigRequest, error) {
	m := new(XDSConfigRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *globalKDSServiceClient) StreamStats(ctx context.Context, opts ...grpc.CallOption) (GlobalKDSService_StreamStatsClient, error) {
	stream, err := c.cc.NewStream(ctx, &GlobalKDSService_ServiceDesc.Streams[1], "/kuma.mesh.v1alpha1.GlobalKDSService/StreamStats", opts...)
	if err != nil {
		return nil, err
	}
	x := &globalKDSServiceStreamStatsClient{stream}
	return x, nil
}

type GlobalKDSService_StreamStatsClient interface {
	Send(*StatsResponse) error
	Recv() (*StatsRequest, error)
	grpc.ClientStream
}

type globalKDSServiceStreamStatsClient struct {
	grpc.ClientStream
}

func (x *globalKDSServiceStreamStatsClient) Send(m *StatsResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *globalKDSServiceStreamStatsClient) Recv() (*StatsRequest, error) {
	m := new(StatsRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *globalKDSServiceClient) StreamClusters(ctx context.Context, opts ...grpc.CallOption) (GlobalKDSService_StreamClustersClient, error) {
	stream, err := c.cc.NewStream(ctx, &GlobalKDSService_ServiceDesc.Streams[2], "/kuma.mesh.v1alpha1.GlobalKDSService/StreamClusters", opts...)
	if err != nil {
		return nil, err
	}
	x := &globalKDSServiceStreamClustersClient{stream}
	return x, nil
}

type GlobalKDSService_StreamClustersClient interface {
	Send(*ClustersResponse) error
	Recv() (*ClustersRequest, error)
	grpc.ClientStream
}

type globalKDSServiceStreamClustersClient struct {
	grpc.ClientStream
}

func (x *globalKDSServiceStreamClustersClient) Send(m *ClustersResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *globalKDSServiceStreamClustersClient) Recv() (*ClustersRequest, error) {
	m := new(ClustersRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GlobalKDSServiceServer is the server API for GlobalKDSService service.
// All implementations must embed UnimplementedGlobalKDSServiceServer
// for forward compatibility
type GlobalKDSServiceServer interface {
	// StreamXDSConfigs is logically a service exposed by Zone CP so Global CP can
	// execute Config Dumps. It is however represented by bi-directional streaming
	// to leverage existing connection from Zone CP to Global CP.
	StreamXDSConfigs(GlobalKDSService_StreamXDSConfigsServer) error
	// StreamStats is logically a service exposed by Zone CP so Global CP can
	// execute kuma-dp stats requests. It is however represented by bi-directional
	// streaming to leverage existing connection from Zone CP to Global CP.
	StreamStats(GlobalKDSService_StreamStatsServer) error
	// StreamStats is logically a service exposed by Zone CP so Global CP can
	// execute kuma-dp clusters request. It is however represented by
	// bi-directional streaming to leverage existing connection from Zone CP to
	// Global CP.
	StreamClusters(GlobalKDSService_StreamClustersServer) error
	mustEmbedUnimplementedGlobalKDSServiceServer()
}

// UnimplementedGlobalKDSServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGlobalKDSServiceServer struct {
}

func (UnimplementedGlobalKDSServiceServer) StreamXDSConfigs(GlobalKDSService_StreamXDSConfigsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamXDSConfigs not implemented")
}
func (UnimplementedGlobalKDSServiceServer) StreamStats(GlobalKDSService_StreamStatsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamStats not implemented")
}
func (UnimplementedGlobalKDSServiceServer) StreamClusters(GlobalKDSService_StreamClustersServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamClusters not implemented")
}
func (UnimplementedGlobalKDSServiceServer) mustEmbedUnimplementedGlobalKDSServiceServer() {}

// UnsafeGlobalKDSServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GlobalKDSServiceServer will
// result in compilation errors.
type UnsafeGlobalKDSServiceServer interface {
	mustEmbedUnimplementedGlobalKDSServiceServer()
}

func RegisterGlobalKDSServiceServer(s grpc.ServiceRegistrar, srv GlobalKDSServiceServer) {
	s.RegisterService(&GlobalKDSService_ServiceDesc, srv)
}

func _GlobalKDSService_StreamXDSConfigs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GlobalKDSServiceServer).StreamXDSConfigs(&globalKDSServiceStreamXDSConfigsServer{stream})
}

type GlobalKDSService_StreamXDSConfigsServer interface {
	Send(*XDSConfigRequest) error
	Recv() (*XDSConfigResponse, error)
	grpc.ServerStream
}

type globalKDSServiceStreamXDSConfigsServer struct {
	grpc.ServerStream
}

func (x *globalKDSServiceStreamXDSConfigsServer) Send(m *XDSConfigRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *globalKDSServiceStreamXDSConfigsServer) Recv() (*XDSConfigResponse, error) {
	m := new(XDSConfigResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GlobalKDSService_StreamStats_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GlobalKDSServiceServer).StreamStats(&globalKDSServiceStreamStatsServer{stream})
}

type GlobalKDSService_StreamStatsServer interface {
	Send(*StatsRequest) error
	Recv() (*StatsResponse, error)
	grpc.ServerStream
}

type globalKDSServiceStreamStatsServer struct {
	grpc.ServerStream
}

func (x *globalKDSServiceStreamStatsServer) Send(m *StatsRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *globalKDSServiceStreamStatsServer) Recv() (*StatsResponse, error) {
	m := new(StatsResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GlobalKDSService_StreamClusters_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GlobalKDSServiceServer).StreamClusters(&globalKDSServiceStreamClustersServer{stream})
}

type GlobalKDSService_StreamClustersServer interface {
	Send(*ClustersRequest) error
	Recv() (*ClustersResponse, error)
	grpc.ServerStream
}

type globalKDSServiceStreamClustersServer struct {
	grpc.ServerStream
}

func (x *globalKDSServiceStreamClustersServer) Send(m *ClustersRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *globalKDSServiceStreamClustersServer) Recv() (*ClustersResponse, error) {
	m := new(ClustersResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GlobalKDSService_ServiceDesc is the grpc.ServiceDesc for GlobalKDSService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GlobalKDSService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kuma.mesh.v1alpha1.GlobalKDSService",
	HandlerType: (*GlobalKDSServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamXDSConfigs",
			Handler:       _GlobalKDSService_StreamXDSConfigs_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamStats",
			Handler:       _GlobalKDSService_StreamStats_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamClusters",
			Handler:       _GlobalKDSService_StreamClusters_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/mesh/v1alpha1/kds.proto",
}

// KDSSyncServiceClient is the client API for KDSSyncService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KDSSyncServiceClient interface {
	// GlobalToZoneSync is logically a service exposed by global control-plane
	// that allows zone control plane to connect and synchronize resources from
	// the global control-plane to the zone control-plane. It uses delta xDS from
	// go-control-plane and responds only with the changes to the resources.
	GlobalToZoneSync(ctx context.Context, opts ...grpc.CallOption) (KDSSyncService_GlobalToZoneSyncClient, error)
	// ZoneToGlobalSync is logically a service exposed by global control-plane
	// that allows zone control plane to connect and synchronize resources to
	// the global control-plane. It uses delta xDS from go-control-plane and
	// responds only with the changes to the resources.
	ZoneToGlobalSync(ctx context.Context, opts ...grpc.CallOption) (KDSSyncService_ZoneToGlobalSyncClient, error)
}

type kDSSyncServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKDSSyncServiceClient(cc grpc.ClientConnInterface) KDSSyncServiceClient {
	return &kDSSyncServiceClient{cc}
}

func (c *kDSSyncServiceClient) GlobalToZoneSync(ctx context.Context, opts ...grpc.CallOption) (KDSSyncService_GlobalToZoneSyncClient, error) {
	stream, err := c.cc.NewStream(ctx, &KDSSyncService_ServiceDesc.Streams[0], "/kuma.mesh.v1alpha1.KDSSyncService/GlobalToZoneSync", opts...)
	if err != nil {
		return nil, err
	}
	x := &kDSSyncServiceGlobalToZoneSyncClient{stream}
	return x, nil
}

type KDSSyncService_GlobalToZoneSyncClient interface {
	Send(*v3.DeltaDiscoveryRequest) error
	Recv() (*v3.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type kDSSyncServiceGlobalToZoneSyncClient struct {
	grpc.ClientStream
}

func (x *kDSSyncServiceGlobalToZoneSyncClient) Send(m *v3.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *kDSSyncServiceGlobalToZoneSyncClient) Recv() (*v3.DeltaDiscoveryResponse, error) {
	m := new(v3.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *kDSSyncServiceClient) ZoneToGlobalSync(ctx context.Context, opts ...grpc.CallOption) (KDSSyncService_ZoneToGlobalSyncClient, error) {
	stream, err := c.cc.NewStream(ctx, &KDSSyncService_ServiceDesc.Streams[1], "/kuma.mesh.v1alpha1.KDSSyncService/ZoneToGlobalSync", opts...)
	if err != nil {
		return nil, err
	}
	x := &kDSSyncServiceZoneToGlobalSyncClient{stream}
	return x, nil
}

type KDSSyncService_ZoneToGlobalSyncClient interface {
	Send(*v3.DeltaDiscoveryResponse) error
	Recv() (*v3.DeltaDiscoveryRequest, error)
	grpc.ClientStream
}

type kDSSyncServiceZoneToGlobalSyncClient struct {
	grpc.ClientStream
}

func (x *kDSSyncServiceZoneToGlobalSyncClient) Send(m *v3.DeltaDiscoveryResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *kDSSyncServiceZoneToGlobalSyncClient) Recv() (*v3.DeltaDiscoveryRequest, error) {
	m := new(v3.DeltaDiscoveryRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KDSSyncServiceServer is the server API for KDSSyncService service.
// All implementations must embed UnimplementedKDSSyncServiceServer
// for forward compatibility
type KDSSyncServiceServer interface {
	// GlobalToZoneSync is logically a service exposed by global control-plane
	// that allows zone control plane to connect and synchronize resources from
	// the global control-plane to the zone control-plane. It uses delta xDS from
	// go-control-plane and responds only with the changes to the resources.
	GlobalToZoneSync(KDSSyncService_GlobalToZoneSyncServer) error
	// ZoneToGlobalSync is logically a service exposed by global control-plane
	// that allows zone control plane to connect and synchronize resources to
	// the global control-plane. It uses delta xDS from go-control-plane and
	// responds only with the changes to the resources.
	ZoneToGlobalSync(KDSSyncService_ZoneToGlobalSyncServer) error
	mustEmbedUnimplementedKDSSyncServiceServer()
}

// UnimplementedKDSSyncServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKDSSyncServiceServer struct {
}

func (UnimplementedKDSSyncServiceServer) GlobalToZoneSync(KDSSyncService_GlobalToZoneSyncServer) error {
	return status.Errorf(codes.Unimplemented, "method GlobalToZoneSync not implemented")
}
func (UnimplementedKDSSyncServiceServer) ZoneToGlobalSync(KDSSyncService_ZoneToGlobalSyncServer) error {
	return status.Errorf(codes.Unimplemented, "method ZoneToGlobalSync not implemented")
}
func (UnimplementedKDSSyncServiceServer) mustEmbedUnimplementedKDSSyncServiceServer() {}

// UnsafeKDSSyncServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KDSSyncServiceServer will
// result in compilation errors.
type UnsafeKDSSyncServiceServer interface {
	mustEmbedUnimplementedKDSSyncServiceServer()
}

func RegisterKDSSyncServiceServer(s grpc.ServiceRegistrar, srv KDSSyncServiceServer) {
	s.RegisterService(&KDSSyncService_ServiceDesc, srv)
}

func _KDSSyncService_GlobalToZoneSync_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(KDSSyncServiceServer).GlobalToZoneSync(&kDSSyncServiceGlobalToZoneSyncServer{stream})
}

type KDSSyncService_GlobalToZoneSyncServer interface {
	Send(*v3.DeltaDiscoveryResponse) error
	Recv() (*v3.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type kDSSyncServiceGlobalToZoneSyncServer struct {
	grpc.ServerStream
}

func (x *kDSSyncServiceGlobalToZoneSyncServer) Send(m *v3.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *kDSSyncServiceGlobalToZoneSyncServer) Recv() (*v3.DeltaDiscoveryRequest, error) {
	m := new(v3.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _KDSSyncService_ZoneToGlobalSync_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(KDSSyncServiceServer).ZoneToGlobalSync(&kDSSyncServiceZoneToGlobalSyncServer{stream})
}

type KDSSyncService_ZoneToGlobalSyncServer interface {
	Send(*v3.DeltaDiscoveryRequest) error
	Recv() (*v3.DeltaDiscoveryResponse, error)
	grpc.ServerStream
}

type kDSSyncServiceZoneToGlobalSyncServer struct {
	grpc.ServerStream
}

func (x *kDSSyncServiceZoneToGlobalSyncServer) Send(m *v3.DeltaDiscoveryRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *kDSSyncServiceZoneToGlobalSyncServer) Recv() (*v3.DeltaDiscoveryResponse, error) {
	m := new(v3.DeltaDiscoveryResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KDSSyncService_ServiceDesc is the grpc.ServiceDesc for KDSSyncService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KDSSyncService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kuma.mesh.v1alpha1.KDSSyncService",
	HandlerType: (*KDSSyncServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GlobalToZoneSync",
			Handler:       _KDSSyncService_GlobalToZoneSync_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ZoneToGlobalSync",
			Handler:       _KDSSyncService_ZoneToGlobalSync_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/mesh/v1alpha1/kds.proto",
}
