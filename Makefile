migrateup:
	migrate -path db/migration -database "postgres://postgres:123456@localhost:5432/base-gin-go?sslmode=disable" -verbose up $(step)
migratedown:
	migrate -path db/migration -database "postgres://postgres:123456@localhost:5432/base-gin-go?sslmode=disable" -verbose down $(step)