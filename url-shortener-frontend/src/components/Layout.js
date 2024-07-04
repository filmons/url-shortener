import React from 'react';
import Container from '@mui/material/Container';
import CssBaseline from '@mui/material/CssBaseline';
import Header from './Header';

const Layout = ({ children }) => {
    return (
        <>
            <CssBaseline />
            <Header />
            <Container>
                {children}
            </Container>
        </>
    );
};

export default Layout;
