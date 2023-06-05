import React, { useEffect, useState } from 'react';
import { useNavigate, } from "react-router-dom";
import './../App.css'


const Login = () =>{
    const navigate = useNavigate();
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const register = () =>{
      navigate("/register");
    };
    const handleSubmit = async (e) => {
        e.preventDefault(); // para que no recarga la página
        if (email === '') {
            document.getElementById("inputEmailLogin").style.borderColor = 'red';
        }
        if (password === '') {
            document.getElementById("inputPasswordLogin").style.borderColor = 'red';
        }
        else{
            try {
                const response = await fetch('http://localhost:8080/login', {
                    method: 'POST', headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({email, password}),
                });
                if (response.ok) {
                    // El usuario está en la base de datos
                    console.log('Usuario válido');
                    navigate("/home")

                } else {
                    // El usuario no está en la base de datos o hay un error en el servidor
                    console.log('Usuario inválido');
                }
            } catch (error) {
                console.log('Error al realizar la solicitud al backend:', error);
            }
        }

    };

    return (
        <div id="body">
            <h1 id="h1Login">Iniciar sesión:</h1>
            <form id="formLogin" onSubmit={handleSubmit} >
                <input id={"inputEmailLogin"}
                    type="email"
                    placeholder="Correo electrónico"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                />
                <input id={"inputPasswordLogin"}
                    type="password"
                    placeholder="Contraseña"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}

                />
                <button id="botonLogin" type="submit">Iniciar sesión</button>
                <br/>
                <button id="botonRegistrarse" onClick={register}>Registrarse</button>
            </form>
        </div>
    );
};

export default Login;