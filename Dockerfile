FROM alpine:3.16.0 AS build
RUN apk update
RUN apk upgrade
RUN apk add --update go=1.18.2-r0 gcc=11.2.1_git20220219-r2 g++=11.2.1_git20220219-r2
WORKDIR /build
COPY app ./
RUN chmod +x "build.sh"
RUN "./build.sh"

FROM alpine:3.16.0

WORKDIR /root/
COPY --from=build /build/rmtcapp /app/rmtcapp

RUN addgroup -S rmtcapp \
  && adduser -S -D -s /sbin/nologin rmtcapp -G rmtcapp \
  && chown -R rmtcapp:rmtcapp /app/rmtcapp

WORKDIR /app/
USER rmtcapp

EXPOSE 8080
CMD ["./rmtcapp"]
