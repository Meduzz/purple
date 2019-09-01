package examples

import (
	"encoding/json"
	"io/ioutil"

	nats "github.com/nats-io/go-nats"

	purple ".."
	"../actor"
)

type (
	// AccountChange describes which account to change, and how much.
	AccountChange struct {
		Account int `json:"account"`
		Amount  int `json:"amount"`
	}
)

/*
	Here we use actors to update balances of a bunch of accounts.
	Each account is stored in a "db".
	The changes to the account balance is recieved through nats.
	For each invocation (change) the account balance will be loaded from the "db".
	It will be put in an object, the js handle method will be bound to this object (hence this.balance) in the js code.
	After the new balance has been calculated in the actor, we put it back in the "db".

	This example completely ignores error handling.
*/
func actormain() {
	// connect to nats
	conn, _ := nats.Connect("nats://localhost:4222")
	// our "db", with account balances
	accounts := make(map[int]int)
	// Our actor body aka the logic, loaded from the aggregate.js file.
	body, _ := ioutil.ReadFile("aggregate.js")
	// actorSystems acts as a repository of your extension, and applies them to all actors you create.
	actorSystem := purple.ActorSystem()
	// otto can handles []byte, strings and more and it can all be used here
	actor, _ := actorSystem.Actor(body)

	// subscribe to add balance events
	conn.QueueSubscribe("account.add", "actors", func(msg *nats.Msg) {
		account := &AccountChange{}
		json.Unmarshal(msg.Data, account)

		actorEvent := make(map[string]interface{})
		actorEvent["type"] = "add"
		actorEvent["amount"] = account.Amount

		// load the account balance
		balanceBefore := accounts[account.Account]
		balanceAfter := callActor(actor, balanceBefore, actorEvent)

		// store it into the "db"
		accounts[account.Account] = balanceAfter
	})

	// subscribe to add balance events
	conn.QueueSubscribe("account.subtract", "actors", func(msg *nats.Msg) {
		account := &AccountChange{}
		json.Unmarshal(msg.Data, account)

		actorEvent := make(map[string]interface{})
		actorEvent["type"] = "subtract"
		actorEvent["amount"] = account.Amount

		balanceBefore := accounts[account.Account]
		balanceAfter := callActor(actor, balanceBefore, actorEvent)

		// store it into the "db"
		accounts[account.Account] = balanceAfter
	})
}

func callActor(actor *actor.Actor, balanceBefore int, actorEvent map[string]interface{}) int {
	// we create an otto.Object here, so we can read the new balance afterwards.
	state, _ := actor.Object()
	state.Set("balance", balanceBefore)
	// we create an otto.Value here because this is just going in, and never comes back.
	event, _ := actor.Value(actorEvent)

	actor.Tell(state, event)

	// dig out the new ballance from state
	ottoBalance, _ := state.Get("balance")
	balanceAfter, _ := ottoBalance.ToInteger()

	// from int64 to int
	return int(balanceAfter)
}
