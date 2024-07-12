FROM golang:1.22.5

# install a nice go linter
ENV LINT_VERSION v1.59.1
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin ${LINT_VERSION}

WORKDIR /app

# Let files be mounted by Docker Compose