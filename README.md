# XXC demo

XXC: https://github.com/vmware-tanzu/cross-cluster-connectivity

## How to run the demo (using Kind as your local cluster)

Create a local kind cluster:
```
kind create cluster --name xcc-demo --config kind.config.yaml
```
__NOTE__: You can find [this config file in the repo](kind.config.yaml)

Deploy contour on your local kind cluster (and a sample kuard application to test):
```
kubectl apply -f https://projectcontour.io/quickstart/contour.yaml --context kind-xcc-demo
kubectl patch daemonsets -n projectcontour envoy -p '{"spec":{"template":{"spec":{"nodeSelector":{"ingress-ready":"true"},"tolerations":[{"key":"node-role.kubernetes.io/master","operator":"Equal","effect":"NoSchedule"}]}}}}' --context kind-xcc-demo
ytt -f values.yaml -f cluster-local/sample-app.yaml | kubectl apply --context kind-xcc-demo -f -
```

__NOTE__: You should be able to access your local application at [http://kuard.127.0.0.1.nip.io](http://kuard.127.0.0.1.nip.io)

Deploy xcc-demo to cluster-b
```
ytt -f values.yaml -f cluster-remote/resources.yaml | kubectl apply --kubeconfig ~/.kube/config.d/kubeconfig-dev-eduk8s-io.yml -f -
```
__NOTE__: Use your own remote kubeconfig file

Deploy xxc (dns and endpointSlice to cluster-a)
```
ytt -f values.yaml -f cluster-local/xcc-dns.yaml | kubectl apply --context kind-xcc-demo -f -
ytt -f values.yaml -f cluster-local/EndpointSlice.yaml | kubectl apply --context kind-xcc-demo -f -
```

Deploy xcc-demo to cluster-a configured to talk to cluster-b mysql
```
ytt -f values.yaml -f cluster-local/resources.yaml | kubectl apply --context kind-xcc-demo -f -
```

Test xcc-demo:
```
curl xcc-demo.127.0.0.1.nip.io
```