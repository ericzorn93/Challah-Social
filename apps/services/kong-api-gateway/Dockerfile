FROM kong/kong-gateway:latest

# Set the user to root
USER root

# Install curl, jq and gettext on UBUNTU
RUN apt update && apt install -y curl jq gettext

# Define the ARG variables
ARG INBOUND_WEBHOOKS_API_URL

# Set the environment variables
ENV INBOUND_WEBHOOKS_API_URL=${INBOUND_WEBHOOKS_API_URL}

# Copy your custom script into the container
COPY ./apps/services/kong-api-gateway/startup-script.sh /usr/local/bin/startup-script.sh

# Copy your declarative configuration file into the container
COPY ./apps/services/kong-api-gateway/kong-template.yml /kong/declarative/kong-template.yml

# Ensure it has the right permissions
RUN chmod +x /usr/local/bin/startup-script.sh

# RUN the script
RUN /bin/sh /usr/local/bin/startup-script.sh

