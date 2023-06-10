import React, { useState } from 'react';

function FormularioLogin() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const handleUsernameChange = (e) => {
    setUsername(e.target.value);
  };

  const handlePasswordChange = (e) => {
    setPassword(e.target.value);
  };

  return (
    <form>
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