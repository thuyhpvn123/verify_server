#!/bin/bash

NGROK_URL="https://microclimatic-darcey-hysterogenic.ngrok-free.dev/webhook/whatsapp"

echo "=== WhatsApp Webhook Debug ==="
echo ""

echo "1. Testing health endpoint..."
curl -s "$NGROK_URL/health" | jq .
echo ""

echo "2. Testing WhatsApp webhook (GET - verification)..."
curl -s "$NGROK_URL/webhook/whatsapp?hub.mode=subscribe&hub.verify_token=YOUR_TOKEN&hub.challenge=test123"
echo ""

echo "3. Testing WhatsApp webhook (POST - message)..."
curl -X POST "$NGROK_URL/webhook/whatsapp" \
  -H "Content-Type: application/json" \
  -d '{
    "object": "whatsapp_business_account",
    "entry": [{
      "changes": [{
        "value": {
          "messages": [{
            "from": "1234567890",
            "text": {"body": "test"},
            "type": "text"
          }]
        }
      }]
    }]
  }'
echo ""

echo "Done! Check server logs."