#!/bin/bash
if [[ $# -lt 1 ]] || [[ "$1" == "--"* ]] || [[ ! -f "$1" ]]; then
    exec /bin/aes-256-cbc-cli "$@"
fi
exec "$@"
