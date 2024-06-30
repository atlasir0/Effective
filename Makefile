run: 
	go run cmd/server/main.go --config=./config/local.yaml 

migrate_up: 
	migrate -source file://migrations -database 'postgres://postgres:123@localhost:5432/todo_list?sslmode=disable' up

migrate_down: 
	migrate -source file://migrations -database 'postgres://postgres:123@localhost:5432/todo_list?sslmode=disable' down 

sqlc: 
	sqlc generate