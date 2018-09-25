import React, { Component } from 'react';
import './ViewPeople.css';

class ViewQuizzes extends Component {
  constructor() {
    super();
    this.state = {       //javascript object
      data: [],
      selectedOption: null,
      submitted: false,
    }
    this.handleOptionChange = this.handleOptionChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  // Lifecycle hook, runs after component has mounted onto the DOM structure
  componentDidMount() {
    const request = new Request('http://127.0.0.1:8080/viewquizzes/');
    fetch(request)
      .then(response => response.json())
      .then(data => this.setState({ data: data }));
  }

  handleOptionChange(event) {
    this.setState({ selectedOption: event.target.value });
  }

  handleSubmit(event) {
    event.preventDefault();
    fetch('http://127.0.0.1:8080/viewquiz/' + this.state.selectedOption, {
      method: 'GET',
    })

      .then(response => {
        if (response.status >= 200 && response.status < 300){
            this.setState({ submitted: true });
        }
        this.setState({ selectedOption: null });

      });
  }


  render() {
    return (
      <div className="App">
        <header className="App-header">
          <h1 className="App-title">View All Quizzes</h1>
        </header>

        <form onSubmit={this.handleSubmit}>
          <table className="table-hover">
            <thead>
              <tr>
                <th>ID</th>
                <th>Genre Name</th>
                <th>Quiz Number</th>
              </tr>
            </thead>
            <tbody>{this.state.data.map(function (item, key) {
              return (
                <tr key={key}>
                  <td>{item.id}</td>
                  <td>{item.genrename}</td>
                  <td>{item.quiznumber}</td>
                  <td><input type="radio" value={item.id} checked={this.state.selectedOption == item.id} onChange={this.handleOptionChange}></input></td>
                </tr>
              )
            }, this)}
            </tbody>
          </table>
          {/* <input type="submit" value="Play Quiz" className="btn btn-default"> <Link to={'/EditPerson'}>Edit Person</Link></input> */}
        </form>
      </div>
    );
  }
}

export default ViewQuizzes;
