import React, { Component } from 'react';
import './NewPerson.css';

class CreateQuiz extends Component {
  constructor() {
    super();
    this.state = {
      formData: {
        genreName: "",
        quizNumber: "",
    },
      submitted: false,
    }
    this.handlegenreChange = this.handlegenreChange.bind(this);
    this.handlenumChange = this.handlenumChange.bind(this);
    // this.handleCChange = this.handleCChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleSubmit (event) {
    event.preventDefault();
    fetch('http://localhost:8080/createquiz/', {
     method: 'POST',
     body: JSON.stringify(this.state.formData),
   })
      .then(response => {
        if(response.status >= 200 && response.status < 300)
          this.setState({submitted: true});
      });
  }

  handlegenreChange(event) {
    this.state.formData.genreName = event.target.value;
  }
  handlenumChange(event) {
    this.state.formData.quizNumber = event.target.value;
  }
  

  render() {

    return (
      <div className="App">
        <header className="App-header">
          <h1 className="App-title">Create a New Quiz</h1>
        </header>
        <br/><br/>
        <div className="formContainer">
          <form onSubmit={this.handleSubmit}>
            <div className="form-group">
                <label>Genre Name</label>
                <input type="text" className="form-control" value={this.state.genreName} onChange={this.handlegenreChange}/>
            </div>
            <div className="form-group">
                <label>Quiz Number</label>
                <input type="text" className="form-control" value={this.state.quizNumber} onChange={this.handlenumChange}/>
            </div>
            {/* <div className="form-group">
                <label>City</label>
                <input type="text" className="form-control" value={this.state.city} onChange={this.handleCChange}/>
            </div> */}
                <button type="submit" className="btn btn-default">Submit</button>
          </form>
        </div>

        {this.state.submitted &&
          <div>
            <h2>
              New quiz successfully added.
            </h2>
             {/* This has been printed using conditional rendering. */}
          </div>
        }

      </div>
    );
  }
}

export default CreateQuiz;
