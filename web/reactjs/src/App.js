import React from 'react';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';
import axios from 'axios';
import Button from '@material-ui/core/Button';

class App extends React.Component {
    render() {
        return (
            <div>
                <Button onClick={async () => {await this.props.fetchUser();}}>Fetch User</Button><br/>
                {
                    JSON.stringify(this.props.users)
                }
            </div>
        );
    }
}

const mapStateToProps = state => {
    return {
        users: state.userReducer.users,
    };
};

const mapDispatchToProps = dispatch => {
    return {
        fetchUser: () => dispatch({type: 'FETCH_USERS', payload: axios.get('http://localhost:9000/users')})
    };
};

App.propTypes = {
    users: PropTypes.object,
    fetchUser: PropTypes.func
};

export default connect(mapStateToProps, mapDispatchToProps)(App);