// Copyright 2019 Paul Sitoh
// SPDX-License-Identifier: Apache-2.0

import React from 'react';
import { Switch, Redirect } from 'react-router-dom';

import {
    RouteWithLayout,
    Main as MainLayout, 
    Minimal as MinimalLayout,
    Dashboard as DashboardView,
    NotFound as NotFoundView
} from './components';

const Routes = () => {
    return (
        <Switch>
            <Redirect exact from="/" to="/dashboard"/>
            <RouteWithLayout component={DashboardView} exact layout={MainLayout} path="/dashboard"/>
            <RouteWithLayout component={NotFoundView} exact layout={MinimalLayout} path="/not-found"/>
            <Redirect to="/not-found" />
        </Switch>
    );
};

export default Routes;
