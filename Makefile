DATABASE_URL := "postgres://root:root@127.0.0.1:5432/banksampah?sslmode=disable"

up:
	migrate -path ./database/migration/files -database $(DATABASE_URL) up

down:
	migrate -path ./database/migration/files -database $(DATABASE_URL) down

version:
	migrate -path ./database/migration/files -database $(DATABASE_URL) version

fix-and-force:
	migrate -path ./database/migration/files -database $(DATABASE_URL) force $(version)
