import React, { useState, useEffect } from 'react';
import { 
  Paper, 
  Typography, 
  Link, 
  IconButton, 
  Box 
} from '@mui/material';
import DeleteIcon from '@mui/icons-material/Delete';
import ContentCopyIcon from '@mui/icons-material/ContentCopy';
import axios from 'axios';

const UrlList = ({ urls: initialUrls, setUrls: parentSetUrls }) => {
  const [urls, setUrls] = useState(initialUrls);

  useEffect(() => {
    setUrls(initialUrls);
  }, [initialUrls]);

  const deleteUrl = async (ID) => {
    console.log("Attempting to delete URL with ID:", ID);
    try {
      const token = localStorage.getItem('token');
      if (!token) {
        console.error('No token found');
        return;
      }
      if (!ID) {
        console.error('No ID provided for deletion');
        return;
      }
      const response = await axios.delete(`http://localhost:8080/del/${ID}`, {
        headers: { Authorization: `Bearer ${token}` }
      });
      console.log("Delete response:", response);
      const updatedUrls = urls.filter(url => url.ID !== ID);
      setUrls(updatedUrls);
      if (parentSetUrls) {
        parentSetUrls(updatedUrls);
      }
    } catch (error) {
      console.error('Error deleting URL:', error.response || error);
    }
  };

  const copyToClipboard = (text) => {
    navigator.clipboard.writeText(text).then(() => {
      console.log('Text copied to clipboard');
    }, (err) => {
      console.error('Could not copy text: ', err);
    });
  };

  return (
    <Box sx={{ width: '80%', margin: '0 auto' }}>
      <Typography 
        variant="h5" 
        sx={{ 
          mb: 2, 
          fontWeight: 'bold',
          textAlign: 'center', 
          width: '100%',
        }}
      >
        Your Shortened URLs
      </Typography>
      {urls.map((url, index) => (
        <Box key={url.ID || index} sx={{ mb: 2 }}>
          <Paper 
            elevation={0}
            sx={{
              display: 'flex',
              borderRadius: '8px',
              overflow: 'hidden',
              backgroundColor: index % 2 === 0 ? '#f5f5f5' : '#ffffff',
              border: '2px solid rgba(138, 43, 226, 0.3)', 
            }}
          >
            <Box 
              sx={{ 
                flex: 2, 
                p: 2,
                display: 'flex',
                flexDirection: 'column',
                justifyContent: 'center',
              }}
            >
              <Typography 
                noWrap 
                sx={{ 
                  mb: 1, 
                  fontSize: '0.9rem', 
                  color: 'text.secondary',
                  maxWidth: '1000px', 
                }}
              >
                {url.long_url}
              </Typography>
              <Link 
                href={`http://localhost:8080/${url.short_url}`} 
                target="_blank" 
                rel="noopener noreferrer"
                sx={{ 
                  textDecoration: 'none',
                  fontSize: '1rem',
                  fontWeight: 'medium',
                  color: 'primary.main',
                }}
              >
                {`localhost:8080/${url.short_url}`}
              </Link>
            </Box>
            <Box 
              sx={{ 
                p: 1, 
                display: 'flex', 
                alignItems: 'center',
                justifyContent: 'center',
              }}
            >
              <IconButton 
                onClick={() => copyToClipboard(`http://localhost:8080/${url.short_url}`)}
                color="primary"
                size="small"
              >
                <ContentCopyIcon />
              </IconButton>
              <IconButton 
                onClick={() => deleteUrl(url.ID)}
                color="error"
                size="small"
              >
                <DeleteIcon />
              </IconButton>
            </Box>
          </Paper>
        </Box>
      ))}
    </Box>
  );
};

export default UrlList;