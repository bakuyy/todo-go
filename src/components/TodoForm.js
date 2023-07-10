import React, { useState, useEffect } from 'react';
import axios from 'axios';

function TodoForm() {
  const [text, setText] = useState('');
  const [tasks, setTasks] = useState([]);

  useEffect(() => {
    getTasks()
  }, []);

  const handleChange = (event) => {
    const newVal = event.target.value;
    setText(newVal)
  };

  const handleSubmitButton = async () => {
    try {
      await axios.post('http://localhost:8080/tasks', { task: text });
      setText('');
      getTasks();
    } catch (error) {
      console.log('Error adding task:', error);
    }
  };

  const getTasks = async () => {
    try {
      const response = await axios.get('http://localhost:8080/tasks');
      setTasks(response.data);
    } catch (error) {
      console.log('Error loading tasks:', error);
    }
  };

  return (
    <div>
      <div className="flex mt-3 mb-12">
        <input
          className="text-sm p-2 border-2 border-gray-400 rounded-xl ml-96 "
          value={text}
          onChange={handleChange}
          placeholder="Add item to do :) ..."
        />
        <button
          className="ml-3 bg-gray-200 p-2 rounded-xl"
          onClick={handleSubmitButton}
        >
          Submit
        </button>
      </div>
      <ul>
        {tasks.map((task) => (
          <li className='bg-red-400 p-4 ml-96 mr-96 rounded-xl my-2' key={task.id}>{task.task}</li>
        ))}
      </ul>
    </div>
  );
}

export default TodoForm;