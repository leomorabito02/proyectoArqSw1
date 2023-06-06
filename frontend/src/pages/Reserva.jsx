import React, { useEffect, useState } from 'react';
import DateRangeComp from '../components/DateRangeComp';
import {useNavigate, useParams} from 'react-router-dom';

const Reserva = () => {
  const { id } = useParams();
  const [hotel, setHotel] = useState(null);
  const [startDate, setStartDate] = useState(null);
  const [endDate, setEndDate] = useState(null);
  const [errorMessage, setErrorMessage] = useState('');
  const navigate = useNavigate();
  const back= () =>{ // funbcion que te redirige a  login
    navigate("/home");
  };
  const handleSubmit = async (e) => {
    e.preventDefault();
    if (startDate === null) {
      console.log('Seleccione fecha de inicio');
    } else if (endDate === null) {
      console.log('Seleccione fecha de salida');
    } else {
      try {
        const startDateString = startDate.toISOString();
        const endDateString = endDate.toISOString();
        const response = await fetch(`http://localhost:8080/reserva/${id}`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ startDate: startDateString, endDate: endDateString }),
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
        const response = await fetch(`http://localhost:8080/hotel/${id}`);
        const data = await response.json();
        setHotel(data);
        //console.log(data)
      } catch (error) {
        console.log('Error al obtener el hotel:', error);
      }
    };

    fetchHotel();
  }, [id]);

  return (
      <div style={{ alignItems: 'left', backgroundColor: '#CBE4DE', minHeight: '100vh' }}>
        {hotel ? (
            <div style={{ display: 'flex', alignItems: 'flex-start', justifyContent: 'flex-start', marginTop: '80px', marginBottom: '20px', marginLeft: '60px' }}>
              <div style={{ alignItems: 'left', maxWidth: '50%' }}>
                <h1 style={{ textAlign: 'left', color: '#0E8388' }}>Hotel {hotel.nombre}</h1>
                <p style={{ textAlign: 'left', color: '#2C3333' }}>Estrellas: {hotel.valoracion}</p>
                <p style={{ textAlign: 'left', color: '#2C3333', maxWidth: '80%' }}>Descripcion: {hotel.descripcion}</p>
                <p style={{ textAlign: 'left', color: '#2C3333' }}>Precio por noche: ${hotel.precio}</p>
                <p style={{ textAlign: 'left', color: '#2C3333' }}> Gimnasio: {hotel.gym ? 'Sí' : 'No'} </p>
                <p style={{ textAlign: 'left', color: '#2C3333' }}> WIFI: {hotel.wifi ? 'Sí' : 'No'} </p>
                <p style={{ textAlign: 'left', color: '#2C3333' }}> Estacionamiento: {hotel.estacionamiento ? 'Sí' : 'No'} </p>
                <p style={{ textAlign: 'left', color: '#2C3333' }}> Bide: {hotel.bide ? 'Sí' : 'No'} </p>
                <p style={{ textAlign: 'left', color: '#2C3333' }}> Pileta: {hotel.pileta ? 'Sí' : 'No'} </p>
              </div>
              <div style={{ alignItems: 'right', maxWidth: '50%', marginTop: '0px' }}>
                <h3 style={{ textAlign: 'left', color: '#0E8388' }}>Seleccione check-in y check-out:</h3>
                <form id="formLogin" onSubmit={handleSubmit}>
                  <DateRangeComp onChange={(item) => { setStartDate(item.startDate); setEndDate(item.endDate); }} />
                  {errorMessage && <p style={{ color: 'red' }}>{errorMessage}</p>}
                  <button id="botonLogin" type="submit" style={{ textAlign: 'right', backgroundColor: '#2E4F4F', marginLeft: '15px' }}>Reservar</button>
                </form>
              </div>
            </div>
        ) : (
            <p>No se encontró el hotel</p>
        )}
        <button id="botonAtras" onClick={back}>Atras</button>
      </div>
  );
};

export default Reserva;
