import {ChangeCart} from "../requests/ChangeCart";
import {GetBook} from "../requests/GetBook";
import {Link} from "react-router-dom";
import {useContext} from "react";
import {ContextCart} from "../requests/GetCart";

export function Cart() {
    const ctx = useContext(ContextCart)
    let Book = GetBook(ctx.BookUUID)

    return (
        <div
            className="border-2 border-teal-200 mx-auto mt-4 w-1/2 h-40 py-5 px-5 rounded-lg grid grid-rows-2 grid-cols-3 bg-white"
        >
            <img src={Book.Image}
                 className="place-self-center object-contain h-20 w-20" alt={Book.Name}
            />



            <p className="text-blue-700 place-self-center text-2xl font-bold">
                Кол-во:{" "}{ctx.Quantity}{" "}{ChangeCart(ctx.BookUUID)}
            </p>


            <p className="text-red-600 place-self-center text-3xl font-bold">
                {Book.Name}
            </p>

            <p className="text-yellow-400 place-self-center text-2xl font-bold">
                {Book.Saleprice} ₽/шт
            </p>

            <Link to={`/payment`}
                  className="border-4 border-orange-500 text-orange-500 hover:bg-orange-500 hover:text-white py-1 px-3 place-self-center rounded-full text-2xl font-bold"
                  state={{Name: Book.Name, Book: ctx.BookUUID, Quantity: ctx.Quantity}}
            >
                Купить
            </Link>
        </div>
    )
}