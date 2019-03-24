let user = {
    isLoggedIn() {
        return sessionStorage.getItem('loggedin') != null;
    },
    login(email) {
        let success = true;

        $.ajax({
            url: 'http://localhost/fu-server',
            data: {
                action: 'login'
            },
            success: function () {
                success = true;
            },
        });
        return success;
    },
    confirmsMail() {
        return window.location.href.indexOf("confirm") > -1
    }
};
