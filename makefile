WORKLOAD = 3

.PHONY: loginApp
loginApp:
	@docker exec -it -w /home/ishocon/webapp/go ishocon2-master_app_1 /bin/bash

.PHONY: up
up:
	@docker-compose up -d

.PHONY: stop
stop:
	@docker-compose stop

.PHONY: upWebApp
upWebApp:
	@docker exec -it -w /home/ishocon/webapp/go ishocon2-master_app_1 ./webapp

.PHONY: buildWebApp
buildWebApp:
	@docker exec -it -w /home/ishocon/webapp/go ishocon2-master_app_1 go build -o webapp

.PHONY: bench
bench:
	@docker exec -it ishocon2-master_bench_1 ./benchmark --workload ${WORKLOAD} --ip app:443

.PHONY: connDB
connDB:
	@docker exec -it -w /home/ishocon/webapp/go ishocon2-master_app_1 mysql -u ishocon -pishocon -D ishocon2

.PHONY: dumpSchema
dumpSchema:
	@docker exec -it -w /home/ishocon/webapp/go ishocon2-master_app_1 mysqldump -u ishocon --no-data ishocon2 -pishocon > dbSchema.sql