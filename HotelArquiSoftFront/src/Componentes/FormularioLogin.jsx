import React, { useState } from 'react';

function FormularioLogin({ handleLogin }) {
  const [user_email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    fetch('http://localhost:8090/users/auth', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ user_email: user_email, password: password }),
    })
      .then((response) => response.json())
      .then((data) => {
        if (data.autenticacion=="true") {
          const tipoUsuario = data.tipo; // Obtener el tipo de usuario autenticado desde la respuesta

          // Llamar a la función handleLogin pasando el tipo de usuario como argumento
          handleLogin(data.tipo);

          alert('Autenticación exitosa');
        } else {
          alert('La autenticación fue incorrecta. Ingrese sus datos nuevamente');
        }
      })
      .catch((error) => {
        console.error('Error:', error);
      });
  };

  return (
    <form onSubmit={handleSubmit}>
      <h3>Iniciar sesión</h3>
      <p>Email:</p>
      <input type="email" className='campoEmail' value={user_email} onChange={(e) => setEmail(e.target.value)} />
      <p>Contraseña:</p>
      <input type="password" className='campoContraseña' value={password}  onChange={(e) => setPassword(e.target.value)} />
      <button type="submit">Iniciar sesión</button>
    </form>
  );
}

export default FormularioLogin;