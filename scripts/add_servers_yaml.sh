#!/bin/bash

# Usage: ./replace_servers.sh openapi.yaml

if [ -z "$1" ]; then
  echo "❌ Usage: $0 <yaml_file>"
  exit 1
fi

FILE="$1"

yq -y -i '.servers = [
  {
    url: "https://dev-cra2.nps-proteantech.in/settlement",
    description: "Development Server"
  },
  {
    url: "https://uat-cra2.nps-proteantech.in/settlement",
    description: "UAT Server"
  },
  {
    url: "/",
    description: "Default Server"
  }
]' "$FILE"

echo "✅ Replaced servers in $FILE"
