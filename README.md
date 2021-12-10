# Swoop

Swoop is a small, minimalist and lightweight golang library designed for removing the boiler plate from writing server rendered HTML web applications, backed by an optional datastore of some kind. Swoop is not a framework for creating API's (web apps producing JSON etc); It's simply a alternative to using just the standard lib for writing webapps, reducing the time needed to get your HTML to your users.

## How Is Swoop Different?

- Minimalism -> Swoop is built upon the belief that it shouldn't be difficult (or require a ton of boiler plate) to pull data from a datastore to render it in a broswer as HTML. While this requires a lot under the hood, these problems have already been solved by the Golang stdlib. Swoop aims to build upon such solutions and make it easy for the developer to create and the user to benefit from.
- Readability -> A Go developer should be able to read and understand a Swoop app from the source code; enabling productivity from day one. The Go syntax and type system helps with this, Swoop builds upon it.
- Scale -> While minimal in how it works, Swoop apps should be able to scale to the point where your hardware becomes the first bottleneck.
- Speed -> Swoop apps are fast! Go offers speed and true parallelism out of the box, which is utilized by Swoop where ever possible.

## How Does It Work?

Swoop doesn't reinvent the wheel, instead the Golang stdlib is used where ever possible. For example:

- HTTP -> `net/http`
- Views -> `template/html`
- Data Model -> Built-in Generics & Interfaces

### Concepts

#### Reading Data

There is only one type in Swoop that produces HTML for the user to view, a `Page`. A page is made up of the following structure:

- -> `Page`
- --> `Layout`
- --> `[]Component` || Just pass a `string` of HTML

A page is made up of one or more components and each component is fetched in parallel (for speed). A page doesn't concern itself with fetching data itself, its components are rendered in turn to make up the page content. Everything that the page needs to work is contained within its definition e.g. route path, list of components to render etc.

Each component fetches some data (from anywhere e.g. a Database) and then renders it as HTML. Simple. Components are reusable across any number of Pages.

There is only one type in Swoop that reads from a Datastore and models its data:

- -> `Query`

A query is used to fetch and return some data (modelled as a struct). Queries are typically called from components which in turn render the queried data as HTML.

#### Writing Data

There is only one type in Swoop that alters a Datastore (and therefore affects the resulting HTML).

- -> `Operation`

An operation is used to tell the server to change some data which'll be reflected the next time its HTML is rendered. Operations are commonly used to create, update and delete data in a Datastore. An operation **must** result in either a `nil` response (the users page doesn't change) or a `redirect` response (allowing the users page to update and reflect the data change).

### Usage

TODO
