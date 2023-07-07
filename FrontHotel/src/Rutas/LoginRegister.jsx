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
      <Header />
      <div className={`container ${isSignUpActive ? 'right-panel-active' : ''}`} id="container">
        <div className="form-container sign-up-container">
          <form action="#">
            <h2>Crear cuenta</h2>
            <input type="text" placeholder="Nombre" />
			<input type="text" placeholder="Apellido" />
            <input type="email" placeholder="Email" />
            <input type="password" placeholder="Contraseña" />
            <button>Registrarse</button>
          </form>
        </div>
        <div className="form-container sign-in-container">
          <form action="#" onSubmit={handleLoginSubmit}>
            <h2>Iniciar sesión</h2>
            <input type="email" placeholder="Email" />
            <input type="password" placeholder="Contraseña" />
            <button type="submit">Loguearse</button>
          </form>
        </div>
        <div className="overlay-container">
          <div className="overlay">
            <div className="overlay-panel overlay-left">
              <h3>Hola de nuevo!</h3>
              <p>Para conectarte, por favor proporciona tus datos de inicio de sesión</p>
              <button className="ghost" onClick={handleSignInClick}>Iniciar sesión</button>
            </div>
            <div className="overlay-panel overlay-right">
              <h3>Bienvenido!</h3>
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