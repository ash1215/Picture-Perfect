const userInfoReducer = (state = {}, action) => {
    switch(action.type) {
        case 'LOGIN':
            return action.info;
        case 'LOGOUT':
            return {};
        default:
            return state;
    }
};

export default userInfoReducer;