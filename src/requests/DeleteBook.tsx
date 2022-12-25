import {deleteBook} from "../modules";


export function DeleteBook(uuid: string) {

    const url = `store`

    function Delete() {
        deleteBook(url, uuid)
    }


    return (
        <form>
            <button onClick={() => Delete()}>Удалить магазин</button>
        </form>
    );

}
