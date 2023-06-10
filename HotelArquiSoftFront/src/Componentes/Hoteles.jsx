    import React from "react"; //para usar react en el componente
    import '../../src/styleSheet/Hoteles.css';

    function Hoteles(props) {
        return (
          <div className='contenedor-hoteles'>
              <img className='imagen-hotel'
             src={(`HotelArquiSoftFront/src/Images/hotel-${this.props.image}.png`)}
             alt={'fotodeHotel1'}/>
              <div className='contenedor-detalle-hoteles'>
                  <p className='nombre-hotel1'> <strong>{this.props.nombreHotel}</strong> </p>
                  <p className='cantidad-piezas'>{this.props.piezas}</p>
             </div>
          </div>
        );
    }

    export default Hoteles;