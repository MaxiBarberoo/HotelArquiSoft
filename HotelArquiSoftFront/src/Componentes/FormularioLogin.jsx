import React from 'react';

function FormularioLogin() {
  return (
    <form>
      <h3>Iniciar sesión</h3>
      <label htmlFor="username">Nombre de usuario:</label>
      <input type="text" id="username" />

      <label htmlFor="password">Contraseña:</label>
      <input type="password" id="password" />

      <button type="submit">Iniciar sesión</button>
    </form>
  );
}

export default FormularioLogin;