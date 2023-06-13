import React, { useState, useEffect } from "react";

function Reservas(props) {
    const [reservas, setReservas] = useState([]);

    useEffect(() => {
        // Llamada al método getReservasByUser utilizando fetch
        fetch(`http://localhost:8090/reservas/reservauser/${props.userId}`)
          .then((response) => response.json())
          .then((data) => {
            // Actualizar el estado con las reservas obtenidas
            setReservas(data);
    
            // Obtener el nombre del hotel para cada reserva
            const fetchHotelName = async () => {
              const reservasConHotel = await Promise.all(
                data.map(async (reserva) => {
                  const response = await fetch(
                    `http://localhost:8090/hotels/${reserva.hotel_id}`
                  );
                  const hotelData = await response.json();
                  return { ...reserva, hotel_nombre: hotelData.name };
                })
              );
              setReservas(reservasConHotel);
            };
    
            fetchHotelName();
          })
          .catch((error) => {
            console.error("Error:", error);
          });
      }, [props.userId]);

    return (
        <div>
            <h2>Reservas</h2>
            {reservas.map((reserva) => (
        <div key={reserva.id}>
          <p>Fecha de Ingreso: {reserva.fecha_ingreso}</p>
          <p>Fecha de Egreso: {reserva.fecha_egreso}</p>
          <p>Hotel: {reserva.hotel_nombre}</p>
        </div>
      ))}
        </div>
    );
}

export default Reservas;