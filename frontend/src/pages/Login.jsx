import React, { useEffect, useState } from 'react';

const Login = () =>{

    let [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const handleSubmit = (e) => {

        if(email === '' || password === '' ){

        }
        console.log('Email:', email);
        console.log('Password:', password);
    };

    return (
        <div style={{ alignItems: 'center', backgroundColor: '#CBE4DE', minHeight: '100vh' }}>
            <h1 style={{ textAlign: 'center', color:'#0E8388'}}>Iniciar sesión:</h1>
            <form onSubmit={handleSubmit} style={{ display: 'grid', flexDirection: 'column', alignItems: 'center' }}>
                <input
                    type="email"
                    placeholder="Correo electrónico"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    style={{ marginBottom: '10px', padding: '5px' }}
                />
                <input
                    type="password"
                    placeholder="Contraseña"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    style={{ marginBottom: '10px', padding: '5px' }}
                />
                <button type="submit" style={{ backgroundColor: '#2E4F4F', color: '#FFF', padding: '8px 16px' }}>Iniciar sesión</button>
                <br/>
                <button type="submit" style={{ backgroundColor: '#2E4F4F', color: '#FFF', padding: '8px 16px' }}>Registrarse</button>
            </form>
        </div>
    );
};

export default Login;