FROM alpine:3.7

RUN apk add curl

RUN curl -fsSL https://deno.land/x/install/install.sh | sh -s v0.2.10

CMD ["deno", "help"]

# TODO(alecmerdler): Use `scratch` for final container