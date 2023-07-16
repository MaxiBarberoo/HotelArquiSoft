import React, { useState, useEffect } from "react"
import '../Stylesheet/Reserve.css'
import Header from '../Componentes/Header'
import DatePicker from 'react-datepicker'
import HotelesR from '../Componentes/HotelesR'
import 'react-datepicker/dist/react-datepicker.css'
import { useParams } from 'react-router-dom';

function Reserve() {
  const [fechaDesde, setFechaDesde] = useState(null);
  const [fechaHasta, setFechaHasta] = useState(null);
  const [hotelesDisponibles, setHotelesDisponibles] = useState([]);
  const [reservas, setReservas] = useState([]);
  const [contadorReserva, setContadorReserva] = useState(1);
  const { token } = useParams();
  const { user_id } = useParams();

  const handleFechaDesdeChange = (date) => {
    setFechaDesde(date);
  };

  const handleFechaHastaChange = (date) => {
    setFechaHasta(date);
  };

  const buscarHotelesDisponibles = () => {
    if (!fechaDesde || !fechaHasta) {
      alert("Debes completar los campos de fecha desde y fecha hasta.");
    } else if (fechaDesde >= fechaHasta) {
      alert("La fecha desde debe ser anterior a la fecha hasta.");
    } else {
      fetch("http://localhost:8090/reservas/hotelsbyfecha", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Authorization": `${token}`,
        },
        body: JSON.stringify({
          fecha_ingreso: fechaDesde,
          fecha_egreso: fechaHasta,
        }),
      })
        .then((response) => response.json())
        .then((data) => {
          const hotelIds = Object.values(data).map((item) => item.hotel_id);
          Promise.all(
            hotelIds.map((hotelId) =>
              fetch(`http://localhost:8090/hotels/${hotelId}`, {
                method: "GET",
                headers: {
                  Authorization: `${token}`,
                },
              }).then((response) => response.json())
            )
          )
            .then((hotelData) => {
              setHotelesDisponibles(hotelData);
            })
            .catch((error) => console.error(error));
        })
        .catch((error) => console.error(error));
    }
  };

  useEffect(() => {
    fetch(`http://localhost:8090/reservas/reservauser/${user_id}`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then((response) => response.json())
      .then((data) => {
        if (data && data.length > 0) {
          obtenerNombres(data);
        } else {
          setReservas([]);
        }
      })
      .catch((error) => console.error(error));
  }, []);

  const obtenerNombres = (reservasData) => {
    const hotelIds = [...new Set(reservasData.map((reserva) => reserva.hotel_id))];
    const userIds = [...new Set(reservasData.map((reserva) => reserva.user_id))];

    const fetchHotelPromises = hotelIds.map((hotelId) =>
      fetch(`http://localhost:8090/hotels/${hotelId}`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      }).then((response) => response.json())
    );

    Promise.all(fetchHotelPromises)
      .then((hotelData) => {
        const hoteles = {};
        hotelData.forEach((hotel) => {
          hoteles[hotel.id] = hotel.name;
        });

        const fetchUserPromises = userIds.map((userId) =>
          fetch(`http://localhost:8090/users/${userId}`, {
            method: "GET",
            headers: {
              "Content-Type": "application/json",
            },
          }).then((response) => response.json())
        );

        Promise.all(fetchUserPromises)
          .then((userData) => {
            const usuarios = {};
            userData.forEach((user) => {
              usuarios[user.id] = `${user.name} ${user.last_name}`;
            });
            const reservasActualizadas = reservasData.map((reserva) => ({
              ...reserva,
              hotel_nombre: hoteles[reserva.hotel_id] || "",
              usuario_nombre: usuarios[reserva.user_id] || "",
            }));

            setReservas(reservasActualizadas);
          })
          .catch((error) => console.error(error));
      })
      .catch((error) => console.error(error));
  };

  const formatDate = (date) => {
    const formattedDate = new Date(date);
    formattedDate.setDate(formattedDate.getDate() + 1); 

    const day = formattedDate.getDate();
    const month = formattedDate.getMonth() + 1;
    const year = formattedDate.getFullYear();

    return `${day.toString().padStart(2, '0')}/${month.toString().padStart(2, '0')}/${year}`;
  };

  return (
    <div className="contenedor-principal">
      <Header />
      <h4 className="titulo-reservas">Mis reservas:</h4>
      <div className="contenedor-reservas-usuario">
        {reservas.length > 0 ? (
          reservas.map((reserva, index) => (
            <div key={reserva.id} className="reserva-item">
              <p className="subtitulo-reserva"> Datos de la reserva {contadorReserva + index}:</p>
              <div className="detalle-reserva">
                <p>Usuario: {reserva.usuario_nombre}</p>
                <p>Hotel: {reserva.hotel_nombre}</p>
                <p>Fecha check-in: {formatDate(reserva.fecha_ingreso)}</p>
                <p>Fecha check-out: {formatDate(reserva.fecha_egreso)}</p>
              </div>
            </div>
          ))
        ) : (
          <p>No tiene reservas realizadas.</p>
        )}
      </div>
      <div className="contenedor-fechas">
        <br></br>
        <h2>INGRESE LAS FECHAS para su estadia</h2>
        <div className="fecha-desde">
          <p>Desde: </p>
          <DatePicker selected={fechaDesde} onChange={handleFechaDesdeChange} />
        </div>
        <div className="fecha-hasta">
          <p>Hasta: </p>
          <DatePicker selected={fechaHasta} onChange={handleFechaHastaChange} />
        </div>
        <button className="boton-buscar" onClick={buscarHotelesDisponibles}>BUSCAR</button>
      </div>
      <div className="contenedor-hoteles-r">
        {hotelesDisponibles.length != null ? (
          hotelesDisponibles.map((hotel) => (
            <div key={hotel.id}>
              <HotelesR
                nombreHotel={hotel.name}
                piezas={hotel.cantHabitaciones}
                descripcion={hotel.descripcion}
                hotelId={hotel.id}
                userId={user_id}
                fechaDesde={fechaDesde}
                fechaHasta={fechaHasta}
                token={token}
              />
            </div>
          ))
        ) : (
          <p>No hay hoteles disponibles en esas fechas.</p>
        )}
      </div>
    </div>
  );
}

export default Reserve;
