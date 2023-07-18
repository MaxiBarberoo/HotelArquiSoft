import React, { useState, useEffect } from "react"
import '../Stylesheet/Admin.css'
import Header from '../Componentes/Header'
import HotelesAdmin from '../Componentes/HotelesAdmin'
import { useParams } from 'react-router-dom'
import DatePicker from 'react-datepicker'
import 'react-datepicker/dist/react-datepicker.css'

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
  const [todosAmenities, setTodosAmenities] = useState([]);
  const [eleccionDeAmenitie, setEleccionDeAmenitie] = useState(null);
  const [filtroBusqueda, setfiltroBusqueda] = useState(1);
  const [fechaDesdeFiltro, setFechaDesdeFiltro] = useState(null);
  const [fechaHastaFiltro, setFechaHastaFiltro] = useState(null);
  const [nombreHotel, setNombreHotel] = useState('');

  const handleFechaDesdeChangeFiltro = (date) => {
    setFechaDesdeFiltro(date);
  }
  const handleFechaHastaChangeFiltro = (date) => {
    setFechaHastaFiltro(date);
  }

  const filtrarReservasFecha = async () => {
    setReservasTotales([]);
    if (fechaDesdeFiltro && fechaHastaFiltro) {
      console.log("entró a filtros por fecha");
      console.log(fechaDesdeFiltro);
      console.log(fechaHastaFiltro);
      try {
        const response = await fetch("http://localhost:8090/reservas/byfecha", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            "Authorization": `${token}`,
          },
          body: JSON.stringify({
            fecha_ingreso: fechaDesdeFiltro,
            fecha_egreso: fechaHastaFiltro,
          }),
        })
          .then((response) => response.json())
          .then((data) => {
            if (data && data.length > 0) {
              obtenerNombres(data);
            } else {
              setReservasTotales([]);
            }
          })
          .catch((error) => console.error(error));
      } catch (error) {
        console.error(error);
      }
    } else {
      alert("Los campos están incompletos. Ingréselos nuevamente");
      window.location.reload();
    }
  };

  const filtrarReservasHotel = async () => {
    setReservasTotales([]);
    if (nombreHotel) {
      const hotelCoincidente = hoteles.find((hotel) => hotel.name === nombreHotel);
      if (hotelCoincidente) {
        const hotelId = hotelCoincidente.id;
        try {
          const response = await fetch(`http://localhost:8090/reservas/hotel/${hotelId}`, {
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
                setReservasTotales([]);
              }
            })
            .catch((error) => console.error(error));

        } catch (error) {
          console.error(error);
        }
      } else {
        alert("El nombre de hotel ingresado es incorrecto y no es posible filtrar. Inténtelo nuevamente.");
        window.location.reload();
      }
    } else {
      alert("El campo está incompleto. Ingréselo nuevamente.");
      window.location.reload();
    }
  };

  const fetchTodasLasReservas = () => {
    fetch("http://localhost:8090/reservas", {
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
          setReservasTotales([]);
        }
      })
      .catch((error) => console.error(error));
  };

  useEffect(() => {
    if (filtroBusqueda) {
      if (fechaDesdeFiltro && fechaHastaFiltro) {
        filtrarReservasFecha();
      } else {
        if (nombreHotel) {
          filtrarReservasHotel();
        } else {
          fetchTodasLasReservas();
        }
      }
    } else {
      fetchTodasLasReservas();
    }
  }, [filtroBusqueda]);


  const HandleOpcionAmenitieChange = (event) => {
    event.preventDefault();
    setEleccionDeAmenitie(event.target.value);
  }

  useEffect(() => {
    const fetchTodosAmenities = async () => {
      try {
        const response = await fetch("http://localhost:8090/amenities");
        if (response.ok) {
          const data = await response.json();
          setTodosAmenities(data);
        } else {
          throw new Error(`Error en la petición GET de las amenidades para el hotel ${props.hotelId}`);
        }
      } catch (error) {
        console.error('Error:', error);
      }
    };
    fetchTodosAmenities();
  }, []);

  const crearNuevoHotel = (event) => {
    event.preventDefault();
    const cantiHabitaciones = parseInt(nuevoHotel.cantidadHabitaciones);
    if (
      nuevoHotel.cantidadHabitaciones <= 0 ||
      eleccionDeAmenitie <= 0 || eleccionDeAmenitie > 6 ||
      nuevoHotel.nombre === "" ||
      nuevoHotel.descripcion === "" ||
      !nuevoHotel.imagenSeleccionada
    ) {
      alert("El nombre, descripción, una imagen y un amenitie son requeridos. Inténtelo nuevamente.");
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
                  const amenitieElegidaId = parseInt(eleccionDeAmenitie);

                  const amenitieData = {
                    hotel_id: hotelId,
                    amenitie_id: amenitieElegidaId,
                  };

                  fetch("http://localhost:8090/amenitiehotel/assign", {
                    method: "POST",
                    headers: {
                      Authorization: `${token}`,
                      "Content-Type": "application/json",
                    },
                    body: JSON.stringify(amenitieData),
                  })
                    .then((response) => response.json())
                    .then((data) => {
                      if (data.error) {
                        console.error(data.error);
                      } else {
                        setEleccionDeAmenitie(null);
                      }
                    })
                    .catch((error) => console.error(error));
                  alert("Se ha creado un nuevo hotel con éxito");
                  setNuevoHotel({
                    nombre: "",
                    descripcion: "",
                    cantidadHabitaciones: 0,
                    imagenSeleccionada: null,
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
      const imagenFile = files[0];
      const reader = new FileReader();
      reader.onloadend = () => {
        const buffer = new Uint8Array(reader.result);
        const byteArray = Array.from(buffer);

        setNuevoHotel((prevNuevoHotel) => ({
          ...prevNuevoHotel,
          imagenSeleccionada: byteArray,
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
      <h2>Filtros de reservas:</h2>
      <form onSubmit={(e) => {
        e.preventDefault();
        setfiltroBusqueda(filtroBusqueda + 1);
      }}>
        <h5>Filtrar por fecha:</h5>
        <p>Fecha desde:</p>
        <DatePicker selected={fechaDesdeFiltro} onChange={handleFechaDesdeChangeFiltro} />
        <p>Fecha hasta:</p>
        <DatePicker selected={fechaHastaFiltro} onChange={handleFechaHastaChangeFiltro} />
        <button type="submit">Filtrar</button>
      </form>
      <form onSubmit={(e) => {
        e.preventDefault();
        setfiltroBusqueda(filtroBusqueda + 1);
      }}>
        <h5>Filtrar por nombre de hotel:</h5>
        <input type="text" placeholder="Nombre del hotel" value={nombreHotel} onChange={(e) => setNombreHotel(e.target.value)} />
        <button type="submit">Filtrar</button>
      </form>
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
        <p>Descripcion:</p>
        <input
          type="text"
          maxLength={500}
          name="descripcion"
          value={nuevoHotel.descripcion}
          onChange={handleDescripcionHotelChange}
          placeholder="Descripción"
        />
        <p>Cantidad de habitaciones:</p>
        <input
          type="number"
          name="cantidadHabitaciones"
          value={nuevoHotel.cantidadHabitaciones}
          onChange={handleCantidadHabitacionesChange}
          placeholder="Cantidad de habitaciones"
        />
        <p>Imagen:</p>
        <input
          type="file"
          name="imagen"
          onChange={handleImagenesChange}
        />
        <p>Amenities para cargar:</p>
        <ol>
          {todosAmenities.sort((a, b) => a.id - b.id).map(amenitie => (
            <li key={amenitie.id}>{amenitie.tipo}</li>
          ))}
        </ol>
        <p>Elija el número de amenitie a cargar para el hotel:</p>
        <input type="number" onChange={HandleOpcionAmenitieChange}></input>
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