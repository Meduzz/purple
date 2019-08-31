package purple

import (
	"./actor"
)

// ActorSystem creates a new actor system
func ActorSystem(extensions ...actor.Extension) *actor.System {
	return actor.NewSystem(extensions...)
}
