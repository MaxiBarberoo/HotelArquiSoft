import React from 'react'
import '../Stylesheet/Hoteles.css'

function HotelesR(props){

    const Reservar = (event) => {
        event.preventDefault();
        // Realizar la solicitud al backend para guardar la reserva
        fetch("http://localhost:8090/reservas", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `${props.token}`
            },
            body: JSON.stringify({
                fecha_ingreso: props.fechaDesde,
                fecha_egreso: props.fechaHasta,
                hotel_id: parseInt(props.hotelId), // Agrega el ID del hotel correspondiente
                user_id: parseInt(props.userId), // Agrega el ID del usuario correspondiente
            }),
        })
            .then((response) => response.json())
            .then((data) => {
                // Mostrar un mensaje de éxito o realizar otras acciones necesarias
                alert("Reserva realizada con éxito");
            })
            .catch((error) => {
                console.error("Error:", error);
            });
    };

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
            <form onSubmit={Reservar} className="boton-reserva">
                <button type="submit">Reservar</button>
            </form>
        </div>
    );
}

export default HotelesR;