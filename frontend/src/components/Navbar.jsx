import { Link } from "react-router-dom";
import "../styles/Navbar.css";


function Navbar(){

    return (

        <nav className="navbar">


            <h2>
                Sistema CRUD
            </h2>


            <div className="nav-links">


                <Link to="/">
                    Home
                </Link>


                <Link to="/categorias">
                    Categorias
                </Link>


                <Link to="/fornecedores">
                    Fornecedores
                </Link>


                <Link to="/produtos">
                    Produtos
                </Link>


            </div>


        </nav>

    )

}


export default Navbar;