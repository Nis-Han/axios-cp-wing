mock:
	mockgen -package mockdb -destination ./internal/database/mock/mock.go github.com/nerd500/axios-cp-wing/internal/database Querier
test_handlers:
	go test -v ./handlers

.PHONY: mock, test_handlers