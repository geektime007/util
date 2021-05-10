# vim:set noet:

CMDS=default test benchmark docs docs-git-add clean
DIRS=lock\
	 log\
	 util\
	 orderedmap\
	 hash\
	 json\


default: test

# Install all packages including vendors to GOPATH/pkg (for autocompeletion of gocode)
install:
	go install ./...
	go install ./vendor/...

dep_ensure:
	dep ensure

deps: dep_ensure install

$(CMDS): $(DIRS)
$(DIRS):
	make -C $@ $(MAKECMDGOALS)

.PHONY: $(CMDS) $(DIRS) install deps
