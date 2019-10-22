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
import {RESTAURANT_VIEW} from "../../../util/Routes"
import {Link} from "@reach/router"
import {inject, observer} from "mobx-react"

class OfferListCurrent extends React.Component {

    constructor(props) {
        super(props)
        this.foodUnit = props.foodUnit
        this.foodUnit.loadOffers()
    }

    render() {
        if (this.foodUnit.offers.length === 0) {
            return <div>No data</div>
        }
        console.log(this.foodUnit.offers[0].id)

        return (
            <div className="mx-0 mx-xl-5 my-4 px-5 py-3 bg-white border rounded-0">
                <h6 className="text-dark text-strong px-0 py-3">Aktuelle Angebote</h6>

                <table className="table table-responsive-xl mb-0 text-center">
                    <tr>
                        <td className="align-middle pl-0 pr-3 py-4">
                            <div className="text-hand text-lg bg-gradient rounded-0 text-dark text-center px-1 py-2">Pizzeria Venezia</div>
                        </td>
                        <td className="align-middle py-3">
                            <p className="text-sm mb-1">Bestellung m&ouml;glich:</p>
                            <p className="text-md text-strong mb-0">Di, 11:15 &ndash; 11:45 Uhr</p>
                        </td>
                        <td className="align-middle py-3">
                            <p className="text-sm mb-1">Angebot erstellt von:</p>
                            <p className="text-md text-strong mb-0">Dominik Braun</p>
                        </td>
                        <td className="align-middle px-0 py-3 text-center">
                            <Link to={RESTAURANT_VIEW} className="btn btn-light rounded-pill text-sm">Angebot ausw&auml;hlen</Link>
                        </td>
                    </tr>

                    <tr>
                        <td className="align-middle pl-0 pr-3 py-4">
                            <div className="text-hand text-lg bg-gradient rounded-0 text-dark text-center px-1 py-2">Imbiss Media</div>
                        </td>
                        <td className="align-middle py-3">
                            <p className="text-sm mb-1">Bestellung m&ouml;glich:</p>
                            <p className="text-md text-strong mb-0">Di, 11:15 &ndash; 11:45 Uhr</p>
                        </td>
                        <td className="align-middle py-3">
                            <p className="text-sm mb-1">Angebot erstellt von:</p>
                            <p className="text-md text-strong mb-0">Karsten Wirler</p>
                        </td>
                        <td className="align-middle px-0 py-3 text-center">
                            <Link to={RESTAURANT_VIEW} className="btn btn-light rounded-pill text-sm">Angebot ausw&auml;hlen</Link>
                        </td>
                    </tr>
                    <tr>
                        <td className="align-middle pl-0 pr-3 py-4">
                            <div className="text-hand text-lg bg-gradient rounded-0 text-dark text-center px-1 py-2">Restaurant</div>
                        </td>
                        <td className="align-middle py-3">
                            <p className="text-sm mb-1">Bestellung m&ouml;glich:</p>
                            <p className="text-md text-strong mb-0">Di, 11:15 &ndash; 11:45 Uhr</p>
                        </td>
                        <td className="align-middle py-3">
                            <p className="text-sm mb-1">Angebot erstellt von:</p>
                            <p className="text-md text-strong mb-0">Adrian Brasin</p>
                        </td>
                        <td className="align-middle px-0 py-3 text-center">
                            <Link to={RESTAURANT_VIEW} className="btn btn-light rounded-pill text-sm">Angebot ausw&auml;hlen</Link>
                        </td>
                    </tr>

                </table>

                <div className="border-top-light text-right pt-3">
                    <a href="create-offer.html" className="btn btn-link rounded-pill text-sm"><i className="fas fa-share mr-2"/>Angebot erstellen</a>
                </div>
            </div>
        )
    }
}

export default inject('foodUnit')(observer(OfferListCurrent))

export function OfferListOld() {
    return <div className="mx-0 mx-xl-5 my-4 px-5 py-3 bg-white border rounded-0">
        <h6 className="text-dark text-strong px-0 py-3">Abgelaufene Angebote</h6>

        <table className="table table-responsive-xl mb-0 text-center text-muted">
            <tr>
                <td className="align-middle pl-0 pr-3 py-4">
                    <div className="text-hand text-lg bg-white rounded-0 text-dark text-center px-1 py-2">Pizzeria Venezia</div>
                </td>
                <td className="align-middle py-3">
                    <p className="text-sm mb-1">Bestellung m&ouml;glich:</p>
                    <p className="text-md text-strong mb-0">Di, 11:15 &ndash; 11:45 Uhr</p>
                </td>
                <td className="align-middle py-3">
                    <p className="text-sm mb-1">Angebot erstellt von:</p>
                    <p className="text-md text-strong mb-0">Angela B&ouml;hm</p>
                </td>
                <td className="align-middle px-0 py-3 text-center">
                    <button className="btn btn-light rounded-pill text-sm" disabled>Angebot ausw&auml;hlen</button>
                </td>
            </tr>

            <tr>
                <td className="align-middle pl-0 pr-3 py-4">
                    <div className="text-hand text-lg bg-white rounded-0 text-dark text-center px-1 py-2">Flying Pizza</div>
                </td>
                <td className="align-middle py-3">
                    <p className="text-sm mb-1">Bestellung m&ouml;glich:</p>
                    <p className="text-md text-strong mb-0">Di, 11:15 &ndash; 11:45 Uhr</p>
                </td>
                <td className="align-middle py-3">
                    <p className="text-sm mb-1">Angebot erstellt von:</p>
                    <p className="text-md text-strong mb-0">Sebastian M&uuml;ller</p>
                </td>
                <td className="align-middle px-0 py-3 text-center">
                    <button className="btn btn-light rounded-pill text-sm" disabled>Angebot ausw&auml;hlen</button>
                </td>
            </tr>
            <tr>
                <td className="align-middle pl-0 pr-3 py-4">
                    <div className="text-hand text-lg bg-white rounded-0 text-dark text-center px-1 py-2">Steffi's Imbiss</div>
                </td>
                <td className="align-middle py-3">
                    <p className="text-sm mb-1">Bestellung m&ouml;glich:</p>
                    <p className="text-md text-strong mb-0">Di, 11:15 &ndash; 11:45 Uhr</p>
                </td>
                <td className="align-middle py-3">
                    <p className="text-sm mb-1">Angebot erstellt von:</p>
                    <p className="text-md text-strong mb-0">Dominik Braun</p>
                </td>
                <td className="align-middle px-0 py-3 text-center">
                    <button className="btn btn-light rounded-pill text-sm" disabled>Angebot ausw&auml;hlen</button>
                </td>
            </tr>
            <tr>
                <td className="align-middle pl-0 pr-3 py-4">
                    <div className="text-hand text-lg bg-white rounded-0 text-dark text-center px-1 py-2">Flying Pizza</div>
                </td>
                <td className="align-middle py-3">
                    <p className="text-sm mb-1">Bestellung m&ouml;glich:</p>
                    <p className="text-md text-strong mb-0">Di, 11:15 &ndash; 11:45 Uhr</p>
                </td>
                <td className="align-middle py-3">
                    <p className="text-sm mb-1">Angebot erstellt von:</p>
                    <p className="text-md text-strong mb-0">Angela B&ouml;hm</p>
                </td>
                <td className="align-middle px-0 py-3 text-center">
                    <button className="btn btn-light rounded-pill text-sm" disabled>Angebot ausw&auml;hlen</button>
                </td>
            </tr>
            <tr>
                <td className="align-middle pl-0 pr-3 py-4">
                    <div className="text-hand text-lg bg-white rounded-0 text-dark text-center px-1 py-2">Steffi's Imbiss</div>
                </td>
                <td className="align-middle py-3">
                    <p className="text-sm mb-1">Bestellung m&ouml;glich:</p>
                    <p className="text-md text-strong mb-0">Di, 11:15 &ndash; 11:45 Uhr</p>
                </td>
                <td className="align-middle py-3">
                    <p className="text-sm mb-1">Angebot erstellt von:</p>
                    <p className="text-md text-strong mb-0">Walter Seethaler</p>
                </td>
                <td className="align-middle px-0 py-3 text-center">
                    <button className="btn btn-light rounded-pill text-sm" disabled>Angebot ausw&auml;hlen</button>
                </td>
            </tr>

        </table>
    </div>
}