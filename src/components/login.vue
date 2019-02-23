<template>
    <div class="row h-100 justify-content-center align-items-center">
        <div class="bg-white-tr rounded-xl p-0 p-sm-3">
            <div class="bg-white shadow-lg rounded-lg" style="width: 21rem;">
                <div class="pt-5 pb-5">
                    <div class="bg-danger mx-auto text-center text-white text-theme h1 p-3 overflow-hidden rounded shadow-sm mt-3"
                         style="width: 8rem;">
                        Food&nbsp;<br/>
                        Unit&nbsp;
                    </div>
                </div>
                <form class="col-12">
                    <div class="form-group w-75 mx-auto my-0">
                        <input type="text" class="form-control border-0 shadow-sm bg-light text-md px-3 py-4 my-2"
                               id="formGroupExampleInput" placeholder="E-Mail-Adresse">
                    </div>
                    <div class="form-group mx-auto text-center my-0 pt-5 pb-5">
                        <button class="btn btn-success border-thick rounded text-bold text-sm px-4 py-2 mx-auto shadow-sm mb-3">
                            LOGIN ANFORDERN
                        </button>
                    </div>
                </form>
            </div>
        </div>
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
