import { BrowserRouter, Routes, Route } from "react-router-dom";
import ReactDOM from "react-dom/client";
import LoginPage from "./pages/login";
import Dashboard from "./pages/dashboard";

export default function App() {
  return (
      <BrowserRouter>
        <Routes>
            <Route path={"/"} element={<LoginPage/>}></Route>
            <Route path={"/login"} element={<LoginPage/>}></Route>
            <Route path={"/dashboard"} element={<Dashboard/>}></Route>
        </Routes>
      </BrowserRouter>
  );
}
