test-integration-recreate:
	docker-compose down --rmi local
	# start all containers detached except the integration container. We
	# start the integration container separately so it's isolated from
	# the exit codes of all other containers.
	docker-compose up -d --force-recreate --scale integration=0
	docker-compose up --exit-code-from integration integration

test-integration:
	docker-compose down
	# start all containers detached except the integration container. We
	# start the integration container separately so it's isolated from
	# the exit codes of all other containers.
	docker-compose up -d --scale integration=0
	docker-compose up --exit-code-from integration integration
