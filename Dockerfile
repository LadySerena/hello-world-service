FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1.18 as builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app/
ADD . .
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-w -s" -o hello-world-service main.go

FROM --platform=${TARGETPLATFORM:-linux/amd64} scratch
WORKDIR /app/
COPY --from=builder /app/hello-world-service /app/hello-world-service
ENTRYPOINT ["/app/hello-world-service"]