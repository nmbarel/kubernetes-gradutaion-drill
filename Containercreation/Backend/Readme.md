# Backend

## Build

run the following:
```
CGO_ENABLED=0 && go build .
```

## create container 
run the following:
```
docker build -t theog75/drillbackend:v1
```


## Running the container

The Backend container requires 1 environment variable named
```
PORT
```