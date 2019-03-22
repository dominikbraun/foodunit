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
                { name: "Hamburger Royal Käse", desc: "Saftiges Rindfleisch mit ausgewählten Zutaten", price: 4.19 },
                { name: "Hamburger Royal TS", desc: "Saftiges Rindfleisch mit ausgewählten Zutaten", price: 4.19 },
                { name: "McRib", desc: "Saftiges Rindfleisch mit ausgewählten Zutaten", price: 3.99 },
                { name: "Pommes kl.", desc: "Saftiges Rindfleisch mit ausgewählten Zutaten mit Rohkostsalat und gebratenem Speck mit frischem Spiegelei", price: 1.79 },
                { name: "McRib", desc: "Saftiges Rindfleisch mit ausgewählten Zutaten", price: 3.99 },
                { name: "Hamburger Royal TS", desc: "Saftiges Rindfleisch mit ausgewählten Zutaten", price: 4.19 },
            ]
        },
        {
            name: "The Signature Collection",
            img: "siganture-collection.jpg",
            dishes: [
                { name: "Hamburger Royal Käse", desc: "Saftiges Rindfleisch mit ausgewählten Zutaten", price: 4.19 },
                { name: "Hamburger Royal TS", desc: "Saftiges Rindfleisch mit ausgewählten Zutaten", price: 4.19 },
                { name: "McRib", desc: "Saftiges Rindfleisch mit ausgewählten Zutaten", price: 3.99 },
                { name: "Pommes kl.", desc: "Saftiges Rindfleisch mit ausgewählten Zutaten", price: 1.79 },
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
