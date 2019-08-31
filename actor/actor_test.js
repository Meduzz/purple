var handle = function (evt) {
    switch (evt.type) {
        case 'add':
            this.sum = add(this.sum, evt.add)
            break;
        case 'get':
            return this.sum
        default:
            console.log('Unhandled event type:', evt.type)
    }
}