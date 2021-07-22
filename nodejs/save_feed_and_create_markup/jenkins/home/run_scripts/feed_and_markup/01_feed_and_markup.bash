curl -X POST -H "Content-Type: application/json" -d '{"category": "devto-latest"}' http://app:3000/feed
sleep 5
curl -X POST -H "Content-Type: application/json" -d '{"category": "devto-latest"}' http://app:3000/feed/markup
