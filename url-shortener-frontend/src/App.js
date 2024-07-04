// import React, { useState, useEffect } from 'react';
// import axios from 'axios';
// import UrlForm from './components/UrlForm';
// import UrlList from './components/UrlList';
// import RegisterForm from './components/RegisterForm';
// import LoginForm from './components/LoginForm';

// const App = () => {
//     const [urls, setUrls] = useState([]);
//     const [isLoggedIn, setIsLoggedIn] = useState(false);

//     useEffect(() => {
//         const token = localStorage.getItem('token');
//         if (token) {
//             setIsLoggedIn(true);
//             fetchUrls(token);
//         }
//     }, []);

//     const fetchUrls = async (token) => {
//         try {
//             const response = await axios.get('http://localhost:8080/urls', {
//                 headers: { Authorization: token }
//             });
//             setUrls(response.data.data);
//         } catch (error) {
//             console.error('Error fetching URLs:', error);
//         }
//     };

//     const handleLogin = () => {
//         setIsLoggedIn(true);
//         const token = localStorage.getItem('token');
//         fetchUrls(token);
//     };

//     const handleUrlShortened = (url) => {
//         setUrls([url, ...urls]);
//     };

//     return (
//         <div>
//             <h1>URL Shortener</h1>
//             {isLoggedIn ? (
//                 <>
//                     <UrlForm onUrlShortened={handleUrlShortened} />
//                     <UrlList urls={urls} />
//                 </>
//             ) : (
//                 <>
//                     <RegisterForm />
//                     <LoginForm onLogin={handleLogin} />
//                 </>
//             )}
//         </div>
//     );
// };
// 
// export default App;
import React, { useState, useEffect } from 'react';
import axios from 'axios';
import UrlForm from './components/UrlForm';
import UrlList from './components/UrlList';
import RegisterForm from './components/RegisterForm';
import LoginForm from './components/LoginForm';

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
        <div>
            <h1>URL Shortener</h1>
            {isLoggedIn ? (
                <>
                    <UrlForm onUrlShortened={handleUrlShortened} />
                    <UrlList urls={urls} />
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
