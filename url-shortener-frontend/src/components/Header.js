import React from 'react';
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import { useNavigate } from 'react-router-dom';

const Header = ({ isLoggedIn, onLogout }) => {
    const navigate = useNavigate();

    return (
        <AppBar position="fixed">
            <Toolbar>
                <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
                    URL Shortener
                </Typography>
                {isLoggedIn ? (
                    <Button color="inherit" onClick={onLogout}>Logout</Button>
                ) : (
                    <>
                        <Button color="inherit" onClick={() => navigate('/login')}>Login</Button>
                        <Button color="inherit" onClick={() => navigate('/register')}>Register</Button>
                    </>
                )}
            </Toolbar>
        </AppBar>
    );
};

export default Header;

