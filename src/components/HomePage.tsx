import {Books} from "./Books"
import {GetBooks, ContextBook} from "../requests/GetBooks";
import {IBook} from "../models";
import React, {useState} from "react";
import {Box, Slider} from "@mui/material";
import {Navbar} from "./Navbar"

export function HomePage() {

    const books= GetBooks()

    const [name, setName] = useState('')

    const filteredStores = books.filter((book: { Name: string }) => {
        return book.Name.toLowerCase().includes(name.toLowerCase())
    })

    const [price, setPrice] = React.useState<number[]>([0, 1000])

    const minDistance = 10;

    const handleChange = (event: Event, newValue: number | number[], activeThumb: number,) => {
        if (!Array.isArray(newValue)) {
            return
        }

        if (activeThumb === 0) {
            setPrice([Math.min(newValue[0], price[1] - minDistance), price[1]])
        } else {
            setPrice([price[0], Math.max(newValue[1], price[0] + minDistance)])
        }
    }

    const marks = [
        {
            value: 0,
            label: '0 ₽',
        },
        {
            value: 250,
            label: '250 ₽',
        },
        {
            value: 500,
            label: '500 ₽',
        },
        {
            value: 750,
            label: '750 ₽',
        },
        {
            value: 1000,
            label: '1000 ₽',
        },
    ];

    function valuetext(price: number) {
        return `${price} Р`;
    }

    return (
        <><Navbar/>
        <div className="bg-yellow-50 min-h-screen">
            <p className="ml-4 text-2xl font-normal text-black">
                Book shop
            </p>

            <p className="text-center text-5xl font-bold text-pink-500">
                Книжный магазин
            </p>

            <div className="mt-5 mob:mt-2">
                <div className="flex place-content-center">
                    <Box sx={{width: 300}}>
                        <Slider
                            aria-label="Price filter"
                            valueLabelDisplay="auto"
                            getAriaValueText={valuetext}
                            value={price}
                            marks={marks}
                            onChange={handleChange}
                            disableSwap
                            step={10}
                            min={0}
                            max={1000}
                        />
                    </Box>
                </div>

                <div className="mx-auto">
                    <form>
                        <input
                            type="text"
                            className="block w-full px-4 py-2 text-gray-500 text-2xl bg-white border rounded-full focus:border-gray-400 focus:ring-gray-400 focus:outline-none focus:ring focus:ring-opacity-40"
                            placeholder="Поиск..."
                            onChange={(event) => setName(event.target.value)}
                        />
                    </form>
                </div>
            </div>


            <div className="pt-5 flex flex-col gap-4 mx-auto container ">
                {filteredStores.filter((book: { Saleprice: number; }) => book.Saleprice >= price[0] && book.Saleprice <= price[1]).map((book: IBook) => {
                    return (
                        <ContextBook.Provider value={book}>
                            <Books/>
                        </ContextBook.Provider>
                    )
                })}
            </div>
        </div>
        </>
    )
}