import React, { useState, useEffect } from 'react';
import '../Stylesheet/LoginRegister.css';
import Header from '../Componentes/Header';
import { useNavigate } from 'react-router-dom';
import jwtDecode from 'jwt-decode';

function LoginRegister() {
  const [isSignUpActive, setIsSignUpActive] = useState(false);
  const [nombre, setNombre] = useState('');
  const [apellido, setApellido] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [token, setToken] = useState('');
  const navigate = useNavigate();

  const handleSignUpClick = () => {
    setIsSignUpActive(true);
  };

  const handleSignInClick = () => {
    setIsSignUpActive(false);
  };

  const fetchToken = async () => {
    const userData = {
      email,
      password
    };
    try {
      const response = await fetch('http://localhost:8090/users/jwt', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ userData }),
      });

      if (response.ok) {
        const tokenRecibido = await response.json();
        console.log(tokenRecibido.token);
        setToken(tokenRecibido.token);
      } else {
        throw new Error('Error al generar el token');
      }
    } catch (error) {
      setError('Error al generar el token');
      console.log(error);
    }
  };

  const handleUserAuth = async () => {
    const userData = {
      email,
      password
    };
    console.log(token);
    
    try {
      const authResponse = await fetch('http://localhost:8090/users/auth', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `${token}`
        },
        body: JSON.stringify(userData)
      });

      if (authResponse.ok) {
        const decodedtoken = await authResponse.json();
        const isAuthenticated = decodedtoken.autenticacion;
        console.log(isAuthenticated);
  
        if (isAuthenticated == 'true') {
          navigate('/reserve');
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

  const handleLoginSubmit = async (event) => {
    event.preventDefault();
    setEmail(event.target.elements.email.value);
    setPassword(event.target.elements.password.value);
    await fetchToken();
  };  

  const handleRegisterSubmit = (event) => {
    event.preventDefault();
    setNombre(event.target.elements.name.value);
    setApellido(event.target.elements.lastname.value);
    setEmail(event.target.elements.email.value);
    setPassword(event.target.elements.password.value);

    const userData = {
      nombre,
      apellido,
      email,
      password,
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
        setNombre('');
        setApellido('');
        setEmail('');
        setPassword('');
      })
      .catch(error => {
        console.error(error);
      });
  }

  useEffect(() => {
    if (token) {
      handleUserAuth();
    }
  }, [token]);

  return (
    <div className="main-loginregister-container">
      <Header />
      <div className={`container ${isSignUpActive ? 'right-panel-active' : ''}`} id="container">
        <div className="form-container sign-up-container">
          <form onSubmit={handleRegisterSubmit}>
            <h2>Crear cuenta</h2>
            <input type="text" name="name" placeholder="Nombre" />
            <input type="text" name="lastname" placeholder="Apellido" />
            <input type="email" name="email" placeholder="Email" />
            <input type="password" name="password" placeholder="Contraseña" />
            <button>Registrarse</button>
          </form>
        </div>
        <div className="form-container sign-in-container">
          <form onSubmit={handleLoginSubmit}>
            <h2>Iniciar sesión</h2>
            <input type="email" name="email" placeholder="Email" />
            <input type="password" name="password" placeholder="Contraseña" />
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
