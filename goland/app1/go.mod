module github.com/kaito2/go-bazel-sample/goland/app1

go 1.18

replace github.com/kaito2/go-bazel-sample/goland/lib/example => ../lib/example

require (
	github.com/kaito2/go-bazel-sample/goland/lib/example v0.0.0-00010101000000-000000000000
	github.com/rs/xid v1.4.0
)
