import './App.css';
import { useState } from 'react';
import Form from "./components/TodoForm"

function App() {
  const [pressed,setPressed] = useState()

  return (
    <div className="App">
      <div className='mt-24 text-xl font-bold'>TO DO!</div>
      <Form/>
    </div>
  );
}

export default App;
