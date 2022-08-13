# dictionaryapi

A lightweight go module to access meanings of words using [dictionaryapi.dev](https://dictionaryapi.dev)

[![Build](https://github.com/atselvan/dictionaryapi/actions/workflows/build.yaml/badge.svg)](https://github.com/atselvan/dictionaryapi/actions/workflows/build.yaml)
[![Release](https://github.com/atselvan/dictionaryapi/actions/workflows/release.yaml/badge.svg)](https://github.com/atselvan/dictionaryapi/actions/workflows/release.yaml)
[![reference](https://img.shields.io/badge/godoc-docs-blue.svg?label=&logo=go)](https://godoc.org/github.com/atselvan/dictionaryapi)
[![Go Report Card](https://goreportcard.com/badge/github.com/atselvan/dictionaryapi)](https://goreportcard.com/report/github.com/atselvan/dictionaryapi)
[![codecov](https://codecov.io/gh/atselvan/dictionaryapi/branch/master/graph/badge.svg)](https://codecov.io/gh/atselvan/dictionaryapi)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=atselvan_dictionaryapi&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=atselvan_dictionaryapi)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fatselvan%2Fdictionaryapi.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fatselvan%2Fdictionaryapi?ref=badge_shield)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Features

* Get the meaning of a word

## Installation

```go
go get github.com/atselvan/dictionaryapi
```

## Usage

Code

```go
import (
	"fmt"
	"log"

	"github.com/atselvan/dictionaryapi"
)

func main() {
	dictionaryapi := dictionaryapi.NewClient()
	word, restErr := dictionaryapi.Word.Get("hello")
	if restErr != nil {
		log.Fatal(restErr)
	}
	fmt.Println(word)
}
```

Output

```log
&{hello  [{ https://api.dictionaryapi.dev/media/pronunciations/en/hello-au.mp3} {/həˈləʊ/ https://api.dictionaryapi.dev/media/pronunciations/en/hello-uk.mp3} {/həˈloʊ/ }]  [{noun [{"Hello!" or an equivalent greeting.  [] []}]} {verb [{To greet with "hello".  [] []}]} {interjection [{A greeting (salutation) said when meeting someone or acknowledging someone’s arrival or presence. Hello, everyone. [] []} {A greeting used when answering the telephone. Hello? How may I help you? [] []} {A call for response if it is not clear if anyone is present or listening, or if a telephone conversation may have been disconnected. Hello? Is anyone there? [] []} {Used sarcastically to imply that the person addressed or referred to has done something the speaker or writer considers to be foolish. You just tried to start your car with your cell phone. Hello? [] []} {An expression of puzzlement or discovery. Hello! What’s going on here? [] []}]}]}
```
