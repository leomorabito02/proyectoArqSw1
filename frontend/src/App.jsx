import './App.css'
import React from 'react';
import Reserva from './pages/Reserva';

const App = () => {
  return (
    <div style={{ backgroundColor: '#CBE4DE', display: 'flex', height: '100vh', paddingBottom:'50px'}}>
      <Reserva/>
    </div>
  );
};

export default App;