import React, {useState,useEffect} from 'react';
import { Typography } from '@material-ui/core';
import Grid from '@material-ui/core/Grid';

function MovieItem({movieDetails}){
    const [Item, setItem] = useState({})

    useEffect(()=> {
        if(movieDetails.Poster === "N/A"){
            movieDetails.Poster = "https://in.bmscdn.com/iedb/movies/images/website/poster/large/canyon-coaster-and-forest-adventure--combo-7d---et00016533-24-03-2017-19-06-06.jpg"
        }
        setItem(movieDetails);
    },[movieDetails])
    
    return(
        <div style={{backgroundColor:'#212121'}}>
            <Grid container spacing={3}>
                <Grid item sm={2} xs={6}>
                    <img src={Item.Poster} style={{maxWidth:'100%'}}></img>
                </Grid>
                <Grid item sm={10} xs={6}>
                    <Typography style={{color:'#fff'}}>
                        {Item.Title} ({Item.Year})
                    </Typography>
                    <Typography style={{color:'#fff'}}>
                        Language : {Item.Language}
                    </Typography>
                    <Typography style={{color:'#fff'}}>
                        IMDb Rating : {Item.imdbRating}/10
                    </Typography>
                </Grid>
            </Grid>
        </div>
    )
}

export default MovieItem;