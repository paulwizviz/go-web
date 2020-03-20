// Copyright 2020 Paul Sitoh
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React from 'react';
import PropTypes from 'prop-types';

import { makeStyles } from '@material-ui/core/styles';

import { Grid, TextField, Button } from '@material-ui/core';

const useStyles = makeStyles(theme => ({
    margin: {
        margin: theme.spacing.unit * 2,
    },
    padding: {
        padding: theme.spacing.unit
    }
}));

const Auth = (props) => {
   
    const classes = useStyles();
    return (
        <div className={classes.root}>
            <Grid container spacing={8} alignItems="flex-end">
                <Grid item md={true} sm={true} xs={true}>
                    <TextField id="username" label="Username" type="email" fullWidth autoFocus required />
                </Grid>
            </Grid>
            <Grid container spacing={8} alignItems="flex-end">
                <Grid item md={true} sm={true} xs={true}>
                    <TextField id="username" label="Password" type="password" fullWidth required />
                </Grid>
            </Grid>
            <Grid container justify="center" style={{ marginTop: '10px' }}>
                <Button onClick={
                    async () => {
                        await props.authenticate();
                        props.history.push('/Dashboard');
                    }
                }
                variant="outlined" 
                color="primary" 
                style={{ textTransform: 'none' }}>Login</Button>
            </Grid>
            <Grid item xs={12}>
                {
                    JSON.stringify(props.user)
                }
            </Grid>
        </div>
    );
};

Auth.propTypes = {
    history: PropTypes.object.isRequired,
    user: PropTypes.object.isRequired,
    authenticate: PropTypes.func.isRequired
};

export default Auth;