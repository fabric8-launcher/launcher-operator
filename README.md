# Launcher Operator

This operator helps enabling the Launcher on an Openshift cluster.

## To run for development using the operator-sdk (on Minishift)

Install the `operator-sdk` [here](https://github.com/operator-framework/operator-sdk).

Register the crd:
```bash
$ oc create -f deploy/crds/launcher_v1alpha1_launcher_crd.yaml  
```

Install dependencies:
```bash 
$ dep ensure -v
```

Start the operator:
```bash 
operator-sdk up local --namespace myproject   
```

## Example Launcher CR

Find a example of the Launcher CR in `deploy/crds/launcher_v1alpha1_launcher_cr.yaml`

## Apply the Launcher CR

The operator and CRD need to be enabled, then:
```bash
$ oc apply -f deploy/crds/launcher_v1alpha1_launcher_cr.yaml
```

## To update the launcher template to it's latest version

```bash
$ ./update-template.sh
```

## Build and push to registry

```bash
$ operator-sdk build fabric8/launcher-operator:v0.0.1
$ docker push fabric8/launcher-operator:v0.0.1
```

## Install the operator

```bash
# Setup Service Account
$ oc create -f deploy/service_account.yaml
# Setup RBAC
$ oc create -f deploy/role.yaml
$ oc create -f deploy/role_binding.yaml
# Setup the CRD
$ oc create -f deploy/crds/launcher_v1alpha1_launcher_crd.yaml
# Deploy the app-operator
$ oc create -f deploy/operator.yaml
```