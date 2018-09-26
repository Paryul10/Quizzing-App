import React, { Component } from 'react';

class Points extends Component {
    render() {
        return (
            <div>
                <header className="App-header">
                    <h1 className="App-title" align = "center">Your Points</h1>
                </header>
                <h3 align = "center"><b>{sessionStorage.getItem("points")}</b></h3>
            </div>
        );
    }
}

export default Points;