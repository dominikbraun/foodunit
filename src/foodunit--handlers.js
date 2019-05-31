const loginBtn = '#login'
const emailInp = '#email'
const saveBtn = '#save-cart'

$(register)

function register() {
    $(loginBtn).on('click', function (e) {
        e.preventDefault()
        login()
    })
    $(saveBtn).on('click', saveCart)
}

function login() {
    $(loginBtn).attr('disabled', true)
    $(loginBtn).html('<span class="spinner-grow spinner-grow-sm" role="status" aria-hidden="true"></span>')

    let email = $(emailInp).val()
    let endpoint = '/sso/' + email

    req(endpoint, function (res) {
        let ok = JSON.parse(res) === true
        if (ok) {
            $(loginBtn).html('E-Mail gesendet')
        } else {
            $(loginBtn).html('Es ist ein Fehler aufgetreten.')
        }
    })
}

function saveCart() {}

function req(endpoint, success) {
    console.log(endpoint)
    $.ajax({
        url: 'api' + endpoint,
        type: 'get',
        success: success
    })
}
