import React, { Component } from 'react';
import DeletePerson from './DeletePerson';
import ViewPeople from './ViewPeople';
import EditPerson from './EditPerson';
// import NewPerson from './NewPerson';
import Home from './Home';
import ViewQuizzes from './ViewQuizzes'

// import Logout from './Logout';

import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';

class App extends Component {

  constructor() {
    super();
    this.state = {
      formData: {
        firstName: "",
        lastName: "",
        userName: "",
        email: "",
        password: "",
      },

      // formData2: {
      //   uName: "",
      //   lpass: "",
      // }
    }
    this.handleFnsuChange = this.handleFnsuChange.bind(this);
    this.handleLnsuChange = this.handleLnsuChange.bind(this);
    this.handleUnsuChange = this.handleUnsuChange.bind(this);
    this.handleEmsuChange = this.handleEmsuChange.bind(this);
    this.handlePasuChange = this.handlePasuChange.bind(this);
    this.handleSignup = this.handleSignup.bind(this);
    this.handleLogin = this.handleLogin.bind(this);
    this.handleLogout = this.handleLogout.bind(this);


    // this.handleUnliChange = this.handleUnliChange.bind(this);
    // this.handlePaliChange = this.handlePaliChange.bind(this);


  }

  handleSignup(event) {
    event.preventDefault();
    fetch('http://localhost:8080/signup/', {
      method: 'POST',
      body: JSON.stringify(this.state.formData),
    })
      .then(response => {
        if (response.status >= 200 && response.status < 300) {
          sessionStorage.setItem("isLoggedIn", true)
          window.location.reload()
        }
      })
  }


  handleLogin(event) {
    console.log("paryul")
    event.preventDefault();
    fetch('http://localhost:8080/login/', {
      method: 'POST',
      body: JSON.stringify(this.state.formData),
    })
      .then(response => {
        if (response.status >= 200 && response.status < 300) {
          sessionStorage.setItem("isLoggedIn", true)
          window.location.reload()
        }
      })
  }

  handleFnsuChange(event) {
    this.state.formData.firstName = event.target.value;
  }
  handleLnsuChange(event) {
    this.state.formData.lastName = event.target.value;
  }
  handleUnsuChange(event) {
    this.state.formData.userName = event.target.value;
  }
  handleEmsuChange(event) {
    this.state.formData.email = event.target.value;
  }
  handlePasuChange(event) {
    this.state.formData.password = event.target.value;
  }


  handleLogout(event) {
    // event.preventDefault();
    sessionStorage.clear()
    window.location.reload()

  }

  render() {
    //const isLoggedIn = false;
    if (sessionStorage.getItem("isLoggedIn")) {
      return (
        <div>
          <Router>
            <div>
              <nav className="navbar navbar-default">
                <div className="container-fluid">
                  <div className="navbar-header">
                    <Link className="navbar-brand" to={'/'}>Quizzer</Link>
                  </div>
                  <ul className="nav navbar-nav">
                    <li><Link to={'/'}>Home</Link></li>
                    {/* <li><Link to={'/NewPerson'}>Create Person</Link></li> */}
                    <li><Link to={'/EditPerson'}>Edit Person</Link></li>
                    <li><Link to={'/DeletePerson'}>Delete Person</Link></li>
                    <li><Link to={'/ViewPeople'}>View People</Link></li>
                    <li><Link to={'/ViewQuizzes'}>View Quizzes</Link></li>

                    <li><Link to={'/'} onClick={this.handleLogout}>Logout</Link></li>
                  </ul>
                </div>
              </nav>
              <Switch>
                <Route exact path='/' component={Home} />
                {/* <Route exact path='/NewPerson' component={NewPerson} /> */}
                <Route exact path='/EditPerson' component={EditPerson} />
                <Route exact path='/DeletePerson' component={DeletePerson} />
                <Route exact path='/ViewPeople' component={ViewPeople} />
                <Route exact path='/ViewQuizzes' component={ViewQuizzes} />                
                {/* <Route exact path='/Logout' component={Logout} />  */}
              </Switch>
            </div>
          </Router>
        </div>
      );
    }
    return (
      <div>
        <div className="container">
          <div className="row">
            <div className="span12">
              <div className="" id="loginModal">
                <div className="modal-header">
                  <button type="button" className="close" data-dismiss="modal" aria-hidden="true">×</button>
                  <h3>Have an Account?</h3>
                </div>
                <div className="modal-body">
                  <div className="well">
                    <ul className="nav nav-tabs">
                      <li className="active"><a href="#login" data-toggle="tab">Login</a></li>
                      <li><a href="#create" data-toggle="tab">Create Account</a></li>
                    </ul>
                    <div id="myTabContent" className="tab-content">
                      <div className="tab-pane active in" id="login">
                        <form className="form-horizontal" onSubmit={this.handleLogin} >
                          <fieldset>
                            <div id="legend">
                              <legend className="">Login</legend>
                            </div>
                            <div className="control-group">
                              <label className="control-label">Username</label>
                              <div className="controls">
                                <input type="text" className="input-xlarge" value={this.state.userName} onChange={this.handleUnsuChange}></input>
                              </div>
                            </div>

                            <div className="control-group">
                              <label className="control-label">Password</label>
                              <div className="controls">
                                <input type="password" id="password" className="input-xlarge" value={this.state.password} onChange={this.handlePasuChange}></input>
                              </div>
                            </div>


                            <div className="control-group">
                              <div className="controls">
                                <button className="btn btn-success">Login</button>
                              </div>
                            </div>
                          </fieldset>
                        </form>
                      </div>
                      <div className="tab-pane fade" id="create">
                        <form id="tab" onSubmit={this.handleSignup}>
                          <div className="control-group">
                            <label className="control-label">First Name</label>
                            <div className="controls">
                              <input type="text" className="input-xlarge" value={this.state.firstName} onChange={this.handleFnsuChange}></input>
                            </div>
                          </div>
                          <div className="control-group">
                            <label className="control-label">Last Name</label>
                            <div className="controls">
                              <input type="text" className="input-xlarge" value={this.state.lastName} onChange={this.handleLnsuChange}></input>
                            </div>
                          </div>
                          <div className="control-group">
                            <label className="control-label">Username</label>
                            <div className="controls">
                              <input type="text" className="input-xlarge" value={this.state.userName} onChange={this.handleUnsuChange}></input>
                            </div>
                          </div>
                          <div className="control-group">
                            <label className="control-label">Email</label>
                            <div className="controls">
                              <input type="text" className="input-xlarge" value={this.state.email} onChange={this.handleEmsuChange}></input>
                            </div>
                          </div>
                          <div className="control-group">
                            <label className="control-label">Password</label>
                            <div className="controls">
                              <input type="text" className="input-xlarge" value={this.state.password} onChange={this.handlePasuChange}></input>
                            </div>
                          </div>


                          <div>
                            <button className="btn btn-primary">Create Account</button>
                          </div>
                        </form>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    );
  }
}
export default App;
