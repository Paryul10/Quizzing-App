import React, { Component } from 'react';
import './NewPerson.css';


class GenrePoints extends Component {
    constructor() {
      super();
      this.state = {
        formData: {
          genrename: "",
          //lastName: "",
          //city: "",
        },
        submitted: false,
      }
      this.handleGenreChange = this.handleGenreChange.bind(this);
      this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleSubmit (event) {
        event.preventDefault();
        console.log(this.state.formData.genrename)
        this.props.history.push('/GenreLeaderBoard/' + this.state.formData.genrename)
      }

      handleGenreChange(event) {
        this.state.formData.genrename = event.target.value;
      }

      render() {

        return (
          <div className="App">
            <header className="App-header">
              <h1 className="App-title">Which Genre??</h1>
            </header>
            <br/><br/>
            <div className="formContainer">
              <form onSubmit={this.handleSubmit}>
                <div className="form-group">
                    <label>Please Enter Genre Name</label>
                    <input type="text" className="form-control" value={this.state.genrename} onChange={this.handleGenreChange}/>
                </div>
                {/* <div className="form-group">
                    <label>Last Name</label>
                    <input type="text" className="form-control" value={this.state.lastName} onChange={this.handleLChange}/>
                </div>
                <div className="form-group">
                    <label>City</label>
                    <input type="text" className="form-control" value={this.state.city} onChange={this.handleCChange}/>
                </div> */}
                    <button type="submit" className="btn btn-default">Submit</button>
              </form>
            </div>
    
            {/* {this.state.submitted &&
              <div>
                <h2>
                  New person successfully added.
                </h2>
                 This has been printed using conditional rendering.
              </div>
            } */}
    
          </div>
        );
      }
    }
    
    export default GenrePoints;