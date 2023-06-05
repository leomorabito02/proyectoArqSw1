import React, { Fragment, useEffect, useState } from 'react';
import WifiIcon from '@mui/icons-material/Wifi';
import FitnessCenterIcon from '@mui/icons-material/FitnessCenter';
import LocalParkingIcon from '@mui/icons-material/LocalParking';
import PoolIcon from '@mui/icons-material/Pool';
import WaterDropIcon from '@mui/icons-material/WaterDrop';

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

    return (
        <Fragment>
            <div style={{ alignItems: 'left', backgroundColor: '#CBE4DE', minHeight: '100vh' }}>
                {hotel ? (
                    <div style={{ display: 'flex', alignItems: 'flex-start', justifyContent: 'flex-start', marginTop: '80px', marginBottom: '20px', marginLeft: '60px' }}>
                        <div style={{ alignItems: 'left', maxWidth: '50%' }}>
                            <h1 style={{ textAlign: 'left', color: '#0E8388' }}>Hotel {hotel.nombre}</h1>
                            <p style={{ textAlign: 'left', color: '#2C3333' }}>Estrellas: {hotel.valoracion}</p>
                            <p style={{ textAlign: 'left', color: '#2C3333', maxWidth: '80%' }}>Descripcion: {hotel.descripcion}</p>
                            <p style={{ textAlign: 'left', color: '#2C3333' }}>Precio por noche: ${hotel.precio}</p>
                            <h3 style={{ textAlign: 'left', color: '#0E8388' }}>Servicios principales:</h3>
                        </div>
                        <div style={{ alignItems: 'right', maxWidth: '50%' }}>
                            <input type="text" className="datepicker" />
                            <p>
                                <button style={{ textAlign: 'right', backgroundColor: '#2E4F4F' }}>Reservar</button>
                            </p>
                        </div>
                    </div>
                ) : (
                    <p>No se encontró el hotel.</p>
                )}

                <div style={{ display: 'flex', alignItems: 'center', marginLeft: '60px' }}>
                    <WifiIcon style={{ color: '#2C3333' }} />
                    <p style={{ color: '#2C3333', marginLeft: '5px' }}>Wifi</p>
                </div>

                {hotel.wifi ? (
                    <div style={{ display: 'flex', alignItems: 'center', marginLeft: '60px' }}>
                        <WifiIcon style={{ color: '#2C3333' }} />
                        <p style={{ color: '#2C3333', marginLeft: '5px' }}>Wifi</p>
                    </div>
                ) : null}

                {hotel.gym ? (
                    <div style={{ display: 'flex', alignItems: 'center', marginLeft: '60px' }}>
                        <FitnessCenterIcon style={{ color: '#2C3333' }} />
                        <p style={{ color: '#2C3333', marginLeft: '5px' }}>Gimnasio</p>
                    </div>
                ) : null}

                {hotel.estacionamiento ? (
                    <div style={{ display: 'flex', alignItems: 'center', marginLeft: '60px' }}>
                        <LocalParkingIcon style={{ color: '#2C3333' }} />
                        <p style={{ color: '#2C3333', marginLeft: '5px' }}>Estacionamiento</p>
                    </div>
                ) : null}

                {hotel.pileta ? (
                    <div style={{ display: 'flex', alignItems: 'center', marginLeft: '60px' }}>
                        <PoolIcon style={{ color: '#2C3333' }} />
                        <p style={{ color: '#2C3333', marginLeft: '5px' }}>Piscina</p>
                    </div>
                ) : null}

                {hotel.bide ? (
                    <div style={{ display: 'flex', alignItems: 'center', marginLeft: '60px' }}>
                        <WaterDropIcon style={{ color: '#2C3333' }} />
                        <p style={{ color: '#2C3333', marginLeft: '5px' }}>Bide</p>
                    </div>
                ) : null}


            </div>
        </Fragment>
    );
};

export default Reserva;