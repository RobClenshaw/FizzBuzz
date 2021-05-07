# FizzBuzz
A distributed FizzBuzz using Go and Kubernetes. This was written for the purposes of learning Kubernetes.

There are three services.
* A Fizz service, which returns "Fizz" when the input is divisible by 3
* A Buzz service, which returns "Buzz" when the input is divisible by 5
* A FizzBuzz service, which concatenates the outputs from the Fizz and Buzz services. This is exposed to the end user.

## You will need

* [Minikube](https://minikube.sigs.k8s.io/docs/start/), a single-node local Kubernetes instance
* [Docker](https://www.docker.com/)

### Optional
* [VS Code](https://code.visualstudio.com/) (or similar)
* [Go](https://golang.org/)
* [Docker Compose](https://docs.docker.com/compose/)

## Quick start
Ensure that Minikube is up and running. If not, run `minikube start`.

To run FizzBuzz you don't even need to clone this repo. To deploy this implementation to Kubernetes, just run 
```
kubectl apply -f https://raw.githubusercontent.com/RobClenshaw/FizzBuzz/main/k8s/fizzbuzz.yaml
```

If you've cloned the repo and want to use your local definition, replace the path in the command above with the path to your local yaml file.

Running `minikube service fizzbuzz` will open the service URL in a browser window. Append `/15` to the URL to get the appropriate response for the number 15.

## Deliberate faults
I have added a liveness probe to the fizz and buzz services that fails after returning once. This is intentional and is to demonstrate how Kubernetes can handle failures.
