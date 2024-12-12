# ghs 📜

[![release](https://github.com/flexwie/ghs/actions/workflows/release.yml/badge.svg)](https://github.com/flexwie/ghs/actions/workflows/release.yml)

A npx-like script runner for GitHub gists

Fetch scripts from GitHub Gists and execute them as if they were native scripts. Uses the GitHub CLI for authentication under the hood.

## Installation

**Homebrew**

```sh
brew tap flexwie/homebrew-flexwie
brew install flexwie/flexwie/ghs
```

**apt**

```sh
echo "deb [trusted=yes] https://apt.fury.io/flexwie/ /" >> /etc/apt/sources.list.d/fury.list
apt update
apt install ghs
```

**nix**  
You can get a nix flake from https://github.com/flexwie/nur-packages

## Usage

```sh
ghs [<user>/]<gist>
```

for example: `ghs flexwie/test.sh`. If no username is provided, the currently logged in user will be assumed.

## Executors

Executors handle different gist types and languages and are mostly used for languages, that can't be invoked with a shebang (Go for example). Currently, there are two executors: one for shebang-style scripts and one for Go scripts.
