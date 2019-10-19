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

class LoginForm extends React.Component {

    constructor(props) {
        super(props);
        this.auth = props.auth;
    }

    handleMailAddressChange = (event) => {
        this.auth.mailAddress = event.target.value
    }

    handlePasswordChange = (event) => {
        this.auth.password = event.target.value
    }

    handleLogin = (event) => {
        event.preventDefault()
        this.auth.login()
    }

    render() {
        return (
            <form className="mx-4 mt-3 pb-4 text-center" spellCheck="false" onSubmit={this.handleLogin}>
                <input className="form-control bg-white rounded-0 border-0 my-2 py-4 text-dark"
                       type="text" name="mail-addr" placeholder="Deine E-Mail-Adresse"
                       value={this.auth.mailAddress}
                       onChange={this.handleMailAddressChange}/>
                <input className="form-control bg-white rounded-0 border-0 my-2 py-4 text-dark"
                       type="password" name="password" placeholder="Dein Passwort"
                       value={this.auth.password}
                       onChange={this.handlePasswordChange}/>
                <button className="btn btn-light border-0 text-strong w-100 rounded-0 border-0 my-2 py-2">
                    einloggen
                </button>

                <p className="text-warning login-error">{this.auth.loginErrorMessage}</p>
            </form>
        );
    }
}

export default inject('auth')(observer(LoginForm));