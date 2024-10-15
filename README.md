# gomponents-starter-kit

<img src="logo.png" alt="Logo" width="300" align="right">

[![GoDoc](https://pkg.go.dev/badge/github.com/maragudk/gomponents-starter-kit)](https://pkg.go.dev/github.com/maragudk/gomponents-starter-kit)
[![Go](https://github.com/maragudk/gomponents-starter-kit/actions/workflows/ci.yml/badge.svg)](https://github.com/maragudk/gomponents-starter-kit/actions/workflows/ci.yml)
[![Go](https://github.com/maragudk/gomponents-starter-kit/actions/workflows/cd.yml/badge.svg)](https://github.com/maragudk/gomponents-starter-kit/actions/workflows/cd.yml)

A starter kit for building a web app with gomponents, HTMX, and TailwindCSS in Go.

Made with ✨sparkles✨ by [maragu](https://www.maragu.dev/).

Does your company depend on this project? [Contact me at markus@maragu.dk](mailto:markus@maragu.dk?Subject=Supporting%20your%20project) to discuss options for a one-time or recurring invoice to ensure its continued thriving.

## Getting started

The easiest way to get started is to [Use this template](https://github.com/new?template_name=gomponents-starter-kit&template_owner=maragudk) to create a new repository. Or you could clone this repository the traditional way:

```shell
git clone git@github.com:maragudk/gomponents-starter-kit.git your-app-name
```

After that, you can start the app with:

```shell
make start
```

If you make style changes, watch the CSS with:

```shell
make watch-css
```

You can run tests and linting with:

```shell
make test lint
```

## Deploying

The [CD workflow](.github/workflows/cd.yml) automatically builds a multi-platform Docker image and pushes it to the Github container registry GHCR.io, tagged with the commit hash as well as `latest`.

You can try building the image locally with:

```shell
make build-docker
```

Note that [you need the containerd image store enabled](https://docs.docker.com/desktop/containerd/#enable-the-containerd-image-store) for this to work.
