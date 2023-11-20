mockDB:
	mockgen -package mockdb -destination ./internal/database/mock/mock.go github.com/nerd500/axios-cp-wing/internal/database Querier
mock_email_client:
	mockgen -package mock_email_client -destination ./client/mock_client/mock_email_client/mock_email_client.go github.com/nerd500/axios-cp-wing/client/email_client EmailClientInterface
test_handlers:
	go test -v ./handlers
migrate_down:
	cd sql/schemas && goose postgres postgres://mydbuser:secretpassword@localhost:5432/axios_cp_wing down && cd ../..
migrate_up:
	cd sql/schemas && goose postgres postgres://mydbuser:secretpassword@localhost:5432/axios_cp_wing up && cd ../..
test:
	go test -v ./handlers/handlers_test
.PHONY: mockDB, test_handlers, migrate_down, migrate_up, test, mock_email_client
