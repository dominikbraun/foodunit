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

import React from "react";
import SidebarLeft from "../SidebarLeft/SidebarLeft";
import {SidebarRightFilled} from "../SidebarRight";
import LoggedIn from "../Auth/LoggedIn";
import Footer from "../Footer";
import {RESTAURANT_VIEW} from "../../util/Routes";
import {Link} from "@reach/router";

export default function OrdersView() {
    return (
        <LoggedIn>
            <div className="row m-0 h-100">
                <SidebarLeft/>
                <div className="col-12 col-lg-6 col-xl-8 px-1 px-md-4 mx-auto">
                    <div className="mx-0 mx-xl-5 my-4 bg-light rounded-0 text-center text-dark">
                        <div className="row m-0">
                            <div className="col-12 col-md-6 col-xl-3 px-2 py-4">
                                <p className="text-sm mb-1">Bestellung m&ouml;glich:</p>
                                <p className="text-md text-strong mb-0">Di, 11:15 &ndash; 11:45 Uhr</p>
                            </div>
                            <div className="col-12 col-md-6 col-xl-3 px-2 py-4">
                                <p className="text-sm mb-1">Abholung durch:</p>
                                <p className="text-md text-strong mb-0">Adrian Brasin</p>
                            </div>
                            <div className="col-12 col-md-6 col-xl-3 px-2 py-4">
                                <p className="text-sm mb-1">Lieferung/Abholung ab:</p>
                                <p className="text-md text-strong mb-0"><i className="far fa-clock mr-2"/>ca. 12 Uhr</p>
                            </div>
                            <div className="col-12 col-md-6 col-xl-3 text-center px-2 py-4">
                                <Link to={RESTAURANT_VIEW} className="btn btn-info rounded-0 text-pmd mt-1 px-3 py-2">&lsaquo; zur&uuml;ck zum Angebot</Link>
                            </div>
                        </div>
                    </div>

                    <div className="mx-0 mx-xl-5 my-4 bg-gradient rounded-0">

                        <div className="py-4">
                            <h1 className="text-hand text-xl text-center text-dark px-2 pt-5 pb-4">Steffi's Imbiss</h1>
                        </div>

                        <div className="row m-0 bg-white-transparent">
                            <div className="col-12 col-lg-6 col-xl-3 text-sm text-dark text-center px-2 py-3">Ittlinger
                                Stra&szlig;e 141
                            </div>
                            <div
                                className="col-12 col-lg-6 col-xl-3 text-sm text-dark text-center px-2 py-3">07:00 &ndash; 01:00
                                Uhr
                            </div>
                            <div className="col-12 col-lg-6 col-xl-3 text-sm text-dark text-center px-2 py-3"><i
                                className="fas fa-phone mr-2"/>09421 3304948
                            </div>
                            <div className="col-12 col-lg-6 col-xl-3 text-dark text-center p-3 p-lg-1 input-group">
                                <div className="input-group-prepend">
                                    <span
                                        className="input-group-text border-0 rounded-0 px-3 bg-white text-dark text-md"><i
                                        className="fas fa-search"/></span>
                                </div>
                                <input type="text"
                                       className="form-control h-100 bg-white text-dark pl-0 pr-3 text-sm border-0 rounded-0"
                                       placeholder="Gericht suchen"/>
                            </div>
                        </div>

                        <ul className="nav nav-dark bg-pre-dark d-block d-lg-flex">
                            <li className="nav-item text-md"><a className="nav-link py-2" href="#">Klassiker</a></li>
                            <li className="nav-item text-md"><a className="nav-link py-2" href="#">F&uuml;r
                                Feinschmecker</a></li>
                            <li className="nav-item text-md"><a className="nav-link py-2" href="#">Burger &amp; Co.</a>
                            </li>
                            <li className="nav-item text-md"><a className="nav-link py-2" href="#">Kinderteller</a></li>
                            <li className="nav-item text-md"><a className="nav-link py-2" href="#">Fisch</a></li>
                            <li className="nav-item text-md"><a className="nav-link py-2" href="#">Salate</a></li>
                            <li className="nav-item text-md"><a className="nav-link py-2" href="#">Alkoholfreie
                                Getr&auml;nke</a></li>
                        </ul>
                    </div>

                    <div className="mx-0 mx-xl-5 my-4">
                        <div className="row">
                            <div className="col-12 col-lg-6 pb-4">
                                <div className="bg-light text-dark text-strong text-md px-3 px-lg-4 py-3">
                                    Dominik Braun (dominik.braun@psbn.ed)
                                </div>
                                <div className="border border-top-0 bg-white">
                                    <div className="border-bottom">
                                        <div className="row m-0">
                                            <div className="col-7 text-dark text-pmd text-strong px-3 px-lg-4 py-4">
                                                Chicken Chips mit Pommes
                                            </div>
                                            <div
                                                className="col-5 text-success text-pmd text-strong text-right px-3 px-lg-4 py-4">
                                                <span
                                                    className="bg-white border px-3 py-1 rounded-pill">Preis: 4.50 &euro;</span>
                                            </div>
                                        </div>
                                        <div className="px-3 px-lg-4 pb-4 text-dark text-pmd">
                                            <p className="mb-1 ml-2"><span
                                                className="text-strong2">Pizza-Durchmesser:</span> <i>28cm</i></p>
                                            <p className="mb-1 ml-2"><span
                                                className="text-strong2">Extra-Zutaten:</span> <i>Obergine, Pepperoni,
                                                Salami</i></p>
                                        </div>
                                    </div>
                                    <div className="border-bottom">
                                        <div className="row m-0">
                                            <div className="col-7 text-dark text-pmd text-strong px-3 px-lg-4 py-4">
                                                Chicken Chips mit Pommes
                                            </div>
                                            <div
                                                className="col-5 text-success text-pmd text-strong text-right px-3 px-lg-4 py-4">
                                                <span
                                                    className="bg-white border px-3 py-1 rounded-pill">Preis: 4.50 &euro;</span>
                                            </div>
                                        </div>
                                        <div className="px-3 px-lg-4 pb-4 text-dark text-pmd">
                                            <p className="mb-1 ml-2"><span
                                                className="text-strong2">Pizza-Durchmesser:</span> <i>28cm</i></p>
                                            <p className="mb-1 ml-2"><span
                                                className="text-strong2">Extra-Zutaten:</span> <i>Obergine, Pepperoni,
                                                Salami</i></p>
                                        </div>
                                    </div>
                                    <div className="px-4 py-3 text-dark text-pmd text-strong">
                                        Gesamt: 3.20 &euro;
                                    </div>
                                </div>
                            </div>
                            <div className="col-12 col-lg-6 pb-4">
                                <div className="bg-light text-dark text-strong text-md px-3 px-lg-4 py-3">
                                    Dominik Braun (dominik.braun@psbn.ed)
                                </div>
                                <div className="border border-top-0 bg-white">
                                    <div className="border-bottom">
                                        <div className="row m-0">
                                            <div className="col-7 text-dark text-pmd text-strong px-3 px-lg-4 py-4">
                                                Chicken Chips mit Pommes
                                            </div>
                                            <div
                                                className="col-5 text-success text-pmd text-strong text-right px-3 px-lg-4 py-4">
                                                <span
                                                    className="bg-white border px-3 py-1 rounded-pill">Preis: 4.50 &euro;</span>
                                            </div>
                                        </div>
                                        <div className="px-3 px-lg-4 pb-4 text-dark text-pmd">
                                            <p className="mb-1 ml-2"><span
                                                className="text-strong2">Pizza-Durchmesser:</span> <i>28cm</i></p>
                                            <p className="mb-1 ml-2"><span
                                                className="text-strong2">Extra-Zutaten:</span> <i>Obergine, Pepperoni,
                                                Salami</i></p>
                                        </div>
                                    </div>
                                    <div className="px-4 py-3 text-dark text-pmd text-strong">
                                        Gesamt: 3.20 &euro;
                                    </div>
                                </div>
                            </div>

                        </div>
                    </div>

                    <div className="mx-0 mx-xl-5 my-5 px-4 py-4 text-center bg-white">
                        <h5 className="">Hast du die Bestellung beim Restaurant aufgegeben?</h5>
                        <p className="text-sm py-3">Trage ein, um wie viel Uhr die Bestellung geliefert wird oder
                            abholbereit ist.</p>
                        <div className="row m-0">
                            <div className="col-12 col-lg-3 col-xl-4 d-none d-lg-block"></div>
                            <div className="col-12 col-lg-6 col-xl-4">
                                <div className="input-group">
                                    <select className="custom-select bg-light border-0 rounded-0 text-md"
                                            id="ready-at-hh">
                                        <option selected>HH</option>
                                        <option value="1">06</option>
                                        <option value="2">07</option>
                                        <option value="3">08</option>
                                        <option value="3">09</option>
                                        <option value="3">09</option>
                                        <option value="3">10</option>
                                        <option value="3">11</option>
                                        <option value="3">12</option>
                                        <option value="3">13</option>
                                        <option value="3">14</option>
                                        <option value="3">15</option>
                                        <option value="3">16</option>
                                        <option value="3">17</option>
                                        <option value="3">18</option>
                                    </select>
                                    <div className="input-group-append input-group-prepend">
                                        <label className="input-group-text bg-light border-0 text-md">:</label>
                                    </div>
                                    <select className="custom-select bg-light border-0 rounded-0 text-md"
                                            id="ready-at-mm">
                                        <option selected>MM</option>
                                        <option>00</option>
                                        <option>05</option>
                                        <option>10</option>
                                        <option>15</option>
                                        <option>20</option>
                                        <option>25</option>
                                        <option>30</option>
                                        <option>35</option>
                                        <option>40</option>
                                        <option>45</option>
                                        <option>50</option>
                                        <option>55</option>
                                    </select>
                                    <button
                                        className="btn btn-dark rounded-0 text-pmd px-3 input-group-append z-1080">speichern
                                    </button>
                                </div>
                            </div>
                            <div className="col-12 col-lg-3 col-xl-4 d-none d-lg-block"></div>
                        </div>
                    </div>

                    <div className="mx-0 mx-xl-5 my-4 bg-white border rounded-0">
                        <div className="p-3 text-dark text-pmd">
                            <i className="fas fa-question-circle text-primary ml-1 mr-3"/>Dies ist
                            die &Uuml;bersicht &uuml;ber alle Bestellungen zu diesem Angebot.
                        </div>
                    </div>

                    <div className="mx-0 mx-xl-5 my-4 px-5 py-3 text-dark text-sm text-center">
                        &copy; 2019 FoodUnit &mdash; <a href="https://github.com/dominikbraun/foodunit"
                                                        className="text-dark link-underlined"
                                                        target="_blank">dominikbraun/foodunit</a><br/>
                        F&uuml;r Feature Requests und Bug-Meldungen, <a className="text-dark link-underlined"
                                                                        href="https://github.com/dominikbraun/foodunit/issues"
                                                                        target="_blank">&ouml;ffne ein Issue</a>.
                    </div>

                </div>

                <div className="col-12 col-lg-3 col-xl-2 p-0 sidebar-right-lg bg-white border-left">
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
                                <button className="btn btn-light rounded-pill text-xs"><i
                                    className="fab fa-paypal mr-2"/>direkt zu PayPal
                                </button>
                            </div>
                        </div>
                    </div>
                    <div className="px-4 pt-0 pb-4">
                        <button className="btn btn-success w-100 text-md rounded-0 text-strong py-2">Einkaufskorb
                            speichern
                        </button>
                    </div>

                    <Footer/>
                </div>
                <SidebarRightFilled/>
            </div>
        </LoggedIn>
    );
}