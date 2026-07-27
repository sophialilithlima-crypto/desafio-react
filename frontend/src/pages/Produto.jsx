import { useEffect, useState } from "react";
import api from "../api/api";


function Produto() {


    const [produtos, setProdutos] = useState([]);


    const [editando, setEditando] = useState(false);


    const [idEditando, setIdEditando] = useState(null);



    const [formulario, setFormulario] = useState({

        nome: "",
        sku: "",
        preco: "",
        categoria_id: ""

    });





    async function buscarProdutos() {

        try {


            const resposta = await api.get("/produtos");


            setProdutos(resposta.data.data || []);



        } catch (erro) {


            console.log(
                "Erro ao buscar produtos:",
                erro
            );


        }

    }






    function alterarCampo(e) {


        setFormulario({

            ...formulario,

            [e.target.name]: e.target.value

        });


    }








    async function salvarProduto(e) {


        e.preventDefault();



        const produto = {


            nome: formulario.nome,

            sku: formulario.sku,

            preco: Number(formulario.preco),

            categoria_id: Number(formulario.categoria_id)


        };




        try {



            if (editando) {



                await api.put(

                    `/produtos/${idEditando}`,

                    produto

                );



                alert(
                    "Produto atualizado com sucesso!"
                );



            } else {



                await api.post(

                    "/produtos",

                    produto

                );



                alert(
                    "Produto cadastrado com sucesso!"
                );


            }






            limparFormulario();


            buscarProdutos();




        } catch (erro) {



            console.log(
                "Erro ao salvar produto:",
                erro
            );



            alert(
                "Erro ao salvar produto"
            );



        }


    }









    function editarProduto(produto) {


        setEditando(true);


        setIdEditando(produto.id);



        setFormulario({

            nome: produto.nome,

            sku: produto.sku,

            preco: produto.preco,

            categoria_id: produto.categoria_id

        });


    }








    function limparFormulario() {


        setFormulario({

            nome: "",

            sku: "",

            preco: "",

            categoria_id: ""

        });



        setEditando(false);


        setIdEditando(null);


    }









    async function excluirProduto(id) {


        const confirmar = window.confirm(
            "Deseja realmente excluir este produto?"
        );



        if (!confirmar) {

            return;

        }






        try {


            await api.delete(

                `/produtos/${id}`

            );



            alert(
                "Produto excluído com sucesso!"
            );



            buscarProdutos();




        } catch (erro) {


            console.log(
                "Erro ao excluir produto:",
                erro
            );



            alert(
                "Erro ao excluir produto"
            );


        }


    }








    useEffect(() => {


        buscarProdutos();


    }, []);









    return (

        <div>


            <h1>

                Produtos

            </h1>





            <h2>

                {
                    editando
                    ?
                    "Editar Produto"
                    :
                    "Cadastrar Produto"
                }

            </h2>






            <form onSubmit={salvarProduto}>


                <input

                    type="text"

                    name="nome"

                    placeholder="Nome"

                    value={formulario.nome}

                    onChange={alterarCampo}

                />




                <input

                    type="text"

                    name="sku"

                    placeholder="SKU"

                    value={formulario.sku}

                    onChange={alterarCampo}

                />




                <input

                    type="number"

                    name="preco"

                    placeholder="Preço"

                    value={formulario.preco}

                    onChange={alterarCampo}

                />




                <input

                    type="number"

                    name="categoria_id"

                    placeholder="Categoria ID"

                    value={formulario.categoria_id}

                    onChange={alterarCampo}

                />






                <button type="submit">


                    {
                        editando
                        ?
                        "Atualizar"
                        :
                        "Cadastrar"
                    }


                </button>





                {

                    editando &&

                    <button

                        type="button"

                        onClick={limparFormulario}

                    >

                        Cancelar

                    </button>


                }



            </form>









            <h2>

                Produtos cadastrados

            </h2>






            <table border="1">


                <thead>


                    <tr>

                        <th>ID</th>

                        <th>Nome</th>

                        <th>SKU</th>

                        <th>Preço</th>

                        <th>Categoria</th>

                        <th>Ações</th>


                    </tr>


                </thead>





                <tbody>


                    {


                        produtos.map((produto)=>(



                            <tr key={produto.id}>


                                <td>

                                    {produto.id}

                                </td>




                                <td>

                                    {produto.nome}

                                </td>




                                <td>

                                    {produto.sku}

                                </td>




                                <td>

                                    R$ {produto.preco}

                                </td>




                                <td>

                                    {produto.categoria_id}

                                </td>





                                <td>


                                    <button

                                        onClick={() =>
                                            editarProduto(produto)
                                        }

                                    >

                                        Editar

                                    </button>





                                    <button

                                        onClick={() =>
                                            excluirProduto(produto.id)
                                        }

                                    >

                                        Excluir

                                    </button>



                                </td>




                            </tr>



                        ))


                    }


                </tbody>


            </table>




        </div>


    );


}



export default Produto;