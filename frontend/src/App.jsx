import './App.css'
import React from 'react';
import Header from './components/Header';
import Home from './pages/Home';


const App = () => {
  return (
    
    <div style={{ backgroundColor: '#CBE4DE', display: 'flex', marginRight:'60px', height: '100vh', paddingBottom:'50px'}}>
      <Header/>
      <Home/>
    </div>
  );
};

export default App;


