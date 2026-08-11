import { Link, useNavigate } from "react-router-dom";
import "../styles/Navbar.css";

function Navbar() {
    const navigate = useNavigate();

    function sair() {
        localStorage.removeItem("token");
        navigate("/login");
    }

    return (
        <nav className="navbar">
            <h2>Sistema CRUD</h2>

            <div className="nav-links">
                <Link to="/">Home</Link>
                <Link to="/categorias">Categorias</Link>
                <Link to="/fornecedores">Fornecedores</Link>
                <Link to="/produtos">Produtos</Link>
                <button onClick={sair}>Sair</button>
            </div>
        </nav>
    );
}

export default Navbar;
