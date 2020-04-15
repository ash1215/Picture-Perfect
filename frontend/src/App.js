import React, { Component } from 'react';
import './App.css';
import axios from 'axios'
import MovieList from './components/MovieList'

class App extends Component{
  state = {
    Movies : [],
    title: "",
    empty: true
  }
  handleChange = event => {
    this.setState({ title: event.target.value });
  }
  handleSubmit = event => {
    event.preventDefault();

    const title = this.state.title

    axios.post('http://192.168.43.56:8500/api',title)
      .then((res) => {
        console.log(res.data.replace(title,""))
        this.setState({Movies: JSON.parse(res.data.replace(title,""))})
      })
      .catch(console.log)
  }
  
  render() {
    return (
      <div className="App">
          <div>
          <h1>Picture Perfect</h1>
            <form onSubmit={this.handleSubmit}>
              <label>
                Movie Title: 
                <input type="text" name="name" onChange={this.handleChange}/>
              </label>
              <button type="submit">Search</button>
            </form>
          </div>
          <MovieList movies = {this.state.Movies} />
      </div>
    );
  }
}

export default App;
