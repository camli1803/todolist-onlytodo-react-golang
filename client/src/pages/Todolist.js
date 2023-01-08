import {useState, useEffect} from "react"
import axios from "axios"

let endpoint = "http://localhost:1710"

function Todolist() {
    // create Todo 
    const todo = {
        id: 0,
        content: "",
        done: false
    }
    
    const handleCreateTodo = () => {
        console.log(todo)
        axios.post(endpoint + "/todos", todo)        
    }

    const handleUpdateTodo = () => {
        console.log(todo)
        console.log(todo.id)
        axios.patch(endpoint + "/todos/" + todo.id, todo)
        
    }
    
    const handleDeleteTodo = (id) => {
        console.log(id)
        axios.delete(endpoint + "/todos/" + id)
    }
    // get Todo
    const [todos, setTodos] = useState([])

    useEffect(()=>
    {  
         const GetAllTodos = async () => {
            const data = await axios.get(endpoint + "/todos")
            setTodos(data.data)
        }
        GetAllTodos()   
    }, [])

    const CreateTodo = () => (
        <div>
            <label htmlFor="content">Task</label>
            <input id="content" type="text" onChange={(e)=> todo.content = e.target.value}/>
            <label htmlFor="status">Done</label>
            <input id="status" type="checkbox" onChange={(e)=> todo.done = e.target.checked}/><br/>
            <button onClick={handleCreateTodo}>Create</button>
        </div> 
    )

    const UpdateTodo = () => (
        <div>
            <label htmlFor="id">ID</label>
            <input id= "id" type= "text" onChange={(e)=> todo.id = e.target.value}/>
            <label htmlFor="content">Task</label>
            <input id="content" type="text" onChange={(e)=> todo.content = e.target.value}/>
            <label htmlFor="status">Done</label>
            <input id="status" type="checkbox" onChange={(e)=> todo.done = e.target.checked}/><br/>
            <button onClick={handleUpdateTodo}>Update</button>
        </div> 
    )

    const ListTodo = () => (
        <ul style={{listStyleType: "none", fontSize: 20}}>
                    {todos.map(todo =>
                        <li key={todo.id}>
                            <pre>
                            <strong>ID:</strong> {todo.id}    
                            <strong>  Task:</strong> {todo.content}   
                            <strong>  Status:</strong> {todo.done === true ? "done" : "doing"} 
                            <br/>
                            <button onClick={() => handleDeleteTodo(todo.id)} style={{fontSize: 15}}>
                                Delete
                            </button>
                            </pre>
                        </li>)}
        </ul>
    ) 

    return (
       <div>       
            <CreateTodo/>
            {   
                todos == null ? <p style={{fontSize: 20}}>Create First Todo !!!</p> : <ListTodo/>
            }
            <UpdateTodo/>       
       </div>

    )
}

export default Todolist
