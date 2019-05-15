let supplier = {}
let cats = []
let catIndex = 0
let dishes = []

let total = 0

$(function () {
    renderSupplier()
    renderCats()
    renderDishes(catIndex)
    renderCart()
})

function renderSupplier() {
    // Ajax request here
    supplier.name = 'McDonald\'s Straubing'
    supplier.addr = 'Ittlingerstr. 119'
    supplier.phone = '09421 3304948'
    supplier.opened = '07:00 - 01:00'

    $('#supplier').text(supplier.name)
    $('#supplier-name').text(supplier.name)
    $('#supplier-addr').text(supplier.addr)
    $('#supplier-opened').text(supplier.opened)
    $('#supplier-phone').text(supplier.phone)
}

function renderCats() {
    // Ajax request here
    cats = ['Alkoholfreie Getr\u00e4nke', 'Burger', 'Klassiker', 'Feinschmecker', 'Salate', 'Kinderteller']

    for (let i = 0; i < cats.length; i++) {
        let html = '<a class="text-dark mx-2" href="">' + cats[i] + '</a>'

        if (i < cats.length - 1) {
            html += '/'
        }
        $('#cats').append(html)
    }
}

function renderDishes(catIndex) {
    let cat = cats[catIndex]

    if (dishes[cat] === undefined) {
        // Ajax request here
        dishes[cat] = [
            {
                id: '0',
                name: 'Hamburger Royal K&auml;se',
                desc: 'Preis und Beschreibung folgen.',
                price: '2.2',
            },
            {
                id: '1',
                name: 'Hamburger',
                desc: 'Preis und Beschreibung folgen.',
                price: '2.8',
            },
            {
                id: '2',
                name: '9 Chicken McNuggets',
                desc: 'Preis und Beschreibung folgen.',
                price: '1.0',
            },
            {
                id: '3',
                name: 'Gro&szlig;e Pommes',
                desc: 'Preis und Beschreibung folgen.',
                price: '2.7',
            },
            {
                id: '4',
                name: 'Hamburger',
                desc: 'Preis und Beschreibung folgen.',
                price: '4.0',
            },
            {
                id: '5',
                name: 'Gro&szlig;e Pommes',
                desc: 'Preis und Beschreibung folgen.',
                price: '3.9',
            },
            {
                id: '6',
                name: 'Hamburger',
                desc: 'Preis und Beschreibung folgen.',
                price: '2.2',
            },
            {
                id: '7',
                name: 'Hamburger Royal K&auml;se',
                desc: 'Preis und Beschreibung folgen.',
                price: '3.9',
            },
            {
                id: '8',
                name: '9 Chicken McNuggets',
                desc: 'Preis und Beschreibung folgen.',
                price: '4.0',
            },
            {
                id: '9',
                name: '9 Chicken McNuggets',
                desc: 'Preis und Beschreibung folgen.',
                price: '3.2',
            },
            {
                id: '5',
                name: 'Gro&szlig;e Pommes',
                desc: 'Preis und Beschreibung folgen.',
                price: '3.9',
            },
            {
                id: '6',
                name: 'Hamburger',
                desc: 'Preis und Beschreibung folgen.',
                price: '2.2',
            },
            {
                id: '7',
                name: 'Hamburger Royal K&auml;se',
                desc: 'Preis und Beschreibung folgen.',
                price: '3.9',
            },
            {
                id: '8',
                name: '9 Chicken McNuggets',
                desc: 'Preis und Beschreibung folgen.',
                price: '4.0',
            },
            {
                id: '9',
                name: '9 Chicken McNuggets',
                desc: 'Preis und Beschreibung folgen.',
                price: '3.2',
            },
            {
                id: '5',
                name: 'Gro&szlig;e Pommes',
                desc: 'Preis und Beschreibung folgen.',
                price: '3.9',
            },
            {
                id: '6',
                name: 'Hamburger',
                desc: 'Preis und Beschreibung folgen.',
                price: '2.2',
            },
            {
                id: '7',
                name: 'Hamburger Royal K&auml;se',
                desc: 'Preis und Beschreibung folgen.',
                price: '3.9',
            },
            {
                id: '8',
                name: '9 Chicken McNuggets',
                desc: 'Preis und Beschreibung folgen.',
                price: '4.0',
            },
            {
                id: '9',
                name: '9 Chicken McNuggets',
                desc: 'Preis und Beschreibung folgen.',
                price: '3.2',
            }
        ]
    }
    $('#cat').text(cat)

    for (let d of dishes[cat]) {
        let html = '<div class="col-12 col-xl-6 px-5 py-3"><div class="row"><div class="col-12 col-lg-9"><span class="text-md text-very-strong">' + d.name + '</span><br /><span class="text-xs"><i>' + d.desc + '</i></span></div><div class="col-12 col-lg-3 text-right">+ <a class="text-white text-strong" href="">' + Number.parseFloat(d.price).toFixed(2) + ' &euro;</a></div></div></div>'
        $('#dishes').append(html)
    }
}

function renderCart() {
    // Ajax request here
    let cart = {
        items: [
            {
                id: '7',
                name: 'Hamburger Royal K&auml;se',
                desc: 'Preis und Beschreibung folgen.',
                price: '3.9',
            },
            {
                id: '8',
                name: '9 Chicken McNuggets',
                desc: 'Preis und Beschreibung folgen.',
                price: '4.0',
            },
            {
                id: '9',
                name: '12 Chicken McNuggets',
                desc: 'Preis und Beschreibung folgen.',
                price: '3.2',
            }
        ],
        remark: 'Bitte mit Ketchup',
    }

    for (let i of cart.items) {
        let html = '<li class="list-inline text-xs text-strong" data-dish-id="' + i.id + '"><div class="row m-0"><div class="col-10 px-2 py-3">' + i.name + '</div><div class="col-2 py-3 remove-item text-center" data-dish-id="' + i.id + '">&#10005;</div></div></li>'
        $('#cart-items').append(html)
    }
    $('#cart-remark').val(cart.remark)
    $('#total').text(Number.parseFloat(total).toFixed(2))
}