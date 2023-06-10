// eslint-disable-next-line no-unused-vars
    import React from "react"; //para usar react en el componente
    import '../../src/styleSheet/Hoteles.css';

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
    }

    export default Hoteles;