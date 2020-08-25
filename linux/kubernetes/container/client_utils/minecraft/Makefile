CURDIR := $(shell pwd)

run-server:
	docker run -e EULA=TRUE -p 25565:25565  -v $(CURDIR)/data:/data:Z --label=disable --name minecraft1 itzg/minecraft-server

