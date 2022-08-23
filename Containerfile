# Copyright (c) 2022 Red Hat, Inc.

# Stage 1: Use image builder to retrieve Go binaries
FROM golang:1.17 AS builder

# Stage 2: Copy Go binaries and run tests on ubi-minimal
FROM registry.access.redhat.com/ubi8/ubi-minimal:latest

RUN microdnf update -y \
        && microdnf install tar \
        && microdnf install gzip \
        && microdnf install git \
        && microdnf install which \
        && microdnf install make \
        && microdnf install findutils \
        && microdnf clean all

# prepare jq
ADD https://github.com/stedolan/jq/releases/download/jq-1.6/jq-linux64 /usr/local/bin/jq
RUN chmod +x /usr/local/bin/jq

# copy go files from builder image
COPY --from=builder /usr/local/go /usr/local/go
COPY --from=builder /go /go

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
WORKDIR $GOPATH
RUN mkdir -p $GOPATH/src/gitlab.cee.redhat.com/mpqe/mps/solutions/openshift_plus/openshift_plus/tests

WORKDIR $GOPATH/src/github.com/liswang89/policy-openshift-plus

COPY download-clis.sh .
COPY go.mod .
#COPY go.sum .
COPY Makefile .
COPY tests ./tests
COPY run_tests.sh .
#COPY /kube/config ./tests/kubeconfig_hub
#COPY /kube/config ./tests/kubeconfig_managed

RUN ./download-clis.sh

RUN make e2e-dependencies

RUN go mod tidy

# Give write permissions for the directory
RUN chmod -R go+wx $GOPATH/src/github.com/liswang89/policy-openshift-plus

WORKDIR $GOPATH/src/github.com/liswang89/policy-openshift-plus

CMD ["./run_tests.sh"]