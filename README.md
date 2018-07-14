# Gotime

A small library of useful extensions to Golang's time package

## Documentation

Use the [Source](https://godoc.org/onwsk8r/gotime)!

## 700mb Overview
So I had this JSON that looked something like this:
```json
"date": "2018-07-14"
```
And I got tired of writing my own date type for like the millionth
time because someone didn't have an ISO date down to the millisecond.
I decided to write a function that could parse anything that looks like
an ISO date, and then I wanted to add simpler parsing for other formats
as well, like when you use them more than once.

## Help Us Grow!
If you have functionality that you'd like to add to this library or
decouple from an app, feel free to open a feature request with some
sample code. If you have a great idea, we'd love to hear about it!
If your timing is serendipitous, I'll write it myself.

## Issues/Contributions
- Fork the repo
- Create a feature branch or don't, it doesn't matter- it's your repo.
- Write your tests or ask for help doing so (I use [Ginkgo](https://onsi.github.io/ginkgo/))
- Write your functionality
- Open a PR

Failure to participate in step 3 will result in me fixing your code and
commenting judiciously on your inability to write decent software in
comments that will appear in godoc. The first result on Google for your
username won't be very flattering.

## TODO
- Add function to parse any ISO8601 date format
- Custom unmarshal function
- Add list of custom formats and parse function that tries them
