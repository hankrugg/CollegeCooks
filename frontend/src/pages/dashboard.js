import React, {useState} from 'react';
import axios from "axios";
import {Button} from "reactstrap";
import Sidebar from "../components/navbar";
import NavBar from "../components/navbar";


export default function Dashboard() {
    const [users, setUsers] = useState([])

    const getUsers = async function () {
        const response = await axios.get("http://localhost:3000/getUsers")
        setUsers(response.data)
        console.log(response.data)
    }

    return (
        <div className="container-fluid">
            <div className="wrapper">
                <NavBar/>
            </div>
            <div className="centered-content">
                <h2>Dashboard</h2>
                <p>This is your dashboard. You can add your content here.</p>
                <Button label="getUsers" onClick={getUsers}>
                    GetUsers
                </Button>
                <div>
                    <ul>
                        {users.map(user => (
                            <li key={user.ID}>{user.email}</li>
                        ))}
                    </ul>
                </div>
            </div>
        </div>
    );
}
