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

import {configure, decorate} from 'mobx'

// enforce the strict mode for actions -> e.g. no state modifying inside of promise without action decorator https://www.leighhalliday.com/mobx-async-actions
configure({ enforceActions: "observed" })

export default class FoodUnitStore {

    constructor(config) {
        this.config = config
    }
}

decorate(FoodUnitStore, {
})