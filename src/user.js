let user = {
    isLoggedIn() {
        return sessionStorage.getItem('loggedin') != null;
    },
    login(email) {
        let success = false;

        $.ajax({
            url: 'http://localhost/fu-server',
            data: {
                action: 'login'
            },
            success: function() {
                alert("suc");
                success = true;
            }
        });
        return success;
    },
    confirmsMail() {
        return window.location.href.indexOf("confirm") > -1
    }
};
