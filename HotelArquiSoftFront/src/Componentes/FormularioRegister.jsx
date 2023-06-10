import React from 'react';

function FormularioRegister() {
  return (
    <form>
      <h3>Registrarse</h3>
      <label htmlFor="username">Nombre de usuario:</label>
      <input type="text" id="username" />

      <label htmlFor="password">Contraseña:</label>
      <input type="password" id="password" />

      <label htmlFor="email">Correo electrónico:</label>
      <input type="email" id="email" />

      <button type="submit">Registrarse</button>
    </form>
  );
}

export default FormularioRegister;