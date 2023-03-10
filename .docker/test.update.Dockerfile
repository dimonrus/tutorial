# Get image from argument
ARG image=image
# Set allias for image
FROM ${image:-image} AS build
# Remove codebase
RUN rm -rf /go/src/tutorial/
# Copy current codebase
COPY ./ /go/src/tutorial/

