let menu = []
let page = 0

let offer = 0

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
    $('#cats').on('click', '.select-cat', function () {
        page = $(this).attr('data-page')
        renderMenuPage(page)
    })
    getOfferId(build)
})

function getOfferId(buildFn) {
    $.ajax({
        url: 'api/offers',
        type: 'get',
        success: function (res) {
            let offers = JSON.parse(res)
            offer = offers[0].supplier_id

            buildFn()
        }
    })
}

function build() {
   buildSupplier(offer)
   buildCart(offer)
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

            window.setTimeout(function () {
                $('#loader-main').removeClass('d-flex').addClass('d-none')
                $('#loaded-content-main').removeClass('d-none')
            }, 450)
        }
    })
}

function buildCart(offerId) {
    $.ajax({
        url: 'api/cart-mono/' + offerId,
        type: 'get',
        success: function (res) {
            let cart = JSON.parse(res)

            renderEmail(cart.email)
            renderRemark(cart.remark)

            for (let item of cart.dishes) {
                item.fromServer = true
                addToCart(item)
            }
            window.setTimeout(function () {
                $('#loader-cart').removeClass('d-flex').addClass('d-none')
                $('#loaded-content-cart').removeClass('d-none')
            }, 450)
        }
    })
}
