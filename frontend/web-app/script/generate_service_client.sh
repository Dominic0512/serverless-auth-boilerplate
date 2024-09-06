#!/bin/bash

# Define the output directory
OUTPUT_DIR="./generated/api"

# Define the API documentation URLs and corresponding file names
declare -a SERVICES=(
  "auth_service"
  "user_service"
)
declare -a URLS=(
  "http://localhost:10001/swagger/doc.json"
  "http://localhost:10002/swagger/doc.json"
)

# Create the output directory (if it does not exist)
mkdir -p "$OUTPUT_DIR"

# Loop through each service
for i in "${!SERVICES[@]}"; do
  SERVICE="${SERVICES[$i]}"
  URL="${URLS[$i]}"
  FILE_NAME="${SERVICE}.ts" # Use the service name as the file name
  echo $SERVICE
  # Generate TypeScript file
  npx swagger-typescript-api -p "$URL" -o "$OUTPUT_DIR/$SERVICE" --axios --modular -n "$FILE_NAME"
done

echo "API files generation completed!"
