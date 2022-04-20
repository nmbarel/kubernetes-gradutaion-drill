# Backend

## Create image 
Run the following:
```
docker build -t <YOUR_NAME>/k8s-final-drill:backend .
```

Or use the existing image in 
```
ghcr.io/expruc/kubernetes-gradutaion-drill/k8s-final-drill:backend
```

## Running the container

The Backend container requires 1 environment variable named
```
PORT
```