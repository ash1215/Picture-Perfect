import React,{useState} from 'react';
import SearchIcon from '@material-ui/icons/Search';
import InputBase from '@material-ui/core/InputBase';
import { Link,useHistory } from 'react-router-dom';

function SearchBar({theme}){

    const [searchText, setsearchText] = useState('');
    let history = useHistory();
    const handleSearch = (e) => {
        e.preventDefault();
        console.log("SearchText: ",searchText)
        if(!searchText) {
            console.log('empty')
            return;
        }
        else{
            return(
                history.push({pathname:"/search",search: searchText,state: {}})
            )
        }
    }


    return(
        <div style={{
        position: 'relative',
        marginRight: theme.spacing(2),
        marginLeft: 0,
        width: '100%',
        [theme.breakpoints.up('sm')]: {
            marginLeft: theme.spacing(3),
            width: 'auto',
        },
        background: '#fff',
        paddingLeft:'25px',
        borderRadius:'25px',
        height:'50px'
        }}>
        <form onSubmit={handleSearch}>
        <InputBase
            placeholder="Search…"
            inputProps={{ 'aria-label': 'search' }}
            type="search"
            fullWidth
            value={searchText}
            onChange = {e => setsearchText(e.target.value)}
        />
        </form>
        <div style={{
            padding: theme.spacing(0, 2),
            height: '100%',
            position: 'absolute',
            pointerEvents: 'none',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
        }}>
            
            <Link to="/SEARCH">HERE
                <SearchIcon color='error'/>
            </Link>
        </div>
        </div>
        // <div />
    )
}

export default SearchBar;