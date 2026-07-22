import { createBrowserRouter } from "react-router-dom";
import Home from "../pages/home";
import LoginPage from "../pages/login";
import RegisterPage from "../pages/register";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <Home />
  },
  {
    path: "/login",
    element: <LoginPage/>
  },
  {
    path: "/register",
    element: <RegisterPage/>
  },
]); 