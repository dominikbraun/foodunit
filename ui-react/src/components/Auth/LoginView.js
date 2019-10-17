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

import {inject, observer} from "mobx-react";
import React from 'react';

class LoginView extends React.Component {

    constructor(props) {
        super(props);
        this.auth = props.auth;
    }

    handleMailAddressChange = (event) => {
        this.auth.changeMailAddress(event.target.value)
    }

    handlePasswordChange = (event) => {
        this.auth.changePassword(event.target.value)
    }

    handleLogin = () => {
        this.auth.login()
    }

    render() {
        return (
            <div className="row m-0 align-items-center h-100">

                <div className="col-12 col-md-3 col-xl-4"></div>

                <div className="col-12 col-md-6 col-xl-4 px-1 px-md-4 px-xl-5 py-0">
                    <div className="mx-0 mx-xl-2 my-5">
                        <div
                            className="bg-dark text-white text-hand text-center text-logo px-4 py-4">FoodUnit&nbsp;</div>
                        <div className="bg-gradient px-3 py-4">
                            <h4 className="text-lg text-dark text-strong text-center py-5">Jetzt einloggen &amp; Essen
                                bestellen</h4>
                            <div className="mx-4 mt-3 pb-4 text-center" spellCheck="false">
                                <input className="form-control bg-white rounded-0 border-0 my-2 py-4 text-dark"
                                       type="text" name="mail-addr" placeholder="Deine E-Mail-Adresse"
                                       value={this.auth.mailAddress}
                                       onChange={this.handleMailAddressChange}/>
                                <input className="form-control bg-white rounded-0 border-0 my-2 py-4 text-dark"
                                       type="text" name="password" placeholder="Dein Passwort"
                                       value={this.auth.password}
                                       onChange={this.handlePasswordChange}/>
                                <button
                                    className="btn btn-light border-0 text-strong w-100 rounded-0 border-0 my-2 py-2"
                                    onClick={this.handleLogin}>einloggen
                                </button>
                            </div>
                            <p className="text-center text-sm">&copy; 2019 FoodUnit &mdash; <a
                                className="text-dark link-underlined" target="_blank"
                                href="https://github.com/dominikbraun/foodunit">dominikbraun/foodunit</a></p>
                        </div>
                        <div className="text-md text-center bg-white my-4 py-3">Du hast noch keinen Account? <a
                            className="text-dark2 text-strong link-underlined2" href="">Jetzt registrieren</a></div>
                    </div>
                </div>

                <div className="col-12 col-md-3 col-xl-4"></div>

            </div>
        );
    }
}

export default inject('auth')(observer(LoginView));