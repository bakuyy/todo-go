import './App.css';
import { useState } from 'react';
import Form from "./components/TodoForm"

function App() {
  const [pressed,setPressed] = useState()

  return (
    <div className="App">
      <div className='mt-24'>To Do!</div>
      <Form/>
    </div>
  );
}

export default App;
