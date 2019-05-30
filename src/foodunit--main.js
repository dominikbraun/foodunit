let menu = []
let page = 0
let total = 0

$(function () {
    $('#prev').on('click', function () {
        if (page > 0) {
            renderMenuPage(--page)
        }
    })
    $('#next').on('click', function () {
        if (page < menu.length - 1) {
            renderMenuPage(++page)
        }
    })
    buildPage()
})

function buildPage() {
    $.ajax({
        url: 'api/offers',
        type: 'get',
        success: function (res) {
            let offers = JSON.parse(res)
            let supplierId = offers[0].supplier_id

            buildSupplier(supplierId)
            buildCart(offers[0].id)
        }
    })
}

function buildSupplier(supplierId) {
    $.ajax({
        url: 'api/supplier-mono/' + supplierId,
        type: 'get',
        success: function (res) {
            let s = JSON.parse(res)
            menu = s.menu
            renderSupplierInfo(s.name, s.address, s.mon, s.phone)
            renderCats(menu.map(c => c.name))
            renderMenuPage(page)
        }
    })
}

function renderSupplierInfo(name, address, opened, phone) {
    $('#supplier').text(name)
    $('#supplier-name').text(name)
    $('#supplier-addr').text(address)
    $('#supplier-opened').text(opened)
    $('#supplier-phone').text(phone)
}

function renderCats(cats) {
    for (let i = 0; i < cats.length; i++) {
        let html = '<a class="text-dark mx-2" href="">' + cats[i] + '</a>'

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
        let html = '<div class="col-12 col-xl-6 px-5 py-3 add-item"><div class="row"><div class="col-12 col-lg-9"><span class="text-md text-very-strong">' + d.name + '</span><br /><span class="text-xs"><i>' + d.desc + '</i></span></div><div class="col-12 col-lg-3 text-right">+ <span class="text-white text-strong price" href="">' + Number.parseFloat(d.price).toFixed(2) + ' &euro;</span></div></div></div>'
        $('#dishes').append(html)
    }
}

function buildCart(offerId) {
    $.ajax({
        url: 'api/cart-mono/' + offerId,
        type: 'get',
        success: function (res) {
            let cart = JSON.parse(res)
            renderCartDishes(cart.dishes)
            renderEmail(cart.email)
            renderRemark(cart.remark)
        }
    })
}

function renderEmail(email) {
    $('#user-email').text(email)
}

function renderCartDishes(dishes) {
    for (let i of dishes) {
        let html = '<li class="list-inline text-xs text-strong" data-dish-id="' + i.dish_id + '"><div class="row m-0"><div class="col-10 px-2 py-3">' + i.name + '</div><div class="col-2 py-3 remove-item text-center" data-dish-id="' + i.id + '">&#10005;</div></div></li>'
        $('#cart-items').append(html)
    }
}

function renderRemark(remark) {
    $('#cart-remark').val(remark)
}
