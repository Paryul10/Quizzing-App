import React, { Component } from 'react';
import './ViewPeople.css';

class ViewQuiz extends Component {
  constructor() {
    super();
    this.state = {       //javascript object
      //data: []
      'items': []
    }
  }

  // Lifecycle hook, runs after component has mounted onto the DOM structure
  componentDidMount() {
    this.getItems();
  }

  getItems() {
    const request = new Request('http://127.0.0.1:8080/viewquiz/' + 9);
    fetch(request)
      .then(response => response.json())
      .then(response => this.setState({ 'items': response }));
    // .then(data => this.setState({data: data}));
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <h1 className="App-title">View All Questions of This Quiz</h1>
        </header>

        <ul>
          {this.state.items.map(function (item, index) {
            return (
              <div>

                <h3> {index + 1}.{item.question} </h3>
                <div>
                  <input type="radio" />
                  <label> A) {item.option1} </label>
                </div>

                <div>
                  <input type="radio" />
                  <label> B) {item.option2} </label>
                </div>

                <div>
                  <input type="radio" />
                  <label> C) {item.option3} </label>
                </div>

                <div>
                  <input type="radio" />
                  <label> D) {item.option4} </label>
                </div>
                {/* if {item.type} {
                  <h6> There May be Multiple Correct Answers</h6>
                } */}
                <h6>{item.type}</h6>
                <input type="submit" value="Submit Question" className="btn btn-default" ></input>
              </div>
            )
          }
          )}
          <div>
            <h6> 0 :-> Single Answer Question </h6>
            <h6> 1 :-> Multiple Answer Question </h6>
          </div>

        </ul>

        {/* <table className="table-hover">
          <thead>
            <tr>
              <th>ID</th>
              <th>First Name</th>
              <th>Last Name</th>
              <th>Username</th>
              <th>Email</th>
              <th>Admin</th>
              <th>Totalpoints</th>
            </tr>
          </thead>
          {/* <tbody>{this.state.data.map(function(item, key) {
               return (
                  <tr key = {key}>
                      <td>{item.id}</td>
                      <td>{item.firstname}</td>
                      <td>{item.lastname}</td>
                      <td>{item.username}</td>
                      <td>{item.email}</td>
                      <td>{item.admin}</td>
                      <td>{item.totalpoints}</td>
                  </tr>
                )
             })}
          </tbody> */}
      </div>
    );
  }
}

export default ViewQuiz;
