import React, { useEffect, useState } from 'react' 
import { useParams } from 'react-router-dom';

const Hotels = () => {

    const [hotel, setHotel] = useState()
    const {id} = useParams()

    const infoHotel = `http://localhost:8080/hotel/${id}`

    const getHotel = async () => {
        const response = await fetch(infoHotel);
        const resolve = await response.json();
        setHotel(resolve)
    }

    useEffect(() => {
        getHotel();
    },[])

    return (
        <div>
            <h1>Detalles del hotel seleccionado: </h1>
            <div>
                <p>{hotel?.name}</p>
                <p>{hotel?.username}</p>
                <p>{hotel?.email}</p>
                <p>{hotel?.phone}</p>
                <p>{hotel?.website}</p>
            </div>
        </div>
    )
}

export default Hotels