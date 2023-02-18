# syntax=docker/dockerfile:1

############################
# STEP 1 build executable binary
############################
# Alpine image is chosen for its small footprint
FROM --platform=linux/amd64 golang:1.19-alpine as builder

# Install SSL ca certificates.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache ca-certificates tzdata && update-ca-certificates

# Create appuser.
ENV USER=appuser
ENV UID=10001
ARG ARG_CONFIG_FILE="/image_verification_in_k8s/config.yaml"

# See https://stackoverflow.com/a/55757473/12429735RUN 
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

WORKDIR $GOPATH/src/image_verification_in_k8s/app
# Copy all the source code
COPY . .

# Build the binary
# Removing debug informations and compile only for linux target and disabling cross compilation.
RUN go mod tidy && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/image_verification_in_k8s

############################
# STEP 2 build a small image
############################
FROM --platform=linux/amd64 alpine:latest

ARG ARG_CONFIG_FILE
ENV APP_CONFIG_FILE=$ARG_CONFIG_FILE

# Import the user and group files from the builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /go/src/image_verification_in_k8s/app/config.yaml /image_verification_in_k8s/config.yaml

ENV HTTP_PORT="8181"
EXPOSE ${HTTP_PORT}
# Copy our static executable.
COPY --from=builder /go/bin/image_verification_in_k8s /go/bin/image_verification_in_k8s

# Use an unprivileged user.
USER appuser:appuser

# Run the binary
ENTRYPOINT [ "/go/bin/image_verification_in_k8s" ]