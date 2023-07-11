import React, { useState } from 'react'
import '../Stylesheet/LoginRegister.css'
import Header from '../Componentes/Header'
import { useNavigate } from 'react-router-dom'

function LoginRegister() {
  const [isSignUpActive, setIsSignUpActive] = useState(false);
  const [nombre, setNombre] = useState('');
  const [apellido, setApellido] = useState('');
  const [email, setEmail] = useState('');
  const [contraseña, setContraseña] = useState('');
  const [token, setToken] = useState('');
  const [error, setError] = useState('');
  const navigate = useNavigate();

  const handleSignUpClick = () => {
    setIsSignUpActive(true);
  };

  const handleSignInClick = () => {
    setIsSignUpActive(false);
  };

  const handleLoginSubmit = async (event) => {
    event.preventDefault();
    const userData = {
      user_email: email,
      password: contraseña,
    };

    try {
      const authResponse = await fetch('http://localhost:8090/users/auth', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(userData),
      });

      if (authResponse.ok) {
        const responseJson = await authResponse.json();
        const tokenRecibido = responseJson.token; // Acceder al campo 'token' en la respuesta
        setToken(tokenRecibido);
        const isAuthenticated = responseJson.autenticacion;
        console.log(isAuthenticated);

        if (isAuthenticated === 'true') {
          navigate(`/reserve/${tokenRecibido}/${responseJson.user_id}`);
        } else {
          setError('Credenciales inválidas');
        }
      } else {
        setError('Error en la autenticación');
      }
    } catch (error) {
      setError('Error en la solicitud');
      console.log(error);
    }
  };

  const handleRegisterSubmit = async (event) => {
    const userData = {
      name: nombre,
      last_name: apellido,
      user_email: email,
      password: contraseña,
    };

    fetch('http://localhost:8090/users', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(userData),
    })
      .then(response => response.json())
      .then(data => {
        console.log(data);
        alert('El usuario ha sido registrado exitosamente');
      })
      .catch(error => {
        console.error(error);
      });
  }

  return (
    <div className="main-loginregister-container">
      <Header />
      <div className={`container ${isSignUpActive ? 'right-panel-active' : ''}`} id="container">
        <div className="form-container sign-up-container">
          <form onSubmit={handleRegisterSubmit}>
            <h2>Crear cuenta</h2>
            <input type="text" name="name" placeholder="Nombre" value={nombre} onChange={(e) => setNombre(e.target.value)} />
            <input type="text" name="lastname" placeholder="Apellido" value={apellido} onChange={(e) => setApellido(e.target.value)} />
            <input type="email" name="email" placeholder="Email" value={email} onChange={(e) => setEmail(e.target.value)} />
            <input type="password" name="contraseña" placeholder="Contraseña" value={contraseña} onChange={(e) => setContraseña(e.target.value)} />
            <button>Registrarse</button>
          </form>
        </div>
        <div className="form-container sign-in-container">
          <form onSubmit={handleLoginSubmit}>
            <h2>Iniciar sesión</h2>
            <input type="email" name="email" placeholder="Email" value={email} onChange={(e) => setEmail(e.target.value)} />
            <input type="password" name="contraseña" placeholder="Contraseña" value={contraseña} onChange={(e) => setContraseña(e.target.value)} />
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
