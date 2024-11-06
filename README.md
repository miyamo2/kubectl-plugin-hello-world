# kubectl-plugin-hello-world

[![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/miyamo2/kubectl-plugin-hello-world?logo=go)](https://img.shields.io/github/go-mod/go-version/miyamo2/kubectl-plugin-hello-world?logo=go)
[![GitHub License](https://img.shields.io/github/license/miyamo2/kubectl-plugin-hello-world?&color=blue)](https://img.shields.io/github/license/miyamo2/kubectl-plugin-hello-world?&color=blue)

Implementation of _Hello World_ as a kubectl plugin

## Usage

### Install

```bash
git clone https://github.com/miyamo2/kubectl-plugin-hello-world.git
cd kubectl-plugin-hello-world
go install
```

### Run

```bash
kubectl hello_world
```