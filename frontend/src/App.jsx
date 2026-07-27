import { BrowserRouter, Routes, Route } from "react-router-dom";

import Navbar from "./components/Navbar";

import Produto from "./pages/Produto";
import Categoria from "./pages/Categoria";
import Fornecedor from "./pages/Fornecedor";

function App() {
  return (
    <BrowserRouter>
      <Navbar />

      <Routes>
        <Route path="/" element={<Produto />} />
        <Route path="/produto" element={<Produto />} />
        <Route path="/categoria" element={<Categoria />} />
        <Route path="/fornecedor" element={<Fornecedor />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;