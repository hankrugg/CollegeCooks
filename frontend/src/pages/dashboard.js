import React, {useEffect, useState} from 'react';
import axios from "axios";
import {Button} from "reactstrap";
import NavBar from "../components/navbar";


export default function Dashboard() {
    const [users, setUsers] = useState([])
    const [user, setUser] = useState("");
    const [recipes, setRecipes] = useState([])



    useEffect(() => {
        getUser()
        getRecipes()
    }, []);


    const getUser = async function () {
        const response = await axios.get("http://192.168.0.98:3000/getUser" , { withCredentials: true })
        setUser(response.data.user)
    }

    const getRecipes = async function () {
        const response = await axios.get("http://192.168.0.98:3000/getRecipes" , { withCredentials: true })
        setRecipes(response.data.recipes)
    }

    return (
        <div className="container-fluid">
            <div className="wrapper">
                <NavBar/>
            </div>
            <div className="centered-content">
                <h2>Hello, {user.first} !</h2>
                <p>This is your dashboard. You can add your content here.</p>
                <div>
                    <ul style={{listStyleType: 'none', padding: 0}}>
                        {recipes.map(recipe => (
                            <li key={recipe.ID} style={{marginBottom: '20px'}}>
                                <div
                                    style={{
                                        border: '1px solid #ccc',
                                        borderRadius: '8px',
                                        boxShadow: '0px 4px 8px rgba(0, 0, 0, 0.1)',
                                        padding: '20px',
                                        backgroundColor: '#fff',
                                    }}
                                >
                                    <h3 style={{margin: 0}}>{recipe.title}</h3>
                                    <p style={{marginTop: '10px'}}>Ingredients: {recipe.ingredients}</p>
                                </div>
                            </li>
                        ))}
                    </ul>
                </div>

            </div>
        </div>
    );
}
