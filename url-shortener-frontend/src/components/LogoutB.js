import React from 'react';

function LogoutB() {
  const handleLogout = () => {
    localStorage.clear();
    window.location.href = '/';
  };

  return <button onClick={handleLogout}>Déconnexion</button>;
}

export default LogoutB;