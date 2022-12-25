import {changeBook} from "../modules";
import React from "react";


export function ChangingBook(uuid: string, name: string, Saleprice: number, year: number, Type: number, Srokgodnost: number, color: string, description: string, image: string) {


    const url = `books`

    function Change() {
        changeBook(uuid, url, name, Saleprice, year, Type, Srokgodnost,color, description, image)
    }


    return (
        <>
            <button
                onClick={() => Change()}
                className="border-4 border-red-500 bg-white text-red-500 hover:bg-red-500 hover:text-white py-1 px-2 place-self-center rounded-full text-2xl font-bold"
            >
                Изменить
            </button>
        </>
    );

}