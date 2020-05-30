import React from 'react';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import { Link } from 'react-router-dom';

import { useDispatch } from 'react-redux';
import { login } from '../actions';

export default function Login() {
  const [open, setOpen] = React.useState(false);
  const [email, setEmail] = React.useState('');
  const [password, setPassword] = React.useState('');
  const [isLogged, setIsLogged] = React.useState(false);
  const dispatch = useDispatch();
  const handleClickOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  const handleLogin = async() => {
    console.log(email,password)
    const info = {
        email,
        password
    }
    var message = ""
    console.log(JSON.stringify(info))
    var data = await fetch('http://192.168.29.57:8500/verifylogin',{
      method: "POST",
      body: JSON.stringify(info)
      })
      .then((res) => {
        return res.text();
      }
    ) 
    data = JSON.parse(data.replace(JSON.stringify(info),""))
    if(data.success){
      const user = {
        Name: data.message,
        Email: email
      }
      dispatch(login(user));
      handleClose();
    }
  }

  return (
    <div>
      <Button color="primary" onClick={handleClickOpen}>
        Login
      </Button>
      <Dialog open={open} onClose={handleClose} aria-labelledby="form-dialog-title">
        <DialogTitle id="form-dialog-title">Log In</DialogTitle>
        <DialogContent>
          
          <TextField
            autoFocus
            margin="dense"
            id="email"
            label="Email Address"
            type="email"
            fullWidth
            value={email}
            onChange = {e => setEmail(e.target.value)}
          />
          <TextField
            margin="dense"
            id="password"
            label="Password"
            type="password"
            fullWidth
            value={password}
            onChange = {e => setPassword(e.target.value)}
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose} color="primary">
            Cancel
          </Button>
          <Button onClick={handleLogin} style={{color:'red'}}>
            Login
          </Button>
        </DialogActions>
        <DialogContentText align='center' style={{paddingTop:'20px'}}>
            Don't have an account?<br/>
            <Button color="primary" onClick={handleClose}>
            <Link to="/register">Sign up</Link>
            </Button>!
        </DialogContentText>
      </Dialog>
    </div>
  );
}
