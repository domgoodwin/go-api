# go-api

Golang API  

## Building the image

```docker build -t go-api```


## Running the image

```docker run -it --rm -v /Users/dom/.aws/credentials:/root/.aws/credentials -p 8080:8080 --name go-api-0 go-api```
