#syntax=docker/dockerfile:1.2
FROM library/golang as builder
# Godep for vendoring
RUN go install github.com/tools/godep@latest
# Recompile the standard library without CGO
RUN CGO_ENABLED=0 go install -a std

ENV APP_DIR /opt/simply
RUN mkdir -p $APP_DIR

# Set the entrypoint
ADD . $APP_DIR

# Compile the binary and statically link
RUN cd $APP_DIR && CGO_ENABLED=0 go build -ldflags '-d -w -s' -o ./build/server ./cmd/server

FROM alpine:latest

ENV APP_DIR /opt/simply

RUN mkdir -p $APP_DIR

COPY --from=builder ${APP_DIR}/build/server ${APP_DIR}/server
COPY --from=builder ${APP_DIR}/config ${APP_DIR}/config

WORKDIR ${APP_DIR}

EXPOSE 8888

ENTRYPOINT (./server)
# CMD [ "./server" ]
