# go-looker

This repository contains a Go client library for Looker's [REST API](https://docs.looker.com/reference/api-and-integration/api-reference/v3.0).


## Generate the client code

1. Download swagger.json from `curl -fSsL https://<looker-instance>/api/3.0/swagger.json -o swagger.json`
1. `make generate`
