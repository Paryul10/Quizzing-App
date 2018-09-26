import React, { Component } from 'react';
import DeletePerson from './DeletePerson';
import ViewPeople from './ViewPeople';
import EditQuestion from './EditQuestion';
// import NewPerson from './NewPerson';
import Home from './Home';
import ViewQuizzes from './ViewQuizzes';
import ViewQuiz from './ViewQuiz';
import LeaderBoard from './LeaderBoard';
import CreateQuiz from './CreateQuiz'

import DeleteQuiz from './DeleteQuiz'

// import Logout from './Logout';

import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';

class Admin extends Component {
    render() {
        //const isLoggedIn = false;

        return (
            <div>
                <div className="App">
                    <header className="App-header">
                        <h1 className="App-title">Admin Panel</h1>
                    </header>
                </div>
                <br></br>
                        <br></br>
                {/* <Router> */}
                <div>
                    <div className="container-fluid">

                        {/* <li><Link to={'/'}>Home</Link></li> */}
                        {/* <li><Link to={'/NewPerson'}>Create Person</Link></li> */}
                        <div align="center"  class="wrapper cf">
                            <Link to={'/ViewQuizzes'}> <button type="button"> View Quizzes </button></Link>
                        </div>
                        <br></br>
                        <br></br>
                        <div align="center">
                            <Link to={'/ViewQuiz'}> <button type="button">View Quiz</button></Link>
                        </div>
                        <br></br>
                        <br></br>
                        <div align="center">
                            <Link to={'/LeaderBoard'}> <button type="button">Leader Board</button></Link>
                        </div>
                        <br></br>
                        <br></br>
                        <div align="center">
                            <Link to={'/EditQuestion'}> <button type="button">Edit Question</button></Link>
                        </div>
                        <br></br>
                        <br></br>
                        <div align="center">
                            <Link to={'/DeletePerson'}> <button type="button">Delete Person</button></Link>
                        </div>
                        <br></br>
                        <br></br>
                        <div align="center">
                            <Link to={'/ViewPeople'}> <button type="button">View People</button></Link>
                        </div>
                        <br></br>
                        <br></br>
                        <div align="center">
                            <Link to={'/CreateQuiz'}> <button type="button">Create Quiz</button></Link>
                        </div>
                        <br></br>
                        <br></br>
                        <div align="center">
                            <Link to={'/DeleteQuiz'}> <button type="button">Delete Quiz</button></Link>
                        </div>

                        {/* <li><Link to={'/'} onClick={this.handleLogout}>Logout</Link></li> */}
                    </div>
                    <Switch>
                        <Route exact path='/' component={Home} />
                        {/* <Route exact path='/NewPerson' component={NewPerson} /> */}
                        <Route exact path='/ViewQuizzes' component={ViewQuizzes} />
                        <Route exact path='/ViewQuiz' component={ViewQuiz} />
                        <Route exact path='/LeaderBoard' component={LeaderBoard} />
                        <Route exact path='/EditQuestion' component={EditQuestion} />
                        <Route exact path='/DeletePerson' component={DeletePerson} />
                        <Route exact path='/ViewPeople' component={ViewPeople} />
                        <Route exact path='/CreateQuiz' component={CreateQuiz} /> 
                        <Route exact path='/DeleteQuiz' component={DeleteQuiz} />
                        {/* <Route exact path='/Logout' component={Logout} />  */}
                    </Switch>
                </div>
                {/* </Router> */}
            </div>
        );
    }
}

export default Admin;
