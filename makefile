.PHONY: loginApp
loginApp:
	@docker exec -it ishocon2-master_app_1 /bin/bash

.PHONY: bench
bench:
	@docker exec -it ishocon2-master_bench_1 ./benchmark --ip app:443