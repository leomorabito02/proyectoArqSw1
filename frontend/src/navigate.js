import { lazy } from "react";
import {Home} from "pages/Home.jsx"

const Home = lazy(() => import("pages/Home"));
const Hotels = lazy(() => import("pages/Hotels"));


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
        Element: Hotels,
    },
    {
        id: 5,
        path: "/reserva", //a chequear
        Element: Reserva,
    },
];

export {Home};