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

import Axios from "axios"
import {checkId} from "../util/Checks"

export default class UserLoader {

    constructor(config) {
        this.config = config
    }

    loadUser(id) {
        checkId(id)

        return Axios.get(this.config.apiUrl +  "/users/" + id,
            {withCredentials: true}
        ).then(function (response) {
            return response.data
        }).catch(function (error) {
            console.log(error)
        })
    }
}