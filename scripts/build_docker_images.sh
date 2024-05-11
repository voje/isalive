#!/bin/bash

cr="podman"
IMG="isalive:local"
CNT="isalive"

templ generate
$cr build . -t "$IMG"

if [ -n "$RUN" ]; then
    $cr run --rm -p 8080:8080 \
        --name "$CNT" \
        --env ISALIVE_SITES="https://www.google.com,http://localhost:8081,http://wikipedia.org" \
        $IMG
fi


if [ -n "$PUSH" ]; then
    $cr push "$IMG"
fi
