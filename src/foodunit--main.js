let cats = []
let catIndex = 0
let total = 0

$(function () {
    $('#prev').on('click', function () {
        if (catIndex > 0) {
            renderMenuPage(--catIndex)
        }
    })
    $('#next').on('click', function () {
        if (catIndex < cats.length - 1) {
            renderMenuPage(++catIndex)
        }
    })
    renderPage()
})

function renderPage() {
    $.ajax({
        url: 'api/offers',
        type: 'get',
        success: function (res) {
            let offers = JSON.parse(res)
            let supplierId = offers[0].supplier_id
            renderSupplier(supplierId)
            renderCart(offers[0].id)
        }
    })
}

function renderSupplier(supplierId) {
    $.ajax({
        url: 'api/supplier/' + supplierId,
        type: 'get',
        success: function (res) {
            let supplier = JSON.parse(res)
            $('#supplier').text(supplier.name)
            $('#supplier-name').text(supplier.name)
            $('#supplier-addr').text(supplier.address)
            $('#supplier-opened').text(supplier.mon)
            $('#supplier-phone').text(supplier.phone)

            renderMenu(supplierId)
        }
    })
}

function renderMenu(supplierId) {
    $.ajax({
        url: 'api/menu/' + supplierId,
        type: 'get',
        success: function (res) {
            cats = JSON.parse(res)

            for (let i = 0; i < cats.length; i++) {
                let html = '<a class="text-dark mx-2" href="">' + cats[i].name + '</a>'

                if (i < cats.length - 1) {
                    html += '/'
                }
                $('#cats').append(html)
            }

            renderMenuPage(catIndex)
        }
    })
}

function renderMenuPage(i) {
    $('#cat').text(cats[i].name)

    $('#bg-container').css('background-image', 'static/img/' + cats[i].img)

    $('#dishes').empty()

    for (let d of cats[i].dishes) {
        let html = '<div class="col-12 col-xl-6 px-5 py-3 add-item"><div class="row"><div class="col-12 col-lg-9"><span class="text-md text-very-strong">' + d.name + '</span><br /><span class="text-xs"><i>' + d.desc + '</i></span></div><div class="col-12 col-lg-3 text-right">+ <span class="text-white text-strong price" href="">' + Number.parseFloat(d.price).toFixed(2) + ' &euro;</span></div></div></div>'
        $('#dishes').append(html)
    }
}

function renderCart(offerId) {
    $.ajax({
        url: 'api/user-order/' + offerId,
        type: 'get',
        success: function (res) {
            let cart = {}
            cart.items = JSON.parse(res)

            for (let i of cart.items) {
                let html = '<li class="list-inline text-xs text-strong" data-dish-id="' + i.dish_id + '"><div class="row m-0"><div class="col-10 px-2 py-3">' + i.name + '</div><div class="col-2 py-3 remove-item text-center" data-dish-id="' + i.id + '">&#10005;</div></div></li>'
                $('#cart-items').append(html)
            }
            $('#cart-remark').val(cart.remark)
            $('#total').text(Number.parseFloat(total).toFixed(2))
        }
    })
    $.ajax({
        url: 'api/email',
        type: 'get',
        success: function (res) {
            let email = JSON.parse(res)
            $('#user-email').text(email)
        }
    })
}
