import React, { useState } from 'react';
import '../Stylesheet/LoginRegister.css'
import Header from '../Componentes/Header'
import { useNavigate } from 'react-router-dom';

function LoginRegister(){
  const [isSignUpActive, setIsSignUpActive] = useState(false);
  const navigate = useNavigate(); 

  const handleSignUpClick = () => {
    setIsSignUpActive(true);
  };

  const handleSignInClick = () => {
    setIsSignUpActive(false);
  };

  const handleLoginSubmit = (event) => {
	navigate('/reserve');
  }

  return (
    <div className="main-loginregister-container">
      <div className="header">
        <h1>SUMMIT LUXURY Hotels</h1>
      </div>
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
          <form action="#" onSubmit={handleLoginSubmit}>
            <h1>Iniciar sesión</h1>
            <input type="email" placeholder="Email" />
            <input type="password" placeholder="Contraseña" />
            <button type="submit">Loguearse</button>
          </form>
        </div>
        <div className="overlay-container">
          <div className="overlay">
            <div className="overlay-panel overlay-left">
              <h2>Hola de nuevo!</h2>
              <p>Para conectarte, por favor proporciona tus datos de inicio de sesión</p>
              <button className="ghost" onClick={handleSignInClick}>Iniciar sesión</button>
            </div>
            <div className="overlay-panel overlay-right">
              <h2>Bienvenido!</h2>
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