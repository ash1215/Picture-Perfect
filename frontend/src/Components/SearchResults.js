import React, {useState,useEffect} from 'react';
import MovieItem from './MovieItem'
import { Grid } from '@material-ui/core';

function MapItems({results}){
    console.log(results)
    if(JSON.stringify(results) === "{}") return( <div></div>)
    return(
        results.map((item,index) => (
            <MovieItem key={index} movieDetails={item} /> 
        ))
    )
}

function SearchResults({location}){
    const [results, setresults] = useState({})
    let moviesList = {}
    console.log(location)
    let searchQuery = location.search;
    searchQuery = searchQuery.replace(searchQuery[0],"")
    useEffect(() => {
        const fetchMovies = () => {
            const data = fetch('http://192.168.29.188:8500/api',{
                method: "POST",
                body: searchQuery
            }).then((res) => res.text())
            .then((res) => res.replace(searchQuery,""))
            .then((res)=>{
                res = JSON.parse(res);
                setresults(res);
                return res;
            })
        };
        fetchMovies();
        
    },[searchQuery])

    useEffect(() => {
        console.log(results)
    },[results])
    return(
        <div>
            <Grid container spacing={3}>
                <Grid item sm={9}>
                    <p style={{color:'#fff'}}>
                        Showing results for "{searchQuery}""
                    </p>
                </Grid>
                <Grid item sm={3}>
                    <p style={{color:'#fff'}}>
                        Filer Results
                    </p>
                </Grid>
                <Grid item sm={9}>
                    <div>
                        <MapItems results={results}/>
                    </div>
                </Grid>
                <Grid item sm={3}>

                </Grid>
            </Grid>
            
        </div>
    )
}

export default SearchResults;