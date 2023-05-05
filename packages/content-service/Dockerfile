# FROM golang:1.19.2 as builder
# WORKDIR /app
# # RUN go mod init content-service
# COPY go.* ./
# RUN go mod download
# COPY *.go ./
# RUN go build -o /service
# # RUN CGO_ENABLED=0 GOOS=linux go build -o /content-service

# FROM gcr.io/distroless/base-debian11
# WORKDIR /
# COPY --from=builder /content-service /content-service
# EXPOSE 8080
# # ENV PORT 8080
# USER nonroot:nonroot
# CMD ["/content-service"]


FROM golang:1.19.2 as builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY cmd/content-service/*.go ./cmd/content-service/
COPY internal/database/*.go ./internal/database/
COPY internal/dto/*.go ./internal/dto/
COPY internal/spotify/*.go ./internal/spotify/
COPY internal/api/*.go ./internal/api/
COPY sample_response.json .
RUN go build -o /cmd/content-service/content-service ./cmd/content-service

FROM gcr.io/distroless/base-debian11
WORKDIR /
COPY --from=builder /cmd/content-service/content-service /cmd/content-service/content-service
EXPOSE 8080
USER nonroot:nonroot
CMD ["/cmd/content-service/content-service"]

