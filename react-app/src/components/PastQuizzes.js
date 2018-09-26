import React, { Component } from 'react';
import './ViewPeople.css';

class PastQuizzes extends Component {
  constructor() {
    super();
    this.state = {       //javascript object
      data: [],
    //   selectedOption: null,
    //   submitted: false,
    //   flag : false,
    }
    // this.handleOptionChange = this.handleOptionChange.bind(this);
    // this.handleSubmit = this.handleSubmit.bind(this);
  }

  // Lifecycle hook, runs after component has mounted onto the DOM structure
  componentDidMount() {
    const request = new Request('http://127.0.0.1:8080/pastquizzes/' + "zt");
    fetch(request)
      .then(response => response.json())
      //console.log(response.json())
      .then(data => this.setState({ data: data }));
      //console.log('data')
      // console.log(this.state.data.genrename)
  }

//   handleOptionChange(event) {
//     this.setState({ selectedOption: event.target.value });
//   }

//   handleSubmit(event) {
//     event.preventDefault();
//     console.log(this.state.selectedOption)
//     this.props.history.push('/ViewQuiz/' + this.state.selectedOption)
//     // fetch('http://127.0.0.1:8080/viewquiz/' + this.state.selectedOption, {
//     //   method: 'GET',
//     // })

//     //   .then(response => {
//     //     if (response.status >= 200 && response.status < 300){
//     //         this.setState({ submitted: true });
//     //     }
//     //     this.setState({ selectedOption: null });

//     //   });
//   }


  render() {
    return (
      <div>
        <div className="App">
          <header className="App-header">
            <h1 className="App-title">View All Past Quizzes</h1>
          </header>

          <form >
            <table className="table-hover">
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Username</th>
                  <th>Genre Name</th>
                  <th>Quiz Number</th>
                  <th>Score</th>
                </tr>
              </thead>
              <tbody>{this.state.data.map(function (item, key) {
                return (
                  <tr key={key}>
                    <td>{item.id}</td>
                    <td>{item.username}</td>
                    <td>{item.genrename}</td>
                    <td>{item.quiznumber}</td>
                    {/* <td> chutiya</td> */}
                    <td>{item.points}</td>
                    {/* <td><input type="radio" value={item.id} checked={this.state.selectedOption == item.id} onChange={this.handleOptionChange}></input></td> */}
                  </tr>
                )
              })}
              </tbody>
            </table>
            {/* {<input type="submit" value="Play Quiz" className="btn btn-default"></input>} */}
          </form>
        </div>
        {/* <Switch>
          <Route exact path='/ViewQuiz' component={ViewQuiz} />
        </Switch> */}
      </div>
    );
  }
}

export default PastQuizzes;
