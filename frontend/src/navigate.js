import { lazy } from "react";
import {Login} from "pages/Login.jsx"
import {Register} from "pages/Register.jsx"
import {Home} from "pages/Home.jsx"

const Home = lazy(() => import("pages/Home.jsx"));
//estas son las rutas de las paginas de nuestro programa
export const navigation = [
    {
        id: 0,
        path: "/",// el login es la raiz porque es la principal, si no pasa el login, no debe entrar al resto
        Element: Login,
    },
    {
        id: 1,
        path: "/login",
        Element: Login,
    },
    {
        id: 2,
        path: "/home",
        Element: Home,
    },
    {
        id: 3,
        path: "/register",
        Element:Register,
    },
    {
        id: 4,
        path: "/hotel/:id",
        Element: Hotel,
    },
    {
        id: 5,
        path: "/reserva", //a chequear
        Element: Reserva,
    },
];