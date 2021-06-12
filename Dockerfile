FROM alpine:3.12

RUN apk update && apk add --no-cache bash

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
COPY aes-256-cbc-cli/aes-256-cbc-cli /bin/aes-256-cbc-cli
RUN chmod +x /bin/aes-256-cbc-cli


ENTRYPOINT ["/entrypoint.sh"]
CMD [ "aes-256-cbc-cli", "--help" ]
