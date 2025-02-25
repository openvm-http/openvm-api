// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: openvm/v1/api.proto

package v1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/openvm-http/openvm-api/gen/openvm/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// ApiServiceName is the fully-qualified name of the ApiService service.
	ApiServiceName = "openvm.v1.ApiService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ApiServiceCloneVMProcedure is the fully-qualified name of the ApiService's CloneVM RPC.
	ApiServiceCloneVMProcedure = "/openvm.v1.ApiService/CloneVM"
	// ApiServiceDeleteVMProcedure is the fully-qualified name of the ApiService's DeleteVM RPC.
	ApiServiceDeleteVMProcedure = "/openvm.v1.ApiService/DeleteVM"
	// ApiServicePauseVMProcedure is the fully-qualified name of the ApiService's PauseVM RPC.
	ApiServicePauseVMProcedure = "/openvm.v1.ApiService/PauseVM"
	// ApiServiceResumeVMProcedure is the fully-qualified name of the ApiService's ResumeVM RPC.
	ApiServiceResumeVMProcedure = "/openvm.v1.ApiService/ResumeVM"
	// ApiServicePowerOffVMProcedure is the fully-qualified name of the ApiService's PowerOffVM RPC.
	ApiServicePowerOffVMProcedure = "/openvm.v1.ApiService/PowerOffVM"
	// ApiServicePowerOnVMProcedure is the fully-qualified name of the ApiService's PowerOnVM RPC.
	ApiServicePowerOnVMProcedure = "/openvm.v1.ApiService/PowerOnVM"
	// ApiServiceResetVMProcedure is the fully-qualified name of the ApiService's ResetVM RPC.
	ApiServiceResetVMProcedure = "/openvm.v1.ApiService/ResetVM"
	// ApiServiceSuspendVMProcedure is the fully-qualified name of the ApiService's SuspendVM RPC.
	ApiServiceSuspendVMProcedure = "/openvm.v1.ApiService/SuspendVM"
	// ApiServiceReadVMVariableProcedure is the fully-qualified name of the ApiService's ReadVMVariable
	// RPC.
	ApiServiceReadVMVariableProcedure = "/openvm.v1.ApiService/ReadVMVariable"
	// ApiServiceWriteVMVariableProcedure is the fully-qualified name of the ApiService's
	// WriteVMVariable RPC.
	ApiServiceWriteVMVariableProcedure = "/openvm.v1.ApiService/WriteVMVariable"
	// ApiServiceUpgradeVMHardwareProcedure is the fully-qualified name of the ApiService's
	// UpgradeVMHardware RPC.
	ApiServiceUpgradeVMHardwareProcedure = "/openvm.v1.ApiService/UpgradeVMHardware"
	// ApiServiceListRunningVMsProcedure is the fully-qualified name of the ApiService's ListRunningVMs
	// RPC.
	ApiServiceListRunningVMsProcedure = "/openvm.v1.ApiService/ListRunningVMs"
	// ApiServiceVMDetailProcedure is the fully-qualified name of the ApiService's VMDetail RPC.
	ApiServiceVMDetailProcedure = "/openvm.v1.ApiService/VMDetail"
	// ApiServiceCreateVMSnapshotProcedure is the fully-qualified name of the ApiService's
	// CreateVMSnapshot RPC.
	ApiServiceCreateVMSnapshotProcedure = "/openvm.v1.ApiService/CreateVMSnapshot"
	// ApiServiceRevertVMSnapshotProcedure is the fully-qualified name of the ApiService's
	// RevertVMSnapshot RPC.
	ApiServiceRevertVMSnapshotProcedure = "/openvm.v1.ApiService/RevertVMSnapshot"
	// ApiServiceRemoveVMSnapshotProcedure is the fully-qualified name of the ApiService's
	// RemoveVMSnapshot RPC.
	ApiServiceRemoveVMSnapshotProcedure = "/openvm.v1.ApiService/RemoveVMSnapshot"
	// ApiServiceVMSnapshotDetailProcedure is the fully-qualified name of the ApiService's
	// VMSnapshotDetail RPC.
	ApiServiceVMSnapshotDetailProcedure = "/openvm.v1.ApiService/VMSnapshotDetail"
	// ApiServiceVMXRegisteredProcedure is the fully-qualified name of the ApiService's VMXRegistered
	// RPC.
	ApiServiceVMXRegisteredProcedure = "/openvm.v1.ApiService/VMXRegistered"
	// ApiServiceReadVMXVariableProcedure is the fully-qualified name of the ApiService's
	// ReadVMXVariable RPC.
	ApiServiceReadVMXVariableProcedure = "/openvm.v1.ApiService/ReadVMXVariable"
	// ApiServiceWriteVMXVariableProcedure is the fully-qualified name of the ApiService's
	// WriteVMXVariable RPC.
	ApiServiceWriteVMXVariableProcedure = "/openvm.v1.ApiService/WriteVMXVariable"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	apiServiceServiceDescriptor                 = v1.File_openvm_v1_api_proto.Services().ByName("ApiService")
	apiServiceCloneVMMethodDescriptor           = apiServiceServiceDescriptor.Methods().ByName("CloneVM")
	apiServiceDeleteVMMethodDescriptor          = apiServiceServiceDescriptor.Methods().ByName("DeleteVM")
	apiServicePauseVMMethodDescriptor           = apiServiceServiceDescriptor.Methods().ByName("PauseVM")
	apiServiceResumeVMMethodDescriptor          = apiServiceServiceDescriptor.Methods().ByName("ResumeVM")
	apiServicePowerOffVMMethodDescriptor        = apiServiceServiceDescriptor.Methods().ByName("PowerOffVM")
	apiServicePowerOnVMMethodDescriptor         = apiServiceServiceDescriptor.Methods().ByName("PowerOnVM")
	apiServiceResetVMMethodDescriptor           = apiServiceServiceDescriptor.Methods().ByName("ResetVM")
	apiServiceSuspendVMMethodDescriptor         = apiServiceServiceDescriptor.Methods().ByName("SuspendVM")
	apiServiceReadVMVariableMethodDescriptor    = apiServiceServiceDescriptor.Methods().ByName("ReadVMVariable")
	apiServiceWriteVMVariableMethodDescriptor   = apiServiceServiceDescriptor.Methods().ByName("WriteVMVariable")
	apiServiceUpgradeVMHardwareMethodDescriptor = apiServiceServiceDescriptor.Methods().ByName("UpgradeVMHardware")
	apiServiceListRunningVMsMethodDescriptor    = apiServiceServiceDescriptor.Methods().ByName("ListRunningVMs")
	apiServiceVMDetailMethodDescriptor          = apiServiceServiceDescriptor.Methods().ByName("VMDetail")
	apiServiceCreateVMSnapshotMethodDescriptor  = apiServiceServiceDescriptor.Methods().ByName("CreateVMSnapshot")
	apiServiceRevertVMSnapshotMethodDescriptor  = apiServiceServiceDescriptor.Methods().ByName("RevertVMSnapshot")
	apiServiceRemoveVMSnapshotMethodDescriptor  = apiServiceServiceDescriptor.Methods().ByName("RemoveVMSnapshot")
	apiServiceVMSnapshotDetailMethodDescriptor  = apiServiceServiceDescriptor.Methods().ByName("VMSnapshotDetail")
	apiServiceVMXRegisteredMethodDescriptor     = apiServiceServiceDescriptor.Methods().ByName("VMXRegistered")
	apiServiceReadVMXVariableMethodDescriptor   = apiServiceServiceDescriptor.Methods().ByName("ReadVMXVariable")
	apiServiceWriteVMXVariableMethodDescriptor  = apiServiceServiceDescriptor.Methods().ByName("WriteVMXVariable")
)

// ApiServiceClient is a client for the openvm.v1.ApiService service.
type ApiServiceClient interface {
	// 克隆虚拟机
	CloneVM(context.Context, *connect.Request[v1.CloneVMRequest]) (*connect.Response[v1.GenericResponse], error)
	// 删除虚拟机
	DeleteVM(context.Context, *connect.Request[v1.DeleteVMRequest]) (*connect.Response[v1.GenericResponse], error)
	// 暂停虚拟机
	PauseVM(context.Context, *connect.Request[v1.GenericRequest]) (*connect.Response[v1.GenericResponse], error)
	// 恢复已暂停的虚拟机
	ResumeVM(context.Context, *connect.Request[v1.GenericRequest]) (*connect.Response[v1.GenericResponse], error)
	// 关闭虚拟机
	PowerOffVM(context.Context, *connect.Request[v1.PowerOffVMRequest]) (*connect.Response[v1.GenericResponse], error)
	// 开启虚拟机
	PowerOnVM(context.Context, *connect.Request[v1.PowerOnVMRequest]) (*connect.Response[v1.GenericResponse], error)
	// 重置虚拟机
	ResetVM(context.Context, *connect.Request[v1.ResetVMRequest]) (*connect.Response[v1.GenericResponse], error)
	// 挂起虚拟机
	SuspendVM(context.Context, *connect.Request[v1.GenericRequest]) (*connect.Response[v1.GenericResponse], error)
	// 获取虚拟机变量
	ReadVMVariable(context.Context, *connect.Request[v1.ReadVMVariableRequest]) (*connect.Response[v1.ReadVMVariableResponse], error)
	// 写入虚拟机变量
	WriteVMVariable(context.Context, *connect.Request[v1.WriteVMVariableRequest]) (*connect.Response[v1.GenericResponse], error)
	// 升级虚拟机硬件版本
	UpgradeVMHardware(context.Context, *connect.Request[v1.GenericRequest]) (*connect.Response[v1.GenericResponse], error)
	// 列出正在运行的虚拟机列表
	ListRunningVMs(context.Context, *connect.Request[v1.GenericRequest]) (*connect.Response[v1.ListRunningVMsResponse], error)
	// 获取虚拟机信息
	VMDetail(context.Context, *connect.Request[v1.GenericRequest]) (*connect.Response[v1.VMDetailResponse], error)
	// 创建虚拟机快照
	CreateVMSnapshot(context.Context, *connect.Request[v1.CreateVMSnapshotRequest]) (*connect.Response[v1.GenericResponse], error)
	// 恢复虚拟机快照
	RevertVMSnapshot(context.Context, *connect.Request[v1.RevertToSnapshotRequest]) (*connect.Response[v1.GenericResponse], error)
	// 删除虚拟机快照
	RemoveVMSnapshot(context.Context, *connect.Request[v1.RemoveVMSnapshotRequest]) (*connect.Response[v1.GenericResponse], error)
	// 获取虚拟机快照详情
	VMSnapshotDetail(context.Context, *connect.Request[v1.VMSnapshotDetailRequest]) (*connect.Response[v1.VMSnapshotDetailResponse], error)
	// 检查VMX是否已经被注册到Vmware
	VMXRegistered(context.Context, *connect.Request[v1.GenericRequest]) (*connect.Response[v1.VMXRegisteredResponse], error)
	// 读取VMX
	ReadVMXVariable(context.Context, *connect.Request[v1.ReadVMXVariableRequest]) (*connect.Response[v1.ReadVMXVariableResponse], error)
	// 写入VMX
	WriteVMXVariable(context.Context, *connect.Request[v1.WriteVMXVariableRequest]) (*connect.Response[v1.GenericResponse], error)
}

// NewApiServiceClient constructs a client for the openvm.v1.ApiService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewApiServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ApiServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &apiServiceClient{
		cloneVM: connect.NewClient[v1.CloneVMRequest, v1.GenericResponse](
			httpClient,
			baseURL+ApiServiceCloneVMProcedure,
			connect.WithSchema(apiServiceCloneVMMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteVM: connect.NewClient[v1.DeleteVMRequest, v1.GenericResponse](
			httpClient,
			baseURL+ApiServiceDeleteVMProcedure,
			connect.WithSchema(apiServiceDeleteVMMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		pauseVM: connect.NewClient[v1.GenericRequest, v1.GenericResponse](
			httpClient,
			baseURL+ApiServicePauseVMProcedure,
			connect.WithSchema(apiServicePauseVMMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		resumeVM: connect.NewClient[v1.GenericRequest, v1.GenericResponse](
			httpClient,
			baseURL+ApiServiceResumeVMProcedure,
			connect.WithSchema(apiServiceResumeVMMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		powerOffVM: connect.NewClient[v1.PowerOffVMRequest, v1.GenericResponse](
			httpClient,
			baseURL+ApiServicePowerOffVMProcedure,
			connect.WithSchema(apiServicePowerOffVMMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		powerOnVM: connect.NewClient[v1.PowerOnVMRequest, v1.GenericResponse](
			httpClient,
			baseURL+ApiServicePowerOnVMProcedure,
			connect.WithSchema(apiServicePowerOnVMMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		resetVM: connect.NewClient[v1.ResetVMRequest, v1.GenericResponse](
			httpClient,
			baseURL+ApiServiceResetVMProcedure,
			connect.WithSchema(apiServiceResetVMMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		suspendVM: connect.NewClient[v1.GenericRequest, v1.GenericResponse](
			httpClient,
			baseURL+ApiServiceSuspendVMProcedure,
			connect.WithSchema(apiServiceSuspendVMMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		readVMVariable: connect.NewClient[v1.ReadVMVariableRequest, v1.ReadVMVariableResponse](
			httpClient,
			baseURL+ApiServiceReadVMVariableProcedure,
			connect.WithSchema(apiServiceReadVMVariableMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		writeVMVariable: connect.NewClient[v1.WriteVMVariableRequest, v1.GenericResponse](
			httpClient,
			baseURL+ApiServiceWriteVMVariableProcedure,
			connect.WithSchema(apiServiceWriteVMVariableMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		upgradeVMHardware: connect.NewClient[v1.GenericRequest, v1.GenericResponse](
			httpClient,
			baseURL+ApiServiceUpgradeVMHardwareProcedure,
			connect.WithSchema(apiServiceUpgradeVMHardwareMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listRunningVMs: connect.NewClient[v1.GenericRequest, v1.ListRunningVMsResponse](
			httpClient,
			baseURL+ApiServiceListRunningVMsProcedure,
			connect.WithSchema(apiServiceListRunningVMsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		vMDetail: connect.NewClient[v1.GenericRequest, v1.VMDetailResponse](
			httpClient,
			baseURL+ApiServiceVMDetailProcedure,
			connect.WithSchema(apiServiceVMDetailMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createVMSnapshot: connect.NewClient[v1.CreateVMSnapshotRequest, v1.GenericResponse](
			httpClient,
			baseURL+ApiServiceCreateVMSnapshotProcedure,
			connect.WithSchema(apiServiceCreateVMSnapshotMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		revertVMSnapshot: connect.NewClient[v1.RevertToSnapshotRequest, v1.GenericResponse](
			httpClient,
			baseURL+ApiServiceRevertVMSnapshotProcedure,
			connect.WithSchema(apiServiceRevertVMSnapshotMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		removeVMSnapshot: connect.NewClient[v1.RemoveVMSnapshotRequest, v1.GenericResponse](
			httpClient,
			baseURL+ApiServiceRemoveVMSnapshotProcedure,
			connect.WithSchema(apiServiceRemoveVMSnapshotMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		vMSnapshotDetail: connect.NewClient[v1.VMSnapshotDetailRequest, v1.VMSnapshotDetailResponse](
			httpClient,
			baseURL+ApiServiceVMSnapshotDetailProcedure,
			connect.WithSchema(apiServiceVMSnapshotDetailMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		vMXRegistered: connect.NewClient[v1.GenericRequest, v1.VMXRegisteredResponse](
			httpClient,
			baseURL+ApiServiceVMXRegisteredProcedure,
			connect.WithSchema(apiServiceVMXRegisteredMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		readVMXVariable: connect.NewClient[v1.ReadVMXVariableRequest, v1.ReadVMXVariableResponse](
			httpClient,
			baseURL+ApiServiceReadVMXVariableProcedure,
			connect.WithSchema(apiServiceReadVMXVariableMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		writeVMXVariable: connect.NewClient[v1.WriteVMXVariableRequest, v1.GenericResponse](
			httpClient,
			baseURL+ApiServiceWriteVMXVariableProcedure,
			connect.WithSchema(apiServiceWriteVMXVariableMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// apiServiceClient implements ApiServiceClient.
type apiServiceClient struct {
	cloneVM           *connect.Client[v1.CloneVMRequest, v1.GenericResponse]
	deleteVM          *connect.Client[v1.DeleteVMRequest, v1.GenericResponse]
	pauseVM           *connect.Client[v1.GenericRequest, v1.GenericResponse]
	resumeVM          *connect.Client[v1.GenericRequest, v1.GenericResponse]
	powerOffVM        *connect.Client[v1.PowerOffVMRequest, v1.GenericResponse]
	powerOnVM         *connect.Client[v1.PowerOnVMRequest, v1.GenericResponse]
	resetVM           *connect.Client[v1.ResetVMRequest, v1.GenericResponse]
	suspendVM         *connect.Client[v1.GenericRequest, v1.GenericResponse]
	readVMVariable    *connect.Client[v1.ReadVMVariableRequest, v1.ReadVMVariableResponse]
	writeVMVariable   *connect.Client[v1.WriteVMVariableRequest, v1.GenericResponse]
	upgradeVMHardware *connect.Client[v1.GenericRequest, v1.GenericResponse]
	listRunningVMs    *connect.Client[v1.GenericRequest, v1.ListRunningVMsResponse]
	vMDetail          *connect.Client[v1.GenericRequest, v1.VMDetailResponse]
	createVMSnapshot  *connect.Client[v1.CreateVMSnapshotRequest, v1.GenericResponse]
	revertVMSnapshot  *connect.Client[v1.RevertToSnapshotRequest, v1.GenericResponse]
	removeVMSnapshot  *connect.Client[v1.RemoveVMSnapshotRequest, v1.GenericResponse]
	vMSnapshotDetail  *connect.Client[v1.VMSnapshotDetailRequest, v1.VMSnapshotDetailResponse]
	vMXRegistered     *connect.Client[v1.GenericRequest, v1.VMXRegisteredResponse]
	readVMXVariable   *connect.Client[v1.ReadVMXVariableRequest, v1.ReadVMXVariableResponse]
	writeVMXVariable  *connect.Client[v1.WriteVMXVariableRequest, v1.GenericResponse]
}

// CloneVM calls openvm.v1.ApiService.CloneVM.
func (c *apiServiceClient) CloneVM(ctx context.Context, req *connect.Request[v1.CloneVMRequest]) (*connect.Response[v1.GenericResponse], error) {
	return c.cloneVM.CallUnary(ctx, req)
}

// DeleteVM calls openvm.v1.ApiService.DeleteVM.
func (c *apiServiceClient) DeleteVM(ctx context.Context, req *connect.Request[v1.DeleteVMRequest]) (*connect.Response[v1.GenericResponse], error) {
	return c.deleteVM.CallUnary(ctx, req)
}

// PauseVM calls openvm.v1.ApiService.PauseVM.
func (c *apiServiceClient) PauseVM(ctx context.Context, req *connect.Request[v1.GenericRequest]) (*connect.Response[v1.GenericResponse], error) {
	return c.pauseVM.CallUnary(ctx, req)
}

// ResumeVM calls openvm.v1.ApiService.ResumeVM.
func (c *apiServiceClient) ResumeVM(ctx context.Context, req *connect.Request[v1.GenericRequest]) (*connect.Response[v1.GenericResponse], error) {
	return c.resumeVM.CallUnary(ctx, req)
}

// PowerOffVM calls openvm.v1.ApiService.PowerOffVM.
func (c *apiServiceClient) PowerOffVM(ctx context.Context, req *connect.Request[v1.PowerOffVMRequest]) (*connect.Response[v1.GenericResponse], error) {
	return c.powerOffVM.CallUnary(ctx, req)
}

// PowerOnVM calls openvm.v1.ApiService.PowerOnVM.
func (c *apiServiceClient) PowerOnVM(ctx context.Context, req *connect.Request[v1.PowerOnVMRequest]) (*connect.Response[v1.GenericResponse], error) {
	return c.powerOnVM.CallUnary(ctx, req)
}

// ResetVM calls openvm.v1.ApiService.ResetVM.
func (c *apiServiceClient) ResetVM(ctx context.Context, req *connect.Request[v1.ResetVMRequest]) (*connect.Response[v1.GenericResponse], error) {
	return c.resetVM.CallUnary(ctx, req)
}

// SuspendVM calls openvm.v1.ApiService.SuspendVM.
func (c *apiServiceClient) SuspendVM(ctx context.Context, req *connect.Request[v1.GenericRequest]) (*connect.Response[v1.GenericResponse], error) {
	return c.suspendVM.CallUnary(ctx, req)
}

// ReadVMVariable calls openvm.v1.ApiService.ReadVMVariable.
func (c *apiServiceClient) ReadVMVariable(ctx context.Context, req *connect.Request[v1.ReadVMVariableRequest]) (*connect.Response[v1.ReadVMVariableResponse], error) {
	return c.readVMVariable.CallUnary(ctx, req)
}

// WriteVMVariable calls openvm.v1.ApiService.WriteVMVariable.
func (c *apiServiceClient) WriteVMVariable(ctx context.Context, req *connect.Request[v1.WriteVMVariableRequest]) (*connect.Response[v1.GenericResponse], error) {
	return c.writeVMVariable.CallUnary(ctx, req)
}

// UpgradeVMHardware calls openvm.v1.ApiService.UpgradeVMHardware.
func (c *apiServiceClient) UpgradeVMHardware(ctx context.Context, req *connect.Request[v1.GenericRequest]) (*connect.Response[v1.GenericResponse], error) {
	return c.upgradeVMHardware.CallUnary(ctx, req)
}

// ListRunningVMs calls openvm.v1.ApiService.ListRunningVMs.
func (c *apiServiceClient) ListRunningVMs(ctx context.Context, req *connect.Request[v1.GenericRequest]) (*connect.Response[v1.ListRunningVMsResponse], error) {
	return c.listRunningVMs.CallUnary(ctx, req)
}

// VMDetail calls openvm.v1.ApiService.VMDetail.
func (c *apiServiceClient) VMDetail(ctx context.Context, req *connect.Request[v1.GenericRequest]) (*connect.Response[v1.VMDetailResponse], error) {
	return c.vMDetail.CallUnary(ctx, req)
}

// CreateVMSnapshot calls openvm.v1.ApiService.CreateVMSnapshot.
func (c *apiServiceClient) CreateVMSnapshot(ctx context.Context, req *connect.Request[v1.CreateVMSnapshotRequest]) (*connect.Response[v1.GenericResponse], error) {
	return c.createVMSnapshot.CallUnary(ctx, req)
}

// RevertVMSnapshot calls openvm.v1.ApiService.RevertVMSnapshot.
func (c *apiServiceClient) RevertVMSnapshot(ctx context.Context, req *connect.Request[v1.RevertToSnapshotRequest]) (*connect.Response[v1.GenericResponse], error) {
	return c.revertVMSnapshot.CallUnary(ctx, req)
}

// RemoveVMSnapshot calls openvm.v1.ApiService.RemoveVMSnapshot.
func (c *apiServiceClient) RemoveVMSnapshot(ctx context.Context, req *connect.Request[v1.RemoveVMSnapshotRequest]) (*connect.Response[v1.GenericResponse], error) {
	return c.removeVMSnapshot.CallUnary(ctx, req)
}

// VMSnapshotDetail calls openvm.v1.ApiService.VMSnapshotDetail.
func (c *apiServiceClient) VMSnapshotDetail(ctx context.Context, req *connect.Request[v1.VMSnapshotDetailRequest]) (*connect.Response[v1.VMSnapshotDetailResponse], error) {
	return c.vMSnapshotDetail.CallUnary(ctx, req)
}

// VMXRegistered calls openvm.v1.ApiService.VMXRegistered.
func (c *apiServiceClient) VMXRegistered(ctx context.Context, req *connect.Request[v1.GenericRequest]) (*connect.Response[v1.VMXRegisteredResponse], error) {
	return c.vMXRegistered.CallUnary(ctx, req)
}

// ReadVMXVariable calls openvm.v1.ApiService.ReadVMXVariable.
func (c *apiServiceClient) ReadVMXVariable(ctx context.Context, req *connect.Request[v1.ReadVMXVariableRequest]) (*connect.Response[v1.ReadVMXVariableResponse], error) {
	return c.readVMXVariable.CallUnary(ctx, req)
}

// WriteVMXVariable calls openvm.v1.ApiService.WriteVMXVariable.
func (c *apiServiceClient) WriteVMXVariable(ctx context.Context, req *connect.Request[v1.WriteVMXVariableRequest]) (*connect.Response[v1.GenericResponse], error) {
	return c.writeVMXVariable.CallUnary(ctx, req)
}

// ApiServiceHandler is an implementation of the openvm.v1.ApiService service.
type ApiServiceHandler interface {
	// 克隆虚拟机
	CloneVM(context.Context, *connect.Request[v1.CloneVMRequest]) (*connect.Response[v1.GenericResponse], error)
	// 删除虚拟机
	DeleteVM(context.Context, *connect.Request[v1.DeleteVMRequest]) (*connect.Response[v1.GenericResponse], error)
	// 暂停虚拟机
	PauseVM(context.Context, *connect.Request[v1.GenericRequest]) (*connect.Response[v1.GenericResponse], error)
	// 恢复已暂停的虚拟机
	ResumeVM(context.Context, *connect.Request[v1.GenericRequest]) (*connect.Response[v1.GenericResponse], error)
	// 关闭虚拟机
	PowerOffVM(context.Context, *connect.Request[v1.PowerOffVMRequest]) (*connect.Response[v1.GenericResponse], error)
	// 开启虚拟机
	PowerOnVM(context.Context, *connect.Request[v1.PowerOnVMRequest]) (*connect.Response[v1.GenericResponse], error)
	// 重置虚拟机
	ResetVM(context.Context, *connect.Request[v1.ResetVMRequest]) (*connect.Response[v1.GenericResponse], error)
	// 挂起虚拟机
	SuspendVM(context.Context, *connect.Request[v1.GenericRequest]) (*connect.Response[v1.GenericResponse], error)
	// 获取虚拟机变量
	ReadVMVariable(context.Context, *connect.Request[v1.ReadVMVariableRequest]) (*connect.Response[v1.ReadVMVariableResponse], error)
	// 写入虚拟机变量
	WriteVMVariable(context.Context, *connect.Request[v1.WriteVMVariableRequest]) (*connect.Response[v1.GenericResponse], error)
	// 升级虚拟机硬件版本
	UpgradeVMHardware(context.Context, *connect.Request[v1.GenericRequest]) (*connect.Response[v1.GenericResponse], error)
	// 列出正在运行的虚拟机列表
	ListRunningVMs(context.Context, *connect.Request[v1.GenericRequest]) (*connect.Response[v1.ListRunningVMsResponse], error)
	// 获取虚拟机信息
	VMDetail(context.Context, *connect.Request[v1.GenericRequest]) (*connect.Response[v1.VMDetailResponse], error)
	// 创建虚拟机快照
	CreateVMSnapshot(context.Context, *connect.Request[v1.CreateVMSnapshotRequest]) (*connect.Response[v1.GenericResponse], error)
	// 恢复虚拟机快照
	RevertVMSnapshot(context.Context, *connect.Request[v1.RevertToSnapshotRequest]) (*connect.Response[v1.GenericResponse], error)
	// 删除虚拟机快照
	RemoveVMSnapshot(context.Context, *connect.Request[v1.RemoveVMSnapshotRequest]) (*connect.Response[v1.GenericResponse], error)
	// 获取虚拟机快照详情
	VMSnapshotDetail(context.Context, *connect.Request[v1.VMSnapshotDetailRequest]) (*connect.Response[v1.VMSnapshotDetailResponse], error)
	// 检查VMX是否已经被注册到Vmware
	VMXRegistered(context.Context, *connect.Request[v1.GenericRequest]) (*connect.Response[v1.VMXRegisteredResponse], error)
	// 读取VMX
	ReadVMXVariable(context.Context, *connect.Request[v1.ReadVMXVariableRequest]) (*connect.Response[v1.ReadVMXVariableResponse], error)
	// 写入VMX
	WriteVMXVariable(context.Context, *connect.Request[v1.WriteVMXVariableRequest]) (*connect.Response[v1.GenericResponse], error)
}

// NewApiServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewApiServiceHandler(svc ApiServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	apiServiceCloneVMHandler := connect.NewUnaryHandler(
		ApiServiceCloneVMProcedure,
		svc.CloneVM,
		connect.WithSchema(apiServiceCloneVMMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	apiServiceDeleteVMHandler := connect.NewUnaryHandler(
		ApiServiceDeleteVMProcedure,
		svc.DeleteVM,
		connect.WithSchema(apiServiceDeleteVMMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	apiServicePauseVMHandler := connect.NewUnaryHandler(
		ApiServicePauseVMProcedure,
		svc.PauseVM,
		connect.WithSchema(apiServicePauseVMMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	apiServiceResumeVMHandler := connect.NewUnaryHandler(
		ApiServiceResumeVMProcedure,
		svc.ResumeVM,
		connect.WithSchema(apiServiceResumeVMMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	apiServicePowerOffVMHandler := connect.NewUnaryHandler(
		ApiServicePowerOffVMProcedure,
		svc.PowerOffVM,
		connect.WithSchema(apiServicePowerOffVMMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	apiServicePowerOnVMHandler := connect.NewUnaryHandler(
		ApiServicePowerOnVMProcedure,
		svc.PowerOnVM,
		connect.WithSchema(apiServicePowerOnVMMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	apiServiceResetVMHandler := connect.NewUnaryHandler(
		ApiServiceResetVMProcedure,
		svc.ResetVM,
		connect.WithSchema(apiServiceResetVMMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	apiServiceSuspendVMHandler := connect.NewUnaryHandler(
		ApiServiceSuspendVMProcedure,
		svc.SuspendVM,
		connect.WithSchema(apiServiceSuspendVMMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	apiServiceReadVMVariableHandler := connect.NewUnaryHandler(
		ApiServiceReadVMVariableProcedure,
		svc.ReadVMVariable,
		connect.WithSchema(apiServiceReadVMVariableMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	apiServiceWriteVMVariableHandler := connect.NewUnaryHandler(
		ApiServiceWriteVMVariableProcedure,
		svc.WriteVMVariable,
		connect.WithSchema(apiServiceWriteVMVariableMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	apiServiceUpgradeVMHardwareHandler := connect.NewUnaryHandler(
		ApiServiceUpgradeVMHardwareProcedure,
		svc.UpgradeVMHardware,
		connect.WithSchema(apiServiceUpgradeVMHardwareMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	apiServiceListRunningVMsHandler := connect.NewUnaryHandler(
		ApiServiceListRunningVMsProcedure,
		svc.ListRunningVMs,
		connect.WithSchema(apiServiceListRunningVMsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	apiServiceVMDetailHandler := connect.NewUnaryHandler(
		ApiServiceVMDetailProcedure,
		svc.VMDetail,
		connect.WithSchema(apiServiceVMDetailMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	apiServiceCreateVMSnapshotHandler := connect.NewUnaryHandler(
		ApiServiceCreateVMSnapshotProcedure,
		svc.CreateVMSnapshot,
		connect.WithSchema(apiServiceCreateVMSnapshotMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	apiServiceRevertVMSnapshotHandler := connect.NewUnaryHandler(
		ApiServiceRevertVMSnapshotProcedure,
		svc.RevertVMSnapshot,
		connect.WithSchema(apiServiceRevertVMSnapshotMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	apiServiceRemoveVMSnapshotHandler := connect.NewUnaryHandler(
		ApiServiceRemoveVMSnapshotProcedure,
		svc.RemoveVMSnapshot,
		connect.WithSchema(apiServiceRemoveVMSnapshotMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	apiServiceVMSnapshotDetailHandler := connect.NewUnaryHandler(
		ApiServiceVMSnapshotDetailProcedure,
		svc.VMSnapshotDetail,
		connect.WithSchema(apiServiceVMSnapshotDetailMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	apiServiceVMXRegisteredHandler := connect.NewUnaryHandler(
		ApiServiceVMXRegisteredProcedure,
		svc.VMXRegistered,
		connect.WithSchema(apiServiceVMXRegisteredMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	apiServiceReadVMXVariableHandler := connect.NewUnaryHandler(
		ApiServiceReadVMXVariableProcedure,
		svc.ReadVMXVariable,
		connect.WithSchema(apiServiceReadVMXVariableMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	apiServiceWriteVMXVariableHandler := connect.NewUnaryHandler(
		ApiServiceWriteVMXVariableProcedure,
		svc.WriteVMXVariable,
		connect.WithSchema(apiServiceWriteVMXVariableMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/openvm.v1.ApiService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ApiServiceCloneVMProcedure:
			apiServiceCloneVMHandler.ServeHTTP(w, r)
		case ApiServiceDeleteVMProcedure:
			apiServiceDeleteVMHandler.ServeHTTP(w, r)
		case ApiServicePauseVMProcedure:
			apiServicePauseVMHandler.ServeHTTP(w, r)
		case ApiServiceResumeVMProcedure:
			apiServiceResumeVMHandler.ServeHTTP(w, r)
		case ApiServicePowerOffVMProcedure:
			apiServicePowerOffVMHandler.ServeHTTP(w, r)
		case ApiServicePowerOnVMProcedure:
			apiServicePowerOnVMHandler.ServeHTTP(w, r)
		case ApiServiceResetVMProcedure:
			apiServiceResetVMHandler.ServeHTTP(w, r)
		case ApiServiceSuspendVMProcedure:
			apiServiceSuspendVMHandler.ServeHTTP(w, r)
		case ApiServiceReadVMVariableProcedure:
			apiServiceReadVMVariableHandler.ServeHTTP(w, r)
		case ApiServiceWriteVMVariableProcedure:
			apiServiceWriteVMVariableHandler.ServeHTTP(w, r)
		case ApiServiceUpgradeVMHardwareProcedure:
			apiServiceUpgradeVMHardwareHandler.ServeHTTP(w, r)
		case ApiServiceListRunningVMsProcedure:
			apiServiceListRunningVMsHandler.ServeHTTP(w, r)
		case ApiServiceVMDetailProcedure:
			apiServiceVMDetailHandler.ServeHTTP(w, r)
		case ApiServiceCreateVMSnapshotProcedure:
			apiServiceCreateVMSnapshotHandler.ServeHTTP(w, r)
		case ApiServiceRevertVMSnapshotProcedure:
			apiServiceRevertVMSnapshotHandler.ServeHTTP(w, r)
		case ApiServiceRemoveVMSnapshotProcedure:
			apiServiceRemoveVMSnapshotHandler.ServeHTTP(w, r)
		case ApiServiceVMSnapshotDetailProcedure:
			apiServiceVMSnapshotDetailHandler.ServeHTTP(w, r)
		case ApiServiceVMXRegisteredProcedure:
			apiServiceVMXRegisteredHandler.ServeHTTP(w, r)
		case ApiServiceReadVMXVariableProcedure:
			apiServiceReadVMXVariableHandler.ServeHTTP(w, r)
		case ApiServiceWriteVMXVariableProcedure:
			apiServiceWriteVMXVariableHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedApiServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedApiServiceHandler struct{}

func (UnimplementedApiServiceHandler) CloneVM(context.Context, *connect.Request[v1.CloneVMRequest]) (*connect.Response[v1.GenericResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("openvm.v1.ApiService.CloneVM is not implemented"))
}

func (UnimplementedApiServiceHandler) DeleteVM(context.Context, *connect.Request[v1.DeleteVMRequest]) (*connect.Response[v1.GenericResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("openvm.v1.ApiService.DeleteVM is not implemented"))
}

func (UnimplementedApiServiceHandler) PauseVM(context.Context, *connect.Request[v1.GenericRequest]) (*connect.Response[v1.GenericResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("openvm.v1.ApiService.PauseVM is not implemented"))
}

func (UnimplementedApiServiceHandler) ResumeVM(context.Context, *connect.Request[v1.GenericRequest]) (*connect.Response[v1.GenericResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("openvm.v1.ApiService.ResumeVM is not implemented"))
}

func (UnimplementedApiServiceHandler) PowerOffVM(context.Context, *connect.Request[v1.PowerOffVMRequest]) (*connect.Response[v1.GenericResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("openvm.v1.ApiService.PowerOffVM is not implemented"))
}

func (UnimplementedApiServiceHandler) PowerOnVM(context.Context, *connect.Request[v1.PowerOnVMRequest]) (*connect.Response[v1.GenericResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("openvm.v1.ApiService.PowerOnVM is not implemented"))
}

func (UnimplementedApiServiceHandler) ResetVM(context.Context, *connect.Request[v1.ResetVMRequest]) (*connect.Response[v1.GenericResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("openvm.v1.ApiService.ResetVM is not implemented"))
}

func (UnimplementedApiServiceHandler) SuspendVM(context.Context, *connect.Request[v1.GenericRequest]) (*connect.Response[v1.GenericResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("openvm.v1.ApiService.SuspendVM is not implemented"))
}

func (UnimplementedApiServiceHandler) ReadVMVariable(context.Context, *connect.Request[v1.ReadVMVariableRequest]) (*connect.Response[v1.ReadVMVariableResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("openvm.v1.ApiService.ReadVMVariable is not implemented"))
}

func (UnimplementedApiServiceHandler) WriteVMVariable(context.Context, *connect.Request[v1.WriteVMVariableRequest]) (*connect.Response[v1.GenericResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("openvm.v1.ApiService.WriteVMVariable is not implemented"))
}

func (UnimplementedApiServiceHandler) UpgradeVMHardware(context.Context, *connect.Request[v1.GenericRequest]) (*connect.Response[v1.GenericResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("openvm.v1.ApiService.UpgradeVMHardware is not implemented"))
}

func (UnimplementedApiServiceHandler) ListRunningVMs(context.Context, *connect.Request[v1.GenericRequest]) (*connect.Response[v1.ListRunningVMsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("openvm.v1.ApiService.ListRunningVMs is not implemented"))
}

func (UnimplementedApiServiceHandler) VMDetail(context.Context, *connect.Request[v1.GenericRequest]) (*connect.Response[v1.VMDetailResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("openvm.v1.ApiService.VMDetail is not implemented"))
}

func (UnimplementedApiServiceHandler) CreateVMSnapshot(context.Context, *connect.Request[v1.CreateVMSnapshotRequest]) (*connect.Response[v1.GenericResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("openvm.v1.ApiService.CreateVMSnapshot is not implemented"))
}

func (UnimplementedApiServiceHandler) RevertVMSnapshot(context.Context, *connect.Request[v1.RevertToSnapshotRequest]) (*connect.Response[v1.GenericResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("openvm.v1.ApiService.RevertVMSnapshot is not implemented"))
}

func (UnimplementedApiServiceHandler) RemoveVMSnapshot(context.Context, *connect.Request[v1.RemoveVMSnapshotRequest]) (*connect.Response[v1.GenericResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("openvm.v1.ApiService.RemoveVMSnapshot is not implemented"))
}

func (UnimplementedApiServiceHandler) VMSnapshotDetail(context.Context, *connect.Request[v1.VMSnapshotDetailRequest]) (*connect.Response[v1.VMSnapshotDetailResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("openvm.v1.ApiService.VMSnapshotDetail is not implemented"))
}

func (UnimplementedApiServiceHandler) VMXRegistered(context.Context, *connect.Request[v1.GenericRequest]) (*connect.Response[v1.VMXRegisteredResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("openvm.v1.ApiService.VMXRegistered is not implemented"))
}

func (UnimplementedApiServiceHandler) ReadVMXVariable(context.Context, *connect.Request[v1.ReadVMXVariableRequest]) (*connect.Response[v1.ReadVMXVariableResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("openvm.v1.ApiService.ReadVMXVariable is not implemented"))
}

func (UnimplementedApiServiceHandler) WriteVMXVariable(context.Context, *connect.Request[v1.WriteVMXVariableRequest]) (*connect.Response[v1.GenericResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("openvm.v1.ApiService.WriteVMXVariable is not implemented"))
}
