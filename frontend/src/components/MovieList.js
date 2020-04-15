import React, { Component } from 'react';
import Movie from './Movie'

class MovieList extends Component{
  state = {
    Movies : [],
    title: "",
    empty: true
  }
  
  render() {
    return (
      this.props.movies.map((movie=>(
          <Movie key={movie.imdbID} movie = {movie} />
      )))
    );
  }
}

export default MovieList;
