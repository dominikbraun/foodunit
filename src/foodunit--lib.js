function req(endpoint, success) {
    console.log(endpoint)
    $.ajax({
        url: 'api' + endpoint,
        type: 'get',
        success: success
    })
}
