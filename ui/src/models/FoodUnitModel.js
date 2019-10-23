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

import {action, configure, decorate, observable} from 'mobx'
import Axios from "axios"

// enforce the strict mode for actions -> e.g. no state modifying inside of promise without action decorator https://www.leighhalliday.com/mobx-async-actions
configure({ enforceActions: "observed" })

export default class FoodUnitModel {

    offers = [];

    constructor(config) {
        this.config = config
    }

    loadOffers() {
        let that = this
        Axios.get(this.config.apiUrl +  "/offers/active",
            {withCredentials: true}
        ).then(function (response) {

            if (Array.isArray(response.data)) {
                that.setOffers(response.data)
            }
        }).catch(function (error) {
            console.log(error)
        })
    }

    setOffers(offers) {
        this.offers = offers
    }
}

decorate(FoodUnitModel, {
    offers: observable,
    setOffers: action,
})