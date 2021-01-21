<p align="center">
  <img alt="GoReleaser Logo" src="images/logo" height="140" />
  <h2 align="center">Crypto CLI</h2>
  <p align="center">Cryptocurrency command-line tool written in Go</p>
</p>

---

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/jasonbirchall/crypto.svg)](https://github.com/jasonbirchall/crypto)
[![GoReportCard example](https://goreportcard.com/badge/github.com/jasonbirchall/crypto)](https://goreportcard.com/report/github.com/jasonbirchall/crypto)
[![GitHub release](https://img.shields.io/github/release/jasonbirchall/crypto.svg)](https://GitHub.com/jasonbirchall/crypto/releases/)
[![GitHub issues](https://img.shields.io/github/issues/jasonbirchall/crypto.svg)](https://GitHub.com/jasonbirchall/crypto/issues/)
[![tests](https://github.com/jasonbirchall/crypto//workflows/tests/badge.svg)](https://github.com/jasonBirchall/crypto/actions?query=workflow%3A%22tests%22)
[![CodeQL](https://github.com/jasonbirchall/crypto//workflows/CodeQL/badge.svg)](https://github.com/jasonBirchall/crypto/actions?query=workflow%3ACodeQL)
[![crypto](https://snapcraft.io/crypto/badge.svg)](https://snapcraft.io/crypto)


Crypto is a simple, robust command-line application that displays the price of popular cryptocurrencies.

### Install

#### Snap

```bash
sudo snap install crypto
```

#### Manual

These installation instructions are written for a Linux system. If you have a different kind of
computer, please amend the steps appropriately.

Please substitute the latest release number. You can see the latest release
number in the badge near the top of this page, and all available releases on
[this page](https://github.com/jasonBirchall/crypto/releases/).

```bash
export RELEASE=0.1.4
wget https://github.com/jasonbirchall/crypto/releases/download/${RELEASE}/crypto_${RELEASE}_linux_amd64.tar.gz
tar xzvf crypto_${RELEASE}_linux_amd64.tar.gz
mv crypto /usr/local/bin/
```

### Usage

Crypto will eventually have a number of subcommands. Execute: `crypto --help` in order to check them out.

#### Track

Track was initially intended for something like [Polybar for i3](https://github.com/polybar/polybar) to display the price of certain coins and their rise/fall in the past five minutes. It is basically a go port of the [polybar-crypto](https://github.com/willHol/polybar-crypto) code base.

The track sub-command has a flag that needs to be set. The `--coin`, or `-c` for short, will allow you to specify which coins you wish to display.

```bash
crypto track --coin btc,eth

BTC £26332.21 | -2.57%   ETH £905.26 | -0.8%
```

#### Graph

Graph uses the same data set to display a coin in an ascii graph in the terminal. This command utilises [asciigraph](https://github.com/guptarohit/asciigraph), which does all the heavy lifting. 

The graph sub-command requires you to specify a coin using the `--coin` or `-c` flag followed by a coin of your choice, for example:

```bash
crypto graph --coin btc -H 10

 26673 ┤           ╭──╮
 26369 ┤           │  ╰╮
 26065 ┤    ╭╮    ╭╯   ╰──╮
 25762 ┤    │╰╮ ╭─╯       │
 25458 ┤   ╭╯ ╰─╯         ╰╮
 25154 ┤  ╭╯               │
 24850 ┤  │                │
 24547 ┤ ╭╯                ╰─
 24243 ┤ │
 23939 ┼╮│
 23636 ┤╰╯

```

The `-H` or `--height` flag allows you to specify the height of the graph displayed.

### Develop

You will need Go installed (version 1.15 or greater).

#### Build locally

Run `make` to create a `crypto` binary.

#### Testing

Run `make test` to run the unit tests.

#### Updating / Publishing

Update the `pkg/version.go` with the tag version and then run `make release`, which creates the release using [GoReleaser](https://github.com/goreleaser/goreleaser).
