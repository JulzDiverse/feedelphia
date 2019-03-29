# Feedelphia - Photofeed 

A Photo-Feed for CF Summit NA 2019 in Philadelphia.

## Demo Purpose

Using the full power of Cloud Foundry and Kubernetes to create your services... wouldn’t that be awesome? With Eirini, you can use the great CF push experience and interact with Kubernetes directly at the same time. This opens up some really interesting possibility for using microservices than span both CF and Kubernetes. Explore the full potential of both platforms combine by deploying this demo application on a Kubernetes cluster with CF Containerization & Eirini deployed.

## Deploy

*Prereqs:*

- Kubernetes Cluster with CF Containerization & Eirini deployed
- NodeJs
- npm
- yarn
- cf cli
- kubectl
- helm

### Setup a CF Org and Space

### Set your Kubeconfig

### Frontend

1. Set the `api_url` in `web/src/config.json` to `http://feed-api.<your-cf-domain>`
1. Build the static content using `yarn`:

   ```command
   $ pushd web 
   $ yarn run build
   ```

1. Run `cf push`
1. The frontend app should be avialable now at `http://feedelphia.<your-cf-domain>`

   At this stage you cannot interact with the UI at all. We will need to setup the backend first.

1. Return to the root directory of `feedelphia`: `$ popd`


### API

1. Navigate to the `api/` directory: `$ pushd api`
1. Run `cf push`
1. Wait for the app to get deployed
1. Navigate back to the root directory of `feedelphia`: `$ popd`

### Deploy the database

1. Navigate to the `db/` directory `$ pushd db`
1. Run `./helm-install-db.sh`
1. Export the `MYSQL_ROOT_PASSWORD` to an environment variable:

   ```command
   $ MYSQL_ROOT_PASSWORD=$(kubectl get secret --namespace feed-db feed-mysql -o jsonpath="{.data.mysql-root-password}" | base64 --decode; echo)
   ```

1. Expose the database deployment using NodePort:

   ```command
   $ kubectl expose deployment feed-db --type=NodePort
   ```

1. Get the node port: `$ kubectl get svc feed-db -o yaml | grep nodePort`
1. Navigate back to the root directory of `feedelphia`: `$ popd`

### Bind the api to the MySql Database

1. Navigate back to the `api`: `$ pushd api`
1. Edit the `db-cups.json` and provide the `MYSQL_ROOT_PASSWORD` value to the `password` property
1. Set the `db_address` to `<node-ip>:<node-port>`
1. Create a User-Provided-Service instance:

   ```command
   $ cf cups feed-db -p db-cups.json
   ```

1. Bind the `feed-api` app to the `feed-db` service instance:

   ```command
   $ cf bind-service feed-api feed-db
   ```

1. Restage the `api` app: `$ cf restage feed-api`


### Feed the App with some Photos

1. Open your browser window and open `http://feedelphia.<cf-app-domain>`
1. Click on `Upload Photo`
1. Choose or take a photo
1. Provide a Photo Title
1. Provide a Author Name
1. Click `upload`
1. Refresh the page

### Monitoring

To get metrics for your apps you can use `prometheus`, you can easily install it by executing:

```command
$ helm install --name prometheus stable/prometheus 
```

Get the Prometheus server URL by running these commands in the same shell:

```command
$ export POD_NAME=$(kubectl get pods --namespace prometheus -l "app=prometheus,component=server" -o jsonpath="{.items[0].metadata.name}")
$ kubectl --namespace prometheus port-forward $POD_NAME 9090
```

Follow the helm instruction after you install `prometheus` and all pods are running and ready. 
