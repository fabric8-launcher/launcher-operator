# Launcher Operator

This operator helps enabling the Launcher on an Openshift cluster.


## Give cluster admin permissions to the user <user>:

The user needs cluster-admin permissions to install the Launcher operator

```bash
$ oc adm policy --as system:admin add-cluster-role-to-user cluster-admin <user>
```

## Install the Launcher operator

Login with the OpenShift client using a user with cluster-admin permissions.
```bash
$ oc login
```

Choose the project that will run the operator and then install all the operator resources:

```bash
oc create -R -f ./deploy 
```

## Install the Launcher (via the installed operator)


1. Log into GitHub and generate a personal access token for the Launcher:
--  https://github.com/settings/tokens
    * Set scopes
        * `repo`
        * `admin:repo_hook`
        * `delete_repo`
2. Copy the example CR based on the [given example](example/launcher_v1alpha1_launcher_cr.yaml):
3. Add your GitHub personal token (using a Secret `valueFrom: secretKeyRef: ...` or directly `valueFrom: text: ...`)
4. Add your CR to OpenShift
```bash
$ oc create -f <your_launcher_cr.yaml>
```
5. Browse the OpenShift Console to see a find the Route to the launcher.



## Example Launcher CR

Find an example of the Launcher CR in `example/launcher_v1alpha1_launcher_cr.yaml`

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

Start the operator (just restart this command to apply your changes):
```bash 
operator-sdk up local --namespace myproject   
```

Run this command when changing the API types (pkg/apis/launcher/v1alpha1/launcher_types.go)
```bash 
operator-sdk generate k8s
```

Then create your launcher CR and watch the logs in the console ouput.


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
