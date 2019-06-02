function renderSupplierInfo(name, address, opened, phone) {
    $('#supplier').text(name)
    $('.supplier-name').text(name)
    $('.supplier-addr').text(address)
    $('.supplier-opened').text(opened)
    $('.supplier-phone').text(phone)
}

function renderCats(cats) {
    for (let i = 0; i < cats.length; i++) {
        let html = '<a class="text-dark mx-2 select-cat" data-page="' + i +'">' + cats[i] + '</a>'

        if (i < cats.length - 1) {
            html += '/'
        }
        $('#cats').append(html)
    }
}

function renderMenuPage(i) {
    $('#cat').text(menu[i].name)

    $('#bg-container').css('background-image', 'static/img/' + menu[i].img)

    $('#dishes').empty()

    for (let d of menu[i].dishes) {
        let html = '<div class="col-12 col-xl-6 px-5 py-3 add-item" data-dish-id="' + d.id + '" data-dish-name="' + d.name + '" data-dish-price="' + d.price + '"><div class="row"><div class="col-12 col-lg-9"><span class="text-md text-very-strong">' + d.name + '</span><br /><span class="text-xs"><i>' + d.desc + '</i></span></div><div class="col-12 col-lg-3 text-right">+ <span class="text-white text-strong price">' + Number.parseFloat(d.price).toFixed(2) + ' &euro;</span></div></div></div>'
        $('#dishes').append(html)
    }
}

function renderEmail(email) {
    $('#user-email').text(email)
}

function renderTotal(total) {
    $('#total').text(total.toFixed(2))
}

function renderCartDishes(dishes, isFromServer) {
    for (let i of dishes) {
        let html = '<li class="list-inline text-xs text-strong" data-dish-id="' + i.dish_id + '"><div class="row m-0"><div class="col-10 px-2 py-3">' + i.name + '</div><div class="col-2 py-3 remove-item text-center" data-dish-id="' + i.dish_id + '" data-from-server="' + isFromServer + '" data-dish-price="' + i.price + '">&#10005;</div></div></li>'
        $('#cart-items').append(html)
    }
}

function renderAddedCartDish(dish) {
    renderCartDishes([dish], false)
}

function renderRemovedCartDish(e) {
    $(e).remove()
}

function renderRemark(remark) {
    $('#cart-remark').val(remark)
}

function renderAllOrders(orders) {
    $('#all-orders-list').empty()

    for (let o of orders) {
        let html = '<div class="border-bottom my-4 pb-2"><p class="text-sm text-strong mb-4"><i class="fa fa-user mr-2"></i>'
        html += o.email
        html += '</p><ul class="p-0">'

        for (let d of o.positions) {
            html += '<li class="list-inline text-xs text-strong"><div class="row m-0"><div class="col-9 px-2 py-3">'
            html += d.name
            html += '</div><div class="col-3 py-3 text-right">'
            html += '0'//d.price
            html += ' &euro;</div></div></li>'
        }

        html += '<li class="list-inline text-xs text-strong"><div class="row m-0"><div class="col-7 px-2 py-3 text-muted"><i>Gesamt</i></div><div class="col-5 py-3 text-right text-muted"><i class="fas fa-shopping-cart mr-2"></i><i>'
        html += parseFloat(o.total).toFixed(2)
        html += ' &euro;</i></div></div></li>'
        html += '</ul><p class="text-xs text-dark text-strong"><i class="fas fa-comment mr-2"></i><i>&quot;'
        html += o.remark
        html += '&quot;</i></p></div>'

        $('#all-orders-list').append(html)
    }

    window.setTimeout(function () {
        $('#loader-all-orders').removeClass('d-flex').addClass('d-none')
        $('#loaded-content-all-orders').removeClass('d-none')
    }, 500)
}
