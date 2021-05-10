# vim:set noet:

GO_TEST=go test -v -timeout=5m
GO_BENCHMARK=$(GO_TEST) -bench=.
GO_DOC=godoc2md -srclink="./%s\#L%d?s=%d:%d"

default: test

test:
	$(GO_TEST)

benchmark:
	$(GO_BENCHMARK)

docs:
	$(GO_DOC) github.com/geektime007/util/log > README.md

docs-git-add: docs
	cd .. && git add log/README.md

.PHONY: default test benchmark docs docs-git-add
