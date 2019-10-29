# [mklink] -- Make Link with Markdown Format

[![Build Status](https://travis-ci.org/spiegel-im-spiegel/mklink.svg?branch=master)](https://travis-ci.org/spiegel-im-spiegel/mklink)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/mklink/master/LICENSE)
[![GitHub release](http://img.shields.io/github/release/spiegel-im-spiegel/mklink.svg)](https://github.com/spiegel-im-spiegel/mklink/releases/latest)

## Declare [mklink] module

See [go.mod](https://github.com/spiegel-im-spiegel/mklink/blob/master/go.mod) file. 

## Usage

```go
link, err := mklink.New("https://git.io/vFR5M")
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(link.Encode(mklink.StyleMarkdown))
// Output:
// [GitHub - spiegel-im-spiegel/mklink: Make Link with Markdown Format](https://github.com/spiegel-im-spiegel/mklink)
```

## Command Line Interface

### Binaries

See [latest release](https://github.com/spiegel-im-spiegel/mklink/releases/latest).

### Usage

```
$ mklink -h
Usage:
  mklink [flags] [URL [URL]...]

Flags:
  -h, --help           help for mklink
  -i, --interactive    interactive mode
      --log string     output log
  -s, --style string   link style (default "markdown")
  -v, --version        output version of mklink
```

```
$ mklink https://git.io/vFR5M
[GitHub - spiegel-im-spiegel/mklink: Make Link with Markdown Format](https://github.com/spiegel-im-spiegel/mklink)
```

```
$ echo https://git.io/vFR5M | mklink
[GitHub - spiegel-im-spiegel/mklink: Make Link with Markdown Format](https://github.com/spiegel-im-spiegel/mklink)
```

```
$ mklink --log log.txt https://git.io/vFR5M
[GitHub - spiegel-im-spiegel/mklink: Make Link with Markdown Format](https://github.com/spiegel-im-spiegel/mklink)

$ cat log.txt
[GitHub - spiegel-im-spiegel/mklink: Make Link with Markdown Format](https://github.com/spiegel-im-spiegel/mklink)
```

### Interactive Mode

```
$ mklink -i
Input 'q' or 'quit' to stop
mklimk> https://git.io/vFR5M
[GitHub - spiegel-im-spiegel/mklink: Make Link with Markdown Format](https://github.com/spiegel-im-spiegel/mklink)
mklimk>
```

### Support Other Style

```
$ mklink -s html https://git.io/vFR5M
<a href="https://github.com/spiegel-im-spiegel/mklink">GitHub - spiegel-im-spiegel/mklink: Make Link with Markdown Format</a>
```

Support: `markdown`, `wiki`, `html`, `csv`

[mklink]: https://github.com/spiegel-im-spiegel/mklink "spiegel-im-spiegel/mklink: Make Link with Markdown Format"
[dep]: https://github.com/golang/dep "golang/dep: Go dependency management tool"
