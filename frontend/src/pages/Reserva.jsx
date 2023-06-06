import React, { Fragment, useEffect, useState } from 'react';
import DateRangeComp from '../components/DateRangeComp';

const Reserva = () => {
  const [hotel, setHotel] = useState(null);
  const [errorMessage, setErrorMessage] = useState('');

  const handleSubmit = async (e, startDate, endDate) => {
    e.preventDefault();
    console.log(startDate)
    console.log(endDate)
    if (selectedStartDate === null) {
      setErrorMessage('Seleccione fecha de inicio');
    } else if (selectedEndDate === null) {
      setErrorMessage('Seleccione fecha de salida');
    } else {
      try {
        const startDateString = startDate.toISOString(); // Serializar fecha de inicio
        const endDateString = endDate.toISOString(); // Serializar fecha de salida
        const response = await fetch('http://localhost:8080/reserva', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ startDate, endDate }),
        });
        if (response.ok) {
          console.log('Se ha registrado su reserva');
        } else {
          console.log('No hay habitaciones disponibles');
        }
      } catch (error) {
        console.log('Error al realizar la solicitud al backend:', error);
      }
    }
  };

  useEffect(() => {
    const fetchHotel = async () => {
      try {
        const response = await fetch('http://localhost:8080/hotel/2');
        const data = await response.json();
        setHotel(data);
      } catch (error) {
        console.log('Error al obtener el hotel:', error);
      }
    };

    fetchHotel();
  }, []);
  console.log(hotel);

  return (
    <div style={{ alignItems: 'left', backgroundColor: '#CBE4DE', minHeight: '100vh' }}>
      {hotel ? (
        <div style={{ display: 'flex', alignItems: 'flex-start', justifyContent: 'flex-start', marginTop: '80px', marginBottom: '20px', marginLeft: '60px' }}>
          <div style={{ alignItems: 'left', maxWidth: '50%' }}>
            <h1 style={{ textAlign: 'left', color: '#0E8388' }}>Hotel {hotel.nombre}</h1>
            <p style={{ textAlign: 'left', color: '#2C3333' }}>Estrellas: {hotel.valoracion}</p>
            <p style={{ textAlign: 'left', color: '#2C3333', maxWidth: '80%' }}>Descripcion: {hotel.descripcion}</p>
            <p style={{ textAlign: 'left', color: '#2C3333' }}>Precio por noche: ${hotel.precio}</p>
          </div>
          <div style={{ alignItems: 'right', maxWidth: '50%', marginTop: '0px' }}>
            <h3 style={{ textAlign: 'left', color: '#0E8388' }}>Seleccione check-in y check-out:</h3>
            <form id="formLogin" onSubmit={(e) => handleSubmit(e, startDate, endDate)}>
              <DateRangeComp onChange={(item) => { setStartDate(item.startDate); setEndDate(item.endDate); }} />
              {errorMessage && <p style={{ color: 'red' }}>{errorMessage}</p>}
              <button id="botonLogin" type="submit" style={{ textAlign: 'right', backgroundColor: '#2E4F4F', marginLeft: '15px' }}>Reservar</button>
            </form>
          </div>
        </div>
      ) : (
        <p>No se encontró el hotel</p>
      )}
    </div>
  );
};

export default Reserva;