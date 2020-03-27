import * as React from 'react';
import * as ReactDom from 'react-dom';
import {BrowserRouter as Router, Route, Switch} from 'react-router-dom';

import Login from './components/Login';
import Loading from './components/Loading'

ReactDom.render(
    <Router>
        <div>
            <Route path='/login' component={Login} />
            <Route path='/loading' component={Loading} />
        </div>
    </Router>,

    document.getElementById('main')
    );


// ReactDom.render(element, document.getElementById('main'));