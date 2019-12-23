// Copyright 2019 Paul Sitoh
// SPDX-License-Identifier: Apache-2.0

import React from 'react';
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';

import { makeStyles } from '@material-ui/core/styles';

const useStyles = makeStyles(theme => ({
    root: {
        padding: theme.spacing(4)
    }
}));

const Dashboard = () => {

    const classes = useStyles();

    return (
        <div className={classes.root}>
            <Grid container spacing={1}>
                <Grid item lg={3} sm={6} xl={3} xs={12}>
                    <Paper className={classes.paper}>
                        <h1>Hello1</h1>
                    </Paper>
                </Grid>
                <Grid item xs={12}>
                    <Paper className={classes.paper}>
                        <h1>Hello</h1>
                    </Paper>
                </Grid>
            </Grid>
        </div>
    );
};
  
export default Dashboard;
