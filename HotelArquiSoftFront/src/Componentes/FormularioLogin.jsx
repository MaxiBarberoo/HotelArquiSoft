import React, { useState } from 'react';

function FormularioLogin({ handleLogin }) {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const handleUsernameChange = (e) => {
    setUsername(e.target.value);
  };

  const handlePasswordChange = (e) => {
    setPassword(e.target.value);
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    if (username == 'usuario' && password == 'contra'){
      handleLogin(); 
    } else{
      alert("Las credenciales no coinciden");
    }
  }

  return (
    <form onSubmit={handleSubmit}>
      <h3>Iniciar sesión</h3>
      <label htmlFor="username">Nombre de usuario:</label>
      <input type="text" id="username" value={username} onChange={handleUsernameChange} />

      <label htmlFor="password">Contraseña:</label>
      <input type="password" id="password" value={password} onChange={handlePasswordChange} />

      <button type="submit">Iniciar sesión</button>
    </form>
  );
}

export default FormularioLogin;