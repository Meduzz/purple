# purple

Based on robertkrimen/otto, which lets you run filthy javascript code
in your go-code. This library takes advantage of javascripts "portability"
and lets you execute it willy nilly.

## Actors

Actors are great, and I miss them in Go. This bit of the lib has something
that reassembles actors, but with a splash of that filthy javascript.

Atm they are kind of wonky to work with, since calling them (successfully)
requires you to pass at least one, pref. two otto.Objects/otto.Values 
(state & evt). Oh, and your js-handler method must be named `handle`.

## Workers

Did you ever build any software, where you wanted the "end user" to 
contribute bits of the logic, without forcing them to recompile the 
whole thing? Or force them into learning go...

This part of the lib is about extendability, under somewhat 
controlled forms.