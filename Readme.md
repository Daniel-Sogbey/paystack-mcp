# Paystack MCP Server
A Model Context Protocol (MCP) tool providing Paystack transaction initialization and verification.

## Features
- `.well-known/mcp-manifest` (discovery)
- `/mcp/paystack/initialize` (POST)
- `/mcp/paystack/verify/:reference` (GET)

## Quickstart (dev)
```bash
export PAYSTACK_SECRET_KEY="sk_test_xxx"
docker run -p 8080:8080 -e PAYSTACK_API_KEY=$PAYSTACK_API_KEY -e PAYSTACK_BASE_URL="https://api.paystack.co" -e PORT="8080"  sogbey/paystack-mcp:latest
curl -X GET "http://localhost:8080/.well-known/mcp-manifest"
