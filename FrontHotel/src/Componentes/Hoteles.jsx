import React, { useEffect, useState } from 'react';
import '../Stylesheet/Hoteles.css';

function Hoteles(props) {

        const [hotels, setHotels] = useState([]);
        const [amenities, setAmenities] = useState([]);

    useEffect(() => {
        // Obtener los IDs de los hoteles mediante una solicitud fetch
        fetch('/hotels')
            .then(response => response.json())
            .then(data => {
                setHotels(data);
            })
            .catch(error => {
                console.error('Error:', error);
            });
    }, []);


    useEffect(() => {
        // Realizar la solicitud fetch para cada hotel y obtener los amenities correspondientes
        hotels.forEach(hotel => {
            fetch(`/amenitiehotel/${hotel.id}`)
                .then(response => response.json())
                .then(data => {
                    // Agregar los amenities al estado general
                    setAmenities(prevAmenities =>[...prevAmenities, ...data]);
                    console.log(data); // Imprimir los datos de amenities en la consola
                })
                .catch(error => {
                    console.error('Error:', error);
                });
        });
    }, [hotels]);

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





