/*
* cart.js manages the shopping cart logic. It holds a cart object
* and registers all event handlers that concern the cart and holds
* all necessary data such as the cart items.
 */

let cart = {
    items: [],
    remark: '',
    add: function (id, dish) {
        this.items.push({ id: id, name: dish })
    },
    remove: function (id) {
        let self = this
        this.items.forEach(function (e, i) {
            if (e.id === id) {
                self.items.splice(i, 1)
            }
        })
    }
}
