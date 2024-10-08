package helper

import (
	"errors"
	"github.com/openvm-http/govix"
)

var vixHostConfig govix.ConnectConfig

func init() {
	vixHostConfig = govix.ConnectConfig{
		Provider: govix.VMWARE_WORKSTATION,
		Options:  govix.HOST_OPTIONS_NONE,
	}
}

func VixNewHost() (*govix.Host, error) {
	return govix.Connect(vixHostConfig)
}

func VixNewVM(vmxFilePath, password string, hostConn *govix.Host) (*govix.VM, *govix.Host, error) {
	if vmxFilePath == "" {
		return nil, nil, errors.New("VIX Wrapper Error: vmx file path is empty")
	}
	var err error
	_hostConn := hostConn
	if hostConn == nil {
		_hostConn, err = VixNewHost()
		if err != nil {
			return nil, nil, err
		}
	}
	vm, err := _hostConn.OpenVM(vmxFilePath, password)
	return vm, hostConn, err
}

func VixFreeHost(hostConn *govix.Host) error {
	if hostConn != nil {
		hostConn.Disconnect()
	}
	return nil
}
