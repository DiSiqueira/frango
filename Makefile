.PHONY: run proto build start

run:
	cd src && for d in ./*/ ; do (cd $$d; $(MAKE) run; cd ..); done

build:
	cd src && for d in ./*/ ; do (cd $$d; $(MAKE) build; cd ..); done

start: build run

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
