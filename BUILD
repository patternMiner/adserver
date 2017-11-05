load("@io_bazel_rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "adserver",
    srcs = [
        "main.go",
    ],
    data = glob([
        "data/**",
        "client/adtag/dist/**"
    ]),
    deps = [
        "//context",
        "//handlers"
    ],
    importpath = "github.com/patternMiner/adserver",
    visibility = ["//visibility:public"],
)
