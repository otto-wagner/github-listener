FROM golang:1.21.9-alpine AS build

WORKDIR /app

COPY . .

RUN go build -a -o github-observer main.go

FROM alpine:3.19.1

ARG USERNAME=observer
ARG USER_UID=1000
ARG GITHUB_TOKEN
ENV GITHUB_TOKEN=$GITHUB_TOKEN

RUN touch github_observer.log && chown $USER_UID:$USER_UID github_observer.log

RUN adduser -u $USER_UID -D $USERNAME
USER $USERNAME

COPY --from=build /app/github-observer /github-observer
COPY --from=build /app/conf /conf

ENTRYPOINT ["/github-observer", "server"]