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

Clone this repository:
```bash
$ git clone https://github.com/fabric8-launcher/launcher-operator
$ cd launcher-operator
```

Choose the project that will run the operator and then install all the operator resources:

```bash
$ oc create -R -f ./deploy 
```

## Install the Launcher (via the installed operator)

1. Log into GitHub and create a OAuth Application for the Launcher named 'launcher':
--  https://github.com/settings/applications/new
    * Using `http://temporary` as 'Homepage URL' an 'Authorization callback URL' (You will know the frontend url later)

2. Create a secret for your GitHub Oauth settings
```bash
$ oc create secret generic launcher-oauth-github --from-literal=clientId=<YOUR_GITHUB_OAUTH_APP_CLIENT_ID> --from-literal=secret=<YOUR_GITHUB_OAUTH_APP_CLIENT_SECRET>
```

3. Customize the Launcher Resource with you OpenShift Console URL and create it
```bash
$ oc create -f example/launcher_cr.yaml
```

4. Get the Launcher URL:
```bash
$ oc get route launcher --template={{.spec.host}}
```

5. Create a OAuth client in OpenShift:
```bash
$ oc create -f <(echo '
  kind: OAuthClient
  apiVersion: oauth.openshift.io/v1
  metadata:
    name: launcher
  secret: <CHANGE_IT>
  redirectURIs:
    - "http://<your frontend hostname>/"
  grantMethod: prompt
  ')
```

6. Edit you GitHub OAuth application with the frontend URL as 'Homepage URL' an 'Authorization callback URL'

## Example Launcher CR

Find an example of the Launcher CR in `example/launcher_cr.yaml`

---

## Develop

Install the `operator-sdk` using [the instructions](https://github.com/operator-framework/operator-sdk).

Register the crd:
```bash
$ oc create -f deploy/crds/launcher_v1alpha2_launcher_crd.yaml  
```

Install dependencies:
```bash 
$ dep ensure -v
```

Start the operator (just restart this command to apply your changes):
```bash 
operator-sdk up local --namespace myproject   
```

Run this command when changing the API types (pkg/apis/launcher/v1alpha2/launcher_types.go)
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
