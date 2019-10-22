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

export default function SidebarRight(props) {
    return <div className="col-12 col-lg-3 col-xl-2 p-0 sidebar-right-lg bg-white border-left">
        <h5 className="text-center border-bottom mb-0 px-2 py-4">Einkaufskorb</h5>
        <p className="text-center text-strong mb-0 pt-5 pb-3 text-sm">Dein Einkaufskorb ist leer.</p>
        <div className="px-4 py-4">
            <div className="row m-0">
                <div className="col-5 text-md text-strong text-muted px-0 py-2">Gesamt: 0.00 &euro;</div>
                <div className="col-7 text-right p-0">
                    <button className="btn btn-light rounded-pill text-xs" disabled><i className="fab fa-paypal mr-2"/>direkt zu PayPal</button>
                </div>
            </div>
        </div>
        <div className="px-4 pt-0 pb-4">
            <button className="btn btn-success w-100 text-md rounded-0 text-strong py-2" disabled>Einkaufskorb speichern</button>
        </div>
    </div>
}

// currently dummy to show how it should look like later
export function SidebarRightFilled(props) {
    return <div className="col-12 col-lg-3 col-xl-2 p-0 sidebar-right-lg bg-white border-left">
        <h5 className="text-center border-bottom mb-0 px-2 py-4">Steffi's Imbiss</h5>
        <div className="px-4 py-4 border-bottom">
            <div className="row m-0">
                <div className="col-10 px-0 pt-2 pb-3 text-md text-strong">4 Chicken Chips mit Pommes</div>
                <div className="col-2 px-0">
                    <button type="button" className="btn btn-white text-md rounded-circle"
                            aria-label="Close">&times;</button>
                </div>
            </div>
            <div className="text-pmd"><i className="fas fa-wrench mr-2"/> <a className="text-muted"
                                                                                href="">konfigurieren</a></div>
            <div className="my-3 bg-white d-none">
                <form className="mb-0">
                    <div className="py-2">
                        <label className="text-md py-1" htmlFor="chara0">Pizza-Durchmesser</label>
                        <select className="custom-select text-md bg-light border-0 rounded-0" id="chara0">
                            <option selected>26 cm</option>
                            <option>28 cm</option>
                            <option>32 cm</option>
                        </select>
                    </div>
                    <div className="py-2">
                        <h6 className="text-md text-normal py-1">Extra-Zutaten</h6>
                        <div className="">
                            <input className="" type="checkbox" value="" id="chara1-0"/>
                                <label className="mb-0 ml-2 text-md" htmlFor="chara1-0">Mozzarella</label>
                        </div>
                        <div className="">
                            <input className="" type="checkbox" value="" id="chara1-1"/>
                                <label className="mb-0 ml-2 text-md" htmlFor="chara1-1">Schinken</label>
                        </div>
                        <div className="">
                            <input className="" type="checkbox" value="" id="chara1-2"/>
                                <label className="mb-0 ml-2 text-md" htmlFor="chara1-2">Pepperoni</label>
                        </div>
                    </div>
                </form>
            </div>
            <div className="text-pmd"><i className="fas fa-pen mr-2"/> <a className="text-muted" href="">Anmerkung
                bearbeiten</a></div>
            <div className="text-md text-strong pt-3 pb-2">Preis: 4.50 &euro;</div>
        </div>
        <div className="px-4 py-4 border-bottom">
            <div className="row m-0">
                <div className="col-10 px-0 pt-2 pb-3 text-md text-strong">Schweineschitzel paniert</div>
                <div className="col-2 px-0">
                    <button type="button" className="btn btn-white text-md rounded-circle"
                            aria-label="Close">&times;</button>
                </div>
            </div>
            <div className="text-pmd"><i className="fas fa-wrench mr-2"/> <a className="text-muted"
                                                                                href="">konfigurieren</a></div>
            <div className="text-pmd"><i className="fas fa-pen mr-2"/> <a className="text-muted" href="">Anmerkung
                bearbeiten</a></div>
            <div className="text-md text-strong pt-3 pb-2">Preis: 3.20 &euro;</div>
        </div>
        <div className="px-4 py-4">
            <div className="row m-0">
                <div className="col-5 text-md text-strong px-0 py-3">Gesamt: 3.20 &euro;</div>
                <div className="col-7 text-right p-2">
                    <button className="btn btn-light rounded-pill text-xs"><i className="fab fa-paypal mr-2"/>direkt zu
                        PayPal
                    </button>
                </div>
            </div>
        </div>
        <div className="px-4 pt-0 pb-4">
            <button className="btn btn-success w-100 text-md rounded-0 text-strong py-2">Einkaufskorb speichern</button>
        </div>
    </div>
}