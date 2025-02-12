migrate: 
	set -o allexport; source .env/server; set +o allexport; go run cmd/migrate/main.go