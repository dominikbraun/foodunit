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

export default function SidebarLeft(props) {
    return <div className="col-12 col-lg-3 col-xl-2 p-0 sidebar-left-lg bg-darker border-right">
        <div className="px-2 py-1 px-xl-4 py-xl-4">
            <div className="bg-primary-gradient text-white text-hand mx-auto text-center text-logo div-logo my-3 px-2 py-2 py-xl-3 rounded-0">Food<wbr/>Unit&nbsp;</div>
        </div>
        <p className="text-light text-center text-strong text-sm mb-0 p-3">Angebot ausw&auml;hlen &amp; bestellen</p>
        <div className="nav flex-column nav-pills side-nav">
            <a className="nav-link text-md px-2 px-xl-3 py-3 rounded-0 active" href="offers.html"><i className="fas fa-pizza-slice ml-1 mr-3"/>Aktuelle Angebote</a>
            <a className="nav-link text-md px-2 px-xl-3 py-3 rounded-0" href="create-offer.html"><i className="fas fa-share ml-1 mr-3"/>Angebot erstellen</a>
            <a className="nav-link text-md px-2 px-xl-3 py-3 rounded-0" href="my-offers.html"><i className="fas fa-layer-group ml-1 mr-3"/>Meine Angebote</a>
        </div>
        <p className="text-light text-center text-strong text-sm mb-0 p-3">Mein Konto</p>
        <div className="nav flex-column nav-pills side-nav">
            <div className="nav-link text-md text-success px-2 px-xl-3 py-3 rounded-0"><i className="fas fa-medal ml-1 mr-3"/>Gesamt: 1200</div>
            <a className="nav-link text-md px-2 px-xl-3 py-3 rounded-0" href="#"><i className="fas fa-user-circle ml-1 mr-3"/>Konto verwalten</a>
            <a className="nav-link text-md px-2 px-xl-3 py-3 rounded-0" href="#"><i className="fas fa-sign-out-alt ml-1 mr-3"/>Ausloggen</a>
        </div>
    </div>
}