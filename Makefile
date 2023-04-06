all:
	go build cmd/main.go
clean-mock:
	rm -rf internal/mocks

gen-repo-mocks:
	mockery --dir internal/services/ --output internal/mocks/repositories --name ChannelRepository
	mockery --dir internal/services/ --output internal/mocks/repositories --name MessageRepository
	mockery --dir internal/services/ --output internal/mocks/repositories --name RoleRepository
	mockery --dir internal/services/ --output internal/mocks/repositories --name ServerRepository
	mockery --dir internal/services/ --output internal/mocks/repositories --name UserRepository
	mockery --dir internal/services/ --output internal/mocks/repositories --name SessionRepository

gen-service-mocks:
	mockery --dir internal/services/ --name RoleService --output internal/mocks/services
	mockery --dir internal/services/ --name ServerProfileChecker --output internal/mocks/services
	mockery --dir internal/services/ --name RoleSystem --output internal/mocks/services


gen-mocks:
	make gen-repo-mocks
	make gen-service-mocks


test-message-service:
	go test ./internal/services/message_service -coverprofile cover.out && go tool cover -func cover.out && rm cover.out

test-auth-service:
	go test ./internal/services/auth_service -coverprofile cover.out && go tool cover -func cover.out && rm cover.out

test-channel-service:
	go test ./internal/services/channel_service -coverprofile cover.out && go tool cover -func cover.out && rm cover.out

test-role-service:
	go test ./internal/services/role_service -coverprofile cover.out && go tool cover -func cover.out && rm cover.out

test-role-system:
	go test ./internal/services/role_system -coverprofile cover.out && go tool cover -func cover.out && rm cover.out

test-server-profile-checker:
	go test ./internal/services/server_profile_checker -coverprofile cover.out && go tool cover -func cover.out && rm cover.out

test-server-service:
	go test ./internal/services/server_service -coverprofile cover.out && go tool cover -func cover.out && rm cover.out

test-all:
	make test-message-service
	make test-auth-service
	make test-channel-service
	make test-role-service
	make test-role-system
	make test-server-profile-checker
	make test-server-service



drop-tables:
	psql -U postgres -d openchat -c "$$(cat migration/drop_tables.sql)"

load-test-data:
	psql -U postgres -d openchat -c "$$(cat migration/test/load_test_data.psql)"



