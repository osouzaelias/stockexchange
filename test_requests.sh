#!/bin/bash

# URL do serviço
SERVICE_URL="http://localhost:8080/order"

# Função para enviar uma requisição POST com um pedido de compra
send_buy_order() {
  curl -X POST -H "Content-Type: application/json" -d '{
    "order_id": '"$1"',
    "price": '"$2"',
    "quantity": '"$3"',
    "buy_sell": "Buy"
  }' $SERVICE_URL
  echo
}

# Função para enviar uma requisição POST com um pedido de venda
send_sell_order() {
  curl -X POST -H "Content-Type: application/json" -d '{
    "order_id": '"$1"',
    "price": '"$2"',
    "quantity": '"$3"',
    "buy_sell": "Sell"
  }' $SERVICE_URL
  echo
}

# Enviando pedidos de teste
echo "Enviando pedidos de teste..."

send_buy_order 1 100 10
send_sell_order 2 100 5
send_buy_order 3 110 15
send_sell_order 4 90 20
send_buy_order 5 105 25
send_sell_order 6 110 10

echo "Testes concluídos."