# gigamoji

generate emoji (like :+1: in slack) banner which support some bitmap font faces

[![PkgGoDev](https://pkg.go.dev/badge/kyoh86/gigamoji)](https://pkg.go.dev/kyoh86/gigamoji)
[![Go Report Card](https://goreportcard.com/badge/github.com/kyoh86/gigamoji)](https://goreportcard.com/report/github.com/kyoh86/gigamoji)
[![Coverage Status](https://img.shields.io/codecov/c/github/kyoh86/gigamoji.svg)](https://codecov.io/gh/kyoh86/gigamoji)
[![Release](https://github.com/kyoh86/gigamoji/workflows/Release/badge.svg)](https://github.com/kyoh86/gigamoji/releases)

## Install

```
go get github.com/kyoh86/gigamoji
```

## Usage

```
$ go run . giga                                                      
□□□□□□□□□□□□□□□□
□□□□□■□□□□□□□□□□
■■■□□□□□■■■□■■■□
■□■□□■□□■□■□□□■□
■□■□□■□□■□■□■■■□
■■■□□■□□■■■□■□■□
□□■□□■□□□□■□■■■□
■■■□□□□□■■■□□□□□
□□□□□□□□□□□□□□□□
```

We can change text size by `--size|s`

```
$ go run . --size 10 giga                                            
□□□□□□□□□□□□□□□□□□□□
□□□□□□□■□□□□□□□□□□□□
□□□□□□□□□□□□□□□□□□□□
□□□□□□□□□□□□□□□□□□□□
□■■■□□■■□□□■■■□□■■□□
■□□■□□□■□□■□□■□□□□■□
■□□■□□□■□□■□□■□□■■■□
■□□■□□□■□□■□□■□■□□■□
□■■■□□■■■□□■■■□□■■■□
□□□■□□□□□□□□□■□□□□□□
□■■□□□□□□□□■■□□□□□□□
□□□□□□□□□□□□□□□□□□□□
□□□□□□□□□□□□□□□□□□□□
```

Dots can be changed by `--foreground|f` and `--background|b`.

```
$ go run . --foreground ○ --background ● giga         
●●●●●●●●●●●●●●●●
●●●●●○●●●●●●●●●●
○○○●●●●●○○○●○○○●
○●○●●○●●○●○●●●○●
○●○●●○●●○●○●○○○●
○○○●●○●●○○○●○●○●
●●○●●○●●●●○●○○○●
○○○●●●●●○○○●●●●●
●●●●●●●●●●●●●●●●
```

Emoji's can be specified like below. (`pbcopy` is the command to copy to clipboard)

```
$ go run . --foreground :red_circle: --background :white_circle: giga | pbcopy
```

And we can paste it in Slack then get big banner like below.

<img width="443" alt="" src="https://user-images.githubusercontent.com/5582459/64483747-0af54a00-d243-11e9-9424-42ac787dcbc2.png">

# LICENSE

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg)](http://www.opensource.org/licenses/MIT)

This is distributed under the [MIT License](http://www.opensource.org/licenses/MIT).

## Fonts

Following fonts are included in thie program.

* [ELISA Font Home Page by YAFO](http://hp.vector.co.jp/authors/VA002310/)
* [JF Dot M+10](http://jikasei.me/font/jf-dotfont/)
* [JF Dot M+12](http://jikasei.me/font/jf-dotfont/)
* [JF Dot Shinonome Gothic 14](http://jikasei.me/font/jf-dotfont/)
* [JF Dot Shinonome Gothic 16](http://jikasei.me/font/jf-dotfont/)

Thanks a lot for authors!
