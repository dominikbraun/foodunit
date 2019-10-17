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

import React from 'react';
import MainStore from "../stores/MainStore";
import SidebarLeft from "./SidebarLeft";
import SidebarRight from "./SidebarRight";
import OfferView from "./Offer/OfferView";
import LoginView from "./Auth/LoginView";
import {observer, Provider} from "mobx-react";

class App extends React.Component {
    mainStore = new MainStore();

    constructor(props) {
        super(props);
        this.mainStore.init.bind(this.mainStore)()
    }

    render() {
        console.log(this.mainStore)

        let view;

        if (!this.mainStore.auth.loggedIn) {
            view = <LoginView/>;

        } else if (this.mainStore.auth.loggedIn) {
            view =  <div className="row m-0 h-100">
                        <SidebarLeft/>

                        <OfferView/>

                        <SidebarRight/>
                    </div>
        } else {
            view = <div>ERROR</div>;
        }


        return (
            <Provider auth={this.mainStore.auth} foodUnit={this.mainStore.foodUnit}>
                {view}
            </Provider>
        );


    }
}

export default observer(App);