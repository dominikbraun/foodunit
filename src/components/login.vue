<template>
    <div class="row h-100">
        <div class="col-12 col-sm-4 h-100 d-none d-sm-block"></div>
        <div class="col-12 col-sm-4 h-100 justify-content-center align-items-center d-flex">
            <div class="bg-white text-white p-5 rounded-lg shadow-lg">
                <div class="my-3">
                    <div class="bg-danger mx-auto text-center text-theme h1 p-3 rounded shadow mb-0" style="width: 8rem;">
                        Food&nbsp;<br />
                        Unit&nbsp;
                    </div>
                </div>
                <div class="my-5">
                    <form autocomplete="off">
                        <div class="form-group">
                            <input type="text" class="form-control border-0 bg-light rounded shadow-sm text-md px-3 py-4" id="user-email" placeholder="Ihre E-Mail-Adresse" />
                        </div>
                    </form>
                </div>
                <div class="mt-3 text-center">
                    <button class="btn btn-success border-thick rounded text-bold text-sm px-4 py-2 mx-auto shadow mb-3">
                        LOGIN ANFORDERN
                    </button>
                </div>
            </div>
        </div>
        <div class="col-12 col-sm-4 h-100 d-none d-sm-block"></div>
    </div>
</template>

<script>
export default {
  name: 'login',
  data () {
    return {}
  }
}
</script>
