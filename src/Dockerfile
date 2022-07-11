FROM golang:alpine AS builder

RUN mkdir /app \
&& go env -w GO111MODULE=on \
&& go env -w GOPROXY=https://goproxy.io,direct

COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp
###
FROM scratch as final
COPY --from=builder /app/myapp .
EXPOSE 8080
ENTRYPOINT [ "/myapp" ]