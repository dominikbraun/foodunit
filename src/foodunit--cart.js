const addItemLink = '.add-item'
const removeItemLink = '.remove-item'
const saveBtn = '#save-cart'

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
}

function addItem() {
    let id = $(this).attr('data-dish-id')
    let name = $(this).attr('data-dish-name')

    if (cart[id] !== undefined) {
        cart[id]++
    } else {
        cart[id] = 1
    }

    let dish = {dish_id: id, name: name}
    renderAddedCartDish(dish)

    $(saveBtn).attr('disabled', false)
    $(saveBtn).html('Bestellung speichern')
}

function removeItem() {
    let id = $(this).attr('data-dish-id')
    let isFromServer = $(this).attr('data-from-server')

    if (isFromServer === "true") {
        explicitRemoves.push(id)
    } else {
        cart[id]--
    }
    renderRemovedCartDish($(this).parent())

    $(saveBtn).attr('disabled', false)
    $(saveBtn).html('Bestellung speichern')
}

let offerId = 0

function saveCart() {
    $.ajax({
        url: 'api/offers',
        type: 'get',
        success: function (res) {
            let offers = JSON.parse(res)
            offerId = offers[0].supplier_id
        }
    })

    let calls = 0

    let remark = $('#cart-remark').val()

    $.ajax({
        url: 'api/remark/' + offer + '/' + remark,
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

    for (let id in cart) {
        for (let i = 0; i < cart[id]; i++) {
            $.ajax({
                url: 'api/add/' + offer + '/' + id,
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

    for (let id of explicitRemoves) {
        $.ajax({
            url: 'api/del/' + offer + '/' + id,
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
