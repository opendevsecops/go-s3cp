[![Codacy Badge](https://api.codacy.com/project/badge/Grade/b07a461e6a9a48fc84226baefff06423)](https://www.codacy.com/app/OpenDevSecOps/go-s3cp?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=opendevsecops/go-s3cp&amp;utm_campaign=Badge_Grade)
[![Follow on Twitter](https://img.shields.io/twitter/follow/opendevsecops.svg?logo=twitter)](https://twitter.com/opendevsecops)

# go-s3cp

Simple utility to copy files from/to s3.

## Rationale

The AWS CLI is the defacto tool of choice for anything AWS related including S3. However, the CLI toolkit is relatively heavy and requires a number of dependencies which will bloat any standard docker image. The s3cp command solves this problem by providing a simple, standalone AWS S3 copy utility fit for minimal docker containers.

## Getting Started

To install s3cp simply do:

```sh
$ go get github.com/opendevsecops/go-s3cp
```

Once the command is installed in your home go/bin folder execute it like this:

```sh
$ ~/go/bin/go-s3cp --help
```

The command is very simple as it serves a single purpose (copying a file to/from/ s3). To copy a local file into s3 do:

```sh
$ ~/go/bin/go-s3cp --from /path/to/file --to s://bucket/key
```

To copy a file from s3 do:

```sh
$ ~/go/bin/go-s3cp --to /path/to/file --from s://bucket/key
````
