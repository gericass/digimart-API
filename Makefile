stop:
	docker kill $(docker ps -q)

rm:
	docker rm $(docker ps -a -q)

rmi:
	docker rmi $(docker images -q)
