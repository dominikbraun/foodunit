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