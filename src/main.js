let validSession = true;

let supplier = {
    name: "McDonald's Straubing",
    st: "Ittlingerstr. 119",
    ph: "+49 9424 3394948",
    opened: "07:00 - 01:00",
    cats: [
        {
            name: "Hamburger & Co.",
            img: "hamburgers.jpg",
            dishes: [
                { name: "Hamburger Royal Käse", price: 4.19 },
                { name: "Hamburger Royal TS", price: 4.19 },
                { name: "McRib", price: 3.99 },
                { name: "Pommes kl.", price: 1.79 },
            ]
        },
        {
            name: "The Signature Collection",
            img: "siganture-collection.jpg",
            dishes: [
                { name: "Hamburger Royal Käse", price: 4.19 },
                { name: "Hamburger Royal TS", price: 4.19 },
                { name: "McRib", price: 3.99 },
                { name: "Pommes kl.", price: 1.79 },
            ]
        }
    ]
};

new Vue({
    el: '#app',
    data: {
        loggedIn: validSession,
        supplier: supplier,
    }
});
