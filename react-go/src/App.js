import React from 'react';
import './App.css';
import axios from "axios"

class App extends React.Component {
  constructor() {
    super()
    this.state = { todoList: []}
  }


  renderTodoList(list){
    const todoList = list.map(todo => {
    return (
      <ul>
      <li>{todo.id}</li>
      <li>{todo.title}</li>
      <li>{todo.contents}</li>
    </ul>
    )
    })
    return todoList
  }

  render() {
    console.log(this.state.todoList)
    return (
      <div>
        <h1>Todo App</h1>
        <h2>list</h2>
        <button onClick={this.getTodos}>Get Todo List</button>
        <div>{this.renderTodoList(this.state.todoList)}</div>
      </div>
    );
  }

  getTodos = () =>{
    const url = "http://localhost:8080/todos/"
    axios.get(url).then(res => {
          this.setState({todoList: res.data})
    })
  }
}

export default App;
