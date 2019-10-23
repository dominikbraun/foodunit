/*
 * Copyright 2019 The FoodUnit Authors. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import Axios from "axios"
import UserLoader from "./UserLoader"
import RestaurantLoader from "./RestaurantLoader"

export default class OfferLoader {

    constructor(config) {
        this.config = config
        this.userLoader = new UserLoader(config)
        this.restaurantLoader = new RestaurantLoader(config)
    }

    /**
     * loadActive promises to loads all active Offers
     * @returns {Promise<AxiosResponse<T>>}
     */
    loadActive() {
        return Axios.get(this.config.apiUrl +  "/offers/active",
            {withCredentials: true}
        ).then((response) => {

            if (Array.isArray(response.data)) {
                let offers = response.data

                // save promises to wait for full finish of all offers
                let toAwait = []
                offers.forEach((offer) => {
                    // complete the offer (add owner and restaurant info)
                    toAwait.push(this.completeOffer(offer))
                })

                // wait for completing all offers and return all offers
                return Promise.all(toAwait).then(() => offers)
            }
        }).catch(function (error) {
            console.log(error)
        })
    }

    async completeOffer(incompleteOffer) {
        let owner = await this.userLoader.loadUser(incompleteOffer.owner.id)
        if (owner)
            incompleteOffer.owner = owner

        let restaurant = await this.restaurantLoader.loadRestaurant(incompleteOffer.restaurant.id)
        if (restaurant)
            incompleteOffer.restaurant = restaurant
    }
}