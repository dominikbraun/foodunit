import {Link, navigate} from "@reach/router";
import {LOGIN_ROUTE} from "../../util/Routes";
import {inject, observer} from "mobx-react";
import React from 'react';

const LoggedIn = inject('auth')(observer(({auth, children}) => {
    if (!auth.loggedIn)
    {
        navigate(LOGIN_ROUTE);

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
}));

export default LoggedIn;