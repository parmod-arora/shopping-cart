ARG BUILDER_IMAGE_NAME=golang
ARG BUILDER_IMAGE_TAG=1.13
ARG RELEASE_IMAGE_NAME=alpine
ARG RELEASE_IMAGE_TAG=3.10

FROM ${BUILDER_IMAGE_NAME}:${BUILDER_IMAGE_TAG} as builder

ARG APP_HOME=/shopping-cart
ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64

RUN apt-get update && apt-get install -y ca-certificates && \
    rm -rf /var/lib/apt/lists/*

WORKDIR $APP_HOME
COPY . $APP_HOME

RUN make build
FROM ${RELEASE_IMAGE_NAME}:${RELEASE_IMAGE_TAG}
RUN apk --no-cache add tzdata ca-certificates

COPY --from=builder /shopping-cart/serverd /
COPY --from=builder /shopping-cart/jwt-cert.yml /shopping-cart/jwt-cert.yml
RUN pwd
RUN ls
CMD ./serverd
