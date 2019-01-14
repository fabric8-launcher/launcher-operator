# Launcher Operator

This operator helps enabling the Launcher on an Openshift cluster.


## Install the Launcher operator

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

## Create your Launcher

1. Create your own launcher resource based on the [given example](./deploy/crds/launcher_v1alpha1_launcher_cr.yaml):
2. Create in Openshift
```bash
$ oc create -f deploy/crds/launcher_v1alpha1_launcher_cr.yaml
```


## Example Launcher CR

Find an example of the Launcher CR in `deploy/crds/launcher_v1alpha1_launcher_cr.yaml`

---

## Develop

Install the `operator-sdk` using [the instructions](https://github.com/operator-framework/operator-sdk).

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

Then create your launcher resource and watch the logs in the console ouput.


### Update the launcher template to the latest version from GitHub

```bash
$ ./update-template.sh
```


### Build and push the operator to registry

```bash
$ operator-sdk build fabric8/launcher-operator:vX.Y.Z
$ docker push fabric8/launcher-operator:vX.Y.Z
```

Update the version in [the operator yaml file](./deploy/operator.yaml) and push it.
