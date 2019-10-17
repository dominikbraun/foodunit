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

export default class AuthModel {
    mailAddress = "";
    password = "";
    loggedIn = false;

    constructor() {
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
            },
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

    changeMailAddress(mailAddress) {
        this.mailAddress = mailAddress;
    }

    changePassword(password) {
        this.password = password;
    }
}

function loggedIn(that) {
    that.password = "";
    that.loggedIn = true;
}

decorate(AuthModel, {
    mailAddress: observable,
    password: observable,
    loggedIn: observable,

    login: action,
    changeMailAddress: action,
    changePassword: action,
});