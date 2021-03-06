
hello-world:
	docker run hello-world

build-web-server:
	cd web_server && docker build -t web-server .

run-web-server:
	docker run -dit --name web-server-instance -p 8080:80 web-server

search-nginx:
	docker search --format "{{.Name}}: {{.StarCount}}" nginx
