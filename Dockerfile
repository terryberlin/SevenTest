FROM alpine:latest
RUN apk add --no-cache ca-certificates
ADD SevenShifts /
ENTRYPOINT ["/SevenShifts"]