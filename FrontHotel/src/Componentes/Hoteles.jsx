import React, { useEffect, useState } from 'react';
import '../Stylesheet/Hoteles.css';

function Hoteles(props) {
  const [imagenURL, setImagenURL] = useState('');

  const fetchImagen = async () => {
    try {
      const response = await axios(`http://localhost:8090/imagenes/${props.hotelId}`);
      console.log(props.hotelId);
      if (response.ok) {
        const jsonResponse = await response.json(); // Obtener el JSON de la respuesta
        console.log(jsonResponse.nombre);
        const blobData = new Blob([jsonResponse.contenido], { type: 'image/jpeg' }); // Crear el Blob a partir del contenido de la imagen
        console.log(blobData);
        const imageURL = URL.createObjectURL(blobData);
        console.log(imageURL);
        setImagenURL(imageURL);
      } else {
        throw new Error('Error en la petición GET de la imagen');
      }
    } catch (error) {
      console.error(error);
    }
  };

  useEffect(() => {
    fetchImagen();
  }, []);

  return (
    <div className="contenedor-hoteles">
      <p className="nombre-hotel1">
        <strong>{props.nombreHotel}</strong>
      </p>
      <img src={imagenURL} alt="Imagen del hotel" />
      <p className="cantidad-piezas">
        Habitaciones: {props.piezas}
      </p>
      <p className="descripcion-hotel">
        Descripción: {props.descripcion}
      </p>
    </div>
  );
}

export default Hoteles;




