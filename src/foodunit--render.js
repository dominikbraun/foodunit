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

function renderEmail(email) {
    $('#user-email').text(email)
}

function renderTotal(total) {
    $('#total').text(total)
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
