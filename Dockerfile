################
# BUILD BINARY #
################
# golang:1.22.3-alpine
FROM golang@sha256:b8ded51bad03238f67994d0a6b88680609b392db04312f60c23358cc878d4902 AS builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

WORKDIR $GOPATH/src/dealls-test
COPY . .

RUN echo "$PWD" && ls -lah

# Fetch dependencies.
# RUN go get -d -v
RUN go mod download
RUN go mod verify

# CMD go build -v
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/dealls-test .

#####################
# MAKE SMALL BINARY #
#####################
FROM alpine:3.14

RUN apk update && apk add --no-cache tzdata
ENV TZ=UTC

# Import from builder.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

# Copy the executable.
COPY --from=builder /go/bin/dealls-test /go/bin/dealls-test
CMD ["/go/bin/dealls-test"]