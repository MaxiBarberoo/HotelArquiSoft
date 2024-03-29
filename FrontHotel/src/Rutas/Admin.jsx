import React, { useState, useEffect } from "react"
import '../Stylesheet/Admin.css'
import Header from '../Componentes/Header'
import HotelesAdmin from '../Componentes/HotelesAdmin'
import { useParams } from 'react-router-dom';

function Admin() {
  const [reservasTotales, setReservasTotales] = useState([]);
  const [nuevoHotel, setNuevoHotel] = useState({
    nombre: '',
    descripcion: '',
    cantidadHabitaciones: 0,
    imagenSeleccionada: null,
  });
  const [hoteles, setHoteles] = useState([]);
  const [amenities, setAmenities] = useState([]);
  const { token } = useParams();
  const { user_id } = useParams();
  const [contadorReserva, setContadorReserva] = useState(1);

  const crearNuevoHotel = (event) => {
    event.preventDefault();
    const cantiHabitaciones = parseInt(nuevoHotel.cantidadHabitaciones);
    if (
      nuevoHotel.cantidadHabitaciones <= 0 ||
      nuevoHotel.nombre === "" ||
      nuevoHotel.descripcion === "" ||
      !nuevoHotel.imagenSeleccionada // Cambiamos la condición a una sola imagen seleccionada
    ) {
      alert("El nombre, descripción y una imagen son requeridos. Inténtelo nuevamente.");
      window.location.reload();
    } else {
      const hotelData = {
        name: nuevoHotel.nombre,
        descripcion: nuevoHotel.descripcion,
        cantHabitaciones: cantiHabitaciones,
      };

      fetch("http://localhost:8090/hotels", {
        method: "POST",
        headers: {
          Authorization: `${token}`,
          "Content-Type": "application/json",
        },
        body: JSON.stringify(hotelData),
      })
        .then((response) => response.json())
        .then((data) => {
          if (data.error) {
            console.error(data.error);
          } else {
            const hotelId = data.id;
            const imagenData = {
              hotel_id: hotelId,
              contenido: nuevoHotel.imagenSeleccionada, 
            };

            fetch("http://localhost:8090/imagenes", {
              method: "POST",
              headers: {
                Authorization: `${token}`,
                "Content-Type": "application/json",
              },
              body: JSON.stringify(imagenData), 
            })
              .then((response) => response.json())
              .then((data) => {
                if (data.error) {
                  console.error(data.error);
                } else {
                  setHoteles((prevHoteles) => [...prevHoteles, data]);
                  alert("Se ha creado un nuevo hotel con éxito");
                  setNuevoHotel({
                    nombre: "",
                    descripcion: "",
                    cantidadHabitaciones: 0,
                    imagenSeleccionada: null, // Limpiamos la imagen seleccionada
                  });
                  window.location.reload();
                }
              })
              .catch((error) => console.error(error));
          }
        })
        .catch((error) => console.error(error));
    }
  };

  const handleImagenesChange = (event) => {
    const { files } = event.target;
    if (files && files.length > 0) {
      const imagenFile = files[0]; // Tomamos solo la primera imagen seleccionada
      const reader = new FileReader();
      reader.onloadend = () => {
        const buffer = new Uint8Array(reader.result); // Obtenemos el Uint8Array
        const byteArray = Array.from(buffer); // Convertimos el Uint8Array a []byte

        setNuevoHotel((prevNuevoHotel) => ({
          ...prevNuevoHotel,
          imagenSeleccionada: byteArray, // Almacenamos el []byte en el estado
        }));
      };
      reader.readAsArrayBuffer(imagenFile);
    }
  };


  const handleNombreHotelChange = (event) => {
    const { value } = event.target;
    setNuevoHotel((prevNuevoHotel) => ({
      ...prevNuevoHotel,
      nombre: value,
    }));
  };

  const handleDescripcionHotelChange = (event) => {
    const { value } = event.target;
    setNuevoHotel((prevNuevoHotel) => ({
      ...prevNuevoHotel,
      descripcion: value,
    }));
  };

  const handleCantidadHabitacionesChange = (event) => {
    const { value } = event.target;
    setNuevoHotel((prevNuevoHotel) => ({
      ...prevNuevoHotel,
      cantidadHabitaciones: parseInt(value),
    }));
  };

  useEffect(() => {
    fetch('http://localhost:8090/hotels')
      .then((response) => response.json())
      .then((data) => setHoteles(data))
      .catch((error) => console.error(error));

    fetch('http://localhost:8090/reservas', {
      method: 'GET',
      headers: {
        "Content-Type": 'application/json',
        "Authorization": `${token}`
      }
    })
      .then((response) => response.json())
      .then((data) => {
        if (data && data.length > 0) {
          obtenerNombres(data);
        }
      })
      .catch((error) => console.error(error));
  }, []);

  useEffect(() => {
    const fetchAmenitiesForHotels = async () => {
      const hotelsWithAmenities = await Promise.all(
          amenities.map(async (amenitie) => {
            const response = await fetch(`http://localhost:8090/amenities/${amenitie.id}`);
            if (response.ok) {
              const amenitiesData = await response.json();
              return { ...amenitie, amenities: amenitiesData };
            } else {
              console.error(`Error en la petición GET de amenities para el hotel ${hotel.id}`);
              return amenitie;
            }
          })
      );
      setHoteles(hotelsWithAmenities);
    };

    if (amenities.length > 0) {
      fetchAmenitiesForHotels();
    }
  }, [hoteles]);

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

            setReservasTotales(reservasActualizadas);
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
    <div>
      <Header />
      {reservasTotales.length > 0 ? (
        reservasTotales.map((reserva, index) => (
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
      <h2>Agregar nuevo hotel</h2>
      <div className="contenedor-crear-hoteles">
        <p>Nombre del hotel:</p>
        <input
          type="text"
          name="nombre"
          value={nuevoHotel.nombre}
          onChange={handleNombreHotelChange}
          placeholder="Nombre del hotel"
        />

        <input
          type="text"
          maxLength={500}
          name="descripcion"
          value={nuevoHotel.descripcion}
          onChange={handleDescripcionHotelChange}
          placeholder="Descripción"
        />

        <input
          type="number"
          name="cantidadHabitaciones"
          value={nuevoHotel.cantidadHabitaciones}
          onChange={handleCantidadHabitacionesChange}
          placeholder="Cantidad de habitaciones"
        />

        <input
          type="file"
          name="imagen"
          onChange={handleImagenesChange}
        />

        <button className="boton-crear-hotel" onClick={crearNuevoHotel}>Crear Hotel</button>
      </div>
      <div className="contenedor-hoteles-admin">
        <h2>Listado de Hoteles</h2>
        {hoteles.map((hotel) => (
          <HotelesAdmin
            key={hotel.id}
            hotelId={hotel.id}
            imagenesURLs={hotel.imagenesURLs}
            nombreHotel={hotel.name}
            piezas={hotel.cantHabitaciones}
            descripcion={hotel.descripcion}
            amenities={hotel.amenities}
            token={token}
          />
        ))}
      </div>
    </div>
  );
}

export default Admin;