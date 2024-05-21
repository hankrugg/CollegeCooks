import { Sidebar, Menu, MenuItem } from 'react-pro-sidebar';
import {Link, useNavigate} from 'react-router-dom';
import axios from "axios";



export default function NavBar() {
    const navigate = useNavigate()


    const logout = async function () {
        const response = await axios.get("http://192.168.0.98:3000/logout" , { withCredentials: true })
    }

    return(
        <div>
            <Sidebar>
                <Menu
                    menuItemStyles={{
                        button: {
                            // the active class will be added automatically by react router
                            [`&.active`]: {
                                backgroundColor: '#13395e',
                                color: '#b6c8d9',
                            },
                        },
                    }}
                >
                    <MenuItem component={<Link to="/pantry"/>}> Pantry</MenuItem>
                    <MenuItem component={<Link to="/dashboard"/>}> Dashboard</MenuItem>
                    <MenuItem component={<Link to="/login"/>}>Logout</MenuItem>
                </Menu>
            </Sidebar>
        </div>
    );
}