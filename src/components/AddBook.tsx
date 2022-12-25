import React, {useState} from "react"
import {Navbar} from "./Navbar";
import {AddingBook} from "../requests/AddBook";
import {Link, useLocation} from "react-router-dom";

export function AddBook() {
    const [name, setName] = useState('');
    const handleChangeName = (event: { target: { value: any; }; }) => {
        setName(event.target.value);
    };

    const [saleprice, setPrice] = useState(200);
    const handleChangeDiscount = (event: { target: { value: any; }; }) => {
        setPrice(Number(event.target.value));
    };

    const [year, setYear] = useState(0);
    const handleChangePrice = (event: { target: { value: any; }; }) => {
        setYear(Number(event.target.value));
    };

    const [srokgodnost, setQuantity] = useState(0);
    const handleChangeQuantity = (event: { target: { value: any; }; }) => {
        setQuantity(Number(event.target.value));
    };
    const [type, setType] = useState('');
    const handleChangeType = (event: { target: { value: any; }; }) => {
        setType((event.target.value));
    };

    const [color, setPromo] = useState('');
    const handleChangePromo = (event: { target: { value: string; }; }) => {
        setPromo((event.target.value));
    };

    const [description, setDesc] = useState('');
    const handleChangeDesc = (event: { target: { value: any; }; }) => {
        setDesc(event.target.value);
    };



    const [image, setImage] = useState('');
    const handleChangeImage = (event: { target: { value: any; }; }) => {
        setImage(event.target.value);
    };

    return (
        <>
            <Navbar/>

            <div className="bg-yellow-50 min-h-screen">
                <p className="ml-4 sm:text-2xl text-1xl font-normal text-black">
                    <Link to="/books" className="mr-2">
                        Book shop
                    </Link>
                    / adding
                </p>

                <p className="text-center sm:text-5xl text-3xl font-bold text-pink-500">
                    Добавление книги
                </p>

                <form className="mt-10 mx-5 bg-white rounded-lg border-2 border-teal-200">
                    <div className="grid grid-cols-4 grid-rows-2 gap-10 p-8">
                        <div className="">
                            <label htmlFor="first-name"
                                   className="block text-base font-medium text-gray-700">
                                Название
                            </label>
                            <input
                                type="text"
                                onChange={handleChangeName}
                                value={name}
                                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-base"
                            />
                        </div>

                        <div className="">
                            <label htmlFor="first-name"
                                   className="block text-base font-medium text-gray-700">
                                Цена
                            </label>
                            <input
                                type="number"
                                min="20"
                                max="2000"
                                onChange={handleChangeDiscount}
                                value={saleprice}
                                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-base"
                            />
                        </div>

                        <div className="">
                            <label htmlFor="first-name"
                                   className="block text-base font-medium text-gray-700">
                                Год издания
                            </label>
                            <input
                                type="number"
                                min="10"
                                max="1000"
                                onChange={handleChangePrice}
                                value={year}
                                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-base"
                            />
                        </div>

                        <div className="">
                            <label htmlFor="first-name"
                                   className="block text-base font-medium text-gray-700">
                                Срок годности
                            </label>
                            <input
                                type="number"
                                min="1"
                                onChange={handleChangeQuantity}
                                value={srokgodnost}
                                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-base"
                            />
                        </div>

                        <div className="col-span-2">
                            <label htmlFor="first-name"
                                   className="block text-base font-medium text-gray-700">
                                Тип книги
                            </label>

                            <input
                                type="text"
                                onChange={handleChangeType}
                                value={type}
                                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-base"
                            />
                        </div>


                        <div className="col-span-2">
                            <label htmlFor="first-name"
                                   className="block text-base font-medium text-gray-700">
                                Цвет
                            </label>

                            <input
                                type="text"
                                onChange={handleChangePromo}
                                value={color}
                                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-base"
                            />
                        </div>

                        <div className="col-span-2">
                            <label htmlFor="first-name"
                                   className="block text-base font-medium text-gray-700">
                                Описание
                            </label>

                            <input
                                type="text"
                                onChange={handleChangeDesc}
                                value={description}
                                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-base"
                            />
                        </div>

                        <div className="col-span-2">
                            <label htmlFor="first-name"
                                   className="block text-base font-medium text-gray-700">
                                Изображение
                            </label>
                            <input
                                type="text"
                                onChange={handleChangeImage}
                                value={image}
                                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-base"
                            />
                        </div>
                    </div>

                    <div className="text-center mb-6">
                        {AddingBook(name, saleprice,year,type,srokgodnost,color,description,image)}
                    </div>
                </form>

                <p className="py-8 text-center">
                    <Link to="/books"
                          className="border-4 border-blue-700 text-blue-700 hover:bg-blue-700 hover:text-white py-1 px-3 rounded-full text-2xl font-bold"
                    >
                        Обратно на главную
                    </Link>
                </p>
            </div>
        </>
    )
}