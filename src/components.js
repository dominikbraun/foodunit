
Vue.component('foodunit-main', {
    template: `<div class="row h-100 m-0">
                    <div class="col-12 col-lg-2 h-100 m-0 p-0 d-none d-sm-block"></div>
                    <div class="col-12 col-lg-8 h-100 m-0 p-0">
                        <div class="w-100">
                            <div class="py-4">
                                <div class="row m-0">
                                    <div class="col-12 col-xl-2 pt-4 pl-0">
                                        <div class="bg-danger mx-auto ml-xl-0 text-center text-white text-theme h1 p-3 rounded shadow mb-0" style="width: 9rem;">
                                            Food&nbsp;<br />Unit&nbsp;
                                        </div>
                                    </div>
                                    <div class="col-12 col-xl-8">
                                        <h1 class="text-dark text-theme text-xxl text-center m-0 py-5">Angebot des Tages</h1>
                                    </div>
                                    <div class="col-12 col-xl-2 m-0 pt-4 px-0">
                                    </div>
                                </div>
                            </div>
                            <template v-for="c in supplier.cats">
                                <div class="bg-dark shadow-lg menu-area mb-5 rounded-lg shadow" :style="{ backgroundImage: 'url(\\'static/img/categories/' + c.img + '\\')' }">
                                    <div class="bg-dark-tr rounded-lg">
                                        <div class="py-5">
                                            <h2 class="text-white text-menu text-center text-xl">{{ c.name }}</h2>
                                        </div>
                                        <div class="row m-0 pb-5">
                                            <template v-for="dish in c.dishes">
                                                <div class="col-12 col-xl-6 text-white">
                                                    <div class="row m-0 dish-row p-3 rounded">
                                                        <div class="col-8 dish-detail">
                                                            <strong class="text-md">{{ dish.name }}</strong><br />
                                                            <small><i>{{ dish.desc }}</i></small>
                                                        </div>
                                                        <div class="col-4 text-right">
                                                            <strong class="text-md">{{ dish.price }} &euro;</strong><br />
                                                            <small>+ <a class="text-white">hinzufügen</a></small>
                                                        </div>
                                                    </div>
                                                </div>
                                            </template>
                                        </div>
                                    </div>
                                </div>
                            </template>
                            <div class="bg-white border border-light menu-area mb-5 rounded-lg">
                                <div class="pt-5 text-center">
                                    <h2 class="bg-danger text-white text-theme px-4 py-3 rounded shadow mx-0 mx-sm-3 my-1 my-sm-3 d-inline-block">
                                        {{ supplier.name }}
                                    </h2>
                                </div>
                                <div class="pt-4 pb-5 text-center">
                                    <p class="text-md2">
                                        <i class="fas fa-map-marker-alt pr-2"></i><span class="text-bold">{{ supplier.st }}<br />94315 Straubing</span>
                                    </p>
                                    <p class="text-md2">
                                        <i class="fas fa-store pr-2"></i><span class="text-bold text-success">{{ supplier.opened }} Uhr</span>
                                    </p>
                                    <p class="text-md">
                                        <i class="fas fa-phone pr-2"></i><span class="text-bold">{{ supplier.ph }}</span>
                                    </p>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="col-12 col-lg-2 h-100 m-0 p-2 p-lg-4">
                        <div class="bg-white border border-light rounded mb-4">
                            <p class="text-sm text-bold text-right mb-0 px-1 py-3">
                                <a class="text-dark text-sm text-bold" data-toggle="collapse" href="#cart" role="button" aria-expanded="false" aria-controls="cart">
                                    <i class="fas fa-shopping-cart pr-2"></i>8.89 &euro;<i class="fas fa-chevron-down px-3"></i>
                                </a>
                            </p>
                            <div class="collapse text-sm px-3" id="cart">
                                <ul class="p-0" id="cart-items">
                                    <li class="list-inline py-1"><a role="button" class="text-dark"><i class="fas fa-times pr-2"></i></a>Hamburger Royal TS</li>
                                    <li class="list-inline py-1"><a role="button" class="text-dark"><i class="fas fa-times pr-2"></i></a>Pommes kl.</li>
                                    <li class="list-inline py-1"><a role="button" class="text-dark"><i class="fas fa-times pr-2"></i></a>Hamburger Royal TS</li>
                                    <li class="list-inline py-1"><a role="button" class="text-dark"><i class="fas fa-times pr-2"></i></a>Hamburger Royal TS</li>
                                    <li class="list-inline py-1"><a role="button" class="text-dark"><i class="fas fa-times pr-2"></i></a>McRib</li>
                                    <li class="list-inline py-1"><a role="button" class="text-dark"><i class="fas fa-times pr-2"></i></a>McRib</li>
                                    <li class="list-inline py-1"><a role="button" class="text-dark"><i class="fas fa-times pr-2"></i></a>Coca-Cola</li>
                                </ul>
                                <input type="text" class="form-control bg-light border-light rounded mb-3 py-4 text-sm" id="remark" placeholder="Extrawunsch" spellcheck="false" />
                                <button class="btn btn-success mb-3 px-3 py-2 rounded text-sm text-bold w-100" id="save-cart">
                                    Bestellung speichern
                                </button>
                            </div>
                        </div>
                    </div>
                </div>
                <footer class="text-sm text-dark text-center text-bold mb-5">
                    &copy; 2019 <a class="text-dark" href="https://github.com/dominikbraun" target="_blank">github.com/dominikbraun</a>
                </footer>`,
    data() {
        return supplier;
    }
})

Vue.component('foodunit-login', {
    template: `<div class="row h-100 m-0">
                    <div class="col-12 col-sm-3 h-100 m-0 p-0 d-none d-sm-block"></div>
                    <div class="col-12 col-sm-6 h-100 m-0 p-0 justify-content-center align-items-center d-flex">
                        <div class="p-4 p-sm-0 w-100" id="login-area">
                            <div class="bg-danger mx-auto text-center text-white text-theme h1 p-3 rounded shadow mb-0" style="width: 9rem;">
                            Food&nbsp;<br />Unit&nbsp;
                            </div>
                            <h1 class="text-dark text-theme text-xl text-center m-0 py-5">Bevor du bestellst...</h1>
                            <div class="">
                                <form autocomplete="off">
                                    <div class="form-group">
                                        <input type="text" class="form-control bg-white border-light rounded px-3 py-4" id="user-name" placeholder="Dein Name" spellcheck="false" />
                                    </div>
                                    <div class="form-group">
                                        <input type="text" class="form-control bg-white border-light rounded px-3 py-4" id="user-email" placeholder="Deine E-Mail" spellcheck="false" />
                                    </div>
                                    <div class="my-5">
                                        <button class="btn btn-success px-3 py-2 rounded w-100">
                                            Login anfordern
                                        </button>
                                    </div>
                                </form>
                            </div>
                        </div>
                    </div>
                    <div class="col-12 col-sm-3 h-100 m-0 p-0 d-none d-sm-block"></div>
                </div>`,
    data() {
        return 0
    }
})