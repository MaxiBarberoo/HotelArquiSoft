import React, { useState } from 'react';

function FormularioLogin({ handleLogin }) {
  const [user_email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    if (user_email == '2100815@ucc.edu.ar' && password == 'Contra'){
      handleLogin();
      setEmail('');
      setPassword('');
      alert("Ha iniciado sesión con éxito");
    }
  }

  return (
    <form onSubmit={handleSubmit}>
      <h3>Iniciar sesión</h3>
      <p>Email:</p>
      <input type="email" className='campoEmail' onChange={(e) => setEmail(e.target.value)} />
      <p>Contraseña:</p>
      <input type="password" className='campoContraseña' onChange={(e) => setPassword(e.target.value)} />
      <button type="submit">Iniciar sesión</button>
    </form>
  );
}

export default FormularioLogin;