import {addBook} from "../modules";
import React from "react";


export function AddingBook( name: string, saleprice: number, year: number, type: string, Srokgodnost: number, color: string, description: string, image: string) {

    const url = `books`

    function Add() {
        addBook(url, name, saleprice, year, type, Srokgodnost, color, description, image)
    }


    return (
        <>
            <button
                onClick={() => Add()}
                className="border-4 border-red-500 bg-white text-red-500 hover:bg-red-500 hover:text-white py-1 px-2 place-self-center rounded-full text-2xl font-bold"
            >
                Добавить
            </button>
        </>
    );

}