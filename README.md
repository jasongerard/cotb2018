# Code on the Beach 2018 Sample App

There's not much here, just a simple Go app and some Kubernetes config files

The Go app runs a simple HTTP server with the following endpoints:

`GET /`
Returns Hello World

`GET /healthz`
Endpoint for livenessProbe

`GET /ready`
Endpoint for readinessProbe

`GET /togglehealth`
Toggle flag to report healthy or not for livenessProbe

The `k8s` folder has 3 files.

`pod.yaml` shows an example of of a pod configuration.

`deployment.yaml` shows an example of a deployment. It will conflict with the pod defined in `pod.yaml`.

`service.yaml` exposes the deployment from `deployment.yaml` using a LoadBalancer service. Use `minikube service cotbhttp` 
if using minikube to run the service from a local web browser.

`curl.Dockerfile` builds a the basic image used to run curl from inside the kubenetes cluster.

Build with `docker build -t curl:1.0.0 -f curl.Dockerfile .` run from the project root. If using minikube, run `eval $(minikube docker-env)` first.

Run in Kubernetes with `kubectl run -i --tty curl --image=curl:1.0.0 -- sh`

`Dockerfile` is the image for the sample app. Run `build.sh` to build a Linux binary and Docker image.

Use `kubectl apply -f k8s/deployment.yaml` to deploy to Kubernetes. Alternatively, you could run `kubectl run cotb2018 --image=cotb2018:1.0.0 --replicas=3`.

