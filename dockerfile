ARG BUILDER_IMAGE_NAME=golang
ARG BUILDER_IMAGE_TAG=1.13
ARG UI_BUILDER_IMAGE_NAME=node
ARG UI_BUILDER_IMAGE_TAG=14.14.0
ARG RELEASE_IMAGE_NAME=alpine
ARG RELEASE_IMAGE_TAG=3.10

FROM ${UI_BUILDER_IMAGE_NAME}:${UI_BUILDER_IMAGE_TAG} as uibuilder
ARG UI_APP_HOME=/shopping-cart-ui
WORKDIR $UI_APP_HOME
COPY ./ui-app $UI_APP_HOME 
RUN yarn install && yarn build
RUN ls
RUN pwd

FROM ${BUILDER_IMAGE_NAME}:${BUILDER_IMAGE_TAG} as builder

ARG APP_HOME=/shopping-cart
ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64

RUN apt-get update && apt-get install -y ca-certificates && \
    rm -rf /var/lib/apt/lists/*

WORKDIR $APP_HOME
COPY . $APP_HOME

RUN make -f Makefile.docker build
FROM ${RELEASE_IMAGE_NAME}:${RELEASE_IMAGE_TAG}
RUN apk --no-cache add tzdata ca-certificates

COPY --from=builder /shopping-cart/serverd /
COPY --from=builder /shopping-cart/jwt-cert.yml /shopping-cart/jwt-cert.yml
COPY --from=uibuilder /shopping-cart-ui/build /shopping-cart/ui-app/build
RUN ls
RUN pwd

CMD ./serverd
