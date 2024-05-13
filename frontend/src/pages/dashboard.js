import React, {useState} from 'react';
import axios from "axios";
import {Button} from "reactstrap";


export default function Dashboard() {
    const [users, setUsers] = useState('')
    const config = {
        headers: {
            "Access-Control-Allow-Origin": "*",
            "Access-Control-Allow-Methods": "GET,PUT,POST,DELETE,PATCH,OPTIONS"
        }
    };

    const getUsers = async function () {
        const response = await axios.get("http://localhost:3000/getUsers")
        console.log(response)
    }

    return (
        <div className="container-fluid">
            <nav className="navbar navbar-expand-lg navbar-light bg-light">
                <div className="container-fluid">
                </div>
            </nav>
            <div className="container mt-5">
                <div className="row">
                    <div className="col">
                        <h2>Dashboard</h2>
                        <p>This is your dashboard. You can add your content here.</p>
                        <Button label="getUsers" onClick={ getUsers }>
                            GetUsers
                        </Button>

                    </div>
                </div>
            </div>
        </div>
    );
}
