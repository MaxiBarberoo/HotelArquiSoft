import React from 'react';
import '../Stylesheet/Hoteles.css';

function Hoteles(props) {
    return (
      <div className='contenedor-hoteles'>
         <div className='contenedor-detalle-hoteles'>
             <p className='titulo-hotel'>
                 <strong>{props.nombreHotel}.</strong>
             </p>
             <p className='cantidad-piezas'>
                 Cantidad de habitaciones: 2.
             </p>
         </div>
             <div className='fecha-desde'>
                 <p className='texto-fecha-desde'>
                     Fecha de ingreso:
                 </p>
                 <input type='date' className='fecha-desde-reserva' />
                 <p className='texto-fecha-hasta'>
                     Fecha de egreso:
                 </p>
                 <input type='date' className='fecha-hasta-reserva' />
             </div>

      </div>
    );
}

export default Hoteles;