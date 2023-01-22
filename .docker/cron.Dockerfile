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
COPY ./etc /
# Copy configs
COPY --from=build /go/src/tutorial/app/config/yaml /tutorial/app/config/yaml
# Copy resource
COPY --from=build /go/src/tutorial/resource/ /tutorial/resource/
# Create folder for symlink
RUN mkdir -p /root/tutorial/app/config
# Create symlink
RUN ln -s /tutorial/app/config/yaml /root/tutorial/app/config/yaml
# Cronbab permission
RUN chmod 0644 cron/crontab
# Cron bash permission
RUN chmod 0777 cron/cronap.sh
# Update crontab
RUN crontab cron/crontab
# Entry
ENTRYPOINT ["sh", "cron/cronap.sh"]

