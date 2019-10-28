import * as React from "react";
import "../style/Login"
import * as Router from 'react-router-dom'

function GetUser() {
    this.setState(user => {
        document.getElementsByName('user');
    })
    return this.user;
}

function GetPass() {

    this.setState(pass => {
        document.getElementsByName('pass');
    })
    return this.pass
}

class Login extends React.Component {
    render (){
        return (
            <div>
                <form action='/loading/'>
                <input placeholder="Username/Email" type="text" id="user" />
                <input placeholder="Password" type="password" id="pass" />

                <button id="login-button" type="submit">Click to be successful</button>
                </form>
            </div>
        );
    }
}

export default Login;