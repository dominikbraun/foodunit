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
import {MY_OFFERS_ROUTE} from "../../util/Routes";
import Footer from "../Footer";

export default function MyOffersView() {
    return (
        <LoggedIn>
            <div className="row m-0 h-100">
                <SidebarLeft currentActiveRoute={MY_OFFERS_ROUTE}/>
                <div className="col-12 col-lg-6 col-xl-8 px-1 px-md-4 mx-auto">

                    <div className="mx-0 mx-xl-5 my-4 px-5 py-3 bg-white border rounded-0">
                        <h6 className="text-dark text-strong px-0 py-3">Meine aktuellen Angebote</h6>

                        <table className="table table-responsive-xl mb-0 text-center">
                            <tr>
                                <td className="align-middle pl-0 pr-3 py-4">
                                    <div
                                        className="text-hand text-lg bg-light rounded-0 text-dark text-center px-1 py-2">Pizzeria
                                        Venezia
                                    </div>
                                </td>
                                <td className="align-middle py-3">
                                    <p className="text-sm mb-1">Bestellung m&ouml;glich:</p>
                                    <p className="text-md text-strong mb-0">Di, 11:15 &ndash; 11:45 Uhr</p>
                                </td>
                                <td className="align-middle py-3">
                                    <p className="text-sm mb-1">Anzahl Bestellungen:</p>
                                    <p className="text-md text-strong mb-0">Bisher 6</p>
                                </td>
                                <td className="align-middle px-0 py-3 text-center">
                                    <a href="restaurant.html" className="btn btn-light rounded-pill text-sm px-3">zum
                                        Angebot</a>
                                </td>
                            </tr>

                            <tr>
                                <td className="align-middle pl-0 pr-3 py-4">
                                    <div
                                        className="text-hand text-lg bg-light rounded-0 text-dark text-center px-1 py-2">Imbiss
                                        Media
                                    </div>
                                </td>
                                <td className="align-middle py-3">
                                    <p className="text-sm mb-1">Bestellung m&ouml;glich:</p>
                                    <p className="text-md text-strong mb-0">Di, 11:15 &ndash; 11:45 Uhr</p>
                                </td>
                                <td className="align-middle py-3">
                                    <p className="text-sm mb-1">Anzahl Bestellungen:</p>
                                    <p className="text-md text-strong mb-0">Bisher 6</p>
                                </td>
                                <td className="align-middle px-0 py-3 text-center">
                                    <a href="restaurant.html" className="btn btn-light rounded-pill text-sm px-3">zum
                                        Angebot</a>
                                </td>
                            </tr>
                            <tr>
                                <td className="align-middle pl-0 pr-3 py-4">
                                    <div
                                        className="text-hand text-lg bg-light rounded-0 text-dark text-center px-1 py-2">McDonald's
                                    </div>
                                </td>
                                <td className="align-middle py-3">
                                    <p className="text-sm mb-1">Bestellung m&ouml;glich:</p>
                                    <p className="text-md text-strong mb-0">Di, 11:15 &ndash; 11:45 Uhr</p>
                                </td>
                                <td className="align-middle py-3">
                                    <p className="text-sm mb-1">Anzahl Bestellungen:</p>
                                    <p className="text-md text-strong mb-0">Bisher 6</p>
                                </td>
                                <td className="align-middle px-0 py-3 text-center">
                                    <a href="restaurant.html" className="btn btn-light rounded-pill text-sm px-3">zum
                                        Angebot</a>
                                </td>
                            </tr>

                        </table>

                        <div className="border-top-light text-right pt-3">
                            <a href="create-offer.html" className="btn btn-link rounded-pill text-sm">
                                <i className="fas fa-share mr-2"/>Angebot erstellen</a>
                        </div>
                    </div>

                    <div className="mx-0 mx-xl-5 my-4 px-5 py-3 bg-white border rounded-0">
                        <h6 className="text-dark text-strong px-0 py-3">Meine abgelaufenen Angebote</h6>

                        <table className="table table-responsive-xl mb-0 text-center">
                            <tr>
                                <td className="align-middle pl-0 pr-3 py-4">
                                    <div
                                        className="text-hand text-lg bg-white rounded-0 text-dark text-center px-1 py-2">Pizzeria
                                        Venezia
                                    </div>
                                </td>
                                <td className="align-middle py-3">
                                    <p className="text-sm mb-1">Bestellung m&ouml;glich:</p>
                                    <p className="text-md text-strong mb-0">Di, 11:15 &ndash; 11:45 Uhr</p>
                                </td>
                                <td className="align-middle py-3">
                                    <p className="text-sm mb-1">Anzahl Bestellungen:</p>
                                    <p className="text-md text-strong mb-0">Insgesamt 6</p>
                                </td>
                                <td className="align-middle px-0 py-3 text-center">
                                    <a href="restaurant.html" className="btn btn-light rounded-pill text-sm px-3">zum
                                        Angebot</a>
                                </td>
                            </tr>

                            <tr>
                                <td className="align-middle pl-0 pr-3 py-4">
                                    <div
                                        className="text-hand text-lg bg-white rounded-0 text-dark text-center px-1 py-2">Imbiss
                                        Media
                                    </div>
                                </td>
                                <td className="align-middle py-3">
                                    <p className="text-sm mb-1">Bestellung m&ouml;glich:</p>
                                    <p className="text-md text-strong mb-0">Di, 11:15 &ndash; 11:45 Uhr</p>
                                </td>
                                <td className="align-middle py-3">
                                    <p className="text-sm mb-1">Anzahl Bestellungen:</p>
                                    <p className="text-md text-strong mb-0">Insgesamt 6</p>
                                </td>
                                <td className="align-middle px-0 py-3 text-center">
                                    <a href="restaurant.html" className="btn btn-light rounded-pill text-sm px-3">zum
                                        Angebot</a>
                                </td>
                            </tr>
                            <tr>
                                <td className="align-middle pl-0 pr-3 py-4">
                                    <div
                                        className="text-hand text-lg bg-white rounded-0 text-dark text-center px-1 py-2">McDonald's
                                    </div>
                                </td>
                                <td className="align-middle py-3">
                                    <p className="text-sm mb-1">Bestellung m&ouml;glich:</p>
                                    <p className="text-md text-strong mb-0">Di, 11:15 &ndash; 11:45 Uhr</p>
                                </td>
                                <td className="align-middle py-3">
                                    <p className="text-sm mb-1">Anzahl Bestellungen:</p>
                                    <p className="text-md text-strong mb-0">Insgesamt 6</p>
                                </td>
                                <td className="align-middle px-0 py-3 text-center">
                                    <a href="restaurant.html" className="btn btn-light rounded-pill text-sm px-3">zum
                                        Angebot</a>
                                </td>
                            </tr>

                        </table>
                    </div>

                    <div className="mx-0 mx-xl-5 my-4 bg-white border rounded-0">
                        <div className="p-3 text-dark text-pmd">
                            <i className="fas fa-question-circle text-primary ml-1 mr-3"></i>Mit einem Angebot
                            bietest du deinen Kollegen an, Essen bei einem Restaurant zu bestellen und alle
                            Bestellungen dort abzuholen.
                        </div>
                    </div>

                    <Footer/>
                </div>
                <SidebarRight/>
            </div>
        </LoggedIn>
    );
}