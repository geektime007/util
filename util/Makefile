# vim:set noet:

GO_TEST=go test -v -timeout=5m
GO_BENCHMARK=$(GO_TEST) -bench=.
GO_DOC=godoc2md -srclink="./%s\#L%d?s=%d:%d"

test:

benchmark:

docs:
	$(GO_DOC) github.com/geektime007/util/util > README.md

docs-git-add: docs
	cd .. && git add util/README.md

.PHONY: docs docs-git-add test benchmark
