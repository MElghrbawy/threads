# syntax=docker/dockerfile:1

ARG GO_VERSION=1.23.2
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS build
WORKDIR /src

# Download dependencies
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

ARG TARGETARCH

# Build the application
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 GOARCH=$TARGETARCH go build -o /bin/server ./cmd/server

# Build the migration tool
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 GOARCH=$TARGETARCH go build -o /bin/migrate ./cmd/migrate

FROM alpine:latest AS final

RUN --mount=type=cache,target=/var/cache/apk \
    apk --update add \
        ca-certificates \
        tzdata \
        && \
        update-ca-certificates

ARG UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    appuser

# Copy the executables from the "build" stage
COPY --from=build /bin/server /bin/
COPY --from=build /bin/migrate /bin/

# Copy migrations
COPY internal/db/migrations /migrations

# Copy the startup script and set permissions
COPY start.sh /start.sh
RUN chmod +x /start.sh

# Switch to non-root user
USER appuser

EXPOSE 8000

ENTRYPOINT ["/start.sh"]