import React, { useState } from 'react'
import {Button, Input} from "reactstrap";
import { useNavigate } from 'react-router-dom'
import "../styles/styles.css"

export default function LoginPage() {
    const [emailError, setEmailError] = useState('')
    const [passwordError, setPasswordError] = useState('')

    const navigate = useNavigate()

    const Submit = async (e) => {
        e.preventDefault()
        const formData = new FormData(e.currentTarget)
        const email = formData.get('email')
        const password = formData.get('password')
        navigate("/dashboard")
    }

    return (
        <div className="container">
                <h1 className="Ubuntu Sans" >College Cooks</h1>
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
            <br />
                <Button
                    value="Submit" label="Login"
                >
                    Login
                </Button>
            </form>
        </div>
    )
}