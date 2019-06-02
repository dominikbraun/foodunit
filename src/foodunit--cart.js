const addItemLink = '.add-item'
const removeItemLink = '.remove-item'
const saveBtn = '#save-cart'
const allOrdersBtn = '#show-all-orders'

let cart = []
let explicitRemoves = []

let total = 0

$(registerCartHandlers)

function registerCartHandlers() {
    $('#dishes').on('click', addItemLink, addItem)

    $('#cart-items').on('click', removeItemLink, removeItem)

    $('#cart-remark').on('keydown', function () {
        $(saveBtn).attr('disabled', false)
        $(saveBtn).html('Bestellung speichern')
    })

    $(saveBtn).on('click', function (e) {
        e.preventDefault()
        $(saveBtn).attr('disabled', true)
        $(saveBtn).html('<span class="spinner-grow spinner-grow-sm" role="status" aria-hidden="true"></span>')
        saveCart()
    })

    $(allOrdersBtn).on('click', loadAllOrders)
}

function addItem() {
    let id = $(this).attr('data-dish-id')
    let name = $(this).attr('data-dish-name')
    let price = $(this).attr('data-dish-price')

    if (cart[id] !== undefined) {
        cart[id]++
    } else {
        cart[id] = 1
    }

    let dish = {dish_id: id, name: name, price: price}
    renderAddedCartDish(dish)

    total += parseFloat(price)
    renderTotal(total)

    $(saveBtn).attr('disabled', false)
    $(saveBtn).html('Bestellung speichern')
}

function removeItem() {
    let id = $(this).attr('data-dish-id')
    let isFromServer = $(this).attr('data-from-server')
    let price = $(this).attr('data-dish-price')

    if (isFromServer === "true") {
        explicitRemoves.push(id)
    } else {
        cart[id]--
    }
    renderRemovedCartDish($(this).parent())

    total -= parseFloat(price)
    renderTotal(total)

    $(saveBtn).attr('disabled', false)
    $(saveBtn).html('Bestellung speichern')
}

let offerId = 0

function saveCart() {
    if (offerId === 0) {
        $.ajax({
            url: 'api/offers',
            type: 'get',
            success: function (res) {
                let offers = JSON.parse(res)
                offerId = offers[0].supplier_id
                continueSaveCart()
            }
        })
    } else {
        continueSaveCart()
    }
}

function continueSaveCart() {
    let calls = 0
    let remark = $('#cart-remark').text()

    if (remark !== '') {
        $.ajax({
            url: 'api/remark/' + offerId + '/' + remark,
            type: 'get',
            beforeSend: function(xhr) {
                calls++;
            },
            success: function (res) {
                calls--;
                if (calls === 0) {
                    window.setTimeout(function () {
                        $(saveBtn).html('Bestellung gespeichert')
                    }, 500)
                }
            }
        })
    }
    for (let id in cart) {
        for (let i = 0; i < cart[id]; i++) {
            $.ajax({
                url: 'api/add/' + offerId + '/' + id,
                type: 'get',
                beforeSend: function(xhr) {
                    calls++;
                },
                success: function (res) {
                    calls--;

                    if (calls === 0) {
                        window.setTimeout(function () {
                            $(saveBtn).html('Bestellung gespeichert')
                        }, 500)
                    }
                }
            })
        }
    }
    for (let id of explicitRemoves) {
        $.ajax({
            url: 'api/del/' + offerId + '/' + id,
            type: 'get',
            beforeSend: function(xhr) {
                calls++;
            },
            success: function (res) {
                calls--;
                if (calls === 0) {
                    $(saveBtn).html('Bestellung gespeichert')
                }
            }
        })
    }
}

function loadAllOrders() {
    if (offerId === 0) {
        $.ajax({
            url: 'api/offers',
            type: 'get',
            success: function (res) {
                let offers = JSON.parse(res)
                offerId = offers[0].supplier_id
                continueAllOrders()
            }
        })
    } else {
        continueAllOrders()
    }
}

function continueAllOrders() {
    $.ajax({
        url: 'api/orders/' + offerId,
        type: 'get',
        success: function (res) {
            let orders = JSON.parse(res)
            renderAllOrders(orders)
        }
    })
}
