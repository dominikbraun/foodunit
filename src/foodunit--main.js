let menu = []
let page = 0

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
    build()
})

function build() {
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
