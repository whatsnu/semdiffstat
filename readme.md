THIS PACKAGE IS EXPERIMENTAL.

Package semdiffstat calculates semantic diffstats.

Semantic diffstats are intended to present readable, meaningful high level summaries of changes to a human.

This is currently just a quick sketch of what a semantic diffstat might look like, and a very crude implementation for generating one out of two Go files.

There's a little demo command in cmd/semdiffstat. It can be invoked manually:

```bash
$ semdiffstat a.go b.go
```

or as an external git diff tool:

```bash
$ GIT_EXTERNAL_DIFF=/path/to/semdiffstat git diff
```
