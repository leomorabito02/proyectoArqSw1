import './App.css'
import {BrowserRouter as Router, Route, Routes} from 'react-router-dom';
import React from 'react';
import Header from './components/Header';
import Login from './pages/Login.jsx';
import Footer from './components/Footer.jsx'
import Register from "./pages/Register.jsx";
import Home from "./pages/Home.jsx";
import Reserva from "./pages/Reserva.jsx";

const App = () => {
  return (
      <div>
          <Router>
              <Header />
               <Routes>
                      <Route path="/" element={<Login />} />
                      <Route path="/login" element={<Login />} />
                      <Route path="/register" element={<Register />} />
                       <Route path="/home" element={<Home/>}/>
                      <Route path="reserva" element={<Reserva />} />
              </Routes>
              <Footer />
          </Router>
      </div>
  );
};

export default App
