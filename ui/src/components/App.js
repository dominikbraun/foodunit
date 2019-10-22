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

import React from 'react'
import MainStore from "../stores/MainStore"
import OfferView from "./Offer/OfferView"
import LoginView from "./Auth/LoginView"
import {observer, Provider} from "mobx-react"
import LogoutView from "./Auth/LogoutView"
import {Router} from "@reach/router"
import CreateOfferView from "./Offer/CreateOfferView"
import {
    CREATE_OFFER_ROUTE,
    LOGIN_ROUTE,
    LOGOUT_ROUTE,
    MY_OFFERS_ROUTE,
    OFFERS_ROUTE,
    ORDERS_VIEW,
    RESTAURANT_VIEW
} from "../util/Routes"
import MyOffersView from "./Offer/MyOffersView"
import RestaurantView from "./Restaurant/RestaurantView"
import OrdersView from "./Restaurant/OrdersView"

class App extends React.Component {
    mainStore = new MainStore();

    constructor(props) {
        super(props)
        this.mainStore.init.bind(this.mainStore)()
    }

    render() {
        return (
            <Provider auth={this.mainStore.auth} foodUnit={this.mainStore.foodUnit}>
                <Router>
                    <LogoutView path={LOGOUT_ROUTE}/>
                    <LoginView default path={LOGIN_ROUTE}/>
                    <OfferView path={OFFERS_ROUTE}/>
                    <CreateOfferView path={CREATE_OFFER_ROUTE}/>
                    <MyOffersView path={MY_OFFERS_ROUTE}/>
                    <RestaurantView path={RESTAURANT_VIEW}/>
                    <OrdersView path={ORDERS_VIEW}/>
                </Router>
            </Provider>
        )
    }
}

export default observer(App)