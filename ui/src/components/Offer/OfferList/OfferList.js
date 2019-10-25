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
import {OFFER_ROUTE} from "../../../util/Routes"
import {Link} from "@reach/router"
import {inject, observer} from "mobx-react"
import Table, {TableConfig} from "../../Base/Table"
import {runInAction} from "mobx"

function getOfferTableConfig(disabled) {
    return [
        new TableConfig("pl-0 py-4", (offer) => (
            <div className={`text-hand text-lg rounded-0 text-dark text-center px-1 py-2 ${disabled ? "bg-white" : "bg-gradient"}`}>
                {offer.restaurant.name}
            </div>)
        ),
        // TODO: format date
        new TableConfig("", (offer) => (
            <React.Fragment>
                <p className={`text-sm mb-1 ${disabled ? "text-muted" : ""}`}>Bestellung m&ouml;glich:</p>
                <p className={`text-md text-strong mb-0 ${disabled ? "text-muted" : ""}`}>{offer.valid_from}</p>
            </React.Fragment>)
        ),
        new TableConfig("", (offer) => (
            <React.Fragment>
                <p className={`text-sm mb-1 ${disabled ? "text-muted" : ""}`}>Angebot erstellt von:</p>
                <p className={`text-md text-strong mb-0 ${disabled ? "text-muted" : ""}`}>{offer.owner.name}</p>
            </React.Fragment>)
        ),
        new TableConfig("px-0 text-center", (offer) => (
            <Link to={OFFER_ROUTE + "/" + offer.id} className={`btn btn-light rounded-pill text-sm ${disabled ? "disabled-all" : ""}`}>
                Angebot ausw&auml;hlen
            </Link>)
        ),
    ]
}

/**
 * OfferListCurrent displays all offers which are currently possible.
 * It requests always the current list and doesn't cache it
 */
class OfferList extends React.Component {

    constructor(props) {
        super(props)
        this.offerLoader = props.offerLoader

        this.tableConfig = getOfferTableConfig(this.props.old)

        this.state = {offers: []}
        this.loadOffers()
    }

    loadOffers() {
        let loadOffersPromise = null
        if (this.props.old)
            loadOffersPromise = this.offerLoader.loadOld()
        else
            loadOffersPromise = this.offerLoader.loadActive()

        loadOffersPromise.then((offers) => {
            runInAction(() => {
                this.setState({offers})
            })
        })
    }

    render() {
        return (

                <Table config={this.tableConfig} rows={this.state.offers}/>

        )
    }
}

export default inject('offerLoader')(observer(OfferList))