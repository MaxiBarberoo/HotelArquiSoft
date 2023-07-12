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

  const fetchData = async () => {
      console.log("Se realizo gethotels")
      const response = await fetch('http://localhost:8090/hotels');
      if (response.ok) {
        const data = await response.json();
        setHoteles(data);
      } 
  };

  useEffect(() => {
    fetchData();
  }, []);

  return (
    <div>
      <Header />
      <form onSubmit={handleRedirectSubmit} className="contenedor-boton-redireccion">
        <button type="submit" className="boton-redireccion">INICIA SESION PARA RESERVAR TU HOTEL</button>
      </form>
      <h2>Hoteles disponibles:</h2>
      <div className="contenedor-de-hoteles">
        {hoteles.map((hotel) => (
          <Hoteles
            key={hotel.id}
            hotelId = {hotel.id}
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


