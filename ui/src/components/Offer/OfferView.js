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
import SidebarLeft from "../SidebarLeft/SidebarLeft"
import {SidebarRightFilled} from "../SidebarRight"
import LoggedIn from "../Auth/LoggedIn"
import Footer from "../Footer"
import {Link} from "@reach/router"
import {ORDERS_VIEW} from "../../util/Routes"

export default function OfferView(props) {
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
                                <p className="text-md text-strong mb-0"><i className="far fa-clock mr-2"/>ca. 12 Uhr
                                </p>
                            </div>
                            <div className="col-12 col-md-6 col-xl-3 text-center px-2 py-4">
                                <Link to={ORDERS_VIEW} className="btn btn-info rounded-0 text-pmd mt-1 px-3 py-2">Alle Bestellungen anzeigen &rsaquo;</Link>
                            </div>
                        </div>
                    </div>

                    <div className="mx-0 mx-xl-5 my-4 bg-gradient rounded-0">

                        <div className="py-4">
                            <h1 className="text-hand text-xl text-center text-dark px-2 pt-5 pb-4">{props.offerId}</h1>
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

                    <div className="mx-0 mx-xl-5 my-4 px-5 py-2 bg-white border rounded-0">
                        <h4 className="text-dark text-strong mb-0 px-0 py-5">Burger &amp; Nuggets</h4>

                        <table className="table table-responsive-xl">
                            <tr>
                                <td className="align-middle px-0 py-4">
                                    <h6 className="text-strong mt-1">4 Chicken Chips mit Pommes</h6>
                                    <p className="text-dark text-pmd mb-1">Frisches Schnitzel vom Schwein mit Pommes,
                                        mit Blaukraut und Feldsalat</p>
                                </td>
                                <td className="align-middle py-3">
                                    <span className="text-success text-strong text-md">Preis: 7.20 &euro;</span>
                                </td>
                                <td className="align-middle px-0 py-3 text-center">
                                    <button className="btn btn-light rounded-pill text-sm"><i
                                        className="fas fa-shopping-basket mr-2"/>hinzuf&uuml;gen
                                    </button>
                                </td>
                            </tr>


                            <tr>
                                <td className="align-middle px-0 py-4">
                                    <h6 className="text-strong mt-1">Pizza Drehspie&szlig;</h6>
                                    <p className="text-dark text-pmd mb-1">Frisches Schnitzel vom Schwein mit Pommes,
                                        mit Blaukraut und Feldsalat</p>
                                </td>
                                <td className="align-middle py-3">
                                    <span className="text-success text-strong text-md">Preis: 7.20 &euro;</span><span
                                    className="text-sm text-muted"><i className="fas fa-wrench ml-3 mr-1"
                                                                      title="Dieses Gericht ist individuell konfigurierbar."/></span>
                                </td>
                                <td className="align-middle px-0 py-3 text-center">
                                    <button className="btn btn-light rounded-pill text-sm"><i
                                        className="fas fa-shopping-basket mr-2"/>hinzuf&uuml;gen
                                    </button>
                                </td>
                            </tr>
                            <tr>
                                <td className="align-middle px-0 py-4">
                                    <h6 className="text-strong mt-1">Drehspie&szlig; XL</h6>
                                    <p className="text-dark text-pmd mb-1">Frisches Schnitzel vom Schwein mit Pommes,
                                        mit Blaukraut und Feldsalat</p>
                                </td>
                                <td className="align-middle py-3">
                                    <span className="text-success text-strong text-md">Preis: 7.20 &euro;</span>
                                </td>
                                <td className="align-middle px-0 py-3 text-center">
                                    <button className="btn btn-light rounded-pill text-sm"><i
                                        className="fas fa-shopping-basket mr-2"/>hinzuf&uuml;gen
                                    </button>
                                </td>
                            </tr>
                            <tr>
                                <td className="align-middle px-0 py-4">
                                    <h6 className="text-strong mt-1">Pommes gro&szlig;</h6>
                                    <p className="text-dark text-pmd mb-1">Frisches Schnitzel vom Schwein mit Pommes,
                                        mit Blaukraut und Feldsalat</p>
                                </td>
                                <td className="align-middle py-3">
                                    <span className="text-success text-strong text-md">Preis: 7.20 &euro;</span>
                                </td>
                                <td className="align-middle px-0 py-3 text-center">
                                    <button className="btn btn-light rounded-pill text-sm"><i
                                        className="fas fa-shopping-basket mr-2"/>hinzuf&uuml;gen
                                    </button>
                                </td>
                            </tr>
                            <tr>
                                <td className="align-middle px-0 py-4">
                                    <h6 className="text-strong mt-1">Hamburger Royal K&auml;se</h6>
                                    <p className="text-dark text-pmd mb-1">Frisches Schnitzel vom Schwein mit Pommes,
                                        mit Blaukraut und Feldsalat</p>
                                </td>
                                <td className="align-middle py-3">
                                    <span className="text-success text-strong text-md">Preis: 7.20 &euro;</span><span
                                    className="text-sm text-muted"><i className="fas fa-wrench ml-3 mr-1"
                                                                      title="Dieses Gericht ist individuell konfigurierbar."/><i
                                    className="fas fa-exclamation-triangle ml-1 mr-1" title=""/></span>
                                </td>
                                <td className="align-middle px-0 py-3 text-center">
                                    <button className="btn btn-light rounded-pill text-sm"><i
                                        className="fas fa-shopping-basket mr-2"/>hinzuf&uuml;gen
                                    </button>
                                </td>
                            </tr>
                            <tr>
                                <td className="align-middle px-0 py-4">
                                    <h6 className="text-strong mt-1">Holzf&auml;llersteak vom Schwein</h6>
                                    <p className="text-dark text-pmd mb-1">Frisches Schnitzel vom Schwein mit Pommes,
                                        mit Blaukraut und Feldsalat</p>
                                </td>
                                <td className="align-middle py-3">
                                    <span className="text-success text-strong text-md">Preis: 7.20 &euro;</span><span
                                    className="text-sm text-muted"><i className="fas fa-leaf ml-3 mr-1"
                                                                      title="Dieses Gericht ist vegetarisch."/></span>
                                </td>
                                <td className="align-middle px-0 py-3 text-center">
                                    <button className="btn btn-light rounded-pill text-sm"><i
                                        className="fas fa-shopping-basket mr-2"/>hinzuf&uuml;gen
                                    </button>
                                </td>
                            </tr>
                            <tr>
                                <td className="align-middle px-0 py-4">
                                    <h6 className="text-strong mt-1">Pizza Drehspie&szlig;</h6>
                                    <p className="text-dark text-pmd mb-1">Frisches Schnitzel vom Schwein mit Pommes,
                                        mit Blaukraut und Feldsalat</p>
                                </td>
                                <td className="align-middle py-3">
                                    <span className="text-success text-strong text-md">Preis: 7.20 &euro;</span>
                                </td>
                                <td className="align-middle px-0 py-3 text-center">
                                    <button className="btn btn-light rounded-pill text-sm"><i
                                        className="fas fa-shopping-basket mr-2"/>hinzuf&uuml;gen
                                    </button>
                                </td>
                            </tr>
                            <tr>
                                <td className="align-middle px-0 py-4">
                                    <h6 className="text-strong mt-1">Drehspie&szlig; XL</h6>
                                    <p className="text-dark text-pmd mb-1">Frisches Schnitzel vom Schwein mit Pommes,
                                        mit Blaukraut und Feldsalat</p>
                                </td>
                                <td className="align-middle py-3">
                                    <span className="text-success text-strong text-md">Preis: 7.20 &euro;</span>
                                </td>
                                <td className="align-middle px-0 py-3 text-center">
                                    <button className="btn btn-light rounded-pill text-sm"><i
                                        className="fas fa-shopping-basket mr-2"/>hinzuf&uuml;gen
                                    </button>
                                </td>
                            </tr>
                            <tr>
                                <td className="align-middle px-0 py-4">
                                    <h6 className="text-strong mt-1">Pommes gro&szlig;</h6>
                                    <p className="text-dark text-pmd mb-1">Frisches Schnitzel vom Schwein mit Pommes,
                                        mit Blaukraut und Feldsalat</p>
                                </td>
                                <td className="align-middle py-3">
                                    <span className="text-success text-strong text-md">Preis: 7.20 &euro;</span><span
                                    className="text-sm text-muted"><i className="fas fa-wrench ml-3 mr-1"/></span>
                                </td>
                                <td className="align-middle px-0 py-3 text-center">
                                    <button className="btn btn-light rounded-pill text-sm"><i
                                        className="fas fa-shopping-basket mr-2"/>hinzuf&uuml;gen
                                    </button>
                                </td>
                            </tr>
                            <tr>
                                <td className="align-middle px-0 py-4">
                                    <h6 className="text-strong mt-1">Hamburger Royal K&auml;se</h6>
                                    <p className="text-dark text-pmd mb-1">Frisches Schnitzel vom Schwein mit Pommes,
                                        mit Blaukraut und Feldsalat</p>
                                </td>
                                <td className="align-middle py-3">
                                    <span className="text-success text-strong text-md">Preis: 7.20 &euro;</span>
                                </td>
                                <td className="align-middle px-0 py-3 text-center">
                                    <button className="btn btn-light rounded-pill text-sm"><i
                                        className="fas fa-shopping-basket mr-2"/>hinzuf&uuml;gen
                                    </button>
                                </td>
                            </tr>
                            <tr>
                                <td className="align-middle px-0 py-4">
                                    <h6 className="text-strong mt-1">Holzf&auml;llersteak vom Schwein</h6>
                                    <p className="text-dark text-pmd mb-1">Frisches Schnitzel vom Schwein mit Pommes,
                                        mit Blaukraut und Feldsalat</p>
                                </td>
                                <td className="align-middle py-3">
                                    <span className="text-success text-strong text-md">Preis: 7.20 &euro;</span>
                                </td>
                                <td className="align-middle px-0 py-3 text-center">
                                    <button className="btn btn-light rounded-pill text-sm"><i
                                        className="fas fa-shopping-basket mr-2"/>hinzuf&uuml;gen
                                    </button>
                                </td>
                            </tr>
                            <tr>
                                <td className="align-middle px-0 py-4">
                                    <h6 className="text-strong mt-1">Pizza Drehspie&szlig;</h6>
                                    <p className="text-dark text-pmd mb-1">Frisches Schnitzel vom Schwein mit Pommes,
                                        mit Blaukraut und Feldsalat</p>
                                </td>
                                <td className="align-middle py-3">
                                    <span className="text-success text-strong text-md">Preis: 7.20 &euro;</span>
                                </td>
                                <td className="align-middle px-0 py-3 text-center">
                                    <button className="btn btn-light rounded-pill text-sm"><i
                                        className="fas fa-shopping-basket mr-2"/>hinzuf&uuml;gen
                                    </button>
                                </td>
                            </tr>
                            <tr>
                                <td className="align-middle px-0 py-4">
                                    <h6 className="text-strong mt-1">Drehspie&szlig; XL</h6>
                                    <p className="text-dark text-pmd mb-1">Frisches Schnitzel vom Schwein mit Pommes,
                                        mit Blaukraut und Feldsalat</p>
                                </td>
                                <td className="align-middle py-3">
                                    <span className="text-success text-strong text-md">Preis: 7.20 &euro;</span>
                                </td>
                                <td className="align-middle px-0 py-3 text-center">
                                    <button className="btn btn-light rounded-pill text-sm"><i
                                        className="fas fa-shopping-basket mr-2"/>hinzuf&uuml;gen
                                    </button>
                                </td>
                            </tr>
                            <tr>
                                <td className="align-middle px-0 py-4">
                                    <h6 className="text-strong mt-1">Pommes gro&szlig;</h6>
                                    <p className="text-dark text-pmd mb-1">Frisches Schnitzel vom Schwein mit Pommes,
                                        mit Blaukraut und Feldsalat</p>
                                </td>
                                <td className="align-middle py-3">
                                    <span className="text-success text-strong text-md">Preis: 7.20 &euro;</span>
                                </td>
                                <td className="align-middle px-0 py-3 text-center">
                                    <button className="btn btn-light rounded-pill text-sm"><i
                                        className="fas fa-shopping-basket mr-2"/>hinzuf&uuml;gen
                                    </button>
                                </td>
                            </tr>
                        </table>
                    </div>

                    <div className="mx-0 mx-xl-5 my-4 px-5 py-2 bg-white border rounded-0">
                        <h4 className="text-dark text-strong mb-0 px-0 py-5">Burger &amp; Nuggets</h4>

                        <table className="table table-responsive-xl">
                            <tr>
                                <td className="align-middle px-0 py-4">
                                    <h6 className="text-strong mt-1">4 Chicken Chips mit Pommes</h6>
                                    <p className="text-dark text-pmd mb-1">Frisches Schnitzel vom Schwein mit Pommes,
                                        mit Blaukraut und Feldsalat</p>
                                </td>
                                <td className="align-middle py-3">
                                    <span className="text-success text-strong text-md">Preis: 7.20 &euro;</span>
                                </td>
                                <td className="align-middle px-0 py-3 text-center">
                                    <button className="btn btn-light rounded-pill text-sm"><i
                                        className="fas fa-shopping-basket mr-2"/>hinzuf&uuml;gen
                                    </button>
                                </td>
                            </tr>


                            <tr>
                                <td className="align-middle px-0 py-4">
                                    <h6 className="text-strong mt-1">Pizza Drehspie&szlig;</h6>
                                    <p className="text-dark text-pmd mb-1">Frisches Schnitzel vom Schwein mit Pommes,
                                        mit Blaukraut und Feldsalat</p>
                                </td>
                                <td className="align-middle py-3">
                                    <span className="text-success text-strong text-md">Preis: 7.20 &euro;</span><span
                                    className="text-sm text-muted"><i className="fas fa-wrench ml-3 mr-1"/><i
                                    className="fas fa-exclamation-triangle ml-1 mr-1"/></span>
                                </td>
                                <td className="align-middle px-0 py-3 text-center">
                                    <button className="btn btn-light rounded-pill text-sm"><i
                                        className="fas fa-shopping-basket mr-2"/>hinzuf&uuml;gen
                                    </button>
                                </td>
                            </tr>
                            <tr>
                                <td className="align-middle px-0 py-4">
                                    <h6 className="text-strong mt-1">Drehspie&szlig; XL</h6>
                                    <p className="text-dark text-pmd mb-1">Frisches Schnitzel vom Schwein mit Pommes,
                                        mit Blaukraut und Feldsalat</p>
                                </td>
                                <td className="align-middle py-3">
                                    <span className="text-success text-strong text-md">Preis: 7.20 &euro;</span>
                                </td>
                                <td className="align-middle px-0 py-3 text-center">
                                    <button className="btn btn-light rounded-pill text-sm"><i
                                        className="fas fa-shopping-basket mr-2"/>hinzuf&uuml;gen
                                    </button>
                                </td>
                            </tr>
                            <tr>
                                <td className="align-middle px-0 py-4">
                                    <h6 className="text-strong mt-1">Pommes gro&szlig;</h6>
                                    <p className="text-dark text-pmd mb-1">Frisches Schnitzel vom Schwein mit Pommes,
                                        mit Blaukraut und Feldsalat</p>
                                </td>
                                <td className="align-middle py-3">
                                    <span className="text-success text-strong text-md">Preis: 7.20 &euro;</span>
                                </td>
                                <td className="align-middle px-0 py-3 text-center">
                                    <button className="btn btn-light rounded-pill text-sm"><i
                                        className="fas fa-shopping-basket mr-2"/>hinzuf&uuml;gen
                                    </button>
                                </td>
                            </tr>
                            <tr>
                                <td className="align-middle px-0 py-4">
                                    <h6 className="text-strong mt-1">Hamburger Royal K&auml;se</h6>
                                    <p className="text-dark text-pmd mb-1">Frisches Schnitzel vom Schwein mit Pommes,
                                        mit Blaukraut und Feldsalat</p>
                                </td>
                                <td className="align-middle py-3">
                                    <span className="text-success text-strong text-md">Preis: 7.20 &euro;</span><span
                                    className="text-sm text-muted"><i className="fas fa-wrench ml-3 mr-1"/></span>
                                </td>
                                <td className="align-middle px-0 py-3 text-center">
                                    <button className="btn btn-light rounded-pill text-sm"><i
                                        className="fas fa-shopping-basket mr-2"/>hinzuf&uuml;gen
                                    </button>
                                </td>
                            </tr>
                            <tr>
                                <td className="align-middle px-0 py-4">
                                    <h6 className="text-strong mt-1">Holzf&auml;llersteak vom Schwein</h6>
                                    <p className="text-dark text-pmd mb-1">Frisches Schnitzel vom Schwein mit Pommes,
                                        mit Blaukraut und Feldsalat</p>
                                </td>
                                <td className="align-middle py-3">
                                    <span className="text-success text-strong text-md">Preis: 7.20 &euro;</span>
                                </td>
                                <td className="align-middle px-0 py-3 text-center">
                                    <button className="btn btn-light rounded-pill text-sm"><i
                                        className="fas fa-shopping-basket mr-2"/>hinzuf&uuml;gen
                                    </button>
                                </td>
                            </tr>
                            <tr>
                                <td className="align-middle px-0 py-4">
                                    <h6 className="text-strong mt-1">Pizza Drehspie&szlig;</h6>
                                    <p className="text-dark text-pmd mb-1">Frisches Schnitzel vom Schwein mit Pommes,
                                        mit Blaukraut und Feldsalat</p>
                                </td>
                                <td className="align-middle py-3">
                                    <span className="text-success text-strong text-md">Preis: 7.20 &euro;</span>
                                </td>
                                <td className="align-middle px-0 py-3 text-center">
                                    <button className="btn btn-light rounded-pill text-sm"><i
                                        className="fas fa-shopping-basket mr-2"/>hinzuf&uuml;gen
                                    </button>
                                </td>
                            </tr>
                            <tr>
                                <td className="align-middle px-0 py-4">
                                    <h6 className="text-strong mt-1">Drehspie&szlig; XL</h6>
                                    <p className="text-dark text-pmd mb-1">Frisches Schnitzel vom Schwein mit Pommes,
                                        mit Blaukraut und Feldsalat</p>
                                </td>
                                <td className="align-middle py-3">
                                    <span className="text-success text-strong text-md">Preis: 7.20 &euro;</span>
                                </td>
                                <td className="align-middle px-0 py-3 text-center">
                                    <button className="btn btn-light rounded-pill text-sm"><i
                                        className="fas fa-shopping-basket mr-2"/>hinzuf&uuml;gen
                                    </button>
                                </td>
                            </tr>
                            <tr>
                                <td className="align-middle px-0 py-4">
                                    <h6 className="text-strong mt-1">Pommes gro&szlig;</h6>
                                    <p className="text-dark text-pmd mb-1">Frisches Schnitzel vom Schwein mit Pommes,
                                        mit Blaukraut und Feldsalat</p>
                                </td>
                                <td className="align-middle py-3">
                                    <span className="text-success text-strong text-md">Preis: 7.20 &euro;</span><span
                                    className="text-sm text-muted"><i className="fas fa-wrench ml-3 mr-1"/></span>
                                </td>
                                <td className="align-middle px-0 py-3 text-center">
                                    <button className="btn btn-light rounded-pill text-sm"><i
                                        className="fas fa-shopping-basket mr-2"/>hinzuf&uuml;gen
                                    </button>
                                </td>
                            </tr>
                            <tr>
                                <td className="align-middle px-0 py-4">
                                    <h6 className="text-strong mt-1">Hamburger Royal K&auml;se</h6>
                                    <p className="text-dark text-pmd mb-1">Frisches Schnitzel vom Schwein mit Pommes,
                                        mit Blaukraut und Feldsalat</p>
                                </td>
                                <td className="align-middle py-3">
                                    <span className="text-success text-strong text-md">Preis: 7.20 &euro;</span>
                                </td>
                                <td className="align-middle px-0 py-3 text-center">
                                    <button className="btn btn-light rounded-pill text-sm"><i
                                        className="fas fa-shopping-basket mr-2"/>hinzuf&uuml;gen
                                    </button>
                                </td>
                            </tr>
                            <tr>
                                <td className="align-middle px-0 py-4">
                                    <h6 className="text-strong mt-1">Holzf&auml;llersteak vom Schwein</h6>
                                    <p className="text-dark text-pmd mb-1">Frisches Schnitzel vom Schwein mit Pommes,
                                        mit Blaukraut und Feldsalat</p>
                                </td>
                                <td className="align-middle py-3">
                                    <span className="text-success text-strong text-md">Preis: 7.20 &euro;</span>
                                </td>
                                <td className="align-middle px-0 py-3 text-center">
                                    <button className="btn btn-light rounded-pill text-sm"><i
                                        className="fas fa-shopping-basket mr-2"/>hinzuf&uuml;gen
                                    </button>
                                </td>
                            </tr>
                            <tr>
                                <td className="align-middle px-0 py-4">
                                    <h6 className="text-strong mt-1">Pizza Drehspie&szlig;</h6>
                                    <p className="text-dark text-pmd mb-1">Frisches Schnitzel vom Schwein mit Pommes,
                                        mit Blaukraut und Feldsalat</p>
                                </td>
                                <td className="align-middle py-3">
                                    <span className="text-success text-strong text-md">Preis: 7.20 &euro;</span>
                                </td>
                                <td className="align-middle px-0 py-3 text-center">
                                    <button className="btn btn-light rounded-pill text-sm"><i
                                        className="fas fa-shopping-basket mr-2"/>hinzuf&uuml;gen
                                    </button>
                                </td>
                            </tr>
                            <tr>
                                <td className="align-middle px-0 py-4">
                                    <h6 className="text-strong mt-1">Drehspie&szlig; XL</h6>
                                    <p className="text-dark text-pmd mb-1">Frisches Schnitzel vom Schwein mit Pommes,
                                        mit Blaukraut und Feldsalat</p>
                                </td>
                                <td className="align-middle py-3">
                                    <span className="text-success text-strong text-md">Preis: 7.20 &euro;</span>
                                </td>
                                <td className="align-middle px-0 py-3 text-center">
                                    <button className="btn btn-light rounded-pill text-sm"><i
                                        className="fas fa-shopping-basket mr-2"/>hinzuf&uuml;gen
                                    </button>
                                </td>
                            </tr>
                            <tr>
                                <td className="align-middle px-0 py-4">
                                    <h6 className="text-strong mt-1">Pommes gro&szlig;</h6>
                                    <p className="text-dark text-pmd mb-1">Frisches Schnitzel vom Schwein mit Pommes,
                                        mit Blaukraut und Feldsalat</p>
                                </td>
                                <td className="align-middle py-3">
                                    <span className="text-success text-strong text-md">Preis: 7.20 &euro;</span>
                                </td>
                                <td className="align-middle px-0 py-3 text-center">
                                    <button className="btn btn-light rounded-pill text-sm"><i
                                        className="fas fa-shopping-basket mr-2"/>hinzuf&uuml;gen
                                    </button>
                                </td>
                            </tr>
                        </table>
                    </div>

                    <div className="mx-0 mx-xl-5 my-4 px-5 py-2 bg-light rounded-0">
                        <h6 className="text-dark text-strong mb-0 px-0 pt-4 pb-3">Legende</h6>
                        <table className="table table-responsive mb-3 text-dark text-pmd w-100">
                            <tr>
                                <td className="px-0 py-1"><i className="fas fa-wrench ml-0 mr-3 text-dark"/></td>
                                <td className="px-0 py-1">Dieses Gericht ist individuell konfigurierbar.</td>
                            </tr>
                            <tr>
                                <td className="px-0 py-1"><i
                                    className="fas fa-exclamation-triangle ml-0 mr-3 text-primary"/></td>
                                <td className="px-0 py-1">Dieses Gericht ist nur unsicher verf&uuml;gbar. Du kannst beim
                                    Bestellen eine Alternative angeben.
                                </td>
                            </tr>
                            <tr>
                                <td className="px-0 py-1"><i className="fas fa-leaf ml-0 mr-3 text-success"/></td>
                                <td className="px-0 py-1">Dieses Gericht ist vegetarisch oder vegan.</td>
                            </tr>
                        </table>
                    </div>

                    <div className="mx-0 mx-xl-5 my-4 bg-white border rounded-0">
                        <div className="p-3 text-dark text-pmd">
                            <i className="fas fa-question-circle text-primary ml-1 mr-3"/>Du kannst deine Bestellung
                            Bar oder per PayPal bei Adrian Brasin bezahlen.
                        </div>
                    </div>

                    <Footer/>
                </div>
                <SidebarRightFilled/>
            </div>
        </LoggedIn>
    )
}