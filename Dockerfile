####################################################################################################
# base
####################################################################################################
FROM alpine:3.20 AS base
ARG TARGETARCH
RUN apk update && apk upgrade && \
    apk add ca-certificates && \
    apk --no-cache add tzdata

COPY dist/custom-transformer-example-${TARGETARCH} /bin/custom-transformer-example
RUN chmod +x /bin/custom-transformer-example

####################################################################################################
# custom-transformer-example
####################################################################################################
FROM scratch AS custom-transformer-example
COPY --from=base /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=base /bin/custom-transformer-example /bin/custom-transformer-example
ENTRYPOINT [ "/bin/custom-transformer-example" ]