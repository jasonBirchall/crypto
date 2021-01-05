<p align="center">
  <h2 align="center">Crypto Tracker</h2>
  <p align="center">Cryptocurrency coin tracker written in Go</p>
</p>

---

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/jasonbirchall/crypto-tracker.svg)](https://github.com/jasonbirchall/crypto-tracker)
[![GoReportCard example](https://goreportcard.com/badge/github.com/jasonbirchall/crypto-tracker)](https://goreportcard.com/report/github.com/jasonbirchall/crypto-tracker)
[![GitHub release](https://img.shields.io/github/release/jasonbirchall/crypto-tracker.svg)](https://GitHub.com/jasonbirchall/crypto-tracker/releases/)
[![GitHub issues](https://img.shields.io/github/issues/jasonbirchall/crypto-tracker.svg)](https://GitHub.com/jasonbirchall/crypto-tracker/issues/)
[![tests](https://github.com/jasonbirchall/crypto-tracker//workflows/Run%20Tests/badge.svg)](https://github.com/jasonBirchall/crypto-tracker/actions?query=workflow%3A%22Run+Tests%22)

Crypto-tracker is a simple, robust command-line application that displays the price of popular cryptocurrencies.

### Install

These installation instructions are written for a Linux system. If you have a different kind of
computer, please amend the steps appropriately.

Please substitute the latest release number. You can see the latest release
number in the badge near the top of this page, and all available releases on
[this page](https://github.com/jasonBirchall/crypto-tracker/releases/).

```bash
export RELEASE=0.1.4
wget https://github.com/jasonbirchall/crypto-tracker/releases/download/${RELEASE}/crypto-tracker_${RELEASE}_linux_amd64.tar.gz
tar xzvf crypto-tracker_${RELEASE}_linux_amd64.tar.gz
mv crypto /usr/local/bin/
```

### Usage

Crypto-tracker will eventually have a number of subcommands. Execute: `crypto --help` in order to check them out.

#### track

Track was initially intended for something like [Polybar for i3](https://github.com/polybar/polybar) to display the price of certain coins and their rise/fall in the past five minutes. It is inspired by [polybar-crypto](https://github.com/willHol/polybar-crypto).

```bash
crypto track --coin btc,eth

£23852.33 | 3.56%  £782.73 | 5.43%
```

The track sub-command has a flag that needs to be set. The --coin, or -c for short, will allow you to specify which coins you wish to display.

### Develop

