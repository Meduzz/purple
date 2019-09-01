package purple

import (
	"./actor"
	"./worker"
)

// ActorSystem creates a new actor system
func ActorSystem(extensions ...actor.Extension) *actor.System {
	return actor.NewSystem(extensions...)
}

// Worker creates a new worker
func Worker(logic interface{}, extensions ...worker.Extension) (*worker.Worker, error) {
	return worker.NewWorker(logic, extensions...)
}
