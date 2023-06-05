import { lazy } from "react";
import {Login} from "pages/Login.jsx"
import {Register} from "pages/Register.jsx"

const Home = lazy(() => import("pages/Home.jsx"));

export const navigation = [
    {
        id: 0,
        path: "/",
        Element: Home,
    },
    {
        id: 1,
        path: "/home",
        Element: Home,
    },
    {
        id: 2,
        path: "/login",
        Element: Login,
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