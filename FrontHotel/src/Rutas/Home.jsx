import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import Header from "../Componentes/Header";
import Hoteles from "../Componentes/Hoteles";
import "../Stylesheet/Home.css";

function Home() {
  const [hoteles, setHoteles] = useState([]);
  const navigate = useNavigate();

  const handleRedirectSubmit = (event) => {
    navigate('/loginandregister');
  }

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch('http://localhost:8090/hotels');

        if (response.ok) {
          const data = await response.json();
          setHoteles(data);
        } else {
          throw new Error('Error en la petición GET');
        }
      } catch (error) {
        console.error(error);
      }
    };

    fetchData();
  }, []);

  return (
    <div>
      <Header />
      <form onSubmit={handleRedirectSubmit} className="contenedor-boton-redireccion">
        <button type="submit" className="boton-redireccion">INICIA SESION PARA RESERVAR TU HOTEL</button>
      </form>
      <div className="contenedor-de-hoteles">
        {hoteles.map((hotel) => (
          <Hoteles
            key={hotel.id}
            nombreHotel={hotel.name}
            piezas={hotel.cantHabitaciones}
            descripcion={hotel.descripcion}
          />
        ))}
      </div>
    </div>
  );
}

export default Home;


