import React, { useState } from 'react';
import FormularioRegister from './FormularioRegister';

function BotonRegister() {
  const [isOpen, setIsOpen] = useState(false);

  const openPopup = () => {
    setIsOpen(true);
  };

  const closePopup = () => {
    setIsOpen(false);
  };

  return (
    <>
      <button onClick={openPopup}>Registrarse</button>
      {isOpen && (
        <div className="popup">
          <FormularioRegister />
          <button onClick={closePopup}>Cerrar</button>
        </div>
      )}
    </>
  );
}

export default BotonRegister;