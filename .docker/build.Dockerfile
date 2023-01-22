# Get image from argument
ARG image=image
# Set allias for image
FROM ${image:-image} AS build
# Install Make and git
RUN apk add --update make git
# Remove codebase
RUN rm -rf /go/src/tutorial/
# Copy current codebase
COPY ./ /go/src/tutorial/
# Set work directory
WORKDIR /go/src/tutorial/
# Build app
RUN make project-build