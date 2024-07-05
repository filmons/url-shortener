// import React, { useState } from 'react';
// import axios from 'axios';

// const UrlForm = ({ onUrlShortened }) => {
//     const [longUrl, setLongUrl] = useState('');
//     const [shortUrl, setShortUrl] = useState('');

//     const handleSubmit = async (e) => {
//         e.preventDefault();
//         try {
//             const response = await axios.post('http://localhost:8080/shorten', { long_url: longUrl });
//             setShortUrl(response.data.data.short_url);
//             onUrlShortened(response.data.data);
//         } catch (error) {
//             console.error('Error shortening URL:', error);
//         }
//     };

//     return (
//         <div>
//             <form onSubmit={handleSubmit}>
//                 <input
//                     type="text"
//                     value={longUrl}
//                     onChange={(e) => setLongUrl(e.target.value)}
//                     placeholder="Enter long URL"
//                     required
//                 />
//                 <button type="submit">Shorten</button>
//             </form>
//             {shortUrl && (
//                 <div>
//                     <p>Short URL: <a href={`http://localhost:8080/${shortUrl}`} target="_blank" rel="noopener noreferrer">{`http://localhost:8080/${shortUrl}`}</a></p>
//                 </div>
//             )}
//         </div>
//     );
// };

// export default UrlForm;


// :)

// import React, { useState } from 'react';
// import axios from 'axios';

// const UrlForm = ({ onUrlShortened }) => {
//     const [longUrl, setLongUrl] = useState('');
//     const [shortUrl, setShortUrl] = useState('');

//     const handleSubmit = async (e) => {
//         e.preventDefault();
//         const token = localStorage.getItem('token');
//         try {
//             const response = await axios.post('http://localhost:8080/shorten', 
//                 { long_url: longUrl },
//                 { headers: { Authorization: token } }
//             );
//             setShortUrl(response.data.data.short_url);
//             onUrlShortened(response.data.data);
//         } catch (error) {
//             console.error('Error shortening URL:', error);
//         }
//     };

//     return (
//         <div>
//             <form onSubmit={handleSubmit}>
//                 <input
//                     type="text"
//                     value={longUrl}
//                     onChange={(e) => setLongUrl(e.target.value)}
//                     placeholder="Enter long URL"
//                     required
//                 />
//                 <button type="submit">Shorten</button>
//             </form>
//             {shortUrl && (
//                 <div>
//                     <p>Short URL: <a href={`http://localhost:8080/${shortUrl}`} target="_blank" rel="noopener noreferrer">{`http://localhost:8080/${shortUrl}`}</a></p>
//                 </div>
//             )}
//         </div>
//     );
// };

// export default UrlForm;


////////
// import React, { useState } from 'react';
// import axios from 'axios';
// import { TextField, Button, Box, Typography } from '@mui/material';

// const UrlForm = ({ onUrlShortened }) => {
//     const [longUrl, setLongUrl] = useState('');
//     const [shortUrl, setShortUrl] = useState('');

//     const handleSubmit = async (e) => {
//         e.preventDefault();
//         const token = localStorage.getItem('token');
//         try {
//             const response = await axios.post('http://localhost:8080/shorten', 
//                 { long_url: longUrl },
//                 { headers: { Authorization: token } }
//             );
//             setShortUrl(response.data.data.short_url);
//             onUrlShortened(response.data.data);
//         } catch (error) {
//             console.error('Error shortening URL:', error);
//         }
//     };

//     return (
//         <Box component="form" onSubmit={handleSubmit} sx={{ mb: 4 }}>
//             <TextField
//                 fullWidth
//                 label="Enter long URL"
//                 variant="outlined"
//                 value={longUrl}
//                 onChange={(e) => setLongUrl(e.target.value)}
//                 required
//                 sx={{ mb: 2 }}
//             />
//             <Button type="submit" variant="contained" color="primary" fullWidth>
//                 Shorten
//             </Button>
//             {shortUrl && (
//                 <Typography variant="body1" sx={{ mt: 2 }}>
//                     Short URL: <a href={`http://localhost:8080/${shortUrl}`} target="_blank" rel="noopener noreferrer">{`http://localhost:8080/${shortUrl}`}</a>
//                 </Typography>
//             )}
//         </Box>
//     );
// };

// export default UrlForm;

// v1
// import React, { useState } from 'react';
// import axios from 'axios';
// import { TextField, Button, Box, Typography } from '@mui/material';

// const UrlForm = ({ onUrlShortened }) => {
//     const [longUrl, setLongUrl] = useState('');
//     const [shortUrl, setShortUrl] = useState('');

//     const handleSubmit = async (e) => {
//         e.preventDefault();
//         const token = localStorage.getItem('token');
//         try {
//             const response = await axios.post('http://localhost:8080/shorten', 
//                 { long_url: longUrl },
//                 { headers: { Authorization: `Bearer ${token}` } }
//             );
//             setShortUrl(response.data.data.short_url);
//             onUrlShortened(response.data.data);
//         } catch (error) {
//             console.error('Error shortening URL:', error);
//         }
//     };

//     return (
//         <Box component="form" onSubmit={handleSubmit} sx={{ mb: 4 }}>
//             <TextField
//                 fullWidth
//                 label="Enter long URL"
//                 variant="outlined"
//                 value={longUrl}
//                 onChange={(e) => setLongUrl(e.target.value)}
//                 required
//                 sx={{ mb: 2 }}
//             />
//             <Button type="submit" variant="contained" color="primary" fullWidth>
//                 Shorten
//             </Button>
//             {shortUrl && (
//                 <Typography variant="body1" sx={{ mt: 2 }}>
//                     Short URL: <a href={`http://localhost:8080/${shortUrl}`} target="_blank" rel="noopener noreferrer">{`http://localhost:8080/${shortUrl}`}</a>
//                 </Typography>
//             )}
//         </Box>
//     );
// };

// export default UrlForm;




import React, { useState } from 'react';
import axios from 'axios';
import { TextField, Button, Box, Typography, Divider } from '@mui/material';

const UrlForm = ({ onUrlShortened }) => {
    const [longUrl, setLongUrl] = useState('');
    const [shortUrl, setShortUrl] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();
        const token = localStorage.getItem('token');
        try {
            const response = await axios.post('http://localhost:8080/shorten', 
                { long_url: longUrl },
                { headers: { Authorization: `Bearer ${token}` } }
            );
            setShortUrl(response.data.data.short_url);
            onUrlShortened(response.data.data);
        } catch (error) {
            console.error('Error shortening URL:', error);
        }
    };

    return (
        <Box sx={{ mb: 4 }}>
            <Box
                component="form"
                onSubmit={handleSubmit}
                sx={{
                    display: 'flex',
                    alignItems: 'center',
                    border: '1px solid #e0e0e0',
                    borderRadius: '20px',
                    padding: '6px 10px',
                    width: '100%',
                    backgroundColor: '#ffffff',
                    boxShadow: '0 3px 10px rgba(0,0,0,0.1)',
                }}
            >
                <TextField
                    fullWidth
                    placeholder="Enter long URL"
                    variant="standard"
                    value={longUrl}
                    onChange={(e) => setLongUrl(e.target.value)}
                    required
                    InputProps={{
                        disableUnderline: true,
                        style: { fontSize: '16px', padding: '8px' }
                    }}
                    sx={{ flexGrow: 1, mr: 1 }}
                />
                <Divider orientation="vertical" flexItem sx={{ mr: 1, my: 0.5 }} />
                <Button 
                    type="submit" 
                    variant="outlined"
                    sx={{
                        height: '40px',
                        minWidth: '100px',
                        borderRadius: '10px',
                        borderColor: '#3f51b5',
                        color: '#3f51b5',
                        backgroundColor: 'transparent',
                        '&:hover': {
                            backgroundColor: 'rgba(63, 81, 181, 0.04)',
                            borderColor: '#3f51b5',
                        },
                        textTransform: 'none',
                        fontWeight: 'bold',
                    }}
                >
                    Shorten
                </Button>
            </Box>
            {/* {shortUrl && (
                <Typography variant="body1" sx={{ mt: 2 }}>
                    Short URL: <a href={`http://localhost:8080/${shortUrl}`} target="_blank" rel="noopener noreferrer">{`http://localhost:8080/${shortUrl}`}</a>
                </Typography>
            )} */}
        </Box>
    );
};

export default UrlForm;