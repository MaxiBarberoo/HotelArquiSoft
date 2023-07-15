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
    const [reservas, setReservas] = useState([]); // Agregar el estado de las reservas
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
                    // Obtener los ID's de los hoteles disponibles
                    const hotelIds = Object.values(data).map((item) => item.hotel_id);
                    // Realizar las solicitudes GET para obtener la información completa de los hoteles
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
                            // Actualizar el estado con la información completa de los hoteles
                            setHotelesDisponibles(hotelData);
                            // Actualizar el estado con las reservas del usuario obtenidas del backend
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
        // Obtener los IDs únicos de hoteles y usuarios en las reservas
        const hotelIds = [...new Set(reservasData.map((reserva) => reserva.hotel_id))];
        const userIds = [...new Set(reservasData.map((reserva) => reserva.user_id))];
      
        // Promesa para obtener los nombres de los hoteles
        const fetchHotelPromises = hotelIds.map((hotelId) =>
          fetch(`http://localhost:8090/hotels/${hotelId}`, {
            method: "GET",
            headers: {
              "Content-Type": "application/json",
            },
          }).then((response) => response.json())
        );
      
        // Obtener el nombre del hotel para cada reserva
        Promise.all(fetchHotelPromises)
          .then((hotelData) => {
            // Crear un objeto con el ID de hotel como clave y el nombre del hotel como valor
            const hoteles = {};
            hotelData.forEach((hotel) => {
              hoteles[hotel.id] = hotel.name;
            });
      
            // Promesa para obtener los nombres y apellidos de los usuarios
            const fetchUserPromises = userIds.map((userId) =>
              fetch(`http://localhost:8090/users/${userId}`, {
                method: "GET",
                headers: {
                  "Content-Type": "application/json",
                },
              }).then((response) => response.json())
            );
      
            // Obtener el nombre y apellido del usuario para cada reserva
            Promise.all(fetchUserPromises)
              .then((userData) => {
                // Crear un objeto con el ID de usuario como clave y el nombre y apellido del usuario como valor
                const usuarios = {};
                userData.forEach((user) => {
                  usuarios[user.id] = `${user.name} ${user.last_name}`;
                });
      
                // Actualizar las reservas con los nombres de hoteles y nombres y apellidos de usuarios
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
        formattedDate.setDate(formattedDate.getDate() + 1); // Sumar 1 día
      
        const day = formattedDate.getDate();
        const month = formattedDate.getMonth() + 1; // Los meses en JavaScript se indexan desde 0
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
