# FROM golang:1.19.2 as builder
# WORKDIR /app
# # RUN go mod init hello-app
# COPY go.* ./
# RUN go mod download
# COPY *.go ./
# RUN go build -o /hello-app
# # RUN CGO_ENABLED=0 GOOS=linux go build -o /hello-app

# FROM gcr.io/distroless/base-debian11
# WORKDIR /
# COPY --from=builder /hello-app /hello-app
# EXPOSE 8080
# # ENV PORT 8080
# USER nonroot:nonroot
# CMD ["/hello-app"]

# FROM golang:1.19.2 AS build

# WORKDIR /app

# COPY go.mod ./
# # COPY go.sum ./
# RUN go mod download

# COPY *.go ./

# RUN go build -o /docker-gs-ping

# ## Deploy
# FROM gcr.io/distroless/base-debian10

# WORKDIR /

# COPY --from=build /docker-gs-ping /docker-gs-ping

# EXPOSE 8080

# USER nonroot:nonroot

# ENTRYPOINT ["/docker-gs-ping"]

FROM golang:1.16-buster AS builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY *.go ./
RUN go build -o /hello_go_http
# Create a new release build stage
FROM gcr.io/distroless/base-debian10
# Set the working directory to the root directory path
WORKDIR /
# Copy over the binary built from the previous stage
COPY --from=builder /hello_go_http /hello_go_http
EXPOSE 8080
ENTRYPOINT ["/hello_go_http"]