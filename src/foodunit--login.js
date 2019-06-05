const loginBtn = '#login'
const emailInp = '#email'

$(function registerLoginHandlers() {
    $(loginBtn).on('click', function (e) {
        e.preventDefault()
        loginHandler()
    })
})

function loginHandler() {
    animateLoginBtn()
    let email = $(emailInp).val()

    $.ajax({
        url: 'api/sso/' + email,
        type: 'get',
        success: function (res) {
            let ok = JSON.parse(res) === true
            if (JSON.parse(res) === true) {
                disableLoginBtn('E-Mail gesendet')
            } else {
                disableLoginBtn('Es ist ein Fehler aufgetreten.')
            }
        }
    })
}

function animateLoginBtn() {
    let markup = '<span class="spinner-grow spinner-grow-sm" role="status" aria-hidden="true"></span>'
    $(loginBtn).attr('disabled', true).html(markup)
}

function disableLoginBtn(msg) {
    $(loginBtn).attr('disabled', true).html(msg)
}
