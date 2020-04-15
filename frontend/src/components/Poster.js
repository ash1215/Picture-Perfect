import React, { Component } from 'react';

class Poster extends Component{
  state = {
    Image: this.props.poster
  }
  componentDidMount(){
      if(this.state.Image === "N/A"){
          this.setState({Image:"https://in.bmscdn.com/iedb/movies/images/website/poster/large/canyon-coaster-and-forest-adventure--combo-7d---et00016533-24-03-2017-19-06-06.jpg"})
      }
  }
  
  render() {
    return (
      <div>
        <img alt="Poster" src={this.state.Image} width="300" height="400"></img>
      </div>
    );
  }
}

export default Poster;
