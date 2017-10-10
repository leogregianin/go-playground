# Number of active repositories in GitHub

## Request

A typical request for one language for one month is:
https://api.github.com/search/repositories?q=size:>10+language:%22Go%22+pushed:2017-10-01..2017-10-31

The program exposes the following resource: http://localhost:8000 and a CSV file is returned.

The program wait 7 seconds between each request to not exceed the current GitHub rate limit (10 requests/minute)

## Install

```bash
go get -u github.com/gorilla/mux
```
