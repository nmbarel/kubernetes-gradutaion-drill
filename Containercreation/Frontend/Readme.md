# Frontend

## Create image 
Run the following:
```
docker build -t <YOUR_NAME>/k8s-final-drill:frontend .
```

Or use the existing image in 
```
ghcr.io/expruc/kubernetes-gradutaion-drill/k8s-final-drill:frontend
```

## Running the container

The Frontend container requires 2 environment variable named
```
PORT
BACKENDURL
```