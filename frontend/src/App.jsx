import './App.css';
import React from 'react';
import Home from './pages/Home';
import Hotels from './pages/Hotels';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';


const App = () => {
  return (
    
    <div style={{ backgroundColor: '#CBE4DE', display: 'flex', marginRight:'60px', height: '100vh', paddingBottom:'50px'}}>
      
      <Home/>
    </div>
  );
};

export default App;


