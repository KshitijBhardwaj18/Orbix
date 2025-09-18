#!/bin/bash

echo "🚀 Orbix Clean Event Monitor"
echo "Choose monitoring mode:"
echo "1) Database Events (db@*)"
echo "2) WebSocket Events (order@*, trade@*, depth@*)"
echo "3) Specific Market WebSocket (e.g., BTC_USD)"
echo "4) All Events"
read -p "Enter choice (1-4): " choice

case $choice in
    1) 
        echo "📊 Monitoring Database Events: db@orderplaced, db@orderupdated, db@trade"
        docker exec -i orbix-broker-1 redis-cli PSUBSCRIBE "db@*"
        ;;
    2) 
        echo "📡 Monitoring WebSocket Events: order@*, trade@*, depth@*"
        docker exec -i orbix-broker-1 redis-cli PSUBSCRIBE "order@*" "trade@*" "depth@*"
        ;;
    3) 
        read -p "Enter market (e.g., BTC_USD): " market
        echo "🎯 Monitoring $market WebSocket Events"
        docker exec -i orbix-broker-1 redis-cli PSUBSCRIBE "order@$market" "trade@$market" "depth@$market"
        ;;
    4) 
        echo "🌐 Monitoring All Events"
        docker exec -i orbix-broker-1 redis-cli PSUBSCRIBE "*@*"
        ;;
    *) 
        echo "🌐 Default: Monitoring All Events"
        docker exec -i orbix-broker-1 redis-cli PSUBSCRIBE "*@*"
        ;;
esac
