import {ICart, IOrder, IBook} from "./models";

export let BookContext: IBook = {
    UUID: "",
    Name: "",
    Saleprice: 0,
    Year: 0,
    Type: "",
    Srokgodnost: 0,
    Color: "",
    Description: "",
    Image: "",
}

export let CartContext: ICart = {
    BookUUID: "",
    Quantity: 0
}
export let OrderContext: IOrder = {
    UUID: "",
    Book: "",
    Quantity: 0,
    UserUUID: "",
    Date: "",
    Status: "",
}