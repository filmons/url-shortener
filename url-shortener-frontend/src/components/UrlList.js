// import React from 'react';

// const UrlList = ({ urls }) => {
//     return (
//         <div>
//             <h2>Shortened URLs</h2>
//             <ul>
//                 {urls.map((url, index) => (
//                     <li key={index}>
//                         Long URL: {url.long_url} - Short URL: <a href={`http://localhost:8080/${url.short_url}`} target="_blank" rel="noopener noreferrer">{`http://localhost:8080/${url.short_url}`}</a>
//                     </li>
//                 ))}
//             </ul>
//         </div>
//     );
// };

// export default UrlList;
import React from 'react';
import { List, ListItem, ListItemText, Link, Typography } from '@mui/material';

const UrlList = ({ urls }) => {
    return (
        <>
            <Typography variant="h6" sx={{ mb: 2 }}>
                Shortened URLs
            </Typography>
            <List>
                {urls.map((url, index) => (
                    <ListItem key={index}>
                        <ListItemText
                            primary={`Long URL: ${url.long_url}`}
                            secondary={
                                <Link href={`http://localhost:8080/${url.short_url}`} target="_blank" rel="noopener noreferrer">
                                    {`Short URL: http://localhost:8080/${url.short_url}`}
                                </Link>
                            }
                        />
                    </ListItem>
                ))}
            </List>
        </>
    );
};

export default UrlList;

