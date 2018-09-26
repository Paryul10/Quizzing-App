import React, { Component } from 'react';
import './ViewPeople.css';

class ViewQuiz extends Component {
  constructor() {
    super();
    this.state = {       //javascript object
      //data: []
      // id: this.props.match.params.id,
      items: [],
      formData: {
        answer1: false,
        answer2: false,
        answer3: false,
        answer4: false,
      },
      selectedQuestion: null,
      points: 0,
    }
    this.handleAChange = this.handleAChange.bind(this)
    this.handleBChange = this.handleBChange.bind(this)
    this.handleCChange = this.handleCChange.bind(this)
    this.handleDChange = this.handleDChange.bind(this)
    this.submitQuestion = this.submitQuestion.bind(this)
    this.submitQuiz = this.submitQuiz.bind(this)
  }

  // Lifecycle hook, runs after component has mounted onto the DOM structure
  componentDidMount() {
    console.log(this.props.match.params.id)
    sessionStorage.setItem("points", 0)
    this.getItems();
  }

  getItems() {
    const request = new Request('http://127.0.0.1:8080/viewquiz/' + this.props.match.params.id);
    fetch(request)
      .then(response => response.json())
      .then(response => this.setState({ 'items': response }));
    // .then(data => this.setState({data: data}));
  }

  handleAChange(e) {
    this.state.formData.answer1 = e.target.checked
    this.state.selectedQuestion = e.target.value
  }

  handleBChange(e) {
    this.state.formData.answer2 = e.target.checked
    this.state.selectedQuestion = e.target.value
  }

  handleCChange(e) {
    this.state.formData.answer3 = e.target.checked
    this.state.selectedQuestion = e.target.value
  }

  handleDChange(e) {
    this.state.formData.answer4 = e.target.checked
    this.state.selectedQuestion = e.target.value
  }

  submitQuestion(event) {
    console.log(this.state.formData)
    event.preventDefault();
    fetch('http://127.0.0.1:8080/evaluatequestion/' + sessionStorage.getItem("username") + '/' + this.state.selectedQuestion, {
      method: 'POST',
      body: JSON.stringify(this.state.formData),
    })
      .then(response => {
        if (response.status == 200){
          console.log("correct")
          this.state.points += 5
          sessionStorage.setItem("points", this.state.points)
        }
        else if (response.status == 201)
          console.log("wrong")
        this.state.selectedQuestion = null;
        this.state.formData.answer1 = false;
        this.state.formData.answer2 = false;
        this.state.formData.answer3 = false;
        this.state.formData.answer4 = false;
        console.log(sessionStorage.getItem("points"))
      });
  }

  submitQuiz() {
    this.props.history.push('/Points')
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
                  <input type="checkbox" value={item.id} onChange={this.handleAChange} />
                  <label> A) {item.option1} </label>
                </div>

                <div>
                  <input type="checkbox" value={item.id} onChange={this.handleBChange} />
                  <label> B) {item.option2} </label>
                </div>

                <div>
                  <input type="checkbox" value={item.id} onChange={this.handleCChange} />
                  <label> C) {item.option3} </label>
                </div>

                <div>
                  <input type="checkbox" value={item.id} onChange={this.handleDChange} />
                  <label> D) {item.option4} </label>
                </div>
                {/* if {item.type} {
                  <h6> There May be Multiple Correct Answers</h6>
                } */}
                <h6>{item.type}</h6>
                <input type="submit" value="Submit Question" className="btn btn-default" onClick={this.submitQuestion}></input>
              </div>
            )
          }, this
          )}
          <div>
            <h6> 0 :-> Single Answer Question </h6>
            <h6> 1 :-> Multiple Answer Question </h6>
          </div>
          <button type="button" onClick={this.submitQuiz} className="btn btn-default">Done</button>

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
