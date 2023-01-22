# Get image from argument
ARG image=image
# Set allias for image
FROM ${image:-image} AS build
# Getting aline image
FROM alpine:latest
# Copy certs
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# Copy bin file
COPY --from=build /go/src/tutorial/tutorial /tutorial/
# Copy etc files
COPY --from=build /go/src/tutorial/etc /
# Copy configs
COPY --from=build /go/src/tutorial/app/config/yaml /tutorial/app/config/yaml
# Copy resource
COPY --from=build /go/src/tutorial/resource/ /tutorial/resource/
# Entry
ENTRYPOINT ["tutorial/tutorial", "-app", "consumer"]

