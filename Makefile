build-sync-db:
	go build -o ./bin/puppet-sync-db ./cmd/puppet-sync-db/

build-promotion-db:
	go build -o ./bin/promotion-service ./cmd/puppet-promotion-service/