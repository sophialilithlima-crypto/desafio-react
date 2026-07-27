import { Link } from "react-router-dom";
import "./Navbar.css";

function Navbar() {
    return (
        <nav>
            <h2>Sistema CRUD</h2>

            <div>
                <Link to="/produto">Produtos</Link>

                <Link to="/categoria">Categorias</Link>

                <Link to="/fornecedor">Fornecedores</Link>
            </div>
        </nav>
    );
}

export default Navbar;