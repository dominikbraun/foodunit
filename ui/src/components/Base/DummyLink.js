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

import React from "react"

/**
 * DummyLink provides a <a href component which reacts like a normal link, but it has no action.
 * So you have to provide a onClick property.
 *
 * @param props are just passed to the link
 * @returns {React.Component}
 * @constructor
 */
export default function DummyLink(props) {
    return (
        <a href="#0" {...props}>
            {props.children}
        </a>
    )
}