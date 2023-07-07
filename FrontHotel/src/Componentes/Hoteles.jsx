import React from 'react'
import '../Stylesheet/Hoteles.css'

function Hoteles(props){

    const handleReservaSubmit = () => {
        alert("Para reservar, primero debe iniciar sesión.");
    }

    return(
    <div className="contenedor-hoteles">
      <div className="contenedor-detalle-hoteles">
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
    </div>
    );
}

export default Hoteles;