import {useEffect, useReducer} from "react";
import {getFromBackendToken} from "../modules";

const initialState = {book: ""}
const success = "Success"

function reducer(state: any, action: { type: any; book: any; }) {
    switch (action.type) {
        case success:
            return {
                book: action.book
            }
        default:
            return state
    }
}

export function GetBook(uuid: string) {
    const [state, dispatch] = useReducer(reducer, initialState)
    const url = `book/${uuid}`

    useEffect(() => {
        getFromBackendToken(url).then((result) => {
            dispatch({type: success, book: result})
        })
    }, [url])

    return state.book

}