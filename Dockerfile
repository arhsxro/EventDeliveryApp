ARG GO_VERSION=1.20
FROM golang:${GO_VERSION}-alpine AS build
 
WORKDIR /app
 
COPY go.mod ./
COPY go.sum ./
 
RUN go mod download
 
COPY ./ ./

RUN CGO_ENABLED=0 go build -o /eventdeliveryapp
 
FROM gcr.io/distroless/static
 
WORKDIR /
USER nonroot:nonroot
 
COPY --from=build --chown=nonroot:nonroot /eventdeliveryapp /eventdeliveryapp
 
ENTRYPOINT ["/eventdeliveryapp"]