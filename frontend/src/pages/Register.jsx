import React, { useEffect, useState } from 'react';
import { useNavigate } from "react-router-dom";
import './../App.css'

//creada solo para testear


//parte funcional
const Register = () =>{

    const navigate = useNavigate(); //permite la navegación entre paginasd con las rutas
    const [nombre, setNombre] = useState(''); //se inicializan las variables vacias
    const [apellido, setApellido] = useState('');
    const [dni, setDni] = useState('');
    const [email, setEmail] = useState(''); 
    const [password, setPassword] = useState('');
    const register = () =>{ // funbcion que te redirige a  home
      navigate("/home");
    };
    const handleSubmit = async (e) => { //recibe los datos del formulario
        e.preventDefault(); // para que no recarga la página
        if (nombre === '') {
            document.getElementById("inputNombreRegister").style.borderColor = 'red';
        }
        if (apellido === '') {
            document.getElementById("inputApellidoRegister").style.borderColor = 'red';
        }
        if (dni === '') {
            document.getElementById("inputDniRegister").style.borderColor = 'red';
        }
        if (email === '') {
            document.getElementById("inputEmailRegister").style.borderColor = 'red';
        }
        if (password === '') {
            document.getElementById("inputPasswordRegister").style.borderColor = 'red';
        }
        else{
            try {   //envía la respuesta al back (postaman basicamente)
                const response = await fetch('http://localhost:8080/login/register', {
                    method: 'POST', headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({nombre,apellido,dni,email, password}),
                });
                if (response.ok) {// si el usuario existe -> usuario creado en la db
                    // El usuario está en la base de datos
                    console.log('Usuario creado');
                    navigate("/home")

                } else {
                    // El usuario no (ya) está en la base de datos o hay un error en el servidor 
                    console.log('Usuario inválido');
                }
            } catch (error) {
                console.log('Error al realizar la solicitud al backend:', error);
            }
        }

    };

    //parte visible
    return(
        <div>
            <div id="body">
            <h1 id="h1Register">Resgistrarse</h1>
            <form id="formRegister" onSubmit={handleSubmit} >
                <input id={"inputNombreRegister"}
                    type="nombre"
                    placeholder="Nombre"
                    value={nombre}
                    onChange={(e) => setNombre(e.target.value)}
                />
                <input id={"inputApellidoRegister"}
                    type="apellido"
                    placeholder="Apellido"
                    value={apellido}
                    onChange={(e) => setApellido(e.target.value)}
                />
                <input id={"inputDniRegister"}
                    type="dni"
                    placeholder="DNI"
                    value={dni}
                    onChange={(e) => setDni(e.target.value)}
                />
                <input id={"inputEmailRegister"}
                    type="email"
                    placeholder="Correo electrónico"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                />
                <input id={"inputPasswordRegister"}
                    type="password"
                    placeholder="Contraseña"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}

                />
                <br/>
                <button id="botonRegistrarse" onClick={register}>Registrarse</button>
                
            </form>
            <br/>
            <button id="botonRegistrarse" onClick={register}>Atras</button>
        </div>
        </div>
        
    );
};

export default Register;