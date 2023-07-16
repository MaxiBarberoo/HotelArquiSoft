import React from 'react';
import '../Stylesheet/Hoteles.css';

function Hoteles(props) {
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
    </div>
  );
}

export default Hoteles;





