# Start from golang v1.11 base image
FROM golang:1.11

# install dep
RUN go get github.com/golang/dep/cmd/dep

# create a working directory
WORKDIR /go/src/candidate_service

# add Gopkg.toml and Gopkg.lock
ADD Gopkg.toml Gopkg.toml
ADD Gopkg.lock Gopkg.lock

# install packages
# --vendor-only is used to restrict dep from scanning source code
# and finding dependencies
RUN dep ensure --vendor-only

# add source code
ADD . .

# expose port
EXPOSE 3000

# run main.go
CMD ["go", "run", "main/main.go"]