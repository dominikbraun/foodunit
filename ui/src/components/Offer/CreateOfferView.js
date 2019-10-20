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
import SidebarRight from "../SidebarRight";
import LoggedIn from "../Auth/LoggedIn";
import {CREATE_OFFER_ROUTE} from "../../util/Routes";
import Footer from "../Footer";

export default function CreateOfferView() {
    return (
        <LoggedIn>
            <div className="row m-0 h-100">
                <SidebarLeft currentActiveRoute={CREATE_OFFER_ROUTE}/>
                <div className="col-12 col-lg-6 col-xl-8 px-1 px-md-4 mx-auto">
                <div className="mx-0 mx-xl-5 my-4 px-5 py-3 bg-white border rounded-0">
                    <h6 className="text-dark text-strong px-0 mb-1 pt-3 pb-0">Angebot erstellen</h6>

                    <div className="row m-0">
                        <div className="col-12 col-md-3"></div>
                            <div className="col-12 col-md-6 text-center">

                                <h4 className="px-2 py-4 text-strong">W&auml;hle ein Restaurant</h4>
                                <form className="my-4" spellCheck="false">
                                    <div className="input-group mb-3">
                                        <div className="input-group-prepend">
                                            <label className="input-group-text bg-light border-0 rounded-0"
                                                   htmlFor="select-restaurant">Bestellen bei:</label>
                                        </div>
                                        <select className="custom-select bg-light border-0 rounded-0"
                                                id="select-restaurant">
                                            <option selected>ausw&auml;hlen ...</option>
                                            <option>Steffi's Imbiss</option>
                                            <option>Imbiss Media</option>
                                            <option>Asia Aroma</option>
                                            <option>McDonald's Straubing</option>
                                            <option>Pizzeria Venezia</option>
                                        </select>
                                    </div>
                                </form>

                                <h4 className="px-2 py-4 text-strong">Lege einen Zeitraum fest</h4>
                                <h6 className="text-normal text-pmd">Dies gibt den Zeitraum an, in dem deine Kollegen
                                    bestellen k&ouml;nnen.</h6>
                                <div className="row mx-0 mt-4">
                                    <div className="col-12 col-lg-5 p-0">
                                        <form className="mx-0 my-4" spellCheck="false">
                                            <div className="input-group w-100">
                                                <select className="custom-select bg-light border-0 rounded-0 text-md"
                                                        id="from-hh">
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
                                                    <label
                                                        className="input-group-text bg-light border-0 text-md">:</label>
                                                </div>
                                                <select className="custom-select bg-light border-0 rounded-0 text-md"
                                                        id="from-mm">
                                                    <option selected>MM</option>
                                                    <option>00</option>
                                                    <option>15</option>
                                                    <option>30</option>
                                                    <option>45</option>
                                                </select>
                                            </div>
                                        </form>
                                    </div>
                                    <div className="col-12 col-lg-2 p-1 text-center">
                                        <p className="m-4">&mdash;</p>
                                    </div>
                                    <div className="col-12 col-lg-5 p-0">
                                        <form className="mx-0 my-4" spellCheck="false">
                                            <div className="input-group w-100">
                                                <select className="custom-select bg-light border-0 text-md" id="to-hh">
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
                                                    <label
                                                        className="input-group-text bg-light border-0 text-md">:</label>
                                                </div>
                                                <select className="custom-select bg-light border-0 rounded-0 text-md"
                                                        id="to-mm">
                                                    <option selected>MM</option>
                                                    <option>00</option>
                                                    <option>15</option>
                                                    <option>30</option>
                                                    <option>45</option>
                                                </select>
                                            </div>
                                        </form>
                                    </div>
                                </div>

                                <h4 className="px-2 py-4 text-strong">Erlaubst du Bezahlung per PayPal?</h4>
                                <form className="mx-5 my-3 text-center" spellCheck="false">
                                    <div className="form-check form-check-inline text-md">
                                        <input className="form-check-input bg-light border-0" type="radio"
                                               name="allow-paypal" id="yes" value="yes" checked/>
                                            <label className="form-check-label" htmlFor="yes">Ja</label>
                                    </div>
                                    <div className="form-check form-check-inline text-md">
                                        <input className="form-check-input bg-light border-0" type="radio"
                                               name="allow-paypal" id="no" value="no"/>
                                            <label className="form-check-label" htmlFor="no">Nein</label>
                                    </div>
                                </form>

                                <button
                                    className="btn btn-success text-md rounded-0 w-100 text-strong mt-4 mb-5 py-2">Angebot
                                    erstellen<span className="text-md d-inline-block ml-2 text-white"><i
                                        className="fas fa-medal mr-1"/>+50</span></button>

                            </div>
                            <div className="col-12 col-md-3"></div>
                        </div>
                    </div>
                    <div className="mx-0 mx-xl-5 my-4 bg-white border rounded-0">
                        <div className="p-3 text-dark text-pmd">
                            <i className="fas fa-question-circle text-primary ml-1 mr-3"/>Mit einem Angebot bietest
                            du deinen Kollegen an, Essen bei einem Restaurant zu bestellen und alle Bestellungen dort
                            abzuholen.
                        </div>
                    </div>

                    <Footer/>
                </div>
                <SidebarRight/>
            </div>
        </LoggedIn>
    );
}