# XXC demo

XXC: https://github.com/vmware-tanzu/cross-cluster-connectivity

## How to run the demo (using Kind as your local cluster)

Create a local kind cluster:
```
kind create cluster --name xxc-demo --config kind.config.yaml
```
__NOTE__: You can find [this config file in the repo](kind.config.yaml)

Deploy contour on your local kind cluster (and a sample kuard application to test):
```
kubectl apply -f https://projectcontour.io/quickstart/contour.yaml --context kind-xxc-demo
kubectl patch daemonsets -n projectcontour envoy -p '{"spec":{"template":{"spec":{"nodeSelector":{"ingress-ready":"true"},"tolerations":[{"key":"node-role.kubernetes.io/master","operator":"Equal","effect":"NoSchedule"}]}}}}' --context kind-xxc-demo
kubectl apply -f cluster-local/sample-app.yaml --context kind-xxc-demo
```

__NOTE__: You should be able to access your local application at [http://kuard.127.0.0.1.nip.io](http://kuard.127.0.0.1.nip.io)

Deploy xxc-demo to cluster-b
```
kubectl apply -f cluster-remote/resources.yaml --kubeconfig ~/.kube/config.d/kubeconfig-dev-eduk8s-io.yml
```

Deploy xxc (dns and endpointSlice to cluster-a)
```
kubectl apply -f cluster-local/xxc-dns.yaml --context kind-xxc-demo
kubectl apply -f cluster-local/EndpointSlice.yaml --context kind-xxc-demo
```

Deploy xxc-demo to cluster-a configured to talk to cluster-b mysql
```
kubectl apply -f cluster-local/resources.yaml --context kind-xxc-demo
```

Test xxc-demo:
```
curl xxc-demo.127.0.0.1.nip.io
```