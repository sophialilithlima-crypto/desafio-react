import { BrowserRouter, Navigate, Route, Routes } from "react-router-dom";

import Navbar from "./components/Navbar";
import Home from "./pages/Home";
import Categoria from "./pages/Categoria";
import Fornecedor from "./pages/Fornecedor";
import Produto from "./pages/Produto";
import Login from "./pages/Login";

function Protegido({ children }) {
    const token = localStorage.getItem("token");

    if (!token) {
        return <Navigate to="/login" replace />;
    }

    return children;
}

function App() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/login" element={<Login />} />

                <Route
                    path="*"
                    element={
                        <Protegido>
                            <>
                                <Navbar />
                                <Routes>
                                    <Route path="/" element={<Home />} />
                                    <Route path="/categorias" element={<Categoria />} />
                                    <Route path="/fornecedores" element={<Fornecedor />} />
                                    <Route path="/produtos" element={<Produto />} />
                                    <Route path="*" element={<Navigate to="/" replace />} />
                                </Routes>
                            </>
                        </Protegido>
                    }
                />
            </Routes>
        </BrowserRouter>
    );
}

export default App;
