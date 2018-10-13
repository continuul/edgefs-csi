# CSI driver for EdgeFS (0.4.0 spec conformance)

[Container Storage Interface (CSI)](https://github.com/container-storage-interface/) driver, provisioner, and attacher for EdgeFS Scale-Out NFS

## Overview

EdgeFS CSI plugins implement an interface between CSI enabled Container
Orchestrator (CO) and EdgeFS local cluster site. It allows dynamic and
static provisioning of EdgeFS NFS exports, and attaching them to application
workloads. With EdgeFS NFS implementation, I/O load can be spread-out across
multiple PODs, thus eliminating I/O bottlenecks of classing single-node NFS.
Current implementation of EdgeFS CSI plugins was tested in Kubernetes
environment (requires Kubernetes 1.11+), but the code is Kubernetes version
agnostic and should be able to run with any CSI enabled CO.

For details about configuration and deployment of NFS and EdgeFS CSI plugin,
see Wiki pages:

* [Quick Start Guide](https://github.com/Nexenta/edgefs-csi/wiki/EdgeFS-CSI-Quick-Start-Guide)

## Troubleshooting

Please submit an issue at: [Issues](https://github.com/Nexenta/edgefs-csi/issues)
