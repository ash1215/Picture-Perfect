import userInfoReducer from './userInfo';
import loggedReducer from './isLogged';
import {combineReducers} from 'redux';

const allReducers = combineReducers({
    userInfo : userInfoReducer,
    logInfo : loggedReducer
});

export default allReducers;