import React, { useState } from 'react';
import FormularioLogin from './FormularioLogin';

function BotonLogin() {
  const [isOpen, setIsOpen] = useState(false);

  const openPopup = () => {
    setIsOpen(true);
  };

  const closePopup = () => {
    setIsOpen(false);
  };

  return (
    <>
      <button onClick={openPopup}>Iniciar sesión</button>

      {isOpen && (
        <div className="popup">
          <FormularioLogin />
          <button onClick={closePopup}>Cerrar</button>
        </div>
      )}
    </>
  );
}

export default BotonLogin;