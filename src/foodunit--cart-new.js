let offerId = -1

let cart = []
let removes = []

let total = 0

$(function () {
    $('#dishes').on('click', '.add-item', addHandler)
    $('#cart-items').on('click', '.remove-item', removeHandler)
    $('#save-cart').on('click', function (e) {
        e.preventDefault()
        saveHandler()
    })
    $('#show-all-orders').on('click', showAllOrders)
    $('#cart-remark').on('keydown', enableSaveBtn)
})

function addHandler() {
    let e = $(this)
    let item = {
        dish_id: e.attr('data-dish-id'),
        name: e.attr('data-dish-name'),
        price: e.attr('data-dish-price'),
        fromServer: false
    }
    addToCart(item)
    enableSaveBtn()
}

function removeHandler() {
    let e = $(this)
    let item = {
        dish_id: e.attr('data-dish-id'),
        price: e.attr('data-dish-price'),
        fromServer: e.attr('data-from-server')
    }
    removeFromCart(item, e.parent())
    enableSaveBtn()
}

function addToCart(item) {
    cart.push(item)
    total += parseFloat(item.price)

    renderAddedCartDish(item)
    renderTotal(total)
}

function removeFromCart(item, element) {
    for (let i = 0; i < cart.length; i++) {
        if (cart[i].dish_id === item.dish_id) {
            cart.splice(i, 1)
            break
        }
    }
    if (item.fromServer) {
        removes.push(item)
    }
    total -= parseFloat(item.price)
    renderRemovedCartDish(element)
    renderTotal(total)
}

function saveHandler() {
    let remark = $('#cart-remark').val()
    saveCart(remark)
}

function saveCart(remark) {
    if (offerId === -1) {
        $.ajax({
            url: 'api/offers',
            type: 'get',
            success: function (res) {
                let offers = JSON.parse(res)
                offerId = offers[0].supplier_id
                continueSaveCart(remark)
            }
        })
    } else {
        continueSaveCart(remark)
    }
}

function continueSaveCart(remark) {
    animateSaveBtn()
    let calls = 0

    for (let item of cart) {
        if (item.fromServer === false) {
            $.ajax({
                url: 'api/add/' + offerId + '/' + item.dish_id,
                type: 'get',
                beforeSend: function () {
                    calls++
                },
                success: function () {
                    if (--calls === 0) {
                        finishSaveCart()
                    }
                }
            })
        }
    }
    for (let item of removes) {
        $.ajax({
            url: 'api/del/' + offerId + '/' + item.dish_id,
            type: 'get',
            beforeSend: function () {
                calls++
            },
            success: function () {
                if (--calls === 0) {
                    finishSaveCart()
                }
            }
        })
    }
    let r = remark === '' ? '@none' : remark

    $.ajax({
        url: 'api/remark/' + offerId + '/' + r,
        type: 'get',
        beforeSend: function () {
            calls++
        },
        success: function () {
            if (--calls === 0) {
                finishSaveCart()
            }
        }
    })
}

function finishSaveCart() {
    for (let i = 0; i < cart.length; i++) {
        cart[i].fromServer = true
    }
    disableSaveBtn()
}

function showAllOrders() {
    if (offerId === -1) {
        $.ajax({
            url: 'api/offers',
            type: 'get',
            success: function (res) {
                let offers = JSON.parse(res)
                offerId = offers[0].supplier_id
                continueShowAllOrders()
            }
        })
    } else {
        continueShowAllOrders()
    }
}

function continueShowAllOrders() {
    $.ajax({
        url: 'api/orders/' + offerId,
        type: 'get',
        success: function (res) {
            let orders = JSON.parse(res)
            renderAllOrders(orders)
        }
    })
}

function enableSaveBtn() {
    let markup = 'Bestellung speichern'
    $('#save-cart').attr('disabled', false).html(markup)
}

function animateSaveBtn() {
    let markup = '<span class="spinner-grow spinner-grow-sm" role="status" aria-hidden="true"></span>'
    $('#save-cart').attr('disabled', true).html(markup)
}

function disableSaveBtn() {
    window.setTimeout(function () {
        let markup = 'Bestellung gespeichert'
        $('#save-cart').attr('disabled', true).html(markup)
    }, 350)
}
