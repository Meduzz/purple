package actor

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/robertkrimen/otto"

	"github.com/robertkrimen/otto/parser"
)

var (
	system  = NewSystem(&Add{})
	subject *Actor
)

type (
	Add struct{}
)

func (a *Add) Register(vm *otto.Otto) {
	vm.Set("add", func(call otto.FunctionCall) otto.Value {
		first, _ := call.Argument(0).ToInteger()
		second, _ := call.Argument(1).ToInteger()

		ret, _ := otto.ToValue(first + second)
		return ret
	})
}

func TestMain(m *testing.M) {
	bs, err := ioutil.ReadFile("actor_test.js")

	if err != nil {
		panic(err)
	}

	ps, err := parser.ParseFile(nil, "", bs, 0)

	if err != nil {
		panic(err)
	}

	subject, err = system.Actor(ps)

	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func TestTell(t *testing.T) {
	data := make(map[string]interface{})
	data["add"] = 5
	data["type"] = "add"

	// It seems to handle maps well, json annotated structs, not so much.
	evt, err := subject.Value(data)

	// Clonky solution, and what's up with the paranteses?
	//	evt, err := subject.vm.Object(`({type:"add", add:5})`)

	if err != nil {
		t.Fatalf("Evt threw an error %s", err)
	}

	state, err := subject.Object()

	if err != nil {
		t.Fatalf("State threw an error %s", err)
	}

	state.Set("sum", 2)

	err = subject.Tell(state, evt)

	if err != nil {
		t.Fatalf("Tell threw an error %s", err)
	}

	value, err := state.Get("sum")

	if err != nil {
		t.Fatalf("Tell threw an error %s", err)
	}

	end, err := value.ToInteger()

	if err != nil {
		t.Fatalf("Tell threw an error %s", err)
	}

	if end != 7 {
		t.Fatalf("State value was not 7 but %d", end)
	}
}

func TestAsk(t *testing.T) {
	evt, err := subject.Object()

	if err != nil {
		t.Fatalf("Evt threw an error %s", err)
	}

	evt.Set("type", "get")

	state, err := subject.Object()

	if err != nil {
		t.Fatalf("State threw an error %s", err)
	}

	state.Set("sum", 7)

	value, err := subject.Ask(state, evt)

	if err != nil {
		t.Fatalf("State threw an error %s", err)
	}

	end, err := value.(otto.Value).ToInteger()

	if err != nil {
		t.Fatalf("State threw an error %s", err)
	}

	if end != 7 {
		t.Fatalf("Return was not 7 but %d", end)
	}
}

func TestInvalidActorLogic(t *testing.T) {
	actor, err := system.Actor("I will burn!")

	if err == nil {
		t.Fatal("Expected an error")
	}

	if actor != nil {
		t.Fatal("There should be no actor here")
	}
}
