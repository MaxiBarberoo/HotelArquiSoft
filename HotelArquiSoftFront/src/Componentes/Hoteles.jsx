import React, { useState } from "react"; 
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
import '../../src/styleSheet/Hoteles.css';

<<<<<<< HEAD
    function Hoteles(props) {
        return (
          <div className='contenedor-hoteles'>
              <img className='imagen-hotel'
                   /* eslint-disable-next-line react/prop-types */
             src={(`src/Images/hotel-${props.image}.png`)}
             alt={'fotodeHotel'}/>
              <div className='contenedor-detalle-hoteles'>
                  {/* eslint-disable-next-line react/prop-types */}
                  <p className='nombre-hotel1'> <strong>{props.nombreHotel}</strong> </p>
                  {/* eslint-disable-next-line react/prop-types */}
                  <p className='cantidad-piezas' >Cantidad de habitaciones: {props.piezas}</p>
             </div>
          </div>
        );
=======
function Hoteles(props) { 
    const [fechaDesde, setFechaDesde] = useState(null);
    const [fechaHasta, setFechaHasta] = useState(null);

    const handleFechaDesdeChange = (date) => {
        setFechaDesde(date);
      };

    const handleFechaHastaChange = (date) => {
        setFechaHasta(date);
    };

    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const [isFormValid, setIsFormValid] = useState(true);

    const handleReservaClick = () => {
        if (!isLoggedIn) {
          alert('Debes iniciar sesión para realizar la reserva.');
        } else if (!fechaDesde || !fechaHasta) {
          setIsFormValid(false);
          alert('Debes completar los campos de fecha desde y fecha hasta.');
        } else if (fechaDesde >= fechaHasta) {
          setIsFormValid(false);
          alert('La fecha desde debe ser anterior a la fecha hasta.');
        } else {
          // Aquí puedes realizar la lógica de reserva si todas las validaciones pasan
          alert('¡Reserva realizada con éxito!');
        }
    };

return (
    <div className='contenedor-hoteles'>
        <img className='imagen-hotel' src={(`src/Images/hotel-${props.image}.png`)} alt={'fotodeHotel'}/>
        <div className='contenedor-detalle-hoteles'>
            <p className='nombre-hotel1'> <strong>{props.nombreHotel}</strong> </p>
            <p className='cantidad-piezas'>{props.piezas}</p>
        </div>
        <div className="contenedor-inputs-fechas">
            <p>Desde: </p>
            <DatePicker selected={fechaDesde} onChange={handleFechaDesdeChange} />
            <p>Hasta: </p>
            <DatePicker selected={fechaHasta} onChange={handleFechaHastaChange} />
        </div>
        <div className="boton-reserva">
            <button onClick={handleReservaClick}>Reservar</button>
        </div>
    </div>
    );
>>>>>>> 6e49f87269bdeb9c23bfb518fb72bf6788cce12f
    }

    export default Hoteles;