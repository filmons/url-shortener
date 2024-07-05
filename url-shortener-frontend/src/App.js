import React, { useState, useEffect } from 'react';
import axios from 'axios';
import UrlForm from './components/UrlForm';
import UrlList from './components/UrlList';
import RegisterForm from './components/RegisterForm';
import LoginForm from './components/LoginForm';
import LogoutB from './components/LogoutB';
import "./App.css";

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

    const handleUrlShortened = (url) => {
        setUrls([url, ...urls]);
    };

    return (
        <div className="app-container">
            {isLoggedIn ? (
                <>
                    <div className='url-header'>
                        <h1>Url_shortener</h1>
                        <p>Transform your lengthy web addresses into concise, manageable links with our quick and efficient URL shortening tool. Shorten links instantly.</p>
                        <LogoutB />
                    </div>
                    <div className='url-form-container'>
                        <div className='url-form'>
                            <UrlForm onUrlShortened={handleUrlShortened} />
                        </div>
                    </div>
                    <div className='url-list'>
                        <UrlList urls={urls} />
                    </div>
                </>
            ) : (
                <> 
                    <RegisterForm />
                    <LoginForm onLogin={handleLogin} />
                </>
            )}
        </div>

    );
};

export default App;
