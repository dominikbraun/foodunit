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

import ReactDOM from 'react-dom'
import './foodunit3--custom.css'
import App from './components/App'
import * as serviceWorker from './serviceWorker'
import React from "react"
import Axios from "axios"

const root = document.getElementById('root')
const url = root.getAttribute('config-url')

/**
 * loads config and starts the app with it
 * @returns {Promise<void>}
 */
async function run() {
    let config = await Axios.get(url).then(response => {
        return response.data
    }).catch(error => {
        console.log("error loading config " + error)
    })

    if (config) {
        ReactDOM.render(<App config={config}/>, document.getElementById('root'))
    } else {
        ReactDOM.render(<div>No config</div>, document.getElementById('root'))
    }

    // If you want your app to work offline and load faster, you can change
    // unregister() to register() below. Note this comes with some pitfalls.
    // Learn more about service workers: https://bit.ly/CRA-PWA
    serviceWorker.unregister()
}

run().then(() => {
    console.log("started")
})