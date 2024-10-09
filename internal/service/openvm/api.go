package openvm

import "C"
import (
	"connectrpc.com/connect"
	"context"
	"errors"
	"fmt"
	"github.com/openvm-http/govix"
	"github.com/openvm-http/govmx/v2"
	"github.com/openvm-http/openvm-api/gen/openvm/v1"
	"github.com/openvm-http/openvm-api/internal/helper"
	"log"
	"os"
)

type ApiServer struct{}

// RegisterVM 注册虚拟机
func (s *ApiServer) RegisterVM(ctx context.Context, req *connect.Request[v1.RegisterVMRequest]) (*connect.Response[v1.GenericResponse], error) {
	connectErr := connect.NewError(connect.CodeInternal, errors.New("API_ERROR"))
	h, err := helper.VixNewHost()
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}
	defer func(hostConn *govix.Host) {
		err := helper.VixFreeHost(hostConn)
		if err != nil {
			log.Printf("helper.VixFreeHost Err: %v\n", err)
		}
	}(h)

	err = h.RegisterVM(req.Msg.VmxFilePath)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}

	return connect.NewResponse(&v1.GenericResponse{}), nil
}

// UnregisterVM 移除虚拟机
func (s *ApiServer) UnregisterVM(ctx context.Context, req *connect.Request[v1.UnregisterVMRequest]) (*connect.Response[v1.GenericResponse], error) {
	connectErr := connect.NewError(connect.CodeInternal, errors.New("API_ERROR"))
	h, err := helper.VixNewHost()
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}
	defer func(hostConn *govix.Host) {
		err := helper.VixFreeHost(hostConn)
		if err != nil {
			log.Printf("helper.VixFreeHost Err: %v\n", err)
		}
	}(h)

	err = h.UnregisterVM(req.Msg.VmxFilePath)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}

	return connect.NewResponse(&v1.GenericResponse{}), nil
}

// CloneVM 克隆虚拟机
func (s *ApiServer) CloneVM(ctx context.Context, req *connect.Request[v1.CloneVMRequest]) (*connect.Response[v1.GenericResponse], error) {
	connectErr := connect.NewError(connect.CodeInternal, errors.New("API_ERROR"))
	v, h, err := helper.VixNewVM(req.Msg.VmxFilePathSrc, "", nil)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}
	defer func(hostConn *govix.Host) {
		err := helper.VixFreeHost(hostConn)
		if err != nil {
			log.Printf("helper.VixFreeHost Err: %v\n", err)
		}
	}(h)

	var cloneType govix.CloneType
	switch req.Msg.Type {
	case v1.CloneVMRequest_FULL:
		cloneType = govix.CLONETYPE_FULL
	case v1.CloneVMRequest_LINKED:
		cloneType = govix.CLONETYPE_LINKED
	}
	_, err = v.Clone(cloneType, req.Msg.VmxFilePathDest)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}

	return connect.NewResponse(&v1.GenericResponse{}), nil
}

// DeleteVM 删除虚拟机
func (s *ApiServer) DeleteVM(ctx context.Context, req *connect.Request[v1.DeleteVMRequest]) (*connect.Response[v1.GenericResponse], error) {
	connectErr := connect.NewError(connect.CodeInternal, errors.New("API_ERROR"))
	v, h, err := helper.VixNewVM(req.Msg.VmxFilePath, "", nil)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}
	defer func(hostConn *govix.Host) {
		err := helper.VixFreeHost(hostConn)
		if err != nil {
			log.Printf("helper.VixFreeHost Err: %v\n", err)
		}
	}(h)

	var deleteType govix.VmDeleteOption
	switch req.Msg.Type {
	case v1.DeleteVMRequest_DELETE_DISK_FILES:
		deleteType = govix.VMDELETE_FORCE
	case v1.DeleteVMRequest_DELETE_KEEP_FILES:
		deleteType = govix.VMDELETE_KEEP_FILES
	}
	err = v.Delete(deleteType)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}

	return connect.NewResponse(&v1.GenericResponse{}), nil
}

// PauseVM 暂停虚拟机
func (s *ApiServer) PauseVM(ctx context.Context, req *connect.Request[v1.GenericRequest]) (*connect.Response[v1.GenericResponse], error) {
	connectErr := connect.NewError(connect.CodeInternal, errors.New("API_ERROR"))
	v, h, err := helper.VixNewVM(req.Msg.VmxFilePath, "", nil)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}
	defer func(hostConn *govix.Host) {
		err := helper.VixFreeHost(hostConn)
		if err != nil {
			log.Printf("helper.VixFreeHost Err: %v\n", err)
		}
	}(h)

	err = v.Pause()
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}

	return connect.NewResponse(&v1.GenericResponse{}), nil
}

// ResumeVM 恢复已暂停的虚拟机
func (s *ApiServer) ResumeVM(ctx context.Context, req *connect.Request[v1.GenericRequest]) (*connect.Response[v1.GenericResponse], error) {
	connectErr := connect.NewError(connect.CodeInternal, errors.New("API_ERROR"))
	v, h, err := helper.VixNewVM(req.Msg.VmxFilePath, "", nil)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}
	defer func(hostConn *govix.Host) {
		err := helper.VixFreeHost(hostConn)
		if err != nil {
			log.Printf("helper.VixFreeHost Err: %v\n", err)
		}
	}(h)

	err = v.Resume()
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}

	return connect.NewResponse(&v1.GenericResponse{}), nil
}

// PowerOffVM 关闭虚拟机
func (s *ApiServer) PowerOffVM(ctx context.Context, req *connect.Request[v1.PowerOffVMRequest]) (*connect.Response[v1.GenericResponse], error) {
	connectErr := connect.NewError(connect.CodeInternal, errors.New("API_ERROR"))
	v, h, err := helper.VixNewVM(req.Msg.VmxFilePath, "", nil)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}
	defer func(hostConn *govix.Host) {
		err := helper.VixFreeHost(hostConn)
		if err != nil {
			log.Printf("helper.VixFreeHost Err: %v\n", err)
		}
	}(h)

	var powerType govix.VMPowerOption
	switch req.Msg.Type {
	case v1.PowerOffVMRequest_FROM_GUEST:
		powerType = govix.VMPOWEROP_FROM_GUEST
	case v1.PowerOffVMRequest_NORMAL:
		powerType = govix.VMPOWEROP_NORMAL
	}
	err = v.PowerOff(powerType)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}

	return connect.NewResponse(&v1.GenericResponse{}), nil
}

// PowerOnVM 开启虚拟机
func (s *ApiServer) PowerOnVM(ctx context.Context, req *connect.Request[v1.PowerOnVMRequest]) (*connect.Response[v1.GenericResponse], error) {
	connectErr := connect.NewError(connect.CodeInternal, errors.New("API_ERROR"))
	v, h, err := helper.VixNewVM(req.Msg.VmxFilePath, "", nil)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}
	defer func(hostConn *govix.Host) {
		err := helper.VixFreeHost(hostConn)
		if err != nil {
			log.Printf("helper.VixFreeHost Err: %v\n", err)
		}
	}(h)

	var powerType govix.VMPowerOption
	switch req.Msg.Type {
	case v1.PowerOnVMRequest_NORMAL:
		powerType = govix.VMPOWEROP_NORMAL
	case v1.PowerOnVMRequest_LAUNCH_GUI:
		powerType = govix.VMPOWEROP_LAUNCH_GUI
	}
	err = v.PowerOn(powerType)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}

	return connect.NewResponse(&v1.GenericResponse{}), nil
}

// ResetVM 重置虚拟机
func (s *ApiServer) ResetVM(ctx context.Context, req *connect.Request[v1.ResetVMRequest]) (*connect.Response[v1.GenericResponse], error) {
	connectErr := connect.NewError(connect.CodeInternal, errors.New("API_ERROR"))
	v, h, err := helper.VixNewVM(req.Msg.VmxFilePath, "", nil)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}
	defer func(hostConn *govix.Host) {
		err := helper.VixFreeHost(hostConn)
		if err != nil {
			log.Printf("helper.VixFreeHost Err: %v\n", err)
		}
	}(h)

	var powerType govix.VMPowerOption
	switch req.Msg.Type {
	case v1.ResetVMRequest_FROM_GUEST:
		powerType = govix.VMPOWEROP_FROM_GUEST
	case v1.ResetVMRequest_NORMAL:
		powerType = govix.VMPOWEROP_NORMAL
	}
	err = v.Reset(powerType)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}

	return connect.NewResponse(&v1.GenericResponse{}), nil
}

// SuspendVM 挂起虚拟机
func (s *ApiServer) SuspendVM(ctx context.Context, req *connect.Request[v1.GenericRequest]) (*connect.Response[v1.GenericResponse], error) {
	connectErr := connect.NewError(connect.CodeInternal, errors.New("API_ERROR"))
	v, h, err := helper.VixNewVM(req.Msg.VmxFilePath, "", nil)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}
	defer func(hostConn *govix.Host) {
		err := helper.VixFreeHost(hostConn)
		if err != nil {
			log.Printf("helper.VixFreeHost Err: %v\n", err)
		}
	}(h)

	err = v.Suspend()
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}

	return connect.NewResponse(&v1.GenericResponse{}), nil
}

// ReadVMVariable 获取虚拟机变量
func (s *ApiServer) ReadVMVariable(ctx context.Context, req *connect.Request[v1.ReadVMVariableRequest]) (*connect.Response[v1.ReadVMVariableResponse], error) {
	connectErr := connect.NewError(connect.CodeInternal, errors.New("API_ERROR"))
	v, h, err := helper.VixNewVM(req.Msg.VmxFilePath, "", nil)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}
	defer func(hostConn *govix.Host) {
		err := helper.VixFreeHost(hostConn)
		if err != nil {
			log.Printf("helper.VixFreeHost Err: %v\n", err)
		}
	}(h)

	resp := &v1.ReadVMVariableResponse{}
	for _, _v := range req.Msg.Vars {
		var varType govix.GuestVarType
		switch _v.Type {
		case v1.ReadVMVariableRequest_ReadVMVariableRequestVar_CONFIG_RUNTIME_ONLY:
			varType = govix.VM_CONFIG_RUNTIME_ONLY
		case v1.ReadVMVariableRequest_ReadVMVariableRequestVar_GUEST_ENVIRONMENT_VARIABLE:
			varType = govix.GUEST_ENVIRONMENT_VARIABLE
		case v1.ReadVMVariableRequest_ReadVMVariableRequestVar_GUEST_VARIABLE:
			varType = govix.VM_GUEST_VARIABLE
		}
		value, err := v.ReadVariable(varType, _v.Name)
		if err != nil {
			return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
				Message: fmt.Sprintf("Read VM Variable[%s | %s] Err: %v", _v.Type, _v.Name, err),
				Module:  v1.ErrDetail_VIX,
			})
		}
		resp.Vars = append(resp.Vars, &v1.ReadVMVariableResponse_ReadVMVariableResponseVar{
			Name:  _v.Name,
			Value: value,
		})
	}

	return connect.NewResponse(resp), nil
}

// WriteVMVariable 写入虚拟机变量
func (s *ApiServer) WriteVMVariable(ctx context.Context, req *connect.Request[v1.WriteVMVariableRequest]) (*connect.Response[v1.GenericResponse], error) {
	connectErr := connect.NewError(connect.CodeInternal, errors.New("API_ERROR"))
	v, h, err := helper.VixNewVM(req.Msg.VmxFilePath, "", nil)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}
	defer func(hostConn *govix.Host) {
		err := helper.VixFreeHost(hostConn)
		if err != nil {
			log.Printf("helper.VixFreeHost Err: %v\n", err)
		}
	}(h)

	for _, _v := range req.Msg.Vars {
		var varType govix.GuestVarType
		switch _v.Type {
		case v1.WriteVMVariableRequest_WriteVMVariableRequestVar_CONFIG_RUNTIME_ONLY:
			varType = govix.VM_CONFIG_RUNTIME_ONLY
		case v1.WriteVMVariableRequest_WriteVMVariableRequestVar_GUEST_ENVIRONMENT_VARIABLE:
			varType = govix.GUEST_ENVIRONMENT_VARIABLE
		case v1.WriteVMVariableRequest_WriteVMVariableRequestVar_GUEST_VARIABLE:
			varType = govix.VM_GUEST_VARIABLE
		}
		err := v.WriteVariable(varType, _v.Name, _v.Value)
		if err != nil {
			return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
				Message: fmt.Sprintf("Write VM Variable[%s | %s] Err: %v", _v.Type, _v.Name, err),
				Module:  v1.ErrDetail_VIX,
			})
		}
	}

	return connect.NewResponse(&v1.GenericResponse{}), nil
}

// UpgradeVMHardware 升级虚拟机硬件版本
func (s *ApiServer) UpgradeVMHardware(ctx context.Context, req *connect.Request[v1.GenericRequest]) (*connect.Response[v1.GenericResponse], error) {
	connectErr := connect.NewError(connect.CodeInternal, errors.New("API_ERROR"))
	v, h, err := helper.VixNewVM(req.Msg.VmxFilePath, "", nil)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}
	defer func(hostConn *govix.Host) {
		err := helper.VixFreeHost(hostConn)
		if err != nil {
			log.Printf("helper.VixFreeHost Err: %v\n", err)
		}
	}(h)

	err = v.UpgradeVHardware()
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}

	return connect.NewResponse(&v1.GenericResponse{}), nil
}

// ListRunningVMs 列出正在运行的虚拟机列表
func (s *ApiServer) ListRunningVMs(ctx context.Context, req *connect.Request[v1.GenericRequest]) (*connect.Response[v1.ListRunningVMsResponse], error) {
	connectErr := connect.NewError(connect.CodeInternal, errors.New("API_ERROR"))
	h, err := helper.VixNewHost()
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}
	defer func(hostConn *govix.Host) {
		err := helper.VixFreeHost(hostConn)
		if err != nil {
			log.Printf("helper.VixFreeHost Err: %v\n", err)
		}
	}(h)

	items, err := h.FindItems(govix.FIND_RUNNING_VMS)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}
	resp := &v1.ListRunningVMsResponse{}
	for _, _v := range items {
		resp.Vms = append(resp.Vms, &v1.ListRunningVMsResponse_ListVMsResponseVM{
			VmxFilePath: _v,
		})
	}

	return connect.NewResponse(resp), nil
}

// VMDetail 获取虚拟机信息
func (s *ApiServer) VMDetail(ctx context.Context, req *connect.Request[v1.GenericRequest]) (*connect.Response[v1.VMDetailResponse], error) {
	connectErr := connect.NewError(connect.CodeInternal, errors.New("API_ERROR"))
	v, h, err := helper.VixNewVM(req.Msg.VmxFilePath, "", nil)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}
	defer func(hostConn *govix.Host) {
		err := helper.VixFreeHost(hostConn)
		if err != nil {
			log.Printf("helper.VixFreeHost Err: %v\n", err)
		}
	}(h)

	resp := &v1.VMDetailResponse{}
	if vcpus, err := v.Vcpus(); err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	} else {
		resp.Vcpu = uint32(vcpus)
	}
	if mem, err := v.MemorySize(); err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	} else {
		resp.Memory = uint64(mem)
	}
	if powerStatus, err := v.PowerState(); err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	} else {
		switch powerStatus {
		case govix.POWERSTATE_POWERING_OFF:
			resp.PowerStatus = v1.VMDetailResponse_VM_POWERING_OFF
		case govix.POWERSTATE_POWERED_OFF:
			resp.PowerStatus = v1.VMDetailResponse_VM_POWERED_OFF
		case govix.POWERSTATE_POWERING_ON:
			resp.PowerStatus = v1.VMDetailResponse_VM_POWERING_ON
		case govix.POWERSTATE_POWERED_ON:
			resp.PowerStatus = v1.VMDetailResponse_VM_POWERED_ON
		case govix.POWERSTATE_SUSPENDING:
			resp.PowerStatus = v1.VMDetailResponse_VM_SUSPENDING
		case govix.POWERSTATE_SUSPENDED:
			resp.PowerStatus = v1.VMDetailResponse_VM_SUSPENDED
		case govix.POWERSTATE_TOOLS_RUNNING:
			resp.PowerStatus = v1.VMDetailResponse_VM_TOOLS_RUNNING
		case govix.POWERSTATE_RESETTING:
			resp.PowerStatus = v1.VMDetailResponse_VM_RESETTING
		case govix.POWERSTATE_BLOCKED_ON_MSG:
			resp.PowerStatus = v1.VMDetailResponse_VM_BLOCKED_ON_MSG
		}
	}
	if vmToolsStatus, err := v.ToolsState(); err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	} else {
		switch vmToolsStatus {
		case govix.TOOLSSTATE_UNKNOWN:
			resp.VmToolsStatus = v1.VMDetailResponse_TOOLS_UNKNOWN
		case govix.TOOLSSTATE_RUNNING:
			resp.VmToolsStatus = v1.VMDetailResponse_TOOLS_RUNNING
		case govix.TOOLSSTATE_NOT_INSTALLED:
			resp.VmToolsStatus = v1.VMDetailResponse_TOOLS_NOT_INSTALLED
		}
	}
	if resp.GuestOS, err = v.GuestOS(); err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}

	return connect.NewResponse(resp), nil
}

// CreateVMSnapshot 创建虚拟机快照
func (s *ApiServer) CreateVMSnapshot(ctx context.Context, req *connect.Request[v1.CreateVMSnapshotRequest]) (*connect.Response[v1.GenericResponse], error) {
	connectErr := connect.NewError(connect.CodeInternal, errors.New("API_ERROR"))
	v, h, err := helper.VixNewVM(req.Msg.VmxFilePath, "", nil)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}
	defer func(hostConn *govix.Host) {
		err := helper.VixFreeHost(hostConn)
		if err != nil {
			log.Printf("helper.VixFreeHost Err: %v\n", err)
		}
	}(h)

	var snapshotType govix.CreateSnapshotOption
	switch req.Msg.Type {
	case v1.CreateVMSnapshotRequest_INCLUDE_MEMORY:
		snapshotType = govix.SNAPSHOT_INCLUDE_MEMORY
	case v1.CreateVMSnapshotRequest_NORMAL:
		snapshotType = 0x0
	}
	_, err = v.CreateSnapshot(req.Msg.Name, req.Msg.Description, snapshotType)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}

	return connect.NewResponse(&v1.GenericResponse{}), nil
}

// RevertVMSnapshot 恢复虚拟机快照
func (s *ApiServer) RevertVMSnapshot(ctx context.Context, req *connect.Request[v1.RevertToSnapshotRequest]) (*connect.Response[v1.GenericResponse], error) {
	connectErr := connect.NewError(connect.CodeInternal, errors.New("API_ERROR"))
	v, h, err := helper.VixNewVM(req.Msg.VmxFilePath, "", nil)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}
	defer func(hostConn *govix.Host) {
		err := helper.VixFreeHost(hostConn)
		if err != nil {
			log.Printf("helper.VixFreeHost Err: %v\n", err)
		}
	}(h)

	_s, err := v.SnapshotByName(req.Msg.Name)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}
	err = v.RevertToSnapshot(_s, govix.VMPOWEROP_SUPPRESS_SNAPSHOT_POWERON)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}

	return connect.NewResponse(&v1.GenericResponse{}), nil
}

// RemoveVMSnapshot 删除虚拟机快照
func (s *ApiServer) RemoveVMSnapshot(ctx context.Context, req *connect.Request[v1.RemoveVMSnapshotRequest]) (*connect.Response[v1.GenericResponse], error) {
	connectErr := connect.NewError(connect.CodeInternal, errors.New("API_ERROR"))
	v, h, err := helper.VixNewVM(req.Msg.VmxFilePath, "", nil)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}
	defer func(hostConn *govix.Host) {
		err := helper.VixFreeHost(hostConn)
		if err != nil {
			log.Printf("helper.VixFreeHost Err: %v\n", err)
		}
	}(h)

	_s, err := v.SnapshotByName(req.Msg.Name)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}
	var removeSnapshotType govix.RemoveSnapshotOption
	switch req.Msg.Type {
	case v1.RemoveVMSnapshotRequest_REMOVE_CHILDREN:
		removeSnapshotType = govix.SNAPSHOT_REMOVE_CHILDREN
	case v1.RemoveVMSnapshotRequest_DEFAULT:
		removeSnapshotType = govix.SNAPSHOT_REMOVE_NONE
	}
	err = v.RemoveSnapshot(_s, removeSnapshotType)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}

	return connect.NewResponse(&v1.GenericResponse{}), nil
}

// VMSnapshotDetail 获取虚拟机快照详情
func (s *ApiServer) VMSnapshotDetail(ctx context.Context, req *connect.Request[v1.VMSnapshotDetailRequest]) (*connect.Response[v1.VMSnapshotDetailResponse], error) {
	connectErr := connect.NewError(connect.CodeInternal, errors.New("API_ERROR"))
	v, h, err := helper.VixNewVM(req.Msg.VmxFilePath, "", nil)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}
	defer func(hostConn *govix.Host) {
		err := helper.VixFreeHost(hostConn)
		if err != nil {
			log.Printf("helper.VixFreeHost Err: %v\n", err)
		}
	}(h)

	_s, err := v.SnapshotByName(req.Msg.Name)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}
	resp := &v1.VMSnapshotDetailResponse{}
	if resp.Name, err = _s.Name(); err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}
	if resp.Description, err = _s.Description(); err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}

	return connect.NewResponse(resp), nil
}

// VMXRegistered 检查VMX是否已经被注册到Vmware
func (s *ApiServer) VMXRegistered(ctx context.Context, req *connect.Request[v1.GenericRequest]) (*connect.Response[v1.VMXRegisteredResponse], error) {
	connectErr := connect.NewError(connect.CodeInternal, errors.New("API_ERROR"))
	h, err := helper.VixNewHost()
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VIX,
		})
	}
	defer func(hostConn *govix.Host) {
		err := helper.VixFreeHost(hostConn)
		if err != nil {
			log.Printf("helper.VixFreeHost Err: %v\n", err)
		}
	}(h)

	resp := &v1.VMXRegisteredResponse{
		Registered: true,
	}
	_, err = h.OpenVM(req.Msg.VmxFilePath, "")
	if err != nil {
		var vixError *govix.Error
		if !errors.As(err, &vixError) {
			return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
				Message: "errors.As(err, &vixError) -> false",
				Module:  v1.ErrDetail_SERVICE,
			})
		}
		if vixError.Code != 4000 {
			return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
				Message: err.Error(),
				Module:  v1.ErrDetail_VIX,
			})
		}
		resp.Registered = false
	}

	return connect.NewResponse(resp), nil
}

// ReadVMXVariable 读取VMX
func (s *ApiServer) ReadVMXVariable(ctx context.Context, req *connect.Request[v1.ReadVMXVariableRequest]) (*connect.Response[v1.ReadVMXVariableResponse], error) {
	connectErr := connect.NewError(connect.CodeInternal, errors.New("API_ERROR"))

	// 检查文件是否存在
	_, err := os.Stat(req.Msg.VmxFilePath)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_FILESYSTEM,
		})
	}
	vmxFileContent, err := os.ReadFile(req.Msg.VmxFilePath)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_FILESYSTEM,
		})
	}

	err, o := govmx.SafeUnmarshal(vmxFileContent, nil, true)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VMX,
		})
	}
	resp := &v1.ReadVMXVariableResponse{}
	for _, v := range req.Msg.Vars {
		respVar := &v1.ReadVMXVariableResponse_ReadVMXVariableResponseVar{
			Name: v.Name,
		}
		respVar.Value, _ = o[v.Name]
		resp.Vars = append(resp.Vars, respVar)
	}

	return connect.NewResponse(resp), nil
}

// WriteVMXVariable 写入VMX
func (s *ApiServer) WriteVMXVariable(ctx context.Context, req *connect.Request[v1.WriteVMXVariableRequest]) (*connect.Response[v1.GenericResponse], error) {
	connectErr := connect.NewError(connect.CodeInternal, errors.New("API_ERROR"))

	// 检查文件是否存在
	vmxFileInfo, err := os.Stat(req.Msg.VmxFilePath)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_FILESYSTEM,
		})
	}
	vmxFileContent, err := os.ReadFile(req.Msg.VmxFilePath)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_FILESYSTEM,
		})
	}

	err, o := govmx.SafeUnmarshal(vmxFileContent, nil, true)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VMX,
		})
	}
	for _, v := range req.Msg.Vars {
		if v.Name == "" {
			continue
		}
		if v.Value == "" {
			if _, ok := o[v.Name]; ok {
				delete(o, v.Name)
			}
		} else {
			o[v.Name] = v.Value
		}
	}
	vmxFileContent, err = govmx.SafeMarshal(nil, o, true)
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_VMX,
		})
	}
	err = os.WriteFile(req.Msg.VmxFilePath, vmxFileContent, vmxFileInfo.Mode())
	if err != nil {
		return nil, helper.WrapConnectErrorDetail(connectErr, &v1.ErrDetail{
			Message: err.Error(),
			Module:  v1.ErrDetail_FILESYSTEM,
		})
	}

	return connect.NewResponse(&v1.GenericResponse{}), nil
}
