mock:
	mockgen -package mockdb -destination ./internal/database/mock/mock.go github.com/nerd500/axios-cp-wing/internal/database Querier
test_handlers:
	go test -v ./handlers
migrate_down:
	cd sql/schemas && goose postgres postgres://mydbuser:secretpassword@localhost:5432/axios_cp_wing down && cd ../..
migrate_up:
	cd sql/schemas && goose postgres postgres://mydbuser:secretpassword@localhost:5432/axios_cp_wing up && cd ../..
test:
	go test -v ./handlers
.PHONY: mock, test_handlers, migrate_down, migrate_up, test