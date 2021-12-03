# Swoop

Swoop is a small, lightweight and very fast HTTP framework designed to create server rendered HTML web applications backed by an optional datastore of some kind. Swoop is not a framework for creating API's (web apps producing JSON etc).

## How is Swoop different?

- Minimalism -> Swoop is built upon the belief that it shouldn't be difficult (or require a ton of boiler plate) to pull data from a datastore to render it in a broswer as HTML. While this requires a lot under the hood, these problems have already been solved. Swoop aims to build upon such solutions and make it easy for the developer to create and the user to benefit from.
- Readability -> A Go developer should be able to read and understand a Swoop app from the source code; enabling productivity from day one. The Go syntax and type system helps with this.
- Scale -> While minimal in how it works, Swoop apps should be able to scale to the point where your hardware becomes the first bottleneck.
- Speed -> Swoop apps are fast! Go offers speed and true parallelism out of the box, which is utilized by Swoop where ever possible.

## How does it work?

### Concepts

#### Reading Data

There are two types in Swoop that produce HTML for the user:

- -> Pages
- --> Components

A page is made up of one or more components and each component is fetched in parallel (for speed). A page doesn't concern itself with fetching data itself, its components are rendered in turn to make up the page content. Everything that the page needs to work is contained within its definition e.g. route path, list of components to render etc.

Each component fetches some data (from anywhere e.g. a Database) and then renders it as HTML. Simple. Components are reusable across any number of Pages.

There is only one type in Swoop that reads from the Datastore and models some of its data:

- -> Queries

A query is used to fetch and return some data (modelled as a struct). Queries are typically called from components which in turn render the queried data as HTML.

#### Writing Data

There is only one type in Swoop that alters the Datastore (and therefore affects the resulting HTML).

- -> Operations

An operation is used to tell the server to change some data which'll be reflected the next time its HTML is rendered. Operations are commonly used to create, update and delete data in a Datastore for example. An operation must result in either a `nil` response (the users page doesn't change) or a redirect response (allowing the users page to update and reflect the data change).

### Usage

First let's create a top level page for the index of our app:

> pages/index.go

```go
// TODO
```

> main.go

```go
package main

import (
	"fmt"
	"log"

	"github.com/michaeltelford/swoop"
	"github.com/michaeltelford/swoop/pages"
)

func main() {
	pages := buildPages()
	srv := swoop.NewServer(":8080", pages)

	fmt.Println("Starting server...")
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}

func buildPages() []pages.IPage {
	return []pages.IPage{
		pages.NewPage("/hello", "Hello, World!"),
	}
}
```
