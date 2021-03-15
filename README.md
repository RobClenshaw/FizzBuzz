# GoFizzBuzz
A distributed FizzBuzz using Go and Kubernetes. This was written for the purposes of learning Kubernetes.

There are three services.
* A Fizz service, which returns "Fizz" when the input is divisible by 3
* A Buzz service, which returns "Buzz" when the input is divisible by 5
* A FizzBuzz service, which concatenates the outputs from the Fizz and Buzz services. This is exposed to the end user.

## You will need

* [VS Code](https://code.visualstudio.com/) (or similar) for code editing (and install the Go extension so you get static analysis)
* [Go](https://golang.org/) (optional) if you want to run the servers locally
* [Docker](https://www.docker.com/) to build the containers
* [Docker Compose](https://docs.docker.com/compose/) (optional) if you want to run the application outside of Kubernetes (one container per service)
* [Minikube](https://minikube.sigs.k8s.io/docs/start/), a single-node local Kubernetes instance]
* A container repository that Kubernetes can access (I created one in Azure - if you do the same you'll need to replace the references to mine in the repo. You'll also need to add a secret to Kubernetes containing the repository credentials)

## Building the container images
The easiest way is to do `docker-compose build` from the repository root directory. After that, doing `docker-compose push` will push the images to the container repository. You might need to do `az acr login --name <repositoryname>` first to login to the Azure container registry.

## Deploying to Kubernetes
Ensure that Minikube is up and running, then from the root directory run `kubectl apply -f .\k8s\`. To access the fizzbuzz service, run `minikube service fizzbuzz` and Minikube will provide a URL to connect to the service.

## Deliberate faults
I have added a liveness probe to the fizz and buzz services that fails after returning once. This is intentional and is to demonstrate how Kubernetes can handle failures.
