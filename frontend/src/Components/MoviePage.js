import React, {useState,useEffect} from 'react';
import {useParams} from 'react-router-dom';

function MoviePage() {
    let {id} = useParams();
    return(
        <div>
            <h3 style={{color:'#fff'}}>TMDb id: {id}</h3>
        </div>
    )
}

export default MoviePage;