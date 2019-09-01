// handle the actor method, that sets the balance.
var handle = function (evt) {
    switch (evt.type) {
        case 'add':
            this.balance += evt.value
        case 'subtract':
            this.balance -= evt.value
        default:
            console.log("Unhandled event type: ", evt.type)
    }
}

// sumarize the worker method, that sums up the value of orders
var sumarize = function (order) {
    var sum = 0
    // we dont have any es6 features :(
    for (a in order) {
        var qty = a.quantity
        var price = a.price

        sum += qty * price
    }

    return sum
}