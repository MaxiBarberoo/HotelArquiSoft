import React, { useState } from 'react';

function FormularioRegister({ setIsRegistered }) {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [email, setEmail] = useState('');

  const handleRegister = (e) => {
    e.preventDefault();
    // Aquí puedes realizar la lógica de verificación de registro
    // Por ejemplo, puedes validar los campos y enviar los datos al servidor

    // Si el registro es exitoso, puedes establecer el estado de registro a true
    setIsRegistered(true);
  };

  return (
    <form onSubmit={handleRegister}>
      <h3>Registrarse</h3>
      <label htmlFor="username">Nombre de usuario:</label>
      <input
        type="text"
        id="username"
        value={username}
        onChange={(e) => setUsername(e.target.value)}
      />

      <label htmlFor="password">Contraseña:</label>
      <input
        type="password"
        id="password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
      />

      <label htmlFor="email">Correo electrónico:</label>
      <input
        type="email"
        id="email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
      />

      <button type="submit">Registrarse</button>
    </form>
  );
}

export default FormularioRegister;