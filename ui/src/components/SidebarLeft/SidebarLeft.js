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

import React from "react"
import AccountSection from "./AccountSection"
import {Link} from "@reach/router"
import {CREATE_OFFER_ROUTE, MY_OFFERS_ROUTE, OFFERS_ROUTE} from "../../util/Routes"

/**
 * SidebarLeft contains navigation links
 * @param currentActiveRoute the route which is currently active. It will be highlighted.
 * @returns {React.Component}
 * @constructor
 */
export default function SidebarLeft({currentActiveRoute}) {

    let navigationLinks = [
        [OFFERS_ROUTE, "Aktuelle Angebote", "fa-pizza-slice"],
        [CREATE_OFFER_ROUTE, "Angebot erstellen", "fa-share"],
        [MY_OFFERS_ROUTE, "Meine Angebote", "fa-layer-group"]
    ]

    navigationLinks = navigationLinks.map((route) => {
        let classes = "nav-link text-md px-2 px-xl-3 py-3 rounded-0"

        if (route[0] === currentActiveRoute) {
            classes += " active"
        }

        return <Link to={route[0]} className={classes}><i className={`fas ${route[2]} ml-1 mr-3`}/>{route[1]}</Link>
    })


    return (
        <div className="col-12 col-lg-3 col-xl-2 p-0 sidebar-left-lg bg-darker border-right">
            <div className="px-2 py-1 px-xl-4 py-xl-4">
                <div className="bg-primary-gradient text-white text-hand mx-auto text-center text-logo div-logo my-3 px-2 py-2 py-xl-3 rounded-0">Food<wbr/>Unit&nbsp;</div>
            </div>
            <p className="text-light text-center text-strong text-sm mb-0 p-3">Angebot ausw&auml;hlen &amp; bestellen</p>
            <div className="nav flex-column nav-pills side-nav">
                {navigationLinks}
            </div>
            <AccountSection/>
        </div>
    )
}