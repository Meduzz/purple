package worker

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/robertkrimen/otto"
)

var (
	subject *Worker
)

type Multi struct{}

func (m *Multi) Register(vm *otto.Otto) {
	err := vm.Set("multi", func(call otto.FunctionCall) otto.Value {
		req := call.Argument(0).Object()

		av, _ := req.Get("a")
		bv, _ := req.Get("b")

		a, _ := av.ToInteger()
		b, _ := bv.ToInteger()

		ret, _ := vm.ToValue(a * b)

		return ret
	})

	if err != nil {
		panic(err)
	}
}

func TestMain(m *testing.M) {
	bs, err := ioutil.ReadFile("worker_test.js")

	if err != nil {
		panic(err)
	}

	subject, err = NewWorker(bs, &Multi{})

	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func TestCallMethods(t *testing.T) {
	req := make(map[string]interface{})
	req["a"] = 5
	req["b"] = 2

	add, err := subject.Call("add", req)

	if err != nil {
		t.Fatalf("Add threw an error %s", err)
	}

	addRes, _ := add.ToInteger()

	if addRes != 7 {
		t.Fatalf("AddRes was not 7 but %d", addRes)
	}

	sub, err := subject.Call("subtract", req)

	if err != nil {
		t.Fatalf("Subtract threw an error %s", err)
	}

	subRes, _ := sub.ToInteger()

	if subRes != 3 {
		t.Fatalf("SubRes was not 3 but %d", subRes)
	}

	multi, err := subject.Call("multi", req)

	if err != nil {
		t.Fatalf("Multi threw an error %s", err)
	}

	multiRes, _ := multi.ToInteger()

	if multiRes != 10 {
		t.Fatalf("MultiRes was not 10 but %d", multiRes)
	}
}

func TestWorkerWithBadCode(t *testing.T) {
	_, err := NewWorker([]byte("I will burn"))

	if err == nil {
		t.Fatal("There was no error!")
		t.Fail()
	}
}
