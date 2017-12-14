# Telegraph

[![Build Status](https://travis-ci.org/dynastymasra/telegraph.svg?branch=master)](https://travis-ci.org/dynastymasra/telegraph)
[![Coverage Status](https://coveralls.io/repos/github/dynastymasra/telegraph/badge.svg?branch=master)](https://coveralls.io/github/dynastymasra/telegraph?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/dynastymasra/telegraph)](https://goreportcard.com/report/github.com/dynastymasra/telegraph)
[![GoDoc](https://godoc.org/github.com/dynastymasra/telegraph?status.svg)](https://godoc.org/github.com/dynastymasra/telegraph)
[![Version](https://img.shields.io/badge/version-2.0.0-orange.svg)](https://github.com/dynastymasra/telegraph/tree/2.0.0)
[![License: MIT](https://img.shields.io/badge/license-MIT-yellow.svg)](https://opensource.org/licenses/MIT)


Telegraph is Telegram bot API SDK for Go(Golang), Can use back off retry request to Telegram bot API

Reference to Telegram bot API [Documentation](https://core.telegram.org/bots/api)


## Installation

```bash
$ go get github.com/dynastymasra/telegraph
```

## How to use

import library `github.com/dynastymasra/telegraph`, 
See Telegram API [Documentation](https://core.telegram.org/bots/api#available-methods) to know available method can used, 
and what params can use in a method.

Create new client with no use back off retry, use params `access token` obtain from telegram bot father.
```go
client := telegraph.NewClient(<access_token>)

res, err := client.SetWebHook("https://www.cubesoft.co.id").SetCertificate("./LICENSE").SetMaxConnection(100).SetAllowedUpdates("1", "2", "3").Commit()
if err != nil {
	// Do something when error
}

info, res, err := client.GetWebHookInfo().Commit()
if err != nil {
	// Do something when error
}
```

Create new client with use back off retry, with params `access token` obtain from telegram bot father and `max interval` and `max elapsed time`
```go
client := telegraph.NewClientWithBackOff(<access_token>, telegraph.NewBackOff(<max_interval>, <max_elapsed_time>))

res, err := client.DeleteWebHook().Commit()
if err != nil {
	// Do something when error
}
```

Parse telegram web hook request, reference to telegram [Documentation](https://core.telegram.org/bots/api#getting-updates)

```go
message, err := telegraph.WebHookParseRequest(<request_in_[]byte>)
if err != nil {
	// Do something when error
}
```

## Contributing

If you find any issue you want to fix it, feel free to send me a pull request. 
And also if you have idea for improvement this library, feel free to send me a pull request.

## Library

* [GoRequest](https://github.com/parnurzeal/gorequest) - Simplified HTTP client ( inspired by famous SuperAgent lib in Node.js )
* [Backoff](https://github.com/cenkalti/backoff) - The exponential backoff algorithm in Go (Golang)
* [Gock](https://github.com/h2non/gock) - HTTP traffic mocking and expectations made easy for Go
* [Testify](https://github.com/stretchr/testify) - A toolkit with common assertions and mocks that plays nicely with the standard library