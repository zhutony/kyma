---
title: Runtime Provisioner
type: Details
---

The Runtime Provisioner is a Compass component responsible for provisioning, installing, and deprovisioning clusters with Kyma (Kyma Runtimes). The relationship between clusters and Runtimes is 1:1.

It is powered by [Hydroform](https://github.com/kyma-incubator/hydroform) and it allows you to provision the clusters in the following ways:
- [through Gardener](#tutorials-provision-clusters-through-gardener) on:
    * GCP
    * Microsoft Azure
    * Amazon Web Services (AWS).
    
Note that the operations of provisioning and deprovisioning are asynchronous. They return the operation ID, which you can use to [check the Runtime Operation Status](#tutorials-check-runtime-operation-status).

The Runtime Provisioner also allows you to [clean up Runtime data](#tutorials-clean-up-runtime-data). This operation removes the data for a given Runtime from the database and frees up the Runtime ID for reuse. It is useful when your cluster has died or when the operation of deprovisioning has failed.

> **CAUTION:** Cleaning up Runtime data does not trigger Runtime deprovisioning and the cluster might still exist after the cleanup.
  
The Runtime Provisioner exposes an API to manage cluster provisioning, installation, and deprovisioning. 

Find the specification of the API [here](https://github.com/kyma-incubator/compass/blob/master/components/provisioner/pkg/gqlschema/schema.graphql).
    
To access the Runtime Provisioner, forward the port that the GraphQL Server is listening on:

```bash
kubectl -n compass-system port-forward svc/compass-provisioner 3000:3000
```