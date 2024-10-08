# openvm-api
openvm-api provides http-api wrappers for vix and vmx, which allows you to remotely call vmware workstation's virtual machine management functions

## Feature
This API supports:

* Managing virtual machines life cycle: power on, power off, reset, pause and resume.
* Cloning VMs
* Creating and removing Snapshots as well as restoring a VM from a Snapshot
* Edit vmx
* etc, please view proto/openvm/v1/api.proto
