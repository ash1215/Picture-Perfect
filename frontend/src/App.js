import React from 'react';
import NavBar from './Components/NavBar'
import SearchResults from './Components/SearchResults'
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
      </div>
    </Router>
  );
}

export default App;
