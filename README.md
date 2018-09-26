# Gotime [![Build Status](https://travis-ci.org/onwsk8r/gotime.svg?branch=master)](https://travis-ci.org/onwsk8r/gotime) [![Coverage Status](https://coveralls.io/repos/github/onwsk8r/gotime/badge.svg?branch=master)](https://coveralls.io/github/onwsk8r/gotime?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/onwsk8r/gotime)](https://goreportcard.com/report/github.com/onwsk8r/gotime)

A small library of useful extensions to Golang's time package

## Documentation

Use the [Source](https://godoc.org/github.com/onwsk8r/gotime)! The package comment has additional details; all in all, we're just trying to keep the docs DRY.

- [gotime](https://godoc.org/github.com/onwsk8r/gotime)
- [holiday](https://godoc.org/github.com/onwsk8r/gotime/holiday)

## 700mb Overview

If there's one thing computers are bad at, it's working with dates. On the one hand, we have libraries such as [moment.js](https://momentjs.com/) which give us basic functionality that is mostly already present in Go. On the other hand, though, there's a lot of functionality that would be nice to have, such as parsing JSON dates that aren't accurate to the millisecond or anything involving workdays. This library also aims to provide convenience functions (aren't all functions really just convenience functions?) for finding and comparing dates.

## Help Us Grow

If you have functionality that you'd like to add to this library or decouple from an app, feel free to open a feature request with some sample code. If you have a great idea, we'd love to hear about it! If your timing is serendipitous, I'll write it myself.

## Issues/Contributions

- Fork the repo
- Create a feature branch or don't, it doesn't matter- it's your repo.
- Write your tests or ask for help doing so (I use [Ginkgo](https://onsi.github.io/ginkgo/))
- Write your functionality
- Open a PR

Failure to participate in step 3 will result in me fixing your code and commenting judiciously on your inability to write decent software in comments that will appear in godoc. The first result on Google for your username won't be very flattering.
