# Swoop

Swoop is a small, minimalist and lightweight golang library designed for removing the boiler plate from writing server rendered HTML web applications. Swoop can be thought of as a wrapper around `net/http` and `html/template`, reducing the time needed to get your HTML to your users.

## How Does It Work?

Swoop doesn't reinvent the wheel, instead the Golang stdlib is used where ever possible. For example:

- HTTP -> `net/http`
- Views -> `html/template`
- Data Model -> `map[string]interface{}` acting as context for your views

## Concepts

There is only one type in Swoop that produces HTML for the user to view, a `Page`. A page is made up of the following structure:

- -> `Page`
- --> `Layout`
- --> `[]Component` || Just pass a `string` of HTML

A page is made up of a `string` or one or more components rendering HTML. A pages components are rendered in turn to make up the page content. Components will be rendered in the order they are defined. Everything that the page needs to work is contained within its definition e.g. route path, list of components to render etc. You create a page using code e.g. `page.NewPageFromString("GET", "/", layout, "Welcome!")`.

A layout is a `html/template` used to define everything around the page content. Layouts are used to define the page's header, footer, and any other HTML that needs to be rendered around the page content. Layouts should be stored at the root of the project in `templates/layouts`.

Each component is a `html/template` and can be passed a context (of `map[string]interface{}`) making it accessible from within the template. Components are reusable across any number of Pages. Components should be stored at the root of the project in `templates/components`.

Swoop doesn't concern itself with how you model or access your data, it just provides a simple way to render your data as HTML. The use of Databases etc. is not required to use Swoop. You can use any data you want to render as HTML by passing it to the `map[string]interface{}` context of a component.

## Usage

Check out this demo application as an example of how Swoop can be used:

https://github.com/michaeltelford/swooper

You can even clone the repo and run the demo application yourself, making changes to the code as you wish.
