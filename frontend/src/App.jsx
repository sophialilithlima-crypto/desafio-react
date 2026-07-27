import { BrowserRouter, Routes, Route } from "react-router-dom";

import Navbar from "./components/Navbar";

import Home from "./pages/Home";
import Categoria from "./pages/Categoria";
import Fornecedor from "./pages/Fornecedor";
import Produto from "./pages/Produto";


function App() {


    return (

        <BrowserRouter>


            <Navbar />


            <Routes>


                <Route 
                    path="/" 
                    element={<Home />} 
                />


                <Route 
                    path="/categorias" 
                    element={<Categoria />} 
                />


                <Route 
                    path="/fornecedores" 
                    element={<Fornecedor />} 
                />


                <Route 
                    path="/produtos" 
                    element={<Produto />} 
                />


            </Routes>


        </BrowserRouter>

    )

}


export default App;