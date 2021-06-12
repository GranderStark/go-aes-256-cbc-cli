FROM alpine:3.12 AS build

RUN apk update && apk add --no-cache bash make go

COPY . /src
RUN cd /src && make all

FROM build AS ready
FROM alpine:3.12 AS base

COPY --from=ready /src/aes-256-cbc-cli/aes-256-cbc-cli /bin/aes-256-cbc-cli
RUN apk update && apk add --no-cache bash
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

RUN chmod +x /bin/aes-256-cbc-cli


ENTRYPOINT ["/entrypoint.sh"]
CMD ["--help" ]
