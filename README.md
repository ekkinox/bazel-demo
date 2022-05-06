# Bazel demo

- Bazel
```shell
bazel query //... --output label_kind | sort | column -t
bazel run //xxx
```

- Gazelle
```shell
bazel run //:gazelle
bazel run //:gazelle -- update-repos
```

