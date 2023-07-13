import React, { useState, useEffect } from 'react';
import '../Stylesheet/LoginRegister.css';
import Header from '../Componentes/Header';
import { useNavigate } from 'react-router-dom';


function LoginRegister() {
  const [nombre, setNombre] = useState('');
  const [apellido, setApellido] = useState('');
  const [email, setEmail] = useState('');
  const [contraseña, setContraseña] = useState('');
  const [token, setToken] = useState('');
  const [error, setError] = useState('');
  const navigate = useNavigate();

  const handleLoginSubmit = async (event) => {
    event.preventDefault();
    if (email.trim() == '' || contraseña.trim() == ''){
      alert("Uno de los campos se encuentra vacío, por favor completelos e intente nuevamente.");
      setEmail('');
      setContraseña('');
    } else{
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
    }
  };

  const handleRegisterSubmit = async (event) => {
    if (email.trim() == '' || contraseña.trim() == '' || nombre.trim() == '' || apellido.trim() == ''){
      alert("Uno de los campos se encuentra vacío, por favor completelos e intente nuevamente.");
      setNombre('');
      setApellido('');
      setEmail('');
      setContraseña('');
    } else{
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
          alert('El usuario ha sido registrado exitosamente');
        })
        .catch(error => {
          console.error(error);
        });
    }
  }

  return (
    <div className="main-loginregister-container">
      <Header />
      <div className='container-principal'>
        <div className="main">
          <input type="checkbox" id="chk" aria-hidden="true" />
          <div className="signup">
            <form onSubmit={handleRegisterSubmit}>
              <label for="chk" aria-hidden="true">Sign up</label>
              <input type="text" name="txt" placeholder="Nombre" required="" value={nombre} onChange={(e) => setNombre(e.target.value)} />
              <input type="text" name="txt" placeholder="Apellido" required="" value={apellido} onChange={(e) => setApellido(e.target.value)}/>
              <input type="email" name="email" placeholder="Email" required="" value={email}  onChange={(e) => setEmail(e.target.value)}/>
              <input type="password" name="pswd" placeholder="Contraseña" required="" value={contraseña} onChange={(e) => setContraseña(e.target.value)}/>
              <button className='sign' type="submit">Sign up</button>
            </form>
          </div>
          <div className="login">
            <form onSubmit={handleLoginSubmit}>
              <label for="chk" aria-hidden="true">Login</label>
              <input type="email" name="email" placeholder="Email" required="" value={email} onChange={(e) => setEmail(e.target.value)}/>
              <input type="password" name="pswd" placeholder="Contraseña" required="" value={contraseña} onChange={(e) => setContraseña(e.target.value)}/>
              <button className='log' type="submit">Login</button>
            </form>
          </div>
        </div>
      </div>
    </div>
  );
}

export default LoginRegister;
