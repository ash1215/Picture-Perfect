import React from 'react';
import TextField from '@material-ui/core/TextField';
import { makeStyles } from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';
import Fab from '@material-ui/core/Fab';
const useStyles = makeStyles((theme) => ({
  root: {
    '& .MuiTextField-root': {
      margin: theme.spacing(1),
      width: 200,
    },
  },
}));

function MaterialUIForm() {
  const classes = useStyles();

  return (
    <form className={classes.root} noValidate autoComplete="off">
      <div>
        <TextField error id="standard-error" label="Name" defaultValue="Hello World" />
        <TextField
          error
          id="standard-error-helper-text"
          label="Error"
          defaultValue="Hello World"
          helperText="Incorrect entry."
        />
      </div>
      <div>
        <TextField
          error
          id="filled-error"
          label="Error"
          defaultValue="Hello World"
          variant="filled"
        />
        <TextField
          error
          id="filled-error-helper-text"
          label="Error"
          defaultValue="Hello World"
          helperText="Incorrect entry."
          variant="filled"
        />
      </div>
      <div>
        <TextField
          error
          id="outlined-error"
          label="Error"
          defaultValue="Hello World"
          variant="outlined"
        />
        <TextField
          error
          id="outlined-error-helper-text"
          label="Error"
          defaultValue="Hello World"
          helperText="Incorrect entry."
          variant="outlined"
        />
      </div>
      <div>
      <Fab variant="extended">
        Sign Up
      </Fab>
      </div>
      
    </form>
  );
}

export default function RegisterPage() {
    return(
        <div style={{paddingTop:'20px'}}>
            
            <Grid container spacing={0} spacing={4}>
                <Grid item md={2} xs={0}></Grid>
                <Grid item md={8} xs={12}>
                    <h1 style={{color:'white'}}>Register</h1>
                </Grid>
                <Grid item md={2} xs={0}></Grid>
                <Grid item md={2} xs={0}></Grid>
                <Grid item md={8} xs={12} alignContent='center'>
                    <MaterialUIForm />
                </Grid>
                
            </Grid>
        </div>
    )
}