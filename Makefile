composeup:
	docker compose up -d
	
composodown:
	docker compose down 

startdocker:
	open --background -a Docker

.PHONY: composeup composdown startdocker