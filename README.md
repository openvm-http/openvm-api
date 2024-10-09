# openvm-api
openvm-api provides http-api wrappers for vix and vmx, which allows you to remotely call vmware workstation's virtual machine management functions

## Why
A component called vmrest was introduced into Vmware Workstation. It is undeniable that vmrest is powerful, but it also has some shortcomings, lacks some functions, and is accompanied by some puzzling low-level technical problems (for example, the spelling errors of the API path in Swagger are a disaster). 

After deep thinking, this project was born. It fills some functional gaps through VIX. Although it cannot completely replace vmrest (in fact, no third-party project can completely replace vmrest because vmrest directly calls the core components of VMware Workstation), it can significantly improve the remote control experience of VMware Workstation when used with vmrest. 

It can be understood that openvm-api is an assistant to vmrest, helping it to further optimize the user experience of VMware Workstation, that's all.

## Feature
This API supports:

* Managing virtual machines life cycle: power on, power off, reset, pause and resume.
* Cloning VMs
* Creating and removing Snapshots as well as restoring a VM from a Snapshot
* Edit vmx
* etc, please view proto/openvm/v1/api.proto
