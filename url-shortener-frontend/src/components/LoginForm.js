
import React, { useState } from 'react';
import axios from 'axios';
import { TextField, Button, Box, Typography} from '@mui/material';

const LoginForm = ({ onLogin}) => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [message, setMessage] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const response = await axios.post('http://localhost:8080/login', { email, password });
            localStorage.setItem('token', response.data.token);
            onLogin();
        } catch (error) {
            setMessage('Error logging in');
            console.error('Error logging in:', error);
        }
    };

    return (
        <Box component="form" onSubmit={handleSubmit} sx={{ mb: 4 }}>
            <Typography variant="h6" sx={{ mb: 2 }}>
                Login
            </Typography>
            <TextField
                fullWidth
                label="Email"
                variant="outlined"
                type="email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                required
                sx={{ mb: 2 }}
            />
            <TextField
                fullWidth
                label="Password"
                variant="outlined"
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                required
                sx={{ mb: 2 }}
            />
            <Button type="submit" variant="contained" color="primary" fullWidth>
                Login
            </Button>
            {message && (
                <Typography variant="body1" sx={{ mt: 2 }}>
                    {message}
                </Typography>
            )}
        </Box>
    );
};

export default LoginForm;
