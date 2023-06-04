import React, { useEffect, useState } from 'react';
import { useNavigate } from "react-router-dom";


const navigate = useNavigate();
const Login = () =>{
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const register = () =>{
        navigate("/Register");
    };
    const handleSubmit = (e) => {


        if(email === '' || password === '' ){

        }

    };

    return (
        <div id="body">
            <h1 id="h1Login">Iniciar sesión:</h1>
            <form id="formLogin" onSubmit={handleSubmit} >
                <input id={"inputLogin"}
                    type="email"
                    placeholder="Correo electrónico"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                />
                <input id={"inputLogin"}
                    type="password"
                    placeholder="Contraseña"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}

                />
                <button id="botonLogin" type="submit">Iniciar sesión</button>
                <br/>
                <button id="botonLogin" onClick={register}>Registrarse</button>
            </form>
        </div>
    );
};

export default Login;