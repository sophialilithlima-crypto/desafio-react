import "../styles/Home.css";


function Home(){

    return (

        <div className="home">


            <h1>
                Sistema de Gerenciamento
            </h1>


            <p>
                Bem-vindo ao sistema CRUD de produtos,
                categorias e fornecedores.
            </p>


            <div className="cards">


                <div className="card-home">

                    <h2>
                        Categorias
                    </h2>

                    <p>
                        Cadastre e gerencie categorias.
                    </p>

                </div>



                <div className="card-home">

                    <h2>
                        Fornecedores
                    </h2>

                    <p>
                        Cadastre e gerencie fornecedores.
                    </p>

                </div>



                <div className="card-home">

                    <h2>
                        Produtos
                    </h2>

                    <p>
                        Controle seus produtos e estoque.
                    </p>

                </div>


            </div>


        </div>

    )

}


export default Home;