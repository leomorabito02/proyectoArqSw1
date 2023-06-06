import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import './../App.css';

const Register = () => {
    const navigate = useNavigate();
    const [nombre, setNombre] = useState('');
    const [apellido, setApellido] = useState('');
    const [dni, setDni] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const register1 = () => {
        navigate('/login');
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        if (nombre === '') {
            document.getElementById('inputNombreRegister').style.borderColor = 'red';
        }
        if (apellido === '') {
            document.getElementById('inputApellidoRegister').style.borderColor = 'red';
        }
        if (dni === '') {
            document.getElementById('inputDniRegister').style.borderColor = 'red';
        }
        if (email === '') {
            document.getElementById('inputEmailRegister').style.borderColor = 'red';
        }
        if (password === '') {
            document.getElementById('inputPasswordRegister').style.borderColor = 'red';
        } else {
            try {
                const response = await fetch('http://localhost:8080/register', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        nombre,
                        apellido,
                        dni: parseInt(dni),
                        email,
                        password,
                    }),
                });
                if (response.ok) {
                    console.log('Usuario creado');
                    navigate('/home');
                } else {
                    console.log('Usuario inválido');
                }
            } catch (error) {
                console.log('Error al realizar la solicitud al backend:', error);
            }
        }
    };

    return (
        <div>
            <div id="body">
                <h1 id="h1Register">Registrarse</h1>
                <form id="formRegister" onSubmit={handleSubmit}>
                    <input
                        id="inputNombreRegister"
                        type="text"
                        placeholder="Nombre"
                        value={nombre}
                        onChange={(e) => setNombre(e.target.value)}
                    />
                    <input
                        id="inputApellidoRegister"
                        type="text"
                        placeholder="Apellido"
                        value={apellido}
                        onChange={(e) => setApellido(e.target.value)}
                    />
                    <input
                        id="inputDniRegister"
                        type="text"
                        placeholder="DNI"
                        value={dni}
                        onChange={(e) => setDni(e.target.value)}
                    />
                    <input
                        id="inputEmailRegister"
                        type="email"
                        placeholder="Correo electrónico"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                    />
                    <input
                        id="inputPasswordRegister"
                        type="password"
                        placeholder="Contraseña"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                    />
                    <br />
                    <button id="botonRegistrarse" type="submit">
                        Registrarse
                    </button>
                </form>
                <br />
                <button id="botonRegistrarse" onClick={register1}>
                    Atrás
                </button>
            </div>
        </div>
    );
};

export default Register;
