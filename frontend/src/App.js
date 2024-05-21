import { BrowserRouter, Routes, Route } from "react-router-dom";
import LoginPage from "./pages/login";
import Dashboard from "./pages/dashboard";
import RegisterPage from "./pages/register";
import Pantry from "./pages/pantry";


export default function App() {
  return (
      <BrowserRouter>
        <Routes>
            <Route path={"/"} element={<LoginPage/>}></Route>
            <Route path={"/login"} element={<LoginPage/>}></Route>
            <Route path={"/dashboard"} element={<Dashboard/>}></Route>
            <Route path={"/register"} element={<RegisterPage/>}></Route>
            <Route path={"/pantry"} element={<Pantry/>}></Route>

        </Routes>
      </BrowserRouter>
  );
}
