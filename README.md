# gowall

A simple Go implementation of the Coderwall API (http://coderwall.com/api)

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

See main.go for an example.

## Changelog

See `CHANGELOG.md` for details.

## License

This is released under the
[MIT license](http://www.opensource.org/licenses/mit-license.php).
See `LICENSE.md` file for details.
