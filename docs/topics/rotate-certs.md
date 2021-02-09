# Rotating Kubernetes Certificates

## Prerequisites

This guide assumes that you already have deployed a cluster using `aks-engine` and the cluster is in a healthy state.

## Certificate Rotation

This document provides guidance on how to rotate certificates on an existing AKS Engine cluster and recommendations for adopting `aks-engine rotate-certs` as a tool.

### Know before you go

In order to ensure that your `aks-engine rotate-certs` operation runs smoothly, there are a few things you should be aware of before getting started.

1. You will need access to the API Model (`apimodel.json`) that was generated by `aks-engine deploy` or `aks-engine generate` (by default this file is placed into a relative directory that looks like `_output/<clustername>/`).

1. An `aks-engine rotate-certs` operation causes API Server downtime.

1. `aks-engine rotate-certs` expects an API model that conforms to the current state of the cluster. `aks-engine rotate-certs` executes remote commands on the cluster nodes and uses the API Model information to establish a secure SSH connection. `aks-engine rotate-certs` also relies on some resources (such as VMs) to be named in accordance with the original `aks-engine` deployment.

1. `aks-engine rotate-certs` relies upon a working connection to the cluster control plane during certificate rotation, both (1) to validate each step of the process, and (2) to restart/recreate cluster resources like `kube-system` pods and service account tokens. If you are rotating the certificates of a **private cluster**, you must run `aks-engine rotate-certs` from a host VM that has network access to the control plane, for example a jumpbox VM that resides in the same VNET as the master VMs. For more information on private clusters [refer to this documentation](features.md#feat-private-cluster).

1. If using `aks-engine rotate-certs` in production, it is recommended to stage a certificate rotation test on an cluster that was built to the same specifications (built with the same cluster configuration + the same version of the `aks-engine` command line tool + the same set of enabled addons) as your production cluster before performing the certificate rotation. The reason for this is that AKS Engine supports many different cluster configurations and the extent of E2E testing that the AKS Engine team runs cannot practically cover every possible configuration. Therefore, it is recommended that you ensure in a staging environment that your specific cluster configuration works with `aks-engine rotate-certs` before attempting this potentially destructive operation on your production cluster.

1. `aks-engine rotate-certs` does **not** guarantees backwards compatibility. If you deployed with `aks-engine` version `0.60.x`, you should prefer executing the certificate rotation process with version `0.60.x`.

### Parameters

|Parameter|Required|Description|
|-----------------|---|---|
|--api-model|yes|Relative path to the API model (cluster definition) that declares the expected cluster configuration.|
|--ssh-host|yes|FQDN, or IP address, of an SSH listener that can reach all nodes in the cluster.|
|--linux-ssh-private-key|yes|Path to a valid private SSH key to access the cluster's Linux nodes.|
|--location|yes|Azure location where the cluster is deployed.|
|--subscription-id|yes|Azure subscription where the cluster infra is deployed.|
|--resource-group|yes|Azure resource group where the cluster infra is deployed.|
|--client-id|depends|The Service Principal Client ID. Required if the auth-method is set to client_secret or client_certificate.|
|--client-secret|depends| The Service Principal Client secret. Required if the auth-method is set to client_secret.|
|--certificate-profile|no|Relative path to a JSON file containing the new set of certificates.|
|--force|no|Force execution even if API Server is not responsive.|

### Simple steps to rotate certificates

Once you have read all the [requirements](#pre-requirements), run `aks-engine rotate-certs` with the appropriate arguments:

```bash
./bin/aks-engine rotate-certs \
  --location <resource-group-location> \
  --api-model <generated-apimodel.json> \
  --linux-ssh-private-key <private-SSH-key> \
  --ssh-host <apiserver-URI> \
  --resource-group <resource-group-name> \
  --client-id <service-principal-id> \
  --client-secret <service-principal-secret> \
  --subscription-id <subscription-id>
```

For example,

```bash
./bin/aks-engine rotate-certs \
  --location "westus2" \
  --api-model "_output/my-cluster/apimodel.json" \
  --linux-ssh-private-key "~/.ssh/id_rsa" \
  --ssh-host "my-cluster.westus2.cloudapp.azure.com"\
  --resource-group "my-cluster" \
  --client-id "12345678-XXXX-YYYY-ZZZZ-1234567890ab" \
  --client-secret "12345678-XXXX-YYYY-ZZZZ-1234567890ab" \
  --subscription-id "12345678-XXXX-YYYY-ZZZZ-1234567890ab"
```

> Fetching a new set of certificates from Key Vault is not supported at this point.

## Under The Hood

A Kubernetes cluster relies on multiple PKIs to secure the communication between its components (apiserver, kubelet, etcd, etc). On an AKS Engine cluster, these multiple PKIs share a single certificate authoritiy (CA). On control plane nodes, `aks-engine rotate-certs` rotates all these PKIs at once and reboots the virtual machine. On agent nodes, only `kubelet` and `kube-proxy` are restarted once the node certificates are replaced.

If the certificate rotation process halts before completion due to a failure or transient issue (e.x.: network connectivity), it is safe to rerun `aks-engine rotate-certs` using the `--force` flag.

At a high level, the `aks-engine rotate-certs` command performs the following tasks:

- backup current set of certificates in directory `_rotate_certs_backup/` (relative to the `--api-model` path)
- generate/load new set of certificate and persist them in local directory `_rotate_certs_output/` (relative to the `--api-model` path)
- distribute certificates to the cluster nodes over SSH
- rotate control plane certificates
- reboot control plane nodes
- rotate agent certificates
- update input `apimodel.json` with new certificates information

### Generating certificates

`aks-engine rotate-certs` is able to generate the new set of certificates that will be deployed to the cluster based on the information found in the API model.

Alternatively, AKS Engine can load a new set of certificates from a JSON file specified in `--certificate-profile`.

```json
{
  "certificateProfile": {
    "caCertificate": "",
    "caPrivateKey": "",
    "apiServerCertificate": "",
    "apiServerPrivateKey": "",
    "clientCertificate": "",
    "clientPrivateKey": "",
    "kubeConfigCertificate": "",
    "kubeConfigPrivateKey": "",
    "etcdServerCertificate": "",
    "etcdServerPrivateKey": "",
    "etcdClientCertificate": "",
    "etcdClientPrivateKey": "",
    "etcdPeerCertificates": ["","",""],
    "etcdPeerPrivateKeys": ["","",""]
  }
}
```

### Certificates distribution

The new certificates are securely copied to each cluster node before the certificates rotation process starts. On Linux nodes, they are located in directory `/etc/kubernetes/rotate-certs/certs`. On Windows nodes, the directory is `$env:temp`.

## Best Practices

### Use a reliable network connection

`aks-engine rotate-certs` requires the execution of multiple remote commands which are subject to potential failures, mostly if the connection to the cluster nodes is not reliable.

Executing `aks-engine rotate-certs` from a VM running on the target cloud (Azure or Azure Stack) can drastically reduce the occurence of transient issues.

## Known Limitations

### Cluster-autoscaler

`aks-engine rotate-certs` will not update the ARM template that `cluster-autoscaler` uses to create new cluster nodes. Because the cluster certificates are embedded in the ARM template, the addon won't be able to produce functioning agent nodes after a certificate rotation operation. An `aks-engine upgrade` operation will take care of updating the "new node" template.

## Troubleshooting

`aks-engine rotate-certs` logs the output of every step in file `/var/log/azure/rotate-certs.log` (Linux) and `c:\k\rotate-certs.log` (Windows).