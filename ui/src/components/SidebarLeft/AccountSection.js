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

import {inject, observer} from "mobx-react"
import React from 'react'
import DummyLink from "../Base/DummyLink"

class AccountSection extends React.Component {

    constructor(props) {
        super(props)
        this.auth = props.auth
    }

    handleLogout = () => {
        this.auth.onLogout()
    }

    render() {
        return (
            <div>
                <p className="text-light text-center text-strong text-sm mb-0 p-3">Mein Konto</p>
                <div className="nav flex-column nav-pills side-nav">
                    <div className="nav-link text-md text-success px-2 px-xl-3 py-3 rounded-0"><i className="fas fa-medal ml-1 mr-3"/>Gesamt: 1200</div>
                    <a className="nav-link text-md px-2 px-xl-3 py-3 rounded-0" href="#"><i className="fas fa-user-circle ml-1 mr-3"/>Konto verwalten</a>
                    <DummyLink className="nav-link text-md px-2 px-xl-3 py-3 rounded-0"
                               onClick={this.handleLogout}>
                            <i className="fas fa-sign-out-alt ml-1 mr-3"/>
                            Ausloggen
                    </DummyLink>
                </div>
            </div>
        )
    }
}

export default inject('auth')(observer(AccountSection))