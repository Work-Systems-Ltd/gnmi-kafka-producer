.PHONY: up down restart logs ps tail-topic build clean

COMPOSE ?= docker compose
TOPIC ?= gnmi.telemetry

up:
	$(COMPOSE) up -d --build

down:
	$(COMPOSE) down -v

restart: down up

build:
	$(COMPOSE) build

ps:
	$(COMPOSE) ps

logs:
	$(COMPOSE) logs -f --tail=100

tail-topic:
	$(COMPOSE) exec kafka /opt/kafka/bin/kafka-console-consumer.sh \
		--bootstrap-server localhost:9092 \
		--topic $(TOPIC) \
		--from-beginning \
		--max-messages 50 \
		--property print.key=true \
		--property key.separator=' | '

clean: down
	docker volume prune -f
