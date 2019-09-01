package actor

import (
	"github.com/robertkrimen/otto"
)

type (
	// System a central point to register extensions and create actors.
	System struct {
		extensions []Extension
	}

	// Extension can add  stuff to the vm before things run
	Extension interface {
		Register(vm *otto.Otto)
	}

	// Actor is a reusable type that wraps some filthy "js"-"logic"
	Actor struct {
		vm *otto.Otto
	}
)

// NewSystem returns a new actor system.
func NewSystem(extensions ...Extension) *System {
	return &System{extensions}
}

// Actor creates a new *Actor
func (s *System) Actor(logic interface{}) (*Actor, error) {
	vm, _, err := otto.Run(logic)

	if err != nil {
		return nil, err
	}

	for _, ext := range s.extensions {
		ext.Register(vm)
	}

	return &Actor{vm}, nil
}

// Tell - tell the actor about an event, dont expect a response.
func (a *Actor) Tell(state, evt interface{}) error {
	_, err := a.vm.Call("handle", state, evt)

	return err
}

// Ask - tell the actor about the event, expect an answer.
func (a *Actor) Ask(state, evt interface{}) (otto.Value, error) {
	return a.vm.Call("handle", state, evt)
}

// Value turn a go map or struct into a js value.
func (a *Actor) Value(any interface{}) (otto.Value, error) {
	return a.vm.ToValue(any)
}

// Object creates an empty js object
func (a *Actor) Object() (*otto.Object, error) {
	return a.vm.Object("({})")
}
