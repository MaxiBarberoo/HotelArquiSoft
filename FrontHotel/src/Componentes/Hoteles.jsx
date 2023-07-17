import React, { useEffect, useState } from 'react';
import '../Stylesheet/Hoteles.css';

function Hoteles(props) {

        const [amenities, setAmenities] = useState([]);

        useEffect(() => {
            fetch(`/amenitiehotel/${props.hotelId}`)
                .then(response => response.json())
                .then(data => {
                    setAmenities(data);
                })
                .catch(error => {
                    console.error('Error:', error);
                });
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
            {amenities.map(amenitie => (
                <li key={amenitie.id}>{amenitie.name}</li>
            ))}
        </ul>


    </div>
  );
}

export default Hoteles;





