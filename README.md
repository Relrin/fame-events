# fame-events

To update dependencies:
```
bazelisk run //:gazelle-update-repos && bazelisk run //:gazelle
```

Test all packages:
```
bazelisk run //pkg:pkg_test
```