import React from 'react'
import { useState } from 'react'
import axios from 'axios'

function TodoForm() {
    const [Text, setText] = useState('')
    const [storedText, setStoredText] = useState('')

    function HandleChange(event) {
        const newVal = event.target.value
        setText(newVal)
    }

    const HandleSubmitButton = async ()=> {
      try {
        await axios.post('http://localhost:8080/tasks',{ task })
        setText('')
      } catch (error) {
        console.log("error adding task", error)
      }

    }

  return (
    <div>
    <div className='flex mt-3 '>
    <input className='text-sm p-2 border-2 border-gray-400 rounded-xl' value={Text} onChange={HandleChange} placeholder='Add item...'></input>
    <button className='ml-3 bg-gray-200 p-2 rounded-xl' onClick={HandleSubmitButton}> Submit</button>
  
  </div>
    </div>
  )
}

export default TodoForm
