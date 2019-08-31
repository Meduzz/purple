# purple

Based on robertkrimen/otto, which lets you run filthy javascript code
in your go-code. This library takes advantage of javascripts "portability"
and lets you execute it willy nilly.

## Actors

Actors are great, and I miss them in Go. This bit of the lib has something
that reassembles actors, but with a splash of that filthy javascript.

Atm they are kind of wonky to work with, since calling them (successfully)
requires you to pass two otto.Objects (state & evt). But that will be addressed sooner or later. Oh, and your js-handler method must be named
`handle`.