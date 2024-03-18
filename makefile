GEN_DIR=internal/gen

swagger-gen:
	if ! [ -d $(GEN_DIR) ]; then \
	    mkdir $(GEN_DIR); \
	elif [ -d $(GEN_DIR) ]; then \
		rm -rf $(GEN_DIR); \
		mkdir $(GEN_DIR); \
	fi && \
	swagger generate server -t internal/gen -f ./api/swagger.yml --exclude-main -A anti-brutForce && \
	go mod tidy && \
	git add $(GEN_DIR)

run:
	go run ./cmd/cmd.go