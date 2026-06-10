.PHONY: up down ps logs-redis clean prune

FILE=docker/redis.docker-compose.yml
DATA_DIR=docker/redis-data

# ======================== Setup ========================

init:
	@echo "📁 Preparing directories..."
	mkdir -p $(DATA_DIR)
	chmod -R 777 $(DATA_DIR) || true
	@echo "✅ Ready"

# ======================== Redis ========================

up: init
	@echo "🐳 Starting Redis..."
	docker compose -f $(FILE) up --force-recreate -d
	@echo "✅ Redis started"

down:
	@echo "🛑 Stopping Redis..."
	docker compose -f $(FILE) down
	@echo "✅ Stopped"

ps:
	@echo "📋 Container status:"
	docker ps -a --filter "name=redis-auction"

logs-redis:
	@echo "📝 Redis logs:"
	docker compose -f $(FILE) logs -f redis

# ======================== Cleanup ========================

clean:
	@echo "🧹 Cleaning project data..."
	rm -rf $(DATA_DIR)
	@echo "✅ Cleaned"

prune:
	@echo "💣 Pruning Docker system (DANGEROUS)..."
	docker system prune -a -f
	@echo "✅ Pruned"