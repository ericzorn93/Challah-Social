#!/bin/sh

# Example script to replace environment variables in the template file
echo "Substituting environment variables..."

# Use envsubst or sed to substitute variables
envsubst < /kong/declarative/kong-template.yml > /kong/declarative/kong.yml

