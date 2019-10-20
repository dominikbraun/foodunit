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

import {action, decorate, observable} from 'mobx';
import Axios from 'axios';
import {navigate} from "@reach/router"
import {LOGOUT_ROUTE, MAIN_ROUTE} from '../util/Routes'

/**
 * AuthModel handles login and logout as well as auto login if there is already an existing session.
 */
export default class AuthModel {
    mailAddress = "";
    password = "";
    loggedIn = false;
    loginErrorMessage = "";

    constructor() {
        // do auto login if session still valid
        let that = this;
        Axios.get("http://localhost:9292/v1/users/is-authenticated",
            {withCredentials: true}
        ).then(function (response) {

            if (response.data === true) {
                loggedIn(that);
            }
        })
        .catch(function (error) {
            console.log(error);
        });
    }

    login() {
        let that = this;
        Axios.post("http://localhost:9292/v1/users/login",
            {
                mail_addr: this.mailAddress,
                password: this.password
            },{
                withCredentials: true
            }).then(function (response) {

            if (response.data === true) {
                loggedIn(that);
            } else {
                that.loginErrorMessage = "Login fehlgeschlagen. E-Mail oder Passwort ist falsch."
            }
        })
        .catch(function (error) {
            console.log(error);
        });
    }

    logout() {
        let that = this;
        Axios.get("http://localhost:9292/v1/users/logout",
            {
                withCredentials: true
            }).then(function (response) {

            if (response.data === true) {
                loggedOut(that);
            }
        })
            .catch(function (error) {
                console.log(error);
            });
    }
}

// some private helper functions

function loggedIn(that) {
    that.password = "";
    that.loginErrorMessage = "";
    that.loggedIn = true;

    navigate(MAIN_ROUTE);
}

function loggedOut(that) {
    that.loggedIn = false;
    navigate(LOGOUT_ROUTE);
}

decorate(AuthModel, {
    mailAddress: observable,
    password: observable,
    loggedIn: observable,
    loginErrorMessage: observable,

    login: action,
});