import React, { useState } from 'react';

function FormularioLogin({ handleLogin }) {
  const [user_email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    fetch('/users/auth', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ UserEmail: user_email, Password: password }),
    })
      .then((response) => response.json())
      .then((data) => {
        if (data.authenticated) {
          //Falta autenticar tipo de usuario en esta parte
          alert('Autenticación exitosa');
        } else {
          // Error de autenticación
          alert('La autenticación fue incorrecta. Ingrese sus datos nuevamente');
        }
      })
      .catch((error) => {
        // Manejar el error de la solicitud
        console.error('Error:', error);
      });
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