#!/bin/bash

# Add Go bin to PATH
export PATH=$PATH:$(go env GOPATH)/bin

# Navigate to project root
cd "$(dirname "$0")/.."

# Generate swagger documentation from code annotations
echo "Generating Swagger documentation..."
swag init \
  --parseDependency \
  --parseInternal \
  --dir ./cmd/server/,./internal/beginning_of_day_BOD/api/handlers,./internal/beginning_of_day_BOD/api/dtos,./internal/matching_booking/api/handlers,./internal/matching_booking/api/dtos,./internal/contribution_data/api/handlers,./internal/contribution_data/api/dtos,./internal/fund_confirmation/api/handlers,./internal/fund_confirmation/api/dtos,./internal/pay_in_process/api/handlers,./internal/pay_in_process/api/dtos,./internal/payout/api/handlers,./internal/payout/api/dtos \
  --output ./docs

if [ $? -eq 0 ]; then
    echo "✅ Swagger documentation generated successfully!"
    echo "📄 Files updated:"
    echo "   - docs/docs.go"
    echo "   - docs/swagger.json"  
    echo "   - docs/swagger.yaml"
    echo ""
    echo "🌐 Access Swagger UI at: http://localhost:8080/swagger/index.html"
else
    echo "❌ Failed to generate swagger documentation"
    exit 1
fi