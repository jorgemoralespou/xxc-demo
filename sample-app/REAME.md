# Sample application
This sample application is a go server that replies hello locally or if an environment variable with a remote URL has been specified, replies what the remote service answers.


## How to test it

Build the app:
```
docker build -t "xxc-demo/sample-go-app" .
```

Test the app:
```
docker-compose up -d
sleep 5
# Query local
curl -v localhost:8081
# Query remote
curl -v localhost:8080
docker-compose down
```

# Image on docker hub
This image is published on: `docker.io/jorgemoralespou/xcc-demo-sample-app:latest`