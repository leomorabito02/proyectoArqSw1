import React, { Fragment, useEffect, useState } from 'react';
import DateRangeComp from '../components/DateRangeComp';


const Reserva = () => {
    const [hotel, setHotel] = useState(null);

    useEffect(() => {
        const fetchHotel = async () => {
            try {
                const response = await fetch('http://localhost:8080/hotel/1');
                const data = await response.json();
                setHotel(data);
            } catch (error) {
                console.log('Error al obtener el hotel:', error);
            }
        };

        fetchHotel();
    }, []);
    console.log(hotel)
    
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
                                <DateRangeComp />
                                <button style={{ textAlign: 'right', backgroundColor: '#2E4F4F', marginLeft: '15px' }}>Reservar</button>
                        </div>
                    </div>
                ) : (
                    <p>No se encontró el hotel</p>
                )}


            </div>
    );
};

export default Reserva;