import React, {useEffect, useState} from 'react';
import axios from "axios";
import {Button, Input} from "reactstrap";
import NavBar from "../components/navbar";


export default function Dashboard() {
    const [ingredients, setIngredients] = useState([])
    const [user, setUser] = useState("");


    useEffect(() => {
        getUser()
        getIngredients()
    }, []);


    const getUser = async function () {
        const response = await axios.get("http://192.168.0.98:3000/getUser", {withCredentials: true})
        setUser(response.data.user)
    }

    const getIngredients = async function () {
        const response = await axios.get("http://192.168.0.98:3000/getIngredients", {withCredentials: true})
        setIngredients(response.data.ingredients)
    }

    const addToPantry = async (e) => {
        e.preventDefault()
        const formData = new FormData(e.currentTarget)
        const name = formData.get('name')
        const quantity = parseInt(formData.get('quantity'));
        const fruit = formData.get('category') === "fruit"
        const vegetable = formData.get('category') === "vegetable"
        const meat = formData.get('category') === "meat"
        const grain = formData.get('category') === "grain"

        const response = await axios.post("http://192.168.0.98:3000/addIngredient",
            {Name:name,
            Quantity:quantity,
            Fruit:fruit,
            Vegetable:vegetable,
            Meat:meat,
            Grain:grain,
            User:user,
                UserID: parseInt(user.ID)},
            {withCredentials: true})
        window.location.reload();
    }

    return (
        <div className="container-fluid">
            <div className="wrapper">
                <NavBar/>
            </div>
            <div className="centered-content">
                <h2>Welcome to your pantry, {user.first}!</h2>
                <form onSubmit={addToPantry}>
                    <div className="container-fluid">
                        <Input type="text" id="name" name="name" label="Name" placeholder="Name"/>
                        <Input type="number" id="quantity" name="quantity" label="Quantity" placeholder="Quantity"/>
                        <select id="category" name="category">
                            <option value="grain">Grain</option>
                            <option value="fruit">Fruit</option>
                            <option value="vegetable">Vegetable</option>
                            <option value="meat">Meat</option>
                        </select>
                    </div>
                    <Button type="submit" label="Add to pantry">Add to pantry</Button>
                </form>

                <div>
                    {ingredients.map(ingredient => (
                        <div key={ingredient.ID} className="card">
                            <div className="card-body">
                                <h5 className="card-title">{ingredient.Name}</h5>
                                <p className="card-text">Quantity: {ingredient.Quantity}</p>
                            </div>
                        </div>
                    ))}
                </div>

            </div>
        </div>
    );
}