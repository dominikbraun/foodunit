<template>
    <div class="col-12 col-lg-3 col-xl-2 p-0 sidebar-left-lg bg-darker border-right">
        <div class="px-2 py-1 px-xl-4 py-xl-4">
            <div class="bg-primary-gradient text-white text-hand mx-auto text-center text-logo div-logo my-3 px-2 py-2 py-xl-3 rounded-0">Food<wbr>Unit&nbsp;</div>
        </div>
        <p class="text-light text-center text-strong text-sm mb-0 p-3">Angebot ausw&auml;hlen &amp; bestellen</p>
        <div class="nav flex-column nav-pills side-nav">
            <a class="nav-link text-md px-2 px-xl-3 py-3 rounded-0 active" href=""><i class="fas fa-pizza-slice ml-1 mr-3"></i>Aktuelle Angebote</a>
            <a class="nav-link text-md px-2 px-xl-3 py-3 rounded-0" href=""><i class="fas fa-share ml-1 mr-3"></i>Angebot erstellen</a>
            <a class="nav-link text-md px-2 px-xl-3 py-3 rounded-0" href=""><i class="fas fa-layer-group ml-1 mr-3"></i>Meine Angebote</a>
        </div>
        <p class="text-light text-center text-strong text-sm mb-0 p-3">Mein Konto</p>
        <div class="nav flex-column nav-pills side-nav">
            <div class="nav-link text-md text-success px-2 px-xl-3 py-3 rounded-0"><i class="fas fa-medal ml-1 mr-3"></i>Gesamt: 1200</div>
            <a class="nav-link text-md px-2 px-xl-3 py-3 rounded-0" href=""><i class="fas fa-user-circle ml-1 mr-3"></i>Konto verwalten</a>
            <a class="nav-link text-md px-2 px-xl-3 py-3 rounded-0" href=""><i class="fas fa-sign-out-alt ml-1 mr-3"></i>Ausloggen</a>
        </div>
    </div>
</template>

<script>
    export default {
        name: "SideNav"
    }
</script>

<style scoped></style>