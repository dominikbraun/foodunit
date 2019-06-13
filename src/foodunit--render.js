function renderSupplierInfo(name, address, opened, phone) {
    $('#supplier').text(name)
    $('.supplier-name').text(name)
    $('.supplier-addr').text(address)
    $('.supplier-opened').text(opened)
    $('.supplier-phone').text(phone)
}

function renderCats(cats) {
    for (let i = 0; i < cats.length; i++) {
        let html = '<a class="text-dark mx-2 py-1 select-cat" data-page="' + i +'">' + cats[i] + '</a>'

        if (i < cats.length - 1) {
            html += '/'
        }
        $('#cats').append(html)
    }
}

function renderMenuPage(i) {
    $('#menu-page').fadeOut(110, function () {
        $('#cat').text(menu[i].name)
        $('.select-cat').css('border-bottom', 'none')
        $('[data-page=' + i + ']').css('border-bottom', '2px solid #212529')

        $('#bg-container').css('background-image', 'url(static/img/' + menu[i].img + ')')

        $('#dishes').empty()

        for (let d of menu[i].dishes) {
            let html = '<div class="col-12 col-xl-6 px-5 py-3 add-item" data-dish-id="' + d.id + '" data-dish-name="' + d.name + '" data-dish-price="' + d.price + '"><div class="row"><div class="col-12 col-lg-9"><span class="text-md text-very-strong">' + d.name + '</span><br /><span class="text-xs"><i>' + d.desc + '</i></span></div><div class="col-12 col-lg-3 text-right">+ <span class="text-white text-strong price">' + Number.parseFloat(d.price).toFixed(2) + ' &euro;</span></div></div></div>'
            $('#dishes').append(html)
        }
        $(this).fadeIn(160)
    })
}

function renderEmail(email) {
    $('#user-email').text(email)
}

function renderTotal(total) {
    if (total < 0 || total === '') {
        total = 0
    }
    $('#total').fadeOut(160, function () {
        $(this).text(parseFloat(total).toFixed(2))
        $(this).fadeIn(100)
    })
}

function renderCartDishes(dishes) {
    for (let i of dishes) {
        let html = '<li class="list-inline text-xs text-strong new-cart-item" data-dish-id="' + i.dish_id + '"><div class="row m-0"><div class="col-8 col-lg-7 col-xl-7 px-2 py-3">' + i.name + '</div><div class="col-2 col-lg-3 col-xl-3 px-1 py-3">' + parseFloat(i.price).toFixed(2) + ' &euro;</div><div class="col-2 col-lg-2 col-xl-2 py-3 remove-item text-center" data-dish-id="' + i.dish_id + '" data-from-server="' + i.fromServer + '" data-dish-price="' + i.price + '">&#10005;</div></div></li>'
        $('#cart-items').append(html)
    }
}

function renderAddedCartDish(dish) {
    renderCartDishes([dish])
}

function renderRemovedCartDish(e) {
    $(e).remove()
}

function renderRemark(remark) {
    $('#cart-remark').val(remark)
}

function renderAllOrders(orders, total) {
    $('#all-orders-list').empty()

    for (let o of orders) {
        let html = '<div class="border-bottom my-4 pb-2"><p class="text-sm text-strong mb-4" data-opt><i class="fa fa-user mr-2"></i>'
        html += o.email
        html += '</p><ul class="p-0">'

        for (let d of o.positions) {
            html += '<li class="list-inline text-xs text-strong"><div class="row m-0"><div class="col-9 px-2 py-3">'
            html += d.name
            html += '</div><div class="col-3 py-3 text-right" data-opt>'
            html += parseFloat(d.price).toFixed(2)
            html += ' &euro;</div></div></li>'
        }

        html += '<li class="list-inline text-xs text-strong" data-opt><div class="row m-0"><div class="col-7 px-2 py-3 text-muted"><i>Gesamt</i></div><div class="col-5 py-3 text-right text-muted"><i class="fas fa-shopping-cart mr-2"></i><i>'
        html += parseFloat(o.total).toFixed(2)
        html += ' &euro;</i></div></div></li>'
        html += '</ul>'

        if (o.remark !== '') {
            html += '<p class="text-xs text-dark text-strong"><i class="fas fa-comment mr-2" data-opt></i><i>&quot;'
            html += o.remark
            html += '&quot;</i></p>'
        }
        html += '</div>'

        $('#all-orders-list').append(html)
    }

    $('#all-orders-total').text(parseFloat(total).toFixed(2))

    window.setTimeout(function () {
        $('#loader-all-orders').removeClass('d-flex').addClass('d-none')
        $('#loaded-content-all-orders').removeClass('d-none')
    }, 500)
}
