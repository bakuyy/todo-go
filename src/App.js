import './App.css';
import { useState } from 'react';

function App() {
  const [pressed,setPressed] = useState()

  return (
    <div className="App">
      <div className='mt-24'>To Do!</div>
      <div className='flex mt-3 '>
        <input className='text-sm p-2 border-2 border-gray-400 rounded-xl  ' placeholder='Add item...'></input>
        <button className='ml-3 bg-gray-200 p-2 rounded-xl' onSubmit={setPressed()}> Submit</button>
      
      </div>
    </div>
  );
}

export default App;
