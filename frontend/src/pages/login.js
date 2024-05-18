import React, { useState } from 'react'
import {Button, Input} from "reactstrap";
import { useNavigate } from 'react-router-dom'
import "../styles/styles.css"
import axios from "axios";

export default function LoginPage() {
    const [emailError, setEmailError] = useState('')
    const [passwordError, setPasswordError] = useState('')

    const navigate = useNavigate()

    const Submit = async (e) => {
        e.preventDefault()
        const formData = new FormData(e.currentTarget)
        const email = formData.get('email')
        const password = formData.get('password')

        axios.post("http://localhost:3000/login", { email: email, password: password })
            .then(response => {
                // Handle successful response
                console.log("Response status:", response.status);
                if (response.status === 200) {
                    // Authentication successful, navigate to dashboard or perform other actions
                    navigate("/dashboard")
                } else {
                    // Handle unexpected response status codes
                    console.error("Unexpected response status:", response.status);
                }
            })
            .catch(error => {
                // Handle error
                console.error("Error:", error);
                if (error.response.status === 401) {
                    // Authentication failed, display error message to the user
                    setPasswordError('Username and password combination not found');
                } else {
                    // Handle other types of errors (e.g., network error)
                    // Display appropriate error message or take necessary action
                    setPasswordError('Error. Please contact administration.');

                }
            });
    }

    return (
        <div className="container">
            <h1 className="Ubuntu Sans">College Cooks</h1>
            <br/>
            <form onSubmit={Submit}>
                <div className="input-group-lg">
                    <Input
                        type="text" id="email" name="email" label="Email" placeholder="Email"
                    />
                </div>
                <br/>
                <div className="input-group-lg">
                    <Input
                        type="password" id="password" name="password" label="Password" placeholder="Password"
                    />
                </div>
                <br/>
                <Button
                    value="Submit" label="Login"
                >
                    Login
                </Button>
            </form>
            <div>
                {passwordError && <span>{passwordError}</span>}
            </div>

        </div>
    )
}