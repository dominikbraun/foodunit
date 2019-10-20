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

import {OfferListCurrent, OfferListOld} from "./OfferList/OfferList";
import React from "react";
import SidebarLeft from "../SidebarLeft/SidebarLeft";
import SidebarRight from "../SidebarRight";
import LoggedIn from "../Auth/LoggedIn";
import {OFFERS_ROUTE} from "../../util/Routes";
import Footer from "../Footer";

export default function OfferView() {
    return (
        <LoggedIn>
            <div className="row m-0 h-100">
                <SidebarLeft currentActiveRoute={OFFERS_ROUTE}/>
                <div className="col-12 col-lg-6 col-xl-8 px-1 px-md-4 mx-auto">
                    <OfferListCurrent/>
                    <OfferListOld/>

                    <div className="mx-0 mx-xl-5 my-4 bg-white border rounded-0">
                        <div className="p-3 text-dark text-pmd">
                            <i className="fas fa-question-circle text-primary ml-1 mr-3"/>Mit einem Angebot bieten dir deine Kollegen an, Essen bei einem Restaurant zu bestellen und alle Bestellungen dort abzuholen.
                        </div>
                    </div>

                    <Footer/>
                </div>
                <SidebarRight/>
            </div>
        </LoggedIn>
    );
}