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
import {Link} from "@reach/router"
import {RESTAURANT_VIEW} from "../../util/Routes"

export class TableConfig {
    constructor(tdAdditionalClasses, getElementCallback) {
        this.tdAdditionalClasses = tdAdditionalClasses
        this.getElement = getElementCallback
    }
}

export default class Table extends React.Component {

    constructor(props) {
        super(props)
    }

    rowsToTableTransformer(rows) {
        return (
            <tr>
                <td className="align-middle pl-0 pr-3 py-4">
                    <div className="text-hand text-lg bg-gradient rounded-0 text-dark text-center px-1 py-2">
                        {this.offer.restaurant.name}
                    </div>
                </td>
                <td className="align-middle py-3">
                    <p className="text-sm mb-1">Bestellung m&ouml;glich:</p>
                    <p className="text-md text-strong mb-0">{this.offer.valid_from}</p>
                </td>
                <td className="align-middle py-3">
                    <p className="text-sm mb-1">Angebot erstellt von:</p>
                    <p className="text-md text-strong mb-0">{this.offer.owner.name}</p>
                </td>
                <td className="align-middle px-0 py-3 text-center">
                    <Link to={RESTAURANT_VIEW + "/" + this.offer.id} className="btn btn-light rounded-pill text-sm">Angebot
                        ausw&auml;hlen</Link>
                </td>
            </tr>
        )
    }

    render() {
        let rows = this.props.rows.map((row) => {
            let columns = this.props.config.map((columnConfig) => {
                return (
                    <td className={`align-middle py-3 ${columnConfig.tdAdditionalClasses}`}>
                        {columnConfig.getElement(row)}
                    </td>
                )
            })

            return (
                <tr>
                    {columns}
                </tr>
            )
        })

        return (
            <table className="table table-responsive-xl mb-0 text-center">
                {rows}
            </table>
        )
    }
}