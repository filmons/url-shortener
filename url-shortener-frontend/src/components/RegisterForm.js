import React, { useState } from 'react';
import axios from 'axios';
import { TextField, Button, Box, Typography} from '@mui/material';

const RegisterForm = () => {
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [message, setMessage] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const response = await axios.post('http://localhost:8080/register', { name, email, password });
            setMessage(response.data.message);
        } catch (error) {
            setMessage('Error registering user');
            console.error('Error registering user:', error);
        }
    };

    return (
        <Box component="form" onSubmit={handleSubmit} sx={{ mb: 4 }}>
            <Typography variant="h6" sx={{ mb: 2 }}>
                Register
            </Typography>
            <TextField
                fullWidth
                label="Name"
                variant="outlined"
                value={name}
                onChange={(e) => setName(e.target.value)}
                required
                sx={{ mb: 2 }}
            />
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
                Register
            </Button>
            {message && (
                <Typography variant="body1" sx={{ mt: 2 }}>
                    {message}
                </Typography>
            )}
        </Box>
    );
};

export default RegisterForm;

