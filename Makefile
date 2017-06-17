.PHONY: run proto

run:
	cd ./src/api/ && $(MAKE) run
	cd ./src/search/ && $(MAKE) run

proto:
	mkdir -p .temp
	cp -R proto .temp
	for f in .temp/proto/**/*.proto; do \
		protoc --go_out=plugins=grpc:. $$f; \
		echo compiled: $$f; \
	done
	find .temp -name "*.proto" -type f -delete
	cd src && for d in ./*/ ; do (cd $$d; rm -rf proto); done
	cd src && for d in ./*/ ; do (cp -R ../.temp/proto $$d); done
	rm -rf .temp
