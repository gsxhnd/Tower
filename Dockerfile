FROM alpine:3.17

WORKDIR /app
COPY ./build/main .
EXPOSE 8080
ENTRYPOINT ["./go_pratice"]
