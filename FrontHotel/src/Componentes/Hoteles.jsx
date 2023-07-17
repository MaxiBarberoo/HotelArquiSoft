import React, { useEffect, useState } from 'react';
import '../Stylesheet/Hoteles.css';

function Hoteles(props) {

        const [hotels, setHotels] = useState([]);
        const [amenities, setAmenities] = useState([]);

    useEffect(() => {
        // Obtener los IDs de los hoteles mediante una solicitud fetch
        fetch('http://localhost:8090/hotels')
            .then(response => response.json())
            .then(data => {
                setHotels(data);
            })
            .catch(error => {
                console.error('Error:', error);
            });
    }, []);


    useEffect(() => {
        const fetchAmenities = async () => {
            try {
                const response = await fetch(`http://localhost:8090/amenitiehotel/${props.hotelId}`);
                if (response.ok) {
                    const data = await response.json();
                    setAmenities(data);
                } else {
                    throw new Error(`Error en la petición GET de las amenidades para el hotel ${props.hotelId}`);
                }
            } catch (error) {
                console.error('Error:', error);
            }
        };

        fetchAmenities();
    }, [props.hotelId]);



    return (
    <div className="contenedor-hoteles">
      <p className="nombre-hotel1">
        <strong>{props.nombreHotel}</strong>
      </p>
      <p className="cantidad-piezas">
        Habitaciones: {props.piezas}
      </p>
      <p className="descripcion-hotel">
        Descripción: {props.descripcion}
      </p>

        <h2>Amenities del hotel:</h2>
        <ul>
            {amenities
                .filter(amenitie => amenitie.hotelId === props.hotelId)
                .map(amenitie => (
                    <li key={amenitie.id}>{amenitie.Tipo}</li>
                ))}
        </ul>

    </div>
  );
}

export default Hoteles;





