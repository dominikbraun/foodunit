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

import React from 'react';
import {Link} from "@reach/router";

export default function(props) {
    return (
        <div className="row m-0 align-items-center h-100">

            <div className="col-12 col-md-3 col-xl-4"></div>

            <div className="col-12 col-md-6 col-xl-4 px-1 px-md-4 px-xl-5 py-0">
                <div className="mx-0 mx-xl-2 my-5">
                    <div className="bg-dark text-white text-hand text-center text-logo px-4 py-4">FoodUnit&nbsp;</div>
                    <div className="bg-gradient px-3 py-4">
                        <h4 className="text-lg text-dark text-strong text-center py-5">Vielen Dank &amp; bis bald!</h4>
                        <p className="text-center text-sm">&copy; 2019 FoodUnit &mdash; <a
                            className="text-dark link-underlined" target="_blank"
                            href="https://github.com/dominikbraun/foodunit">dominikbraun/foodunit</a></p>
                    </div>

                    <div className="text-md text-center bg-white my-4 py-3">
                        Noch Hunger?
                        <Link to="/login" className="text-dark2 text-strong link-underlined2">Jetzt einloggen</Link>
                    </div>
                </div>
            </div>

            <div className="col-12 col-md-3 col-xl-4"></div>

        </div>
    );
}