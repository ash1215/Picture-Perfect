export const login = (userInfo) => {
    return {
        type: 'LOGIN',
        info: userInfo
    }; 
};

export const logout = () => {
    return {
        type: 'LOGOUT'
    }; 
};

