run-go:
	AMQP_HOST=localhost AMQP_PORT=5672 AMQP_LOGIN=test AMQP_PASSWORD=test AMQP_VHOST=/ DB_HOST=localhost DB_PORT=5432 DB_NAME=postgres DB_USER=drag DB_PASSWORD=terra GRPC_PORT=5673 go run cmd/web/main.go
migration-up:
	go run migration/migration.go