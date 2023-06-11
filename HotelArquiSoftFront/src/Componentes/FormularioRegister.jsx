import React, { useState } from 'react';

function FormularioRegister() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [user_email, setEmail] = useState('');

  const handleRegister = (e) => {
    e.preventDefault();
    setUsername('');
    setEmail('');
    setPassword('');
    alert("Se ha registrado correctamente");
  };

  return (
    <form onSubmit={handleRegister}>
      <h3>Registrarse</h3>
      <p>Nombre de usuario:</p>
      <input type="text" className='campoUsuario' value={username} onChange={(e) => setUsername(e.target.value)}/>
      <p>Contraseña:</p>
      <input type="password" className='campoContraseña' value={password} onChange={(e) => setPassword(e.target.value)}/>
      <p>Correo electrónico:</p>
      <input type="email" className='campoEmail' value={user_email} onChange={(e) => setEmail(e.target.value)}/>
      <button type="submit">Registrarse</button>
    </form>
  );
}

export default FormularioRegister;