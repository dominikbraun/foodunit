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

import {Link} from "@reach/router";
import {LOGIN_ROUTE} from "../../util/Routes";
import {inject, observer} from "mobx-react";
import React from 'react';

const LoggedIn = inject('auth')(({auth, children}) => {
    if (!auth.loggedIn)
    {
        // Todo: Better HTML / CSS for this
        return (
            <div>
                Du bist nicht angemeldet. Bitte <Link to={LOGIN_ROUTE}>einloggen</Link>.
            </div>
        );
    }

    return (
        <React.Fragment>
            {children}
        </React.Fragment>
    );
});

export default LoggedIn;