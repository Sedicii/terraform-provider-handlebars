#!/usr/bin/env bash

OS=$(uname | tr '[:upper:]' '[:lower:]')
VERSION="v0.1.0"
PROVIDER_DOWNLOAD_URL="https://github.com/Sedicii/terraform-provider-handlebars/releases/download/${VERSION}/terraform-provider-handlebars_${OS}-amd64_${VERSION}"
TF_PLUGINS_PATH="${HOME}/.terraform.d/plugins/${OS}_amd64/"

mkdir -p "${TF_PLUGINS_PATH}"
curl ${PROVIDER_DOWNLOAD_URL} --output "${TF_PLUGINS_PATH}/terraform-provider-handlebars_${VERSION}"
