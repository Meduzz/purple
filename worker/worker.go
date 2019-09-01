package worker

import (
	"github.com/robertkrimen/otto"
)

type (
	// Worker a type that encapsulate a bunch of filthy logic.
	Worker struct {
		vm *otto.Otto
	}

	// Extension lets you add stuff to the vm before stuff runs.
	Extension interface {
		Register(vm *otto.Otto)
	}
)

// NewWorker creates a new worker with the logic and extensions installed.
func NewWorker(logic interface{}, extensions ...Extension) (*Worker, error) {
	vm := otto.New()

	for _, ext := range extensions {
		ext.Register(vm)
	}

	_, err := vm.Run(logic)

	if err != nil {
		return nil, err
	}

	return &Worker{vm}, nil
}

// Call execute @param method with @param request. request will be turned into a otto.Value before calling method.
func (w *Worker) Call(method string, request interface{}) (otto.Value, error) {
	param, err := w.vm.ToValue(request)

	if err != nil {
		return otto.Value{}, err
	}

	return w.vm.Call(method, nil, param)
}
