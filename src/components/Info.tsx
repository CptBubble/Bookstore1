import {Link} from "react-router-dom";
import {Navbar} from "./Navbar";
import React from "react";
export function Info() {
    return (
        <>
        <Navbar/>
        <div className="bg-yellow-50 min-h-screen">
            <p className="ml-4 text-2xl font-normal text-black">
                <Link to="/books" className="mr-2">
                   Book shop
                </Link>
                / info
            </p>

            <p className="text-center font-bold text-5xl text-pink-500">
                Book shop
            </p>

            <p className="text-center mt-4 mx-8 font-medium text-3xl text-indigo-700">
                Это магазин книг, где вы можете купить книги.
                Читайте вместе с нами!
            </p>

            <img src="https://res.cloudinary.com/dh4qv3hob/image/upload/v1667665906/Promos/Discount_hgs7rl.png" width="29%" className="mx-auto" alt="Discount"/>
        </div>
        </>
    )
}