import React, { useState, useEffect } from "react";
import { json, useNavigate } from "react-router-dom";
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
    const fetchImagenes = async (hotelId) => {
      try {
        const response = await fetch(`http://localhost:8090/imagenes/${hotelId}`);
        if (response.ok) {
          const imagenesData = await response.json();
          if (Array.isArray(imagenesData)) {
            const imagenesURLs = await Promise.all(imagenesData.map(async (imagen) => {
              if (imagen && imagen.contenido) {
                const blobData = new Blob([imagen.contenido], { type: 'image/jpeg' });
                return URL.createObjectURL(blobData);
              }
              return null;
            }));
            return imagenesURLs.filter((url) => url !== null);
          } else {
            throw new Error(`Error en el formato de las imágenes del hotel ${hotelId}`);
          }
        } else {
          throw new Error(`Error en la petición GET de las imágenes del hotel ${hotelId}`);
        }
      } catch (error) {
        console.error(error);
        return [];
      }
    };

    const fetchData = async () => {
      try {
        const response = await fetch('http://localhost:8090/hotels');
        if (response.ok) {
          const data = await response.json();
          const hotelesConImagenes = await Promise.all(data.map(async (hotel) => {
            const imagenesURLs = await fetchImagenes(hotel.id);
            return { ...hotel, imagenesURLs };
          }));
          setHoteles(hotelesConImagenes);
        } else {
          throw new Error('Error en la petición GET de hoteles');
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
      <h2>Hoteles disponibles:</h2>
      <div className="contenedor-de-hoteles">
        {hoteles.map((hotel) => (
          <Hoteles
            key={hotel.id}
            hotelID={hotel.id}
            imagenesURLs={hotel.imagenesURLs}
            nombreHotel={hotel.name}
            piezas={hotel.cantHabitaciones}
            descripcion={hotel.descripcion}
            amenities={hotel.amenities}

          />
        ))}
      </div>
    </div>
  );
}

export default Home;











