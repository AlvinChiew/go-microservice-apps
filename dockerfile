### 7mb (almost the size of the statically built binary / exe)
FROM scratch  
WORKDIR /app
COPY src/main ./
ENTRYPOINT [ "./main" ]
EXPOSE 9090

### 359mb
# FROM golang:1.19-alpine  

### 27.26mb Multi-Stage Build
# FROM golang:1.19-alpine as builder
# WORKDIR /app
# COPY src/go.mod src/go.sum ./
# RUN go mod download
# COPY src/ .
# RUN go build -o main

# FROM gcr.io/distroless/base-debian11
# COPY --from=builder /app/main .
# EXPOSE 90 
# CMD ["/main"]