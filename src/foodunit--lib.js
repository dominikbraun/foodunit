function req(endpoint, success) {
    $.ajax({
        url: 'api' + endpoint,
        type: 'get',
        success: success
    })
}
