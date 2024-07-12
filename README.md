# hacker-rank-go

> Hacker Rank Problems solved in GoLang :)

[![Go/Test](https://github.com/gugacavalieri/hacker-rank-go/actions/workflows/go-test.yaml/badge.svg)](https://github.com/gugacavalieri/hacker-rank-go/actions/workflows/go-test.yaml)
[![Go/Lint](https://github.com/gugacavalieri/hacker-rank-go/actions/workflows/go-lint.yaml/badge.svg)](https://github.com/gugacavalieri/hacker-rank-go/actions/workflows/go-lint.yaml)


# Required Tools

* Docker, and Only Docker 🙂 🐳

# Getting Started

```bash
# Running Tests (All Challenges)
docker compose run --rm this go test ./...

# Running Linter (All Folders)
docker compose run --rm this golangci-lint run
```

# Future Improvements

* [x] Add Dockerfile and Compose to facilitate testing ✅
