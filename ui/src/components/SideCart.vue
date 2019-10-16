<template>
    <div class="col-12 col-lg-3 col-xl-2 p-0 sidebar-right-lg bg-white border-left">
        <h5 class="text-center border-bottom mb-0 px-2 py-4">Einkaufskorb</h5>

        <p v-if="!items.length" class="text-center text-strong mb-0 pt-5 pb-3 text-sm">Dein Einkaufskorb ist leer.</p>

        <div v-for="item in items" :key="item" class="px-4 py-4 border-bottom">
            <div class="row m-0">
                <div class="col-10 px-0 pt-2 pb-3 text-md text-strong">{{ item.Name }}</div>
                <div class="col-2 px-0">
                    <button type="button" class="btn btn-white text-md rounded-circle" aria-label="Close">&times;</button>
                </div>
            </div>
            <div class="text-pmd"><i class="fas fa-wrench mr-2"></i><a class="text-muted" href="">konfigurieren</a></div>
            <div class="my-3 bg-white d-none">
                <form class="mb-0">
                    <div class="py-2">
                        <label class="text-md py-1" for="chara0">Pizza-Durchmesser</label>
                        <select class="custom-select text-md bg-light border-0 rounded-0" id="chara0">
                            <option selected>26 cm</option>
                            <option>28 cm</option>
                            <option>32 cm</option>
                        </select>
                    </div>
                    <div class="py-2">
                        <h6 class="text-md text-normal py-1">Extra-Zutaten</h6>
                        <div class="">
                            <input class="" type="checkbox" value="" id="chara1-0">
                            <label class="mb-0 ml-2 text-md" for="chara1-0">Mozzarella</label>
                        </div>
                        <div class="">
                            <input class="" type="checkbox" value="" id="chara1-1">
                            <label class="mb-0 ml-2 text-md" for="chara1-1">Schinken</label>
                        </div>
                        <div class="">
                            <input class="" type="checkbox" value="" id="chara1-2">
                            <label class="mb-0 ml-2 text-md" for="chara1-2">Pepperoni</label>
                        </div>
                    </div>
                </form>
            </div>
            <div class="text-pmd"><i class="fas fa-pen mr-2"></i><a class="text-muted" href="">Anmerkung bearbeiten</a></div>
            <div class="text-md text-strong pt-3 pb-2">Preis: {{ item.price }} &euro;</div>
        </div>

        <div class="px-4 py-4">
            <div class="row m-0">
                <div class="col-5 text-md text-strong text-muted px-0 py-2">Gesamt: {{ total }} &euro;</div>
                <div class="col-7 text-right p-0">
                    <button class="btn btn-light rounded-pill text-xs" :disabled="!items.length"><i class="fab fa-paypal mr-2"></i>direkt zu PayPal</button>
                </div>
            </div>
        </div>
        <div class="px-4 pt-0 pb-4">
            <button class="btn btn-success w-100 text-md rounded-0 text-strong py-2" :disabled="!items.length">Einkaufskorb speichern</button>
        </div>
    </div>
</template>

<script>
    export default {
        name: "SideCart",
        props: {
            items: Array,
            total: Number
        }
    }
</script>

<style scoped></style>