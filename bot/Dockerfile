ARG  BUILDER_IMAGE=golang:1.20
ARG  DISTROLESS_IMAGE=gcr.io/distroless/base
############################
# STEP 1 build executable binary
############################
FROM ${BUILDER_IMAGE} as builder

# Ensure ca-certficates are up to date
RUN update-ca-certificates

WORKDIR /build

# use modules
COPY . .

ENV GO111MODULE=on
RUN go mod download
RUN go mod verify


# Build the binary.
RUN CGO_ENABLED=1 go build cmd/flaggame/flaggame.go

############################
# STEP 2 build a small image
############################
# using base nonroot image
# user:group is nobody:nobody, uid:gid = 65534:65534
FROM ${DISTROLESS_IMAGE}

# Copy our static executable.
COPY --from=builder /build/flaggame /go/bin/flaggame

# Run the hello binary.
ENTRYPOINT ["/go/bin/flaggame"]
