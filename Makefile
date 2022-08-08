.PHONY: accountapi unit-tests integration-tests clean

accountapi: $(call print-target)
	@docker-compose up -d postgresql
	@docker-compose up -d accountapi

unit-tests: $(call print-target)
	@docker-compose up --build --no-start unit-tests
	@docker-compose up --abort-on-container-exit --exit-code-from unit-tests unit-tests

integration-tests: $(call print-target) accountapi
	@docker-compose up --build --no-start integration-tests
	@docker-compose up --abort-on-container-exit --exit-code-from integration-tests integration-tests

clean: $(call print-target)
	@docker-compose down


define print-target
    @printf "Executing target: $@"
endef