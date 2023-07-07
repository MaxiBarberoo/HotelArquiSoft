import React from 'react'

function Admin(){
    const [reservasTotales, setReservasTotales] = useState([]);

    const handleNuevoHotelChange = (event) => {
        const { name, value } = event.target;
        setNuevoHotel((prevNuevoHotel) => ({
            ...prevNuevoHotel,
            [name]: value
        }));
    };

    const crearNuevoHotel = () => {
        const cantHabitaciones = parseInt(nuevoHotel.cantidadHabitaciones);
        fetch('http://localhost:8090/hotels', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                name: nuevoHotel.nombre,
                cantHabitaciones: cantHabitaciones
            })
        })
            .then((response) => response.json())
            .then((data) => {
                if (data.error) {
                    console.error(data.error);
                } else {
                    setHoteles((prevHoteles) => [...prevHoteles, data]);
                    alert('Se ha creado un nuevo hotel con éxito');
                    setNuevoHotel({
                        nombre: '',
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


    return(
        <div>

        </div>
    );
}

export default Admin;