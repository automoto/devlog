### Contributing

<!-- toc -->

* [Building Locally](#building-locally)
* [Running the tests](#running-the-tests)
* [Testing the gorelease builds](#testing-the-gorelease-builds)

<!-- tocstop -->

The `Makefile` has most of the commands you need to build, lint, and run the tests locally.

Fork and create a pull-request if you wish to contribute a fix or improvement to Devlog.

A contributor will build your PR in circle-ci and review the change. Please include unit tests and a description of the change in your pull request.
Consider submitting a github issue or a work in progress pull request first if you desire to make significant design changes.

#### Building Locally

*Prerequisites:*
  - go lang version 1.13

Simply clone this repository and run the following command to build the binary:
```shell
make build
```

This will create a binary locally you can run commands against already, like so:

``` shell
./devlog
```

Build and copy the binary to your local bin to access the CLI anywhere. It will likely prompt you for a password since it's needed to install things to your `/usr/local/bin`.

```shell
make install
```

Now you can run the command `devlog` from anywhere to generate a new devlog file:

``` shell
devlog
2019/09/02 22:00:32 Successfully saved dev log to directory: /home/dev/null
```

#### Running the tests
Run the tests is done via a simple `go test` command that also measures coverage. Run it through the make command:
``` shell
make test
```

Linting is done via:
```shell
make lint
```

#### Testing the gorelease builds
You may need to run the gorelease command locally for debugging. To do this, run this command:

``` shell
DEVLOG_VERSION=$(git describe --tags) goreleaser --rm-dist --snapshot --skip-publish
```

#### Update the TOC for the README and CONTRIBUTING

``` shell
# Readme
markdown-toc --bullets "*" --no-firsth1 -i README.md

# Contributing guide
markdown-toc --bullets "*" -i CONTRIBUTING.md
```