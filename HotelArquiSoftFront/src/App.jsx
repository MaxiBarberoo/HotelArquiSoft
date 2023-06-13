import React, {useState, useEffect} from 'react';
import BotonLogin from './Componentes/BotonLogin';
import BotonRegister from './Componentes/BotonRegister';
import Hoteles from './Componentes/Hoteles.jsx';
import Reservas from './Componentes/Reservas.jsx';

function App() {
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const [isAdmin, setIsAdmin] = useState(false);
    const [hoteles, setHoteles] = useState([]);
    const [userId, setUserId] = useState(null);
    const [mostrarReservas, setMostrarReservas] = useState(false);
    const [reservas, setReservas] = useState([]);
    const [nuevoHotel, setNuevoHotel] = useState({
        nombre: "",
        cantidadHabitaciones: 0
    });
    const [reservasTotales, setReservasTotales] = useState([]);

    const handleNuevoHotelChange = (event) => {
        const {name, value} = event.target;
        setNuevoHotel((prevNuevoHotel) => ({
            ...prevNuevoHotel,
            [name]: value
        }));
    };

    const crearNuevoHotel = () => {
        const cantHabitaciones = parseInt(nuevoHotel.cantidadHabitaciones);
        fetch("http://localhost:8090/hotels", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                name: nuevoHotel.nombre,
                cantHabitaciones: cantHabitaciones
            })
        })
            .then((response) => response.json())
            .then((data) => {
                // Verificar si se produjo un error en la respuesta
                if (data.error) {
                    // Manejar el error, mostrar mensaje, etc.
                    console.error(data.error);
                } else {
                    // Actualizar la lista de hoteles con el nuevo hotel agregado
                    setHoteles((prevHoteles) => [...prevHoteles, data]);
                    alert("Se ha creado un nuevo hotel con exito");
                    setNuevoHotel({
                        nombre: "",
                        cantidadHabitaciones: 0
                    });
                }
            })
            .catch((error) => console.error(error));
    };

    const handleLogin = (tipoUsuario, userId) => {
        setIsLoggedIn(true);
        setIsAdmin(tipoUsuario === 1);
        setUserId(userId);
    };

    const toggleReservas = () => {
        setMostrarReservas(!mostrarReservas);

        if (isAdmin) {
            fetch(`http://localhost:8090/reservas`, {
                method: "GET",
                headers: {
                    "Content-Type": "application/json",
                },
            })
                .then((response) => response.json())
                .then((data) => setReservasTotales(data))
                .catch((error) => console.error(error));
        }
    };

    useEffect(() => {
        fetch(`http://localhost:8090/hotels`)
            .then((response) => response.json())
            .then((data) => setHoteles(data))
            .catch((error) => console.error(error));

        if (isLoggedIn) {
            fetch(`http://localhost:8090/reservas/reservauser/${userId}`, {
                method: "GET",
                headers: {
                    "Content-Type": "application/json",
                },
            })
                .then((response) => response.json())
                .then((data) => setReservas(data))
                .catch((error) => console.error(error));
        }
    }, [isLoggedIn, userId]);

    return (
        <div className="App">
            <div>
                <h1>ENCONTRA LA MEJOR OPCION</h1>
                {!isLoggedIn && <BotonLogin handleLogin={handleLogin}/>}
                {!isLoggedIn && <BotonRegister/>}
                {isLoggedIn && (
                    <div>
                        <button onClick={toggleReservas}>Reservas</button>
                        {mostrarReservas && <Reservas reservas={reservas}/>}
                    </div>)}
                {isAdmin && (
                    <div>
                        <h2>Agregar nuevo hotel</h2>
                        <input
                            type="text"
                            name="nombre"
                            value={nuevoHotel.nombre}
                            onChange={handleNuevoHotelChange}
                            placeholder="Nombre del hotel"
                        />
                        <input
                            type="number"
                            name="cantidadHabitaciones"
                            value={nuevoHotel.cantidadHabitaciones}
                            onChange={handleNuevoHotelChange}
                            placeholder="Cantidad de habitaciones"
                        />
                        <button onClick={crearNuevoHotel}>Crear Hotel</button>
                        {isAdmin && (
                            <div>
                                <button onClick={toggleReservas}>Reservas</button>
                                {mostrarReservas && (
                                    <Reservas reservas={reservas} reservasTotales={reservasTotales}/>
                                )}
                            </div>
                        )}
                        <div>
                            <h2>Listado de Hoteles</h2>
                            {hoteles.length > 0 ? (
                                <ul>
                                    {hoteles.map((hotel) => (
                                        <li key={hotel.id}>
                                            <p>Nombre: {hotel.name}</p>
                                            <p>Cantidad de Habitaciones: {hotel.cantHabitaciones}</p>
                                        </li>
                                    ))}
                                </ul>
                            ) : (
                                <p>No se han creado hoteles todavía.</p>
                            )}
                        </div>
                    </div>
                )}
                {!isAdmin && hoteles.map(hotel => (
                    <Hoteles
                        key={hotel.id}
                        nombreHotel={hotel.name}
                        piezas={hotel.cantHabitaciones}
                        isLoggedIn={isLoggedIn}
                        hotelId={hotel.id}
                        userId={userId}
                    />
                ))}
            </div>
        </div>
    );
}

export default App;
