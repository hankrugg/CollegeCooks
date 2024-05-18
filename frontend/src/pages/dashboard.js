import React, {useEffect, useState} from 'react';
import axios from "axios";
import {Button} from "reactstrap";
import NavBar from "../components/navbar";


export default function Dashboard() {
    const [users, setUsers] = useState([])
    const [user, setUser] = useState("");


    useEffect(() => {
        getUser()
    }, []);


    const getUser = async function () {
        const response = await axios.get("http://localhost:3000/getUser" , { withCredentials: true })
        setUsers(response.data.users)
        setUser(response.data.user)
    }

    return (
        <div className="container-fluid">
            <div className="wrapper">
                <NavBar/>
            </div>
            <div className="centered-content">
                <h2>Hello, {user.first}!</h2>
                <p>This is your dashboard. You can add your content here.</p>
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
