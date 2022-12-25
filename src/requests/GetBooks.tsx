import {createContext, useEffect, useReducer} from "react";
import {getFromBackend} from "../modules";
import {BookContext} from "../context";


export const ContextBook = createContext(BookContext);
const initialState = {books: []}
const success = "Success"

function reducer(state: any, action: { type: any; books: any; }) {
    switch (action.type) {
        case success:
            return {
                books: action.books
            }
        default:
            return state
    }
}

export function GetBooks() {
    const [state, dispatch] = useReducer(reducer, initialState)
    const url = `books`

    useEffect(() => {
        getFromBackend(url).then((result) => {
            dispatch({type: success, books: result})
        })
    }, [url])

    return state.books
}