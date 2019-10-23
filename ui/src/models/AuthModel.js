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

import {action, configure, decorate, observable} from 'mobx'
import Axios from 'axios'
import {navigate} from "@reach/router"
import {LOGOUT_ROUTE, MAIN_ROUTE} from '../util/Routes'

// enforce the strict mode for actions -> e.g. no state modifying inside of promise without action decorator https://www.leighhalliday.com/mobx-async-actions
configure({ enforceActions: "always" })

/**
 * AuthModel handles login and logout as well as auto login if there is already an existing session.
 */
export default class AuthModel {
    mailAddress = ""
    password = ""
    loggedIn = false
    loginErrorMessage = ""

    constructor(config) {
        this.config = config

        // do auto login if session still valid
        let that = this
        Axios.get(this.config.apiUrl +  "/users/is-authenticated",
            {withCredentials: true}
        ).then(function (response) {

            if (response.data === true) {
                that.setLoggedIn()
                navigate(MAIN_ROUTE)
            }
        })
        .catch(function (error) {
            console.log(error)
        })
    }

    onLogin() {
        let that = this
        Axios.post(this.config.apiUrl +  "/users/login",
            {
                mail_addr: this.mailAddress,
                password: this.password
            },{
                withCredentials: true
            }).then(function (response) {

            if (response.data === true) {

                that.setLoggedIn()
                navigate(MAIN_ROUTE)
            } else {
                that.loginErrorMessage = "Login fehlgeschlagen. E-Mail oder Passwort ist falsch."
            }
        }).catch(function (error) {
            console.log(error)
        })
    }

    onLogout() {
        let that = this
        Axios.get(this.config.apiUrl +  "/users/logout",
            {
                withCredentials: true
            }).then(function (response) {

                if (response.data === true) {
                    that.setLoggedOut()
                    navigate(LOGOUT_ROUTE)
                }
            })
            .catch(function (error) {
                console.log(error)
            })
    }

    // actions
    setMailAddress(mailAddress) {
        this.mailAddress = mailAddress
    }

    setPassword(password) {
        this.password = password
    }

    setLoggedIn() {
        this.password = ""
        this.loginErrorMessage = ""
        this.loggedIn = true
    }

    setLoggedOut() {
        this.loggedIn = false
    }
}

decorate(AuthModel, {
    mailAddress: observable,
    password: observable,
    loggedIn: observable,
    loginErrorMessage: observable,

    setMailAddress: action,
    setPassword: action,
    setLoggedIn: action,
    setLoggedOut: action,
})