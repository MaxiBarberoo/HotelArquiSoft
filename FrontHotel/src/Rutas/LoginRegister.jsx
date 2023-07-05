import React, { useState } from 'react';
import '../Stylesheet/LoginRegister.css'
import Header from '../Componentes/Header'

function LoginRegister(){
  const [isSignUpActive, setIsSignUpActive] = useState(false);

  const handleSignUpClick = () => {
    setIsSignUpActive(true);
  };

  const handleSignInClick = () => {
    setIsSignUpActive(false);
  };

  return (
    <div className="main-loginregister-container">
      <div className={`container ${isSignUpActive ? 'right-panel-active' : ''}`} id="container">
        <div className="form-container sign-up-container">
          <form action="#">
            <h1>Crear cuenta</h1>
            <input type="text" placeholder="Nombre" />
			<input type="text" placeholder="Apellido" />
            <input type="email" placeholder="Email" />
            <input type="password" placeholder="Contraseña" />
            <button>Registrarse</button>
          </form>
        </div>
        <div className="form-container sign-in-container">
          <form action="#">
            <h1>Iniciar sesión</h1>
            <input type="email" placeholder="Email" />
            <input type="password" placeholder="Contraseña" />
            <button>Loguearse</button>
          </form>
        </div>
        <div className="overlay-container">
          <div className="overlay">
            <div className="overlay-panel overlay-left">
              <h1>Hola de nuevo!</h1>
              <p>Para conectarte, por favor proporciona tus datos de inicio de sesión</p>
              <button className="ghost" onClick={handleSignInClick}>Iniciar sesión</button>
            </div>
            <div className="overlay-panel overlay-right">
              <h1>Bienvenido!</h1>
              <p>Ingresa tus datos personales y comienza tu travesía con nosotros</p>
              <button className="ghost" onClick={handleSignUpClick}>Registrarse</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default LoginRegister;