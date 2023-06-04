import './App.css'
import React from 'react';
import Header from './components/Header';
import Login from './pages/Login';
import Footer from './components/Footer.jsx'


const App = () => {
    return (

        <div >
            <Header/>
            <Login/>
            <Footer/>
        </div>
    );
};

export default App;