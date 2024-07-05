import React, { useState, useEffect } from 'react';
import { BrowserRouter as Router, Route, Routes, Navigate } from 'react-router-dom';
import axios from 'axios';
import UrlForm from './components/UrlForm';
import UrlList from './components/UrlList';
import RegisterForm from './components/RegisterForm';
import LoginForm from './components/LoginForm';
import './App.css';
import Header from './components/Header';
import { Box } from '@mui/material';

const App = () => {
    const [urls, setUrls] = useState([]);
    const [isLoggedIn, setIsLoggedIn] = useState(false);

    useEffect(() => {
        const token = localStorage.getItem('token');
        if (token) {
            setIsLoggedIn(true);
            fetchUrls(token);
        }
    }, []);

    const fetchUrls = async (token) => {
        try {
            const response = await axios.get('http://localhost:8080/urls', {
                headers: { Authorization: token }
            });
            setUrls(response.data.data);
        } catch (error) {
            console.error('Error fetching URLs:', error);
        }
    };

    const handleLogin = () => {
        setIsLoggedIn(true);
        const token = localStorage.getItem('token');
        fetchUrls(token);
    };

    const handleLogout = () => {
        localStorage.removeItem('token');
        setIsLoggedIn(false);
        setUrls([]);
    };

    const handleUrlShortened = (url) => {
        setUrls([url, ...urls]);
    };

    return (
        <Router>
            <Header isLoggedIn={isLoggedIn} onLogout={handleLogout} />
            <Box sx={{ paddingTop: '64px' }}>
                <Routes>
                    <Route path="/login" element={!isLoggedIn ? <LoginForm onLogin={handleLogin} /> : <Navigate to="/" />} />
                    <Route path="/register" element={!isLoggedIn ? <RegisterForm /> : <Navigate to="/" />} />
                    <Route path="/" element={
                        isLoggedIn ? (
                            <>
                                <UrlForm onUrlShortened={handleUrlShortened} />
                                <UrlList urls={urls} />
                            </>
                        ) : (
                            <Navigate to="/login" />
                        )
                    } />
                </Routes>
            </Box>
        </Router>
    );
};

export default App;
