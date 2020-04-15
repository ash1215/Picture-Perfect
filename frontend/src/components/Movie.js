import React, { Component } from 'react';
import Poster from './Poster'

class Movie extends Component{
  state = {
    Movies : [],
    title: "",
    empty: true
  }
  
  render() {
    return (
      <div>
          <br/> <h2>{this.props.movie.Title}</h2> <br/>
          <Poster poster={this.props.movie.Poster}/>
          <div>
              <h3>
                  IMDb Rating: {this.props.movie.imdbRating}
              </h3>
          </div>
      </div>
    );
  }
}

export default Movie;
