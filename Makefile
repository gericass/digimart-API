stop:
	docker kill $(docker ps -q)

rm:
	docker rm $(docker ps -a -q)

rmi:
	docker rmi $(docker images -q)

migrate/up:
	sql-migrate up

migrate/down:
	sql-migrate down
