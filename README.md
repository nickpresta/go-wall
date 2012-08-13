# gowall

A simple Go implementation of the Coderwall API (http://coderwall.com/api).

Current build status: <a href="http://goci.me/project/go-wall/gowall">
    <img src="http://goci.me/project/image/go-wall/gowall" />
</a>

## Getting the Code

    go get github.com/NickPresta/go-wall/gowall

## Usage

Import this library in your code to use it:

    import (
        "github.com/NickPresta/go-wall/gowall"
    )

Fetch your user details:

    user, err := gowall.FetchUser("NickPresta")
    /* Do something with user */

See `main.go` for an example.

## Documentation

View the documentation on
[GoPkgDoc](http://go.pkgdoc.org/github.com/NickPresta/go-wall/gowall).

## Tests

A very simple test is available via `go test`.

## Changelog

See `CHANGELOG.md` for details.

## License

This is released under the
[MIT license](http://www.opensource.org/licenses/mit-license.php).
See `LICENSE.md` for details.
