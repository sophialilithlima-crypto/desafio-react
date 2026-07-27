import { useEffect, useState } from "react";
import api from "../api/api";
import "../styles/Categoria.css";


function Categoria() {


    const [categorias, setCategorias] = useState([]);

    const [nome, setNome] = useState("");

    const [editando, setEditando] = useState(null);

    const [mensagem, setMensagem] = useState("");





    async function carregarCategorias() {

        try {

            const response = await api.get("/categorias");

            setCategorias(response.data.data || []);

        } catch (error) {

            console.log(error);

        }

    }






    useEffect(() => {

        carregarCategorias();

    }, []);







    async function salvar(e) {

        e.preventDefault();


        if(!nome.trim()){

            setMensagem("Informe o nome da categoria");

            return;

        }



        try {


            if(editando){


                await api.put(`/categorias/${editando}`, {

                    nome

                });


                setMensagem("Categoria atualizada com sucesso");


            }else{


                await api.post("/categorias", {

                    nome

                });


                setMensagem("Categoria cadastrada com sucesso");


            }




            limparFormulario();

            carregarCategorias();



        } catch(error){


            console.log(error);


            setMensagem("Erro ao salvar categoria");


        }


    }








    function editar(categoria){


        setNome(categoria.nome);

        setEditando(categoria.id);

        setMensagem("");

    }








    async function excluir(id){


        const confirmar = window.confirm(
            "Deseja realmente excluir essa categoria?"
        );


        if(!confirmar){

            return;

        }




        try{


            await api.delete(`/categorias/${id}`);


            setMensagem(
                "Categoria removida com sucesso"
            );


            carregarCategorias();



        }catch(error){


            console.log(error);


            setMensagem(
                "Não foi possível excluir"
            );


        }


    }








    function limparFormulario(){


        setNome("");

        setEditando(null);


    }








    return (

        <div className="pagina">


            <h1>
                Categorias
            </h1>




            <div className="card">


                <form onSubmit={salvar}>


                    <input

                        type="text"

                        placeholder="Nome da categoria"

                        value={nome}

                        onChange={(e)=>setNome(e.target.value)}

                    />



                    <div className="botoes">


                        <button className="btn salvar">

                            {
                                editando 
                                ? 
                                "Atualizar"
                                :
                                "Cadastrar"
                            }

                        </button>




                        {
                            editando && (

                                <button

                                    type="button"

                                    className="btn cancelar"

                                    onClick={limparFormulario}

                                >

                                    Cancelar

                                </button>

                            )
                        }



                    </div>



                </form>




                {
                    mensagem &&

                    <p className="mensagem">

                        {mensagem}

                    </p>

                }


            </div>







            <div className="tabela-container">


                <table>


                    <thead>


                        <tr>


                            <th>ID</th>

                            <th>Nome</th>

                            <th>Ações</th>


                        </tr>


                    </thead>




                    <tbody>


                        {
                            categorias.map((categoria)=>(


                                <tr key={categoria.id}>


                                    <td>

                                        {categoria.id}

                                    </td>



                                    <td>

                                        {categoria.nome}

                                    </td>



                                    <td>


                                        <button

                                            className="btn editar"

                                            onClick={()=>
                                                editar(categoria)
                                            }

                                        >

                                            Editar

                                        </button>





                                        <button

                                            className="btn excluir"

                                            onClick={()=>
                                                excluir(categoria.id)
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





        </div>


    )



}



export default Categoria;