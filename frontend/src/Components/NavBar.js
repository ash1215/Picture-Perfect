import React from 'react';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import { createMuiTheme } from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';
import logo from './logowhite.png';
import { Link } from 'react-router-dom';
import SearchBar from './SearchBar';
import Login from './Login';

import { useSelector, useDispatch } from 'react-redux';
import { login, logout, getUserInfo } from '../actions';

function LogTab({logInfo,userInfo}) {
  const dispatch = useDispatch();
  const user1 = {
    Name: 'Ashwini',
    ID: '123'
  }
  console.log(logInfo)
  if(logInfo == true) {
    return (
      <Link to="/profile" onClick={() => dispatch(logout())}>
        <Typography align='left'>{userInfo.Name}</Typography>
      </Link>
    )
  }
  else {
    return (
      <Login/>
    )
  }
}

export default function NavBar() {

  const logInfo = useSelector(state => state.logInfo);
  const userInfo = useSelector(state => state.userInfo);
  const theme = createMuiTheme({
      palette: {
          primary: {
              light: '#484848',
              main: '#212121',
              dark: '#000000',
              contrastText: '#fff',
          },
          secondary: {
              light: '#484848',
              main: '#212121',
              dark: '#000000',
              contrastText: '#fff',
          },
      }
  })

  return (
    <div>
      <AppBar position="static" style={{ background: theme.palette.primary.dark }}>
        <Toolbar>
          <Grid container spacing={0}>
            <Grid item xs={4}>
              <div style={{height:'50px'}}>
                <Link to="/" exact>
                  <img src={logo} alt="logo" style={{maxHeight: '100%'}}></img>
                </Link>
              </div>
            </Grid>
            <Grid item xs={4}>
              <SearchBar theme={theme}/>
            </Grid>
            <Grid item xs={1}></Grid>
            <Grid item xs={1}>
              <Link to="/cineplex">
                <Typography align='left'>CINEPLEX</Typography>
              </Link>
            </Grid>
            <Grid item xs={1}>
              <Link to="/movies">
              <Typography align='left'>MOVIES</Typography>
              </Link>
            </Grid>
            <Grid item xs={1}>
              <LogTab logInfo = {logInfo} userInfo = {userInfo}/>
            </Grid>
          </Grid>
        </Toolbar>
      </AppBar>
    </div>
  );
}