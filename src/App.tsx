import {Routes, Route, BrowserRouter} from 'react-router-dom'
import {HomePage} from "./components/HomePage"
import {Info} from "./components/Info"
import {Payment} from "./components/Payment"
import {NotFound} from "./components/NotFound";
import {CartPage} from "./components/CartPage";
import {Registration} from "./components/RegisterPage";
import {LoginPage} from "./components/LoginPage";
import {ProfilePage} from "./components/ProfilePage";
import {AddBook} from "./components/AddBook";
import {OrderPage} from "./components/OrderPage";
import {ChangeBook} from "./components/ChangeBook";

export const ENDPOINT = "http://localhost:8080"

export default function App() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/books" element={<HomePage/>}/>
                <Route path="/cart" element={<CartPage/>}/>
                <Route path="/info" element={<Info/>}/>
                <Route path="/payment" element={<Payment/>}/>
                <Route path="*" element={<NotFound/>}/>
                <Route path="/login" element={<LoginPage/>}/>
                <Route path="/registration" element={<Registration/>}/>
                <Route path="/add" element={<AddBook/>}/>
                <Route path="/profile" element={<ProfilePage/>}/>
                <Route path="/orders" element={<OrderPage/>}></Route>
                <Route path="/change" element={<ChangeBook/>}/>
            </Routes>
        </BrowserRouter>
    )
}