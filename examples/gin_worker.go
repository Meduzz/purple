package examples

import (
	"io/ioutil"

	purple ".."
	"github.com/gin-gonic/gin"
)

/*
	In this example we sumarize the total amount of a bunch of orders.
	Each order should have a quantity and an price property.
	The worker will returned the total sum of the order.
	It should be printed in the response.
*/
func ginmain() {
	// create a gin server
	srv := gin.Default()
	// load the worker "logic"
	logic, _ := ioutil.ReadFile("aggregate.js")
	// create the worker
	worker, _ := purple.Worker(logic)

	srv.POST("/order", func(ctx *gin.Context) {
		// get the raw post body should be something like:
		/*
			[
				{price:5,quantiy:2},
				{price:2,quantity:5}
			]
		*/
		body, _ := ctx.GetRawData()
		// call the worker, and collect the response
		resp, _ := worker.Call("sumarize", string(body))

		// dig out an integer from the response.
		sum, _ := resp.ToInteger()
		ctx.String(200, "%d", sum)
	})

	srv.Run(":8080")
}
