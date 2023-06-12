import React, { useState } from 'react';

function FormularioLogin({ handleLogin }) {
  const [User_email, setEmail] = useState('');
  const [Password, setPassword] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();

      const userData = {
          user_email: User_email,
          password: Password,
      };
    fetch('http://localhost:8090/users/auth', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(userData),
    })
      .then((response) => response.json())
      .then((data) => {
        if (data.autenticacion==="true") {
          const tipoUsuario = data.tipo; // Obtener el tipo de usuario autenticado desde la respuesta

          // Llamar a la función handleLogin pasando el tipo de usuario como argumento
          handleLogin(tipoUsuario);

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
      <input type="email" className='campoEmail' value={User_email} onChange={(e) => setEmail(e.target.value)} />
      <p>Contraseña:</p>
      <input type="password" className='campoContraseña' value={Password}  onChange={(e) => setPassword(e.target.value)} />
      <button type="submit">Iniciar sesión</button>
    </form>
  );
}

export default FormularioLogin;