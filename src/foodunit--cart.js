const addItemLink = '.add-item'
const removeItemLink = '.remove-item'
const saveBtn = '#save-cart'

let cart = []

$(registerCartHandlers)

function registerCartHandlers() {
    $('#dishes').on('click', addItemLink, addItem)

    $('#cart-items').on('click', removeItemLink, removeItem)

    $(saveBtn).on('click', function (e) {
        e.preventDefault()
        saveCart()
    })
}

function addItem() {
    let id = $(this).attr('data-dish-id')
    let name = $(this).attr('data-dish-name')

    addToCart(id)

    let dish = {id: id, name: name}
    renderAddedCartDish(dish)
    console.log(cart)
}

function removeItem() {
    let id = $(this).attr('data-dish-id')
    cart[id]--
    renderRemovedCartDish($(this).parent())
    console.log(cart)
}

function addToCart(id) {
    if (cart[id] !== undefined) {
        cart[id]++
    } else {
        cart[id] = 1
    }
}

function saveCart() {
}
