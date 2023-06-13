import React, { useEffect, useState } from "react";

function Reservas(props) {
  const { reservas } = props;
  const [reservasConNombreHotel, setReservasConNombreHotel] = useState([]);

  useEffect(() => {
    const fetchReservasConNombreHotel = async () => {
      const reservasConNombre = await Promise.all(
        reservas.map(async (reserva) => {
          const hotel = await getHotelById(reserva.hotel_id);
          return { ...reserva, hotel_nombre: hotel.name };
        })
      );
      setReservasConNombreHotel(reservasConNombre);
    };

    fetchReservasConNombreHotel();
  }, [reservas]);

  const getHotelById = async (hotelId) => {
    try {
      const response = await fetch(`http://localhost:8090/hotels/${hotelId}`);
      if (!response.ok) {
        throw new Error("Error al obtener el hotel");
      }
      const hotel = await response.json();
      return hotel;
    } catch (error) {
      console.error("Error:", error);
      return null;
    }
  };

  return (
    <div>
      <h2>Reservas</h2>
      {reservasConNombreHotel.length > 0 ? (
        <ul>
          {reservasConNombreHotel.map((reserva) => (
            <li key={reserva.id}>
              <p>Fecha de ingreso: {reserva.fecha_ingreso}</p>
              <p>Fecha de egreso: {reserva.fecha_egreso}</p>
              <p>Hotel: {reserva.hotel_nombre}</p>
            </li>
          ))}
        </ul>
      ) : (
        <p>No tienes reservas realizadas.</p>
      )}
    </div>
  );
}

export default Reservas;