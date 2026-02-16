ARG GO_BASE_IMAGE=golang:1.26.0-alpine
ARG DISTROLESS_BASE_IMAGE=gcr.io/distroless/static:nonroot
ARG BUILDPLATFORM=linux/amd64
FROM --platform=${BUILDPLATFORM} $GO_BASE_IMAGE AS builder

WORKDIR /workspace

# 1. install dependencies
COPY go.mod go.sum ./
RUN go mod download

# 2. copy source codes
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /workspace/bin/kube-scheduler ./cmd/scheduler/main.go

FROM --platform=${BUILDPLATFORM} $DISTROLESS_BASE_IMAGE

WORKDIR /bin
COPY --from=builder /workspace/bin/kube-scheduler .
USER 65532:65532

ENTRYPOINT ["/bin/kube-scheduler"]