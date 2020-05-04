import React from 'react';
import NavBar from './Components/NavBar'
import SearchResults from './Components/SearchResults'
import MoviePage from './Components/MoviePage'
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from "react-router-dom";
function App() {
  return (
    <Router>
      <div>
        <NavBar></NavBar>
        <Route path="/search" component={SearchResults}/>
        <Route path="/movie:id" component={MoviePage} />
      </div>
    </Router>
  );
}

export default App;
